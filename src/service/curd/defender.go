package curd

/* @MT-TPL-IMPORT-START */
import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-CREATE-START */
func CreateDefenderService(
	user *master_types.User,
	name string,
	game_id master_types.PrimaryId,
) *master_types.MasterResponse {
	// check if the user has permission of creating
	access_controll := false
	if !access_controll {
		if user.IsAdmin() {
			access_controll = true
		}
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Access Denied")
	}

	now := time.Now()

	// create model
	defender := &master_types.Defender{
		Name:     name,
		CreateAt: now.Unix(),
		Owner:    user.Id,
		GameId:   game_id,
		Score:    10000,
	}

	// save model
	err := model.ModelInsert(defender, nil)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// return response
	return master_types.SuccessResponse(defender)
}

/* @MT-TPL-CREATE-END */

/* @MT-TPL-UPDATE-START */
func UpdateDefenderService(
	user *master_types.User,
	id master_types.PrimaryId,
	name string,
	game_id master_types.PrimaryId,
	score int64,
) *master_types.MasterResponse {
	// check if the user has permission of updating
	var object *master_types.Defender
	get_object := func() (*master_types.Defender, error) {
		if object != nil {
			return object, nil
		}
		var err error
		object, err = model.ModelGet[master_types.Defender](
			model.NewMongoFilter(
				model.IdFilter(
					id,
				),
				model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
			),
		)
		if err != nil {
			return nil, err
		}
		return object, nil
	}

	access_controll := false
	if !access_controll {
		if user.IsAdmin() {
			access_controll = true
		}
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Access Denied")
	}

	// get object and update
	object, err := get_object()
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	object.Name = name
	object.GameId = game_id
	object.Score = score

	// save model
	err = model.ModelUpdate(
		model.NewMongoFilter(
			model.IdFilter(
				id,
			),
			model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
		),
		object,
	)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// return response
	return master_types.SuccessResponse(object)
}

/* @MT-TPL-UPDATE-END */

/* @MT-TPL-DELETE-START */
func DeleteDefenderService(
	user *master_types.User,
	id master_types.PrimaryId,
) *master_types.MasterResponse {
	// check if the user has permission of deleting
	var object *master_types.Defender
	get_object := func() (*master_types.Defender, error) {
		if object != nil {
			return object, nil
		}
		var err error
		object, err = model.ModelGet[master_types.Defender](
			model.NewMongoFilter(
				model.IdFilter(
					id,
				),
			),
		)
		if err != nil {
			return nil, err
		}
		return object, nil
	}

	access_controll := false
	if !access_controll {
		if user.IsAdmin() {
			access_controll = true
		}
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Access Denied")
	}

	// get object and delete
	object, err := get_object()
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// delete model
	object.Delete()
	// save model
	err = model.ModelUpdate(
		model.NewMongoFilter(
			model.IdFilter(
				id,
			),
			model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
		),
		object,
	)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// return response
	return master_types.SuccessResponse(nil)
}

/* @MT-TPL-DELETE-END */

/* @MT-TPL-GET-START */
func GetDefenderService(
	user *master_types.User,
	id master_types.PrimaryId,
) *master_types.MasterResponse {
	// check if the user has permission of getting
	var object *master_types.Defender
	get_object := func() (*master_types.Defender, error) {
		if object != nil {
			return object, nil
		}
		var err error
		object, err = model.ModelGet[master_types.Defender](
			model.NewMongoFilter(
				model.IdFilter(
					id,
				),
				model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
			),
		)
		if err != nil {
			return nil, err
		}
		return object, nil
	}

	access_controll := false
	if !access_controll {
		access_controll = true
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Access Denied")
	}

	// get object
	object, err := get_object()
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// return response
	return master_types.SuccessResponse(object)
}

/* @MT-TPL-GET-END */

/* @MT-TPL-LIST-START */
func ListDefenderService(
	user *master_types.User,
	index, limit int64,
) *master_types.MasterResponse {
	// check if the user has permission of listing
	access_controll := false
	if !access_controll {
		access_controll = true
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Access Denied")
	}

	// get object
	objects, err := model.ModelGetAll[master_types.Defender](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
		),
		&model.MongoOptions{
			Skip:  &index,
			Limit: &limit,
		},
	)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// return response
	return master_types.SuccessResponse(objects)
}

/* @MT-TPL-LIST-END */

/* @MT-TPL-SEARCH-START */
func SearchDefenderService(
	user *master_types.User,
	name string,
	index, limit int64,
) *master_types.MasterResponse {
	// check if the user has permission of listing
	access_controll := false
	if !access_controll {
		access_controll = true
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Access Denied")
	}

	// get object
	objects, err := model.ModelGetAll[master_types.Defender](
		model.NewMongoFilter(
			model.MongoKeyFilter(
				"name",
				name,
			),
			model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
		),
		&model.MongoOptions{
			Skip:  &index,
			Limit: &limit,
		},
	)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// return response
	return master_types.SuccessResponse(objects)
}

/* @MT-TPL-SEARCH-END */
