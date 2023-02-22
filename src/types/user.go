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
	USER_FLAG_ROOT  = 1 << 1 // super admin, can do anything, only one user can be root

	/* @MT-TPL-PERMISSION-START */
	/* @MT-TPL-PERMISSION-END */

	USER_DEFAULT_FLAG             = 0
	USER_DEFAULT_MODEL_PERMISSION = 0
)

func (u *User) IsAdmin() bool {
	return (u.UserFlag&USER_FLAG_ADMIN == USER_FLAG_ADMIN) || (u.UserFlag&USER_FLAG_ROOT == USER_FLAG_ROOT)
}

/* @MT-TPL-PERMISSION-FUNC-START */

/* @MT-TPL-PERMISSION-FUNC-END */
