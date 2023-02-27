package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/report_section Service 获取报告中各种类型的允许范围，如漏洞类型，攻击手法类型等
func ApiCustomAttackerReportSectionService(
	user *master_types.User,
	GameId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAttackerReportSectionResponse struct {
		Success  bool `json:"success"`
		Sections any  `json:"sections"`
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
	var sections = map[string]any{
		"vuln_type":        master_types.GetReportVulnTypes(),
		"attack_type":      master_types.GetReportAttackTypes(),
		"level":            master_types.GetReportLevels(),
		"status":           master_types.GetReportStates(),
		"isolation_break":  master_types.GetReportIsolationBreaks(),
		"achievement_type": master_types.GetReportAchievementTypes(),
	}

	apiCustomAttackerReportSectionResponse.Sections = sections

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAttackerReportSectionResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
