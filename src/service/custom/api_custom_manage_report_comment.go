package custom

/* @MT-TPL-IMPORT-START */
import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/report/comment Service 评论红队提交的报告
func ApiCustomManageReportCommentService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    ReportId master_types.PrimaryId,
    Comment string,
) (*master_types.MasterResponse) {
    var apiCustomManageReportCommentResponse struct {
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

    
    // get Report


    {
_, err := model.ModelGet[master_types.Report](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", ReportId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }


    }

    // create ReportComment
    var report_comment *master_types.ReportComment = &master_types.ReportComment{
        GameId: GameId,
        ReportId: ReportId,
        Content: Comment,
        CreateAt: time.Now().Unix(),
        Owner: user.Id,
    }

    err := model.ModelInsert(report_comment, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
        
    // set response directly
    apiCustomManageReportCommentResponse.Success = true
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageReportCommentResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
