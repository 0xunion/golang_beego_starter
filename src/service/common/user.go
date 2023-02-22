package common

import (
	"time"

	"github.com/0xunion/exercise_back/src/model"
	"github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/auth"
	"github.com/0xunion/exercise_back/src/util/strings"
)

/*
	CustomErrCode:
		-1001: create user failed
		-1002: create user password failed
		-1003: create user email failed
		-1004: create user phone failed

		-2001: no email found
		-2002: no phone found
		-2003: no password found
		-2004: password is not correct
		-2005: no user found

		-3001: captcha is invalid
*/

// apply create user service
func UserApplyByPasswordService(user *types.User, password string) *types.MasterResponse {
	if !user.IsAdmin() {
		return types.ErrorResponse(-403, "permission denied")
	}

	new_user := &types.User{
		UserFlag:        types.USER_DEFAULT_FLAG,
		ModelPermission: types.USER_DEFAULT_MODEL_PERMISSION,
		Username:        "新用户_" + strings.RandomAlphaString(8),
		Avatar:          "",
		Parent:          user.Id,
	}

	// create user
	if err := model.ModelInsert(new_user, nil); err != nil {
		return types.ErrorResponse(-1001, "create user failed")
	}

	// create user password
	if err := model.ModelInsert(&types.Password{
		Uid:      new_user.Id,
		Password: auth.HashPassword(password),
		CreateAt: time.Now().Unix(),
	}, nil); err != nil {
		return types.ErrorResponse(-1002, "create user failed")
	}

	return types.SuccessResponse(new_user)
}

// apply create user service
func UserApplyByEmailService(user *types.User, email string) *types.MasterResponse {
	if !user.IsAdmin() {
		return types.ErrorResponse(-403, "permission denied")
	}

	types.CreateUserWithEmailMutex()
	defer types.CreateUserWithEmailMutexEnd()
	if _, err := model.ModelGet[types.Email](model.NewMongoFilter(
		model.MongoKeyFilter("email", email),
	)); err == nil {
		return types.ErrorResponse(-1003, "email is already used")
	}

	new_user := &types.User{
		UserFlag:        types.USER_DEFAULT_FLAG,
		ModelPermission: types.USER_DEFAULT_MODEL_PERMISSION,
		Username:        "新用户_" + strings.RandomAlphaString(8),
		Avatar:          "",
		Parent:          user.Id,
	}

	// create user
	if err := model.ModelInsert(new_user, nil); err != nil {
		return types.ErrorResponse(-1001, "create user failed")
	}

	// create user password
	if err := model.ModelInsert(&types.Email{
		Uid:      new_user.Id,
		Email:    email,
		CreateAt: time.Now().Unix(),
	}, nil); err != nil {
		return types.ErrorResponse(-1003, "create user failed")
	}

	return types.SuccessResponse(new_user)
}

// apply create user service
func UserApplyByPhoneService(user *types.User, phone string) *types.MasterResponse {
	if !user.IsAdmin() {
		return types.ErrorResponse(-403, "permission denied")
	}

	new_user := &types.User{
		UserFlag:        types.USER_DEFAULT_FLAG,
		ModelPermission: types.USER_DEFAULT_MODEL_PERMISSION,
		Username:        "新用户_" + strings.RandomAlphaString(8),
		Avatar:          "",
		Parent:          user.Id,
	}

	types.CreateUserWithPhoneMutex()
	defer types.CreateUserWithPhoneMutexEnd()
	if _, err := model.ModelGet[types.Phone](model.NewMongoFilter(
		model.MongoKeyFilter("phone", phone),
	)); err == nil {
		return types.ErrorResponse(-1004, "phone is already used")
	}

	// create user
	if err := model.ModelInsert(new_user, nil); err != nil {
		return types.ErrorResponse(-1001, "create user failed")
	}

	// create user password
	if err := model.ModelInsert(&types.Phone{
		Uid:      new_user.Id,
		Phone:    phone,
		CreateAt: time.Now().Unix(),
	}, nil); err != nil {
		return types.ErrorResponse(-1004, "create user failed")
	}

	return types.SuccessResponse(new_user)
}

