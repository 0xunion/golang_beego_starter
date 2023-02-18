package types

type ReadTeam struct {
	BasicType
	Id     PrimaryId `json:"id" bson:"_id,omitempty"`
	Name   string    `json:"name"`
	Score  int       `json:"score"`
	Gid    PrimaryId `json:"gid"`
	GameId PrimaryId `json:"game_id" bson:"game_id"`
}
