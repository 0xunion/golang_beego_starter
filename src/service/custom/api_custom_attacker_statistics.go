package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/statistics Service 红队获取统计信息
func ApiCustomAttackerStatisticsService(
    user *master_types.User,
    GameId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomAttackerStatisticsResponse struct {
        Success bool `json:"success"`
        Statistics any `json:"statistics"`
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
	type statistics struct {
		Reports struct {
			Total      int64 `json:"total" bson:"total"`
			Accepted   int64 `json:"accepted" bson:"accepted"`
			Rejected   int64 `json:"rejected" bson:"rejected"`
			Unverified int64 `json:"unverified" bson:"unverified"`
		} `json:"reports" bson:"reports"`
		AttackTypes struct {
			Gimmack []struct {
				Name  string `json:"name" bson:"name"`
				Count int64  `json:"count" bson:"count"`
				Value int64  `json:"value" bson:"value"`
			} `json:"gimmack" bson:"gimmack"`
			AttackerRank []struct {
				Name   string `json:"name" bson:"name"`
				Count  int64  `json:"count" bson:"count"`
				Source int64  `json:"value" bson:"value"`
			} `json:"attacker_rank" bson:"attacker_rank"`
			DefenderRank []struct {
				Name   string `json:"name" bson:"name"`
				Count  int64  `json:"count" bson:"count"`
				Source int64  `json:"value" bson:"value"`
			} `json:"defender_rank" bson:"defender_rank"`
			VulnType []struct {
				Name  string `json:"name" bson:"name"`
				Count int64  `json:"count" bson:"count"`
				Type  int64  `json:"value" bson:"value"`
			} `json:"vuln_type" bson:"vuln_type"`
			VulnLevel []struct {
				Name  string `json:"name" bson:"name"`
				Count int64  `json:"count" bson:"count"`
				Level int64  `json:"value" bson:"value"`
			} `json:"vuln_level" bson:"vuln_level"`
		}
	}

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAttackerStatisticsResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
