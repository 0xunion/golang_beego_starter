/* @MT-TPL-PACKAGE-START */
package permission
/* @MT-TPL-PACKAGE-END */

/* @MT-TPL-IMPORT-START */
import (
    "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-TYPE-START */
type DefenderPermission struct {
    Id types.PrimaryId `json:"id" bson:"_id,omitempty"` // primary id
    Type int `json:"type" bson:"type"` // which type of permission, e.g. "user", "group", "role"
    OwnerId types.PrimaryId `json:"owner_id" bson:"owner_id"` // who owns this permission
    Permission int `json:"permission" bson:"permission"` // permission value 
    TargetId types.PrimaryId `json:"target_id" bson:"target_id"` // which target this permission belongs to
}
/* @MT-TPL-TYPE-END */

/* @MT-TPL-PERMISSION-START */
const (
    Defender_PERMISSION_TYPE_USER = 1
    Defender_PERMISSION_TYPE_GROUP = 2

    Defender_PERMISSION_READ = 1
    Defender_PERMISSION_WRITE = 2
    Defender_PERMISSION_DELETE = 4
    Defender_PERMISSION_ALL = Defender_PERMISSION_READ | Defender_PERMISSION_WRITE | Defender_PERMISSION_DELETE
)
/* @MT-TPL-PERMISSION-END */

/* @MT-TPL-FUNC-START */
func (p *DefenderPermission) GetId() types.PrimaryId {
    return p.Id
}

func (p *DefenderPermission) GroupAccessRead(gid types.PrimaryId) bool {
    return p.Type == Defender_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Defender_PERMISSION_READ != 0
}

func (p *DefenderPermission) GroupAccessWrite(gid types.PrimaryId) bool {
    return p.Type == Defender_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Defender_PERMISSION_WRITE != 0
}

func (p *DefenderPermission) GroupAccessDelete(gid types.PrimaryId) bool {
    return p.Type == Defender_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Defender_PERMISSION_DELETE != 0
}

func (p *DefenderPermission) UserAccessRead(uid types.PrimaryId) bool {
    return p.Type == Defender_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Defender_PERMISSION_READ != 0
}

func (p *DefenderPermission) UserAccessWrite(uid types.PrimaryId) bool {
    return p.Type == Defender_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Defender_PERMISSION_WRITE != 0
}

func (p *DefenderPermission) UserAccessDelete(uid types.PrimaryId) bool {
    return p.Type == Defender_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Defender_PERMISSION_DELETE != 0
}

func (p *DefenderPermission) SetRead() {
    p.Permission |= Defender_PERMISSION_READ
}

func (p *DefenderPermission) SetWrite() {
    p.Permission |= Defender_PERMISSION_WRITE
}

func (p *DefenderPermission) SetDelete() {
    p.Permission |= Defender_PERMISSION_DELETE
}

func (p *DefenderPermission) SetAll() {
    p.Permission |= Defender_PERMISSION_ALL
}

func (p *DefenderPermission) UnsetRead() {
    p.Permission &= ^Defender_PERMISSION_READ
}

func (p *DefenderPermission) UnsetWrite() {
    p.Permission &= ^Defender_PERMISSION_WRITE
}

func (p *DefenderPermission) UnsetDelete() {
    p.Permission &= ^Defender_PERMISSION_DELETE
}

func (p *DefenderPermission) UnsetAll() {
    p.Permission &= ^Defender_PERMISSION_ALL
}
/* @MT-TPL-FUNC-END */