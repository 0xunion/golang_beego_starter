package types

type Attack struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	Defender PrimaryId `json:"defender" bson:"defender"`
	Reason   string    `json:"reason" bson:"reason"`
	GameId   PrimaryId `json:"game_id" bson:"game_id"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
	Commnt   string    `json:"comment" bson:"comment"`
	State    int       `json:"state" bson:"state"`
}

type AttackSupportType struct {
	Name  string `json:"name"`
	CN    string `json:"cn"`
	Value int    `json:"value"`
}

func GetAttackStates() []AttackSupportType {
	return []AttackSupportType{
		{Name: "Unverified", CN: "未审核", Value: ATTACK_STATE_UNVERIFIED},
		{Name: "Accepted", CN: "已通过", Value: ATTACK_STATE_ACCEPTED},
		{Name: "Rejected", CN: "已拒绝", Value: ATTACK_STATE_REJECTED},
	}
}

const (
	ATTACK_STATE_UNVERIFIED = iota
	ATTACK_STATE_ACCEPTED
	ATTACK_STATE_REJECTED
)
