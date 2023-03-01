package custom

import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/report/commit Service 提交报告
func ApiCustomAttackerReportCommitService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    DefenderId master_types.PrimaryId,
    Content string,
    IsolationBreak int,
    VulnType int,
    AchievementType int,
    AttackType int,
    Uri string,
    VulnLevel int,
    Name string,
) (*master_types.MasterResponse) {
    var apiCustomAttackerReportCommitResponse struct {
        Success bool `json:"success"`
        Report any `json:"report"`
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

    
    // get Gamer

    var D_attacker *master_types.Gamer

    {
value, err := model.ModelGet[master_types.Gamer](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("owner", user.Id),
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_ATTACKER),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }

        D_attacker = value

    }

    // get RedTeam

    var D_attacker_team *master_types.RedTeam

    {
value, err := model.ModelGet[master_types.RedTeam](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("gid", D_attacker.GroupId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }

        D_attacker_team = value

    }

    // get Defender


    {
_, err := model.ModelGet[master_types.Defender](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", DefenderId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }


    }

    // create Report
    var report *master_types.Report = &master_types.Report{
        Owner: user.Id,
        Content: Content,
        Name: Name,
        Uri: Uri,
        Level: VulnLevel,
        DefenderId: DefenderId,
        AttackTeamId: D_attacker_team.Id,
        IsolationBreak: IsolationBreak,
        VulnType: VulnType,
        AttackType: AttackType,
        CreateAt: time.Now().Unix(),
        GameId: GameId,
    }

    err := model.ModelInsert(report, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAttackerReportCommitResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
