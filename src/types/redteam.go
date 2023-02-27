package types

type RedTeam struct {
	BasicType
	Id     PrimaryId `json:"id" bson:"_id,omitempty"`
	Name   string    `json:"name" bson:"name"`
	Score  int       `json:"score" bson:"score"`
	Gid    PrimaryId `json:"gid" bson:"gid"`
	GameId PrimaryId `json:"game_id" bson:"game_id"`
}
