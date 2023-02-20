package types

type Defender struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	Industry string    `json:"industry" bson:"industry"`
	Name     string    `json:"name" bson:"name"`
	Score    int64     `json:"score" bson:"score"`
	GameId   PrimaryId `json:"game_id" bson:"game_id"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

type Asset struct {
	BasicType
	Id         PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner      PrimaryId `json:"owner" bson:"owner"`
	Industry   string    `json:"industry" bson:"industry"`
	Uri        string    `json:"uri" bson:"uri"`
	Name       string    `json:"name" bson:"name"`
	DefenderId PrimaryId `json:"defender_id" bson:"defender_id"`
	GameId     PrimaryId `json:"game_id" bson:"game_id"`
	CreateAt   int64     `json:"create_at" bson:"create_at"`
}
