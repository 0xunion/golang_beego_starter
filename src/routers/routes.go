package routers

/* @MT-TPL-IMPORT-LIST-START */
import (
	common "github.com/0xunion/exercise_back/src/controller/common"
	curd "github.com/0xunion/exercise_back/src/controller/curd"
	beego "github.com/beego/beego/v2/server/web"
)

/* @MT-TPL-IMPORT-LIST-END */

/* @MT-TPL-ROUTE-CURD-FUNC-DECL-START */
func registerCURDApi() {
	/* @MT-TPL-ROUTE-CURD-FUNC-DECL-END */
	/* @MT-TPL-ROUTE-CURD-START */
	beego.Router("/init", &curd.InitController{})
	/* @MT-TPL-ROUTE-CURD-END */
	/* @MT-TPL-ROUTE-CURD-FUNC-END-START */
}

/* @MT-TPL-ROUTE-CURD-FUNC-END-END */

func registerProd() {
}

func registerCommon() {
	beego.Router("/api/common/captcha/image/login", &common.ImageMathmaticalCaptchaController{})
	beego.Router("/api/common/auth/login/email_password", &common.UserLoginByEmailAndPasswordController{})
	beego.Router("/api/common/auth/login/phone_password", &common.UserLoginByPhoneAndPasswordController{})

	// group
	beego.Router("/api/common/auth/group/create", &common.CreateGroupController{})
	beego.Router("/api/common/auth/group/joined", &common.ListMyGroupsController{})
	beego.Router("/api/common/auth/group/list", &common.ListGroupsController{})
	beego.Router("/api/common/auth/group/info", &common.InfoGroupController{})
	beego.Router("/api/common/auth/group/update", &common.UpdateGroupController{})
	beego.Router("/api/common/auth/group/delete", &common.DeleteGroupController{})
	beego.Router("/api/common/auth/group/user/list", &common.ListGroupMembersController{})
	beego.Router("/api/common/auth/group/user/create/email_password", &common.CreateUserInGroupByEmailAndPasswordController{})
	beego.Router("/api/common/auth/group/user/create/phone_password", &common.CreateUserInGroupByPhoneAndPasswordController{})
	beego.Router("/api/common/auth/group/user/create/excel", &common.CreateUserInGroupByExcelController{})
	beego.Router("/api/common/auth/group/user/permission/update", &common.UpdateGroupMemberRoleController{})
	beego.Router("/api/common/auth/group/user/delete", &common.DeleteGroupMemberController{})
}
