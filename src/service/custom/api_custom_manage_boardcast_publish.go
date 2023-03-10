package custom

import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/boardcast/publish Service 发布公告
func ApiCustomManageBoardcastPublishService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    Content string,
) (*master_types.MasterResponse) {
    var apiCustomManageBoardcastPublishResponse struct {
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

    
    // create Boardcast
    var boardcast *master_types.Boardcast = &master_types.Boardcast{
        GameId: GameId,
        Content: Content,
        CreateAt: time.Now().Unix(),
    }

    err := model.ModelInsert(boardcast, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageBoardcastPublishResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
