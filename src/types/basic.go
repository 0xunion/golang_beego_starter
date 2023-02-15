package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type PrimaryId = primitive.ObjectID

type BasicType struct {
	Flag int `json:"flag" bson:"flag"`
}

const (
	BASIC_TYPE_FLAG_DELETED = 1 << 0
)

func (t *BasicType) Delete() {
	t.Flag |= BASIC_TYPE_FLAG_DELETED
}

var zeroPrimaryId = primitive.NilObjectID

func ZeroPrimaryId() PrimaryId {
	return zeroPrimaryId
}
