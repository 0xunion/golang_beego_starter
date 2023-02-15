package types

import "sync"

type User struct {
	BasicType
	Id              PrimaryId `json:"id" bson:"_id,omitempty"`
	UserFlag        int64     `json:"user_flag" bson:"user_flag"`
	ModelPermission int64     `json:"model_permission" bson:"model_permission"`
	Username        string    `json:"username" bson:"username"`
	Avatar          string    `json:"avatar" bson:"avatar"`
	Parent          PrimaryId `json:"parent" bson:"parent"` // who create this user
}

var create_user_mutex struct {
	by_email sync.Mutex
	by_phone sync.Mutex
}

func CreateUserWithEmailMutex() {
	create_user_mutex.by_email.Lock()
}

func CreateUserWithEmailMutexEnd() {
	create_user_mutex.by_email.Unlock()
}

func CreateUserWithPhoneMutex() {
	create_user_mutex.by_phone.Lock()
}

func CreateUserWithPhoneMutexEnd() {
	create_user_mutex.by_phone.Unlock()
}

const (
	USER_FLAG_ADMIN = 1 << 0

	/* @MT-TPL-PERMISSION-START */
	USER_MODEL_PERMISSION_WebFinger = 1 << 2
	USER_MODEL_PERMISSION_Dir       = 1 << 1
	USER_MODEL_PERMISSION_Cdn       = 1 << 0
	/* @MT-TPL-PERMISSION-END */

	USER_DEFAULT_FLAG             = 0
	USER_DEFAULT_MODEL_PERMISSION = 0
)

func (u *User) IsAdmin() bool {
	return u.UserFlag&USER_FLAG_ADMIN == USER_FLAG_ADMIN
}

/* @MT-TPL-PERMISSION-FUNC-START */
func (u *User) AllowManageWebFinger() bool {
	return u.ModelPermission&USER_MODEL_PERMISSION_WebFinger == USER_MODEL_PERMISSION_WebFinger
}
func (u *User) AllowManageDir() bool {
	return u.ModelPermission&USER_MODEL_PERMISSION_Dir == USER_MODEL_PERMISSION_Dir
}
func (u *User) AllowManageCdn() bool {
	return u.ModelPermission&USER_MODEL_PERMISSION_Cdn == USER_MODEL_PERMISSION_Cdn
}

/* @MT-TPL-PERMISSION-FUNC-END */
