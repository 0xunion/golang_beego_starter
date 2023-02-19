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
	Score      int64     `json:"score" bson:"score"`
	Permission int64     `json:"permission" bson:"permission"`
	GroupId    PrimaryId `json:"group_id" bson:"group_id"`
}

const (
	GAMER_IDENTITY_ATTACKER = iota
	GAMER_IDENTITY_DEFENDER
	GAMER_IDENTITY_JUDGEMENT
	GAMER_IDENTITY_CUSTOMER
)

func (g *Gamer) IsAttacker() bool {
	return g.Identity == GAMER_IDENTITY_ATTACKER
}

func (g *Gamer) IsDefender() bool {
	return g.Identity == GAMER_IDENTITY_DEFENDER
}

func (g *Gamer) IsJudgement() bool {
	return g.Identity == GAMER_IDENTITY_JUDGEMENT
}

func (g *Gamer) IsCustomer() bool {
	return g.Identity == GAMER_IDENTITY_CUSTOMER
}

func (g *Gamer) SetAttacker() {
	g.Identity = GAMER_IDENTITY_ATTACKER
}

func (g *Gamer) SetDefender() {
	g.Identity = GAMER_IDENTITY_DEFENDER
}

func (g *Gamer) SetJudgement() {
	g.Identity = GAMER_IDENTITY_JUDGEMENT
}

func (g *Gamer) SetCustomer() {
	g.Identity = GAMER_IDENTITY_CUSTOMER
}