// login service
func UserLoginByEmailAndPasswordService(email, password, captcha_token, captcha string) *types.MasterResponse {
	// check captcha
	captcha_token_obj := auth.NewAuthTokenFromToken[Captcha](captcha_token)
	if !captcha_token_obj.Check(time.Now().Unix()) {
		return types.ErrorResponse(-3001, "captcha is invalid")
	}
	// get captcha
	captcha_obj := captcha_token_obj.Info()
	if !captcha_obj.Try(captcha) {
		return types.ErrorResponse(-3001, "captcha is wrong")
	}

	// hash password
	password = auth.HashPassword(password)

	// get user by email
	email_obj, err := model.ModelGet[types.Email](model.NewMongoFilter(
		model.MongoKeyFilter("email", email),
	))
	if err != nil {
		return types.ErrorResponse(-2001, "no email found")
	}

	// get user by password
	password_obj, err := model.ModelGet[types.Password](model.NewMongoFilter(
		model.MongoKeyFilter("uid", email_obj.Uid),
	))

	if err != nil {
		return types.ErrorResponse(-2003, "no email found")
	}

	if password_obj.Password != password {
		return types.ErrorResponse(-2004, "password is not correct")
	}

	// get user by id
	user_obj, err := model.ModelGet[types.User](model.NewMongoFilter(
		model.IdFilter(email_obj.Uid),
	))

	if err != nil {
		return types.ErrorResponse(-2005, "password is not correct")
	}

	token := auth.NewAuthTokenWithUid(user_obj.Id)
	login_token := token.GenerateToken(uint32(time.Now().Unix()))

	return types.SuccessResponse(map[string]string{
		"token": login_token,
	})
}

// login service by phone
func UserLoginByPhoneAndPasswordService(phone, password, captcha_token, captcha string) *types.MasterResponse {
	// check captcha
	captcha_token_obj := auth.NewAuthTokenFromToken[Captcha](captcha_token)
	if !captcha_token_obj.Check(time.Now().Unix()) {
		return types.ErrorResponse(-3001, "captcha is invalid")
	}
	// get captcha
	captcha_obj := captcha_token_obj.Info()
	if !captcha_obj.Try(captcha) {
		return types.ErrorResponse(-3002, "captcha is invalid")
	}

	// hash password
	password = auth.HashPassword(password)

	// get user by phone
	phone_obj, err := model.ModelGet[types.Phone](model.NewMongoFilter(
		model.MongoKeyFilter("phone", phone),
	))
	if err != nil {
		return types.ErrorResponse(-2002, "no phone found")
	}

	// get user by password
	password_obj, err := model.ModelGet[types.Password](model.NewMongoFilter(
		model.MongoKeyFilter("uid", phone_obj.Uid),
	))

	if err != nil {
		return types.ErrorResponse(-2003, "no password found")
	}

	if password_obj.Password != password {
		return types.ErrorResponse(-2004, "password is not correct")
	}

	// get user by id
	user_obj, err := model.ModelGet[types.User](model.NewMongoFilter(
		model.IdFilter(phone_obj.Uid),
	))

	if err != nil {
		return types.ErrorResponse(-2005, "password is not correct")
	}

	token := auth.NewAuthTokenWithUid(user_obj.Id)
	login_token := token.GenerateToken(uint32(time.Now().Unix()))

	return types.SuccessResponse(map[string]string{
		"token": login_token,
	})
}

func InitRootUserService(email string, password string) *types.MasterResponse {
	// check if root user exists
	if _, err := model.ModelGet[types.User](model.NewMongoFilter(
		model.MongoHasBitFilter("user_flag", types.USER_FLAG_ROOT),
	)); err == nil {
		return types.ErrorResponse(-1002, "root user already exists")
	}

	// create root user
	hash_password := auth.HashPassword(password)

	new_user := &types.User{
		UserFlag:        types.USER_FLAG_ROOT,
		ModelPermission: types.USER_DEFAULT_MODEL_PERMISSION,
		Username:        "root",
		Avatar:          "",
	}

	// create user
	var uid types.PrimaryId
	if err := model.ModelInsert(new_user, &uid); err != nil {
		return types.ErrorResponse(-1001, "create user failed")
	}

	new_user.Id = uid

	// create user password
	if err := model.ModelInsert(&types.Password{
		Uid:      uid,
		Password: hash_password,
		CreateAt: time.Now().Unix(),
	}, nil); err != nil {
		return types.ErrorResponse(-1004, "create user failed")
	}

	// create user email
	if err := model.ModelInsert(&types.Email{
		Uid:      uid,
		Email:    email,
		CreateAt: time.Now().Unix(),
	}, nil); err != nil {
		return types.ErrorResponse(-1004, "create user failed")
	}

	return types.SuccessResponse(new_user)
}
