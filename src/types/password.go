package types

type Password struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Password string    `json:"password" bson:"password"`
	Uid      PrimaryId `json:"uid" bson:"uid"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}
