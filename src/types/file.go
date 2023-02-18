package types

type File struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	Name     string    `json:"name" bson:"name"`
	Hash     string    `json:"hash" bson:"hash"` // random hash for file as key
	Size     int64     `json:"size" bson:"size"`
	Path     string    `json:"path" bson:"path"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}
