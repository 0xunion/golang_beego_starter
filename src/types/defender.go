package types

type Defender struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	Name     string    `json:"name" bson:"name"`
	Score    int64     `json:"score" bson:"score"`
	GameId   PrimaryId `json:"game_id" bson:"game_id"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}
