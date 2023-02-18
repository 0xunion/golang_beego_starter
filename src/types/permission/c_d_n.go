/* @MT-TPL-PACKAGE-START */
package permission
/* @MT-TPL-PACKAGE-END */

/* @MT-TPL-IMPORT-START */
import (
	"github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-TYPE-START */
type CDNPermission struct {
    Id types.PrimaryId `json:"id" bson:"_id,omitempty"` // primary id
    Type int `json:"type" bson:"type"` // which type of permission, e.g. "user", "group", "role"
    OwnerId types.PrimaryId `json:"owner_id" bson:"owner_id"` // who owns this permission
    Permission int `json:"permission" bson:"permission"` // permission value 
    TargetId types.PrimaryId `json:"target_id" bson:"target_id"` // which target this permission belongs to
}
/* @MT-TPL-TYPE-END */

/* @MT-TPL-PERMISSION-START */
const (
    CDN_PERMISSION_TYPE_USER = 1
    CDN_PERMISSION_TYPE_GROUP = 2

    CDN_PERMISSION_READ = 1
    CDN_PERMISSION_WRITE = 2
    CDN_PERMISSION_DELETE = 4
    CDN_PERMISSION_ALL = CDN_PERMISSION_READ | CDN_PERMISSION_WRITE | CDN_PERMISSION_DELETE
)
/* @MT-TPL-PERMISSION-END */

/* @MT-TPL-FUNC-START */
func (p *CDNPermission) GetId() types.PrimaryId {
    return p.Id
}

func (p *CDNPermission) GroupAccessRead(gid types.PrimaryId) bool {
    return p.Type == CDN_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & CDN_PERMISSION_READ != 0
}

func (p *CDNPermission) GroupAccessWrite(gid types.PrimaryId) bool {
    return p.Type == CDN_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & CDN_PERMISSION_WRITE != 0
}

func (p *CDNPermission) GroupAccessDelete(gid types.PrimaryId) bool {
    return p.Type == CDN_PERMISSION_TYPE_GROUP && p.OwnerId == gid && p.Permission & CDN_PERMISSION_DELETE != 0
}

func (p *CDNPermission) UserAccessRead(uid types.PrimaryId) bool {
    return p.Type == CDN_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & CDN_PERMISSION_READ != 0
}

func (p *CDNPermission) UserAccessWrite(uid types.PrimaryId) bool {
    return p.Type == CDN_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & CDN_PERMISSION_WRITE != 0
}

func (p *CDNPermission) UserAccessDelete(uid types.PrimaryId) bool {
    return p.Type == CDN_PERMISSION_TYPE_USER && p.OwnerId == uid && p.Permission & CDN_PERMISSION_DELETE != 0
}

func (p *CDNPermission) SetRead() {
    p.Permission |= CDN_PERMISSION_READ
}

func (p *CDNPermission) SetWrite() {
    p.Permission |= CDN_PERMISSION_WRITE
}

func (p *CDNPermission) SetDelete() {
    p.Permission |= CDN_PERMISSION_DELETE
}

func (p *CDNPermission) SetAll() {
    p.Permission |= CDN_PERMISSION_ALL
}

func (p *CDNPermission) UnsetRead() {
    p.Permission &= ^CDN_PERMISSION_READ
}

func (p *CDNPermission) UnsetWrite() {
    p.Permission &= ^CDN_PERMISSION_WRITE
}

func (p *CDNPermission) UnsetDelete() {
    p.Permission &= ^CDN_PERMISSION_DELETE
}

func (p *CDNPermission) UnsetAll() {
    p.Permission &= ^CDN_PERMISSION_ALL
}
/* @MT-TPL-FUNC-END */
