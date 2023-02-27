package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/attack/reject Service 驳回红队的攻击
func ApiCustomManageAttackRejectService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    AttackId master_types.PrimaryId,
    Comment string,
) (*master_types.MasterResponse) {
    var apiCustomManageAttackRejectResponse struct {
        Success bool `json:"success"`
    }

    access_controll := false
    if !access_controll && user.IsAdmin() {
        access_controll = true
    }
    if !access_controll {
        model_instance, err := model.ModelGet[master_types.Gamer](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("owner", user.Id),
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_JUDGEMENT),
            ),
        )
        if err == nil && model_instance != nil {
            access_controll = true
        }
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }

    
    // get Attack


    {
_, err := model.ModelGet[master_types.Attack](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", AttackId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }


    }

    // update Attack
    
    {
        err := model.ModelUpdateField[master_types.Attack](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", AttackId),
                model.MongoKeyFilter("state", master_types.ATTACK_STATE_UNVERIFIED),
            ),



            model.MongoSetField(
                "state", 
                master_types.ATTACK_STATE_REJECTED,
            ),
            model.MongoSetField(
                "comment", 
                Comment,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }

        
    // set response directly
    apiCustomManageAttackRejectResponse.Success = true
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageAttackRejectResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
