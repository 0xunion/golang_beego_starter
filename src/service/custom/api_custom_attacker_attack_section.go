package custom

/* @MT-TPL-IMPORT-START */
import (
    master_types "github.com/0xunion/exercise_back/src/types"
    model "github.com/0xunion/exercise_back/src/model"
    permission_type "github.com/0xunion/exercise_back/src/types/permission"
    "go.mongodb.org/mongo-driver/bson/primitive"


)
/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/attack/section Service 获取攻击申请中的状态对照表，如0-未审核，1-已通过等
func ApiCustomAttackerAttackSectionService(
    user *master_types.User,
    GameId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomAttackerAttackSectionResponse struct {
        Success bool `json:"success"`
        Sections any `json:"sections"`
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
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_ATTACKER),
            ),
        )
        if err == nil && model_instance != nil {
            access_controll = true
        }
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }
/* @MT-TPL-SERVICE-END */

    // TODO: add service code here, do what you want to do
   
    /* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAttackerAttackSectionResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */