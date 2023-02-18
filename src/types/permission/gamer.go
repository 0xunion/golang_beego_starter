/* @MT-TPL-PACKAGE-START */
package permission
/* @MT-TPL-PACKAGE-END */

/* @MT-TPL-IMPORT-START */
import (
    "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-TYPE-START */
type GamerPermission struct {
    Id types.PrimaryId `json:"id" bson:"_id,omitempty"` // primary id
    Type int `json:"type" bson:"type"` // which type of permission, e.g. "user", "group", "role"
    OwnerId types.PrimaryId `json:"owner_id" bson:"owner_id"` // who owns this permission
    Permission int `json:"permission" bson:"permission"` // permission value 
    TargetId types.PrimaryId `json:"target_id" bson:"target_id"` // which target this permission belongs to
}
/* @MT-TPL-TYPE-END */

/* @MT-TPL-PERMISSION-START */
const (
    Gamer_PERMISSION_TYPE_USER = 1
    Gamer_PERMISSION_TYPE_GROUP = 2

    Gamer_PERMISSION_READ = 1
    Gamer_PERMISSION_WRITE = 2
    Gamer_PERMISSION_DELETE = 4
    Gamer_PERMISSION_ALL = Gamer_PERMISSION_READ | Gamer_PERMISSION_WRITE | Gamer_PERMISSION_DELETE
)
/* @MT-TPL-PERMISSION-END */

/* @MT-TPL-FUNC-START */
func (p *GamerPermission) GetId() types.PrimaryId {
    return p.Id
}

func (p *GamerPermission) GroupAccessRead(gid types.PrimaryId) bool {
    return p.Type == Gamer_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Gamer_PERMISSION_READ != 0
}

func (p *GamerPermission) GroupAccessWrite(gid types.PrimaryId) bool {
    return p.Type == Gamer_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Gamer_PERMISSION_WRITE != 0
}

func (p *GamerPermission) GroupAccessDelete(gid types.PrimaryId) bool {
    return p.Type == Gamer_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Gamer_PERMISSION_DELETE != 0
}

func (p *GamerPermission) UserAccessRead(uid types.PrimaryId) bool {
    return p.Type == Gamer_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Gamer_PERMISSION_READ != 0
}

func (p *GamerPermission) UserAccessWrite(uid types.PrimaryId) bool {
    return p.Type == Gamer_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Gamer_PERMISSION_WRITE != 0
}

func (p *GamerPermission) UserAccessDelete(uid types.PrimaryId) bool {
    return p.Type == Gamer_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Gamer_PERMISSION_DELETE != 0
}

func (p *GamerPermission) SetRead() {
    p.Permission |= Gamer_PERMISSION_READ
}

func (p *GamerPermission) SetWrite() {
    p.Permission |= Gamer_PERMISSION_WRITE
}

func (p *GamerPermission) SetDelete() {
    p.Permission |= Gamer_PERMISSION_DELETE
}

func (p *GamerPermission) SetAll() {
    p.Permission |= Gamer_PERMISSION_ALL
}

func (p *GamerPermission) UnsetRead() {
    p.Permission &= ^Gamer_PERMISSION_READ
}

func (p *GamerPermission) UnsetWrite() {
    p.Permission &= ^Gamer_PERMISSION_WRITE
}

func (p *GamerPermission) UnsetDelete() {
    p.Permission &= ^Gamer_PERMISSION_DELETE
}

func (p *GamerPermission) UnsetAll() {
    p.Permission &= ^Gamer_PERMISSION_ALL
}
/* @MT-TPL-FUNC-END */