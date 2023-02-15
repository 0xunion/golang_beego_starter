package types

type Group struct {
	BasicType
	Id          PrimaryId `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Parent      PrimaryId `json:"parent" bson:"parent"` // who create this group
	CreateAt    int64     `json:"create_at" bson:"create_at"`
	Description string    `json:"description" bson:"description"`
}

type GroupMember struct {
	BasicType
	Id       PrimaryId `json:"id" bson:"_id,omitempty"`
	Gid      PrimaryId `json:"gid" bson:"gid"`
	Uid      PrimaryId `json:"uid" bson:"uid"`
	Role     int64     `json:"role" bson:"role"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

const (
	GROUP_MEMBER_ROLE_ADMIN  = 1 << 0
	GROUP_MEMBER_ROLE_USER   = 1 << 1
	GROUP_MEMBER_ROLE_LEADER = 1 << 2
)

func (g *GroupMember) ClearRole() {
	g.Role = 0
}

func (g *GroupMember) IsAdmin() bool {
	return g.Role&GROUP_MEMBER_ROLE_ADMIN == GROUP_MEMBER_ROLE_ADMIN
}

func (g *GroupMember) IsUser() bool {
	return g.Role&GROUP_MEMBER_ROLE_USER == GROUP_MEMBER_ROLE_USER
}

func (g *GroupMember) IsLeader() bool {
	return g.Role&GROUP_MEMBER_ROLE_LEADER == GROUP_MEMBER_ROLE_LEADER
}

func (g *GroupMember) SetAdmin() {
	g.Role = g.Role | GROUP_MEMBER_ROLE_ADMIN
}

func (g *GroupMember) SetUser() {
	g.Role = g.Role | GROUP_MEMBER_ROLE_USER
}

func (g *GroupMember) SetLeader() {
	g.Role = g.Role | GROUP_MEMBER_ROLE_LEADER
}

func (g *GroupMember) UnsetAdmin() {
	g.Role = g.Role &^ GROUP_MEMBER_ROLE_ADMIN
}

func (g *GroupMember) UnsetUser() {
	g.Role = g.Role &^ GROUP_MEMBER_ROLE_USER
}

func (g *GroupMember) UnsetLeader() {
	g.Role = g.Role &^ GROUP_MEMBER_ROLE_LEADER
}
