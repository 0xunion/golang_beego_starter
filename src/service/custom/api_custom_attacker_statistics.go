package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	"github.com/0xunion/exercise_back/src/types"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/strings"
	"go.mongodb.org/mongo-driver/bson"
	/* @MT-TPL-IMPORT-TIME-START */ /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/statistics Service 红队获取统计信息
func ApiCustomAttackerStatisticsService(
	user *master_types.User,
	GameId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAttackerStatisticsResponse struct {
		Success    bool `json:"success"`
		Statistics any  `json:"statistics"`
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
	var statistics struct {
		Reports struct {
			Total          int64 `json:"total" bson:"total"`
			Accepted       int64 `json:"accepted" bson:"accepted"`
			Rejected       int64 `json:"rejected" bson:"rejected"`
			Unverified     int64 `json:"unverified" bson:"unverified"`
			IsolationBreak int64 `json:"isolation_break" bson:"isolation_break"`
		} `json:"reports" bson:"reports"`
		AttackTypes struct {
			Gimmack []struct {
				Value int64 `json:"value" bson:"value"`
				Count int64 `json:"count" bson:"count"`
			} `json:"gimmack" bson:"gimmack"`
			VulnType []struct {
				Value int64 `json:"value" bson:"value"`
				Count int64 `json:"count" bson:"count"`
			} `json:"vuln_type" bson:"vuln_type"`
			VulnLevel []struct {
				Value int64 `json:"value" bson:"value"`
				Count int64 `json:"count" bson:"count"`
			} `json:"vuln_level" bson:"vuln_level"`
			Industry []struct {
				Value string `json:"value" bson:"value"`
				Count int64  `json:"count" bson:"count"`
			} `json:"industry" bson:"industry"`
		} `json:"attack_types" bson:"attack_types"`
		Rank struct {
			RedTeam  []types.RedTeam  `json:"red_team" bson:"red_team"`
			BlueTeam []types.Defender `json:"blue_team" bson:"blue_team"`
		} `json:"rank" bson:"rank"`
	}

	// get all reports
	reports, err := model.ModelGetAll[master_types.Report](
		model.NewMongoFilter(
			model.MongoKeyFilter("game_id", GameId),
			model.MongoKeyFilter("report_type", master_types.REPORT_TYPE_ATTACK),
		),
		&model.MongoOptions{
			Projection: bson.M{"content": 0},
		},
	)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	// init statistics, get all supported attack types
	attack_types := master_types.GetReportAttackTypes() // gimmack
	vuln_types := master_types.GetReportVulnTypes()     // vuln_type
	vuln_levels := master_types.GetReportLevels()       // vuln_level

	// init statistics
	statistics.Reports.Total = int64(len(reports))
	statistics.Reports.Accepted = 0
	statistics.Reports.Rejected = 0
	statistics.Reports.Unverified = 0
	statistics.Reports.IsolationBreak = 0

	for attack_types_index := range attack_types {
		statistics.AttackTypes.Gimmack = append(statistics.AttackTypes.Gimmack, struct {
			Value int64 `json:"value" bson:"value"`
			Count int64 `json:"count" bson:"count"`
		}{
			Value: int64(attack_types[attack_types_index].Value),
			Count: 0,
		})
	}

	for vuln_types_index := range vuln_types {
		statistics.AttackTypes.VulnType = append(statistics.AttackTypes.VulnType, struct {
			Value int64 `json:"value" bson:"value"`
			Count int64 `json:"count" bson:"count"`
		}{
			Value: int64(vuln_types[vuln_types_index].Value),
			Count: 0,
		})
	}

	for vuln_levels_index := range vuln_levels {
		statistics.AttackTypes.VulnLevel = append(statistics.AttackTypes.VulnLevel, struct {
			Value int64 `json:"value" bson:"value"`
			Count int64 `json:"count" bson:"count"`
		}{
			Value: int64(vuln_levels[vuln_levels_index].Value),
			Count: 0,
		})
	}

	var defender_ids []master_types.PrimaryId

	// get statistics
	for reports_index := range reports {
		report := reports[reports_index]

		// get defender ids
		defender_ids = append(defender_ids, report.DefenderId)

		// get attack type
		for attack_types_index := range attack_types {
			if attack_types[attack_types_index].Value == report.AttackType {
				statistics.AttackTypes.Gimmack[attack_types_index].Count++
				break
			}
		}

		// get vuln type
		for vuln_types_index := range vuln_types {
			if vuln_types[vuln_types_index].Value == report.VulnType {
				statistics.AttackTypes.VulnType[vuln_types_index].Count++
				break
			}
		}

		// get vuln level
		for vuln_levels_index := range vuln_levels {
			if vuln_levels[vuln_levels_index].Value == report.Level {
				statistics.AttackTypes.VulnLevel[vuln_levels_index].Count++
				break
			}
		}

		// get report state
		switch report.State {
		case master_types.REPORT_STATE_ACCEPTED:
			statistics.Reports.Accepted++
		case master_types.REPORT_STATE_REJECTED:
			statistics.Reports.Rejected++
		case master_types.REPORT_STATE_UNVERIFIED:
			statistics.Reports.Unverified++
		}

		// get isolation break
		if report.IsolationBreak == 1 && report.State == master_types.REPORT_STATE_ACCEPTED {
			statistics.Reports.IsolationBreak++
		}
	}

	// get all defenders
	defenders, err := model.ModelGetAll[master_types.Defender](
		model.NewMongoFilter(
			model.MongoKeyFilter("game_id", GameId),
			model.MongoArrayContainsFilter("_id", strings.DepduplicationStringArray(defender_ids)),
		),
		&model.MongoOptions{
			Sort: model.MongoSort("score", -1),
		},
	)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	statistics.Rank.BlueTeam = defenders

	var industry_count_map = make(map[string]int64)
	var defender_id_industry_map = make(map[master_types.PrimaryId]string)

	// init defender industry map
	for defenders_index := range defenders {
		defender := defenders[defenders_index]
		defender_id_industry_map[defender.Id] = defender.Industry
	}

	// get industry count
	for reports_index := range reports {
		report := reports[reports_index]

		// get industry
		industry := defender_id_industry_map[report.DefenderId]

		// get industry count
		industry_count_map[industry]++
	}

	// set industry count
	for industry, count := range industry_count_map {
		statistics.AttackTypes.Industry = append(statistics.AttackTypes.Industry, struct {
			Value string `json:"value" bson:"value"`
			Count int64  `json:"count" bson:"count"`
		}{
			Value: industry,
			Count: count,
		})
	}

	// get red team
	red_team, err := model.ModelGetAll[master_types.RedTeam](
		model.NewMongoFilter(
			model.MongoKeyFilter("game_id", GameId),
		),
		&model.MongoOptions{
			Sort: model.MongoSort("score", -1),
		},
	)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	statistics.Rank.RedTeam = red_team

	apiCustomAttackerStatisticsResponse.Statistics = statistics

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAttackerStatisticsResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
