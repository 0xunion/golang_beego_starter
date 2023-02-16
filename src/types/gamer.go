package types

type Gamer struct {
	BasicType
	Id         PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner      PrimaryId `json:"owner" bson:"owner"`
	Name       string    `json:"name" bson:"name"`
	Phone      string    `json:"phone" bson:"phone"`
	CreateAt   int64     `json:"create_at" bson:"create_at"`
	Identity   int64     `json:"identity" bson:"identity"` // attacher, defender, judgement, customer
	GameId     PrimaryId `json:"game_id" bson:"game_id"`
	Permission int64     `json:"permission" bson:"permission"`
	GroupId    PrimaryId `json:"group_id" bson:"group_id"`
}
