package types

type Email struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Email    string    `json:"email" bson:"email"`
	Uid      PrimaryId `json:"uid" bson:"uid"`
	CreateAt int64     `json:"create" bson:"create_at"`
}
