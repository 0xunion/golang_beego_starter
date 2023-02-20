package types

type File struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Owner    PrimaryId `json:"owner" bson:"owner"`
	Name     string    `json:"name" bson:"name"`
	Hash     string    `json:"hash" bson:"hash"` // random hash for file as key
	Role     int64     `json:"role" bson:"role"` // permission control
	Size     int64     `json:"size" bson:"size"`
	Path     string    `json:"path" bson:"path"`
	GameId   PrimaryId `json:"game_id" bson:"game_id"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

const (
	FILE_ROLE_PUBLIC = 1 << iota
	FILE_ROLE_RED_TEAM
	FILE_ROLE_BLUE_TEAM
	FILE_ROLE_JUDGEMENT
	FILE_ROLE_CUSTOMER
	FILE_ROLE_PARTA
)

func (f *File) PublicAccess() bool {
	return f.Role&FILE_ROLE_PUBLIC != 0
}

func (f *File) RedTeamAccess() bool {
	return f.Role&FILE_ROLE_RED_TEAM != 0
}

func (f *File) BlueTeamAccess() bool {
	return f.Role&FILE_ROLE_BLUE_TEAM != 0
}

func (f *File) JudgementAccess() bool {
	return f.Role&FILE_ROLE_JUDGEMENT != 0
}

func (f *File) CustomerAccess() bool {
	return f.Role&FILE_ROLE_CUSTOMER != 0
}

func (f *File) PartAAccess() bool {
	return f.Role&FILE_ROLE_PARTA != 0
}

func (f *File) SetPublicAccess() {
	f.Role |= FILE_ROLE_PUBLIC
}

func (f *File) SetRedTeamAccess() {
	f.Role |= FILE_ROLE_RED_TEAM
}

func (f *File) SetBlueTeamAccess() {
	f.Role |= FILE_ROLE_BLUE_TEAM
}

func (f *File) SetJudgementAccess() {
	f.Role |= FILE_ROLE_JUDGEMENT
}

func (f *File) SetCustomerAccess() {
	f.Role |= FILE_ROLE_CUSTOMER
}

func (f *File) SetPartAAccess() {
	f.Role |= FILE_ROLE_PARTA
}

func (f *File) ClearPublicAccess() {
	f.Role &= ^FILE_ROLE_PUBLIC
}

func (f *File) ClearRedTeamAccess() {
	f.Role &= ^FILE_ROLE_RED_TEAM
}

func (f *File) ClearBlueTeamAccess() {
	f.Role &= ^FILE_ROLE_BLUE_TEAM
}

func (f *File) ClearJudgementAccess() {
	f.Role &= ^FILE_ROLE_JUDGEMENT
}

func (f *File) ClearCustomerAccess() {
	f.Role &= ^FILE_ROLE_CUSTOMER
}

func (f *File) ClearPartAAccess() {
	f.Role &= ^FILE_ROLE_PARTA
}

func (f *File) UserAccess(user_identity int64) bool {
	switch user_identity {
	case GAMER_IDENTITY_ATTACKER:
		return f.RedTeamAccess()
	case GAMER_IDENTITY_DEFENDER:
		return f.BlueTeamAccess()
	case GAMER_IDENTITY_JUDGEMENT:
		return f.JudgementAccess()
	case GAMER_IDENTITY_CUSTOMER:
		return f.CustomerAccess()
	case GAMER_IDENTITY_PARTA:
		return f.PartAAccess()
	}

	return false
}
