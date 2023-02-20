package types

type Report struct {
	BasicType
	Id              PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner           PrimaryId `json:"owner" bson:"owner"`
	Content         string    `json:"content" bson:"content"`
	Name            string    `json:"name" bson:"name"`
	Uri             string    `json:"uri" bson:"uri"`
	Level           int       `json:"level" bson:"level"`
	DefenderId      PrimaryId `json:"defender_id" bson:"defender_id"`
	AttackTeamId    PrimaryId `json:"attack_team_id" bson:"attack_team_id"`
	State           int       `json:"state" bson:"state"`
	IsolationBreak  int       `json:"isolation_break" bson:"isolation_break"`
	AchievementType int       `json:"achievement_type" bson:"achievement_type"`
	Score           int       `json:"score" bson:"score"`
	CreateAt        int64     `json:"create_at" bson:"create_at"`
	GameId          PrimaryId `json:"game_id" bson:"game_id"`
	ReportType      int       `json:"report_type" bson:"report_type"` // 0: attack report, 1: defense report
}

type ReportSupportType struct {
	Name  string `json:"name"`
	CN    string `json:"cn"`
	Value int    `json:"value"`
}

const (
	REPORT_STATE_UNVERIFIED = iota
	REPORT_STATE_ACCEPTED
	REPORT_STATE_REJECTED
)

func GetReportStates() []ReportSupportType {
	return []ReportSupportType{
		{Name: "Unverified", CN: "未审核", Value: REPORT_STATE_UNVERIFIED},
		{Name: "Accepted", CN: "已通过", Value: REPORT_STATE_ACCEPTED},
		{Name: "Rejected", CN: "已拒绝", Value: REPORT_STATE_REJECTED},
	}
}

const (
	REPORT_TYPE_ATTACK = iota
	REPORT_TYPE_DEFENSE
)

func GetReportTypes() []ReportSupportType {
	return []ReportSupportType{
		{Name: "Attack", CN: "攻击报告", Value: REPORT_TYPE_ATTACK},
		{Name: "Defense", CN: "防御报告", Value: REPORT_TYPE_DEFENSE},
	}
}

const (
	REPORT_LEVEL_LOW = iota
	REPORT_LEVEL_MEDIUM
	REPORT_LEVEL_HIGH
	REPORT_LEVEL_CRITICAL
)

func GetReportLevels() []ReportSupportType {
	return []ReportSupportType{
		{Name: "Low", CN: "低", Value: REPORT_LEVEL_LOW},
		{Name: "Medium", CN: "中", Value: REPORT_LEVEL_MEDIUM},
		{Name: "High", CN: "高", Value: REPORT_LEVEL_HIGH},
		{Name: "Critical", CN: "严重", Value: REPORT_LEVEL_CRITICAL},
	}
}

const (
	REPORT_ACHIEVEMENT_TYPE_NONE = iota
)

func GetReportAchievementTypes() []ReportSupportType {
	return []ReportSupportType{
		{Name: "None", CN: "无", Value: REPORT_ACHIEVEMENT_TYPE_NONE},
	}
}

const (
	REPORT_ISOLATION_BREAK_NONE = iota
	REPORT_ISOLATION_BREAK_INTERANET
)

func GetReportIsolationBreaks() []ReportSupportType {
	return []ReportSupportType{
		{Name: "None", CN: "无", Value: REPORT_ISOLATION_BREAK_NONE},
		{Name: "Interanet", CN: "内网", Value: REPORT_ISOLATION_BREAK_INTERANET},
	}
}

type ReportComment struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	GameId   PrimaryId `json:"game_id" bson:"game_id"`
	ReportId PrimaryId `json:"report_id" bson:"report_id"`
	Content  string    `json:"content" bson:"content"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}
