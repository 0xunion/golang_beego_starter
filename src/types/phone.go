package types

type Phone struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Phone    string    `json:"phone" bson:"phone"`
	Uid      PrimaryId `json:"uid" bson:"uid"`
	CreateAt int64     `json:"create" bson:"create_at"`
}
