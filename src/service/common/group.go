package common

import (
	"time"

	"github.com/0xunion/exercise_back/src/model"
	"github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/auth"
	"github.com/0xunion/exercise_back/src/util/strings"
)

func CreateGroupService(user *types.User, name string, desc string) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	// create group
	group := &types.Group{
		Name:        name,
		Description: desc,
		Parent:      user.Id,
		CreateAt:    time.Now().Unix(),
	}

	var gid types.PrimaryId
	if err := model.ModelInsert(group, &gid); err != nil {
		return types.ErrorResponse(-1, err.Error())
	}

	// create group admin
	admin := &types.GroupMember{
		Gid:      gid,
		Uid:      user.Id,
		CreateAt: time.Now().Unix(),
	}
	admin.SetAdmin()

	if err := model.ModelInsert(admin, nil); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(gid)
}

func ListMyGroupsService(user *types.User, index int64, limit int64) *types.MasterResponse {
	members, err := model.ModelGetAll[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("uid", user.Id),
		),
		&model.MongoOptions{
			Skip:  &index,
			Limit: &limit,
		},
	)

	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	var groups []*types.Group
	for _, member := range members {
		group, err := model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(member.Gid),
			),
		)
		if err != nil {
			return types.ErrorResponse(-500, err.Error())
		}

		groups = append(groups, group)
	}

	return types.SuccessResponse(groups)
}

func ListGroupsService(user *types.User, index int64, limit int64) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	list, err := model.ModelGetAll[types.Group](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
		),
		&model.MongoOptions{
			Skip:  &index,
			Limit: &limit,
		},
	)

	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(list)
}

func InfoGroupService(user *types.User, id types.PrimaryId) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() != user.Id.String() {
		allowed = true
	}

	if _, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		allowed = true
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	return types.SuccessResponse(group)
}

func UpdateGroupService(user *types.User, id types.PrimaryId, field string, value interface{}) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	if member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		if member.IsAdmin() {
			allowed = true
		}
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	switch field {
	case "name":
		// check type
		if _, ok := value.(string); !ok {
			return types.ErrorResponse(-400, "invalid value type")
		}
		group.Name = value.(string)
	case "description":
		// check type
		if _, ok := value.(string); !ok {
			return types.ErrorResponse(-400, "invalid value type")
		}
		group.Description = value.(string)
	default:
		return types.ErrorResponse(-400, "invalid field")
	}

	if err := model.ModelUpdate(model.NewMongoFilter(
		model.IdFilter(id),
	), group); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(true)
}

func DeleteGroupService(user *types.User, id types.PrimaryId) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	if member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		if member.IsAdmin() {
			allowed = true
		}
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	group.Delete()
	if err := model.ModelUpdate(model.NewMongoFilter(
		model.IdFilter(id),
	), group); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(true)
}

func ListGroupMembersService(user *types.User, id types.PrimaryId, index int64, limit int64) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	if _, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		allowed = true
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	list, err := model.ModelGetAll[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
		),
		&model.MongoOptions{
			Skip:  &index,
			Limit: &limit,
		},
	)

	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(list)
}

func CreateUserInGroupByEmailAndPasswordService(user *types.User, id types.PrimaryId, email string, password string) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	if member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		if member.IsAdmin() {
			allowed = true
		}
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	password = auth.HashPassword(password)

	types.CreateUserWithEmailMutex()
	defer types.CreateUserWithEmailMutexEnd()
	// check email
	if _, err := model.ModelGet[types.User](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("email", email),
		),
	); err == nil {
		return types.ErrorResponse(-400, "email already exists")
	}

	new_user := &types.User{
		UserFlag:        types.USER_DEFAULT_FLAG,
		ModelPermission: types.USER_DEFAULT_MODEL_PERMISSION,
		Username:        "新用户_" + strings.RandomAlphaString(8),
		Avatar:          "",
		Parent:          user.Id,
	}

	// create user
	var uid types.PrimaryId
	if err = model.ModelInsert(new_user, &uid); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	// create password
	password_obj := &types.Password{
		Password: password,
		Uid:      uid,
		CreateAt: time.Now().Unix(),
	}
	if err = model.ModelInsert(password_obj, nil); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	// create group member
	group_member := &types.GroupMember{
		Gid:      id,
		Uid:      uid,
		CreateAt: time.Now().Unix(),
	}
	group_member.SetUser()
	if err = model.ModelInsert(group_member, nil); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(true)
}

