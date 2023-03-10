package types

type Boardcast struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Content  string    `json:"content" bson:"content"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	GameId   PrimaryId `json:"game_id" bson:"game_id"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}
