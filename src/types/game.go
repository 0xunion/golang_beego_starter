package types

type Game struct {
	BasicType
	Id          PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner       PrimaryId `json:"owner" bson:"owner"`
	Description string    `json:"description" bson:"description"`
	HeaderHtml  string    `json:"header_html" bson:"header_html"`
	Name        string    `json:"name" bson:"name"`
	CreateAt    int64     `json:"create_at" bson:"create_at"`
	StartTime   int64     `json:"start_time" bson:"start_time"`
	EndTime     int64     `json:"end_time" bson:"end_time"`
}