func CreateUserInGroupByPhoneAndPasswordService(user *types.User, id types.PrimaryId, phone string, password string) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	if member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", id),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		if member.IsAdmin() {
			allowed = true
		}
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	password = auth.HashPassword(password)

	types.CreateUserWithPhoneMutex()
	defer types.CreateUserWithPhoneMutexEnd()
	// check phone
	if _, err := model.ModelGet[types.User](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("phone", phone),
		),
	); err == nil {
		return types.ErrorResponse(-400, "phone already exists")
	}

	new_user := &types.User{
		UserFlag:        types.USER_DEFAULT_FLAG,
		ModelPermission: types.USER_DEFAULT_MODEL_PERMISSION,
		Username:        "新用户_" + strings.RandomAlphaString(8),
		Avatar:          "",
		Parent:          user.Id,
	}

	// create user
	var uid types.PrimaryId
	if err = model.ModelInsert(new_user, &uid); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	// create password
	password_obj := &types.Password{
		Password: password,
		Uid:      uid,
		CreateAt: time.Now().Unix(),
	}

	if err = model.ModelInsert(password_obj, nil); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	// create group member
	group_member := &types.GroupMember{
		Gid:      id,
		Uid:      uid,
		CreateAt: time.Now().Unix(),
	}
	group_member.SetUser()
	if err = model.ModelInsert(group_member, nil); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(true)
}

func UpdateGroupMemberRoleService(user *types.User, gid types.PrimaryId, uid types.PrimaryId, role string) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(gid)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	if member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", gid),
			model.MongoKeyFilter("uid", user.Id),
		),
	); err == nil {
		if member.IsAdmin() {
			allowed = true
		}
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	// update group member
	if member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", gid),
			model.MongoKeyFilter("uid", uid),
		),
	); err == nil {
		switch role {
		case "admin":
			member.ClearRole()
			member.SetAdmin()
		case "user":
			member.ClearRole()
			member.SetUser()
		case "leader":
			member.ClearRole()
			member.SetLeader()
		default:
			return types.ErrorResponse(-400, "invalid role")
		}

		if err = model.ModelUpdate(
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(member.Id),
			), member,
		); err != nil {
			return types.ErrorResponse(-500, err.Error())
		}
	} else {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(true)
}

func DeleteGroupMemberService(user *types.User, gid types.PrimaryId, uid types.PrimaryId) *types.MasterResponse {
	// permission check
	allowed := false
	if user.IsAdmin() {
		allowed = true
	}

	var group *types.Group
	get_group := func(id types.PrimaryId) (*types.Group, error) {
		if group != nil {
			return group, nil
		}
		var err error
		group, err = model.ModelGet[types.Group](
			model.NewMongoFilter(
				model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
				model.IdFilter(id),
			),
		)
		if err != nil {
			return nil, err
		}

		return group, nil
	}

	group, err := get_group(gid)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	if group.Parent.String() == user.Id.String() {
		allowed = true
	}

	member, err := model.ModelGet[types.GroupMember](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.MongoKeyFilter("gid", gid),
			model.MongoKeyFilter("uid", user.Id),
		),
	)

	if err == nil {
		if member.IsAdmin() {
			allowed = true
		}
	}

	if !allowed {
		return types.ErrorResponse(-403, "permission denied")
	}

	// delete group member
	member.Delete()
	if err = model.ModelUpdate(
		model.NewMongoFilter(
			model.MongoNoFlagFilter(types.BASIC_TYPE_FLAG_DELETED),
			model.IdFilter(member.Id),
		), member,
	); err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(true)
}

func CreateUserInGroupByExcelService(user *types.User, gid types.PrimaryId, file types.PrimaryId) *types.MasterResponse {
	panic("implement me")
}
