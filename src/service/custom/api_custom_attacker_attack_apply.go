package custom

/* @MT-TPL-IMPORT-START */
import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/attack/apply Service 申请高危攻击
func ApiCustomAttackerAttackApplyService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    DefenderId master_types.PrimaryId,
    Reason string,
) (*master_types.MasterResponse) {
    var apiCustomAttackerAttackApplyResponse struct {
        Success bool `json:"success"`
        AttackId any `json:"attack_id"`
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

    
    // create Attack
    var attack *master_types.Attack = &master_types.Attack{
        GameId: GameId,
        Owner: user.Id,
        Reason: Reason,
        Defender: DefenderId,
        CreateAt: time.Now().Unix(),
        State: master_types.ATTACK_STATE_UNVERIFIED,
    }

    err := model.ModelInsert(attack, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAttackerAttackApplyResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
