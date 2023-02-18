/* @MT-TPL-PACKAGE-START */
package permission
/* @MT-TPL-PACKAGE-END */

/* @MT-TPL-IMPORT-START */
import (
    "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-TYPE-START */
type GamePermission struct {
    Id types.PrimaryId `json:"id" bson:"_id,omitempty"` // primary id
    Type int `json:"type" bson:"type"` // which type of permission, e.g. "user", "group", "role"
    OwnerId types.PrimaryId `json:"owner_id" bson:"owner_id"` // who owns this permission
    Permission int `json:"permission" bson:"permission"` // permission value 
    TargetId types.PrimaryId `json:"target_id" bson:"target_id"` // which target this permission belongs to
}
/* @MT-TPL-TYPE-END */

/* @MT-TPL-PERMISSION-START */
const (
    Game_PERMISSION_TYPE_USER = 1
    Game_PERMISSION_TYPE_GROUP = 2

    Game_PERMISSION_READ = 1
    Game_PERMISSION_WRITE = 2
    Game_PERMISSION_DELETE = 4
    Game_PERMISSION_ALL = Game_PERMISSION_READ | Game_PERMISSION_WRITE | Game_PERMISSION_DELETE
)
/* @MT-TPL-PERMISSION-END */

/* @MT-TPL-FUNC-START */
func (p *GamePermission) GetId() types.PrimaryId {
    return p.Id
}

func (p *GamePermission) GroupAccessRead(gid types.PrimaryId) bool {
    return p.Type == Game_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Game_PERMISSION_READ != 0
}

func (p *GamePermission) GroupAccessWrite(gid types.PrimaryId) bool {
    return p.Type == Game_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Game_PERMISSION_WRITE != 0
}

func (p *GamePermission) GroupAccessDelete(gid types.PrimaryId) bool {
    return p.Type == Game_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & Game_PERMISSION_DELETE != 0
}

func (p *GamePermission) UserAccessRead(uid types.PrimaryId) bool {
    return p.Type == Game_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Game_PERMISSION_READ != 0
}

func (p *GamePermission) UserAccessWrite(uid types.PrimaryId) bool {
    return p.Type == Game_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Game_PERMISSION_WRITE != 0
}

func (p *GamePermission) UserAccessDelete(uid types.PrimaryId) bool {
    return p.Type == Game_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & Game_PERMISSION_DELETE != 0
}

func (p *GamePermission) SetRead() {
    p.Permission |= Game_PERMISSION_READ
}

func (p *GamePermission) SetWrite() {
    p.Permission |= Game_PERMISSION_WRITE
}

func (p *GamePermission) SetDelete() {
    p.Permission |= Game_PERMISSION_DELETE
}

func (p *GamePermission) SetAll() {
    p.Permission |= Game_PERMISSION_ALL
}

func (p *GamePermission) UnsetRead() {
    p.Permission &= ^Game_PERMISSION_READ
}

func (p *GamePermission) UnsetWrite() {
    p.Permission &= ^Game_PERMISSION_WRITE
}

func (p *GamePermission) UnsetDelete() {
    p.Permission &= ^Game_PERMISSION_DELETE
}

func (p *GamePermission) UnsetAll() {
    p.Permission &= ^Game_PERMISSION_ALL
}
/* @MT-TPL-FUNC-END */