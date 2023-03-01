package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/update Service 更新一个比赛
func ApiCustomAdminGameUpdateService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    Name string,
    Description string,
    HeaderHtml string,
    StartTime int64,
    EndTime int64,
) (*master_types.MasterResponse) {
    var apiCustomAdminGameUpdateResponse struct {
        Success bool `json:"success"`
    }

    access_controll := false
    if !access_controll && user.IsAdmin() {
        access_controll = true
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }

    
    // update Game
    
    {
        err := model.ModelUpdateField[master_types.Game](
            model.NewMongoFilter(
                model.MongoKeyFilter("_id", GameId),
            ),



            model.MongoSetField(
                "name", 
                Name,
            ),
            model.MongoSetField(
                "description", 
                Description,
            ),
            model.MongoSetField(
                "header_html", 
                HeaderHtml,
            ),
            model.MongoSetField(
                "start_time", 
                StartTime,
            ),
            model.MongoSetField(
                "end_time", 
                EndTime,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAdminGameUpdateResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
