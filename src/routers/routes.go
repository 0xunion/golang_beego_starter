package routers

/* @MT-TPL-IMPORT-LIST-START */
import (
	common "github.com/0xunion/exercise_back/src/controller/common"
	curd "github.com/0xunion/exercise_back/src/controller/curd"
	custom "github.com/0xunion/exercise_back/src/controller/custom"
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
	beego.Router("/api/common/auth/init/root", &common.InitRootUserController{})
	beego.Router("/api/common/auth/self/info", &common.InfoSelfController{})
	beego.Router("/api/common/auth/reset/password", &common.ResetPasswordController{})

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

	// file
	beego.Router("/api/common/attacker/file/upload", &common.AttackerCreateFileController{})
	beego.Router("/api/common/admin/file/upload", &common.AdminUploadFileController{})
	beego.Router("/api/common/file/download", &common.GetFileController{})
	beego.Router("/api/common/image/upload", &common.UploadImageController{})
	beego.Router("/api/common/image/download", &common.GetImageController{})
}

func registerCustom() {
	/* @MT-TPL-ROUTE-CUSTOM-START */
    beego.Router("/api/custom/defender/self", &custom.ApiCustomDefenderSelfController{})
    beego.Router("/api/custom/attacker/self", &custom.ApiCustomAttackerSelfController{})
    beego.Router("/api/custom/defender/attacker/list", &custom.ApiCustomDefenderAttackerListController{})
	beego.Router("/api/custom/manage/trace_report/reject", &custom.ApiCustomManageTraceReportRejectController{})
	beego.Router("/api/custom/manage/trace_report/accept", &custom.ApiCustomManageTraceReportAcceptController{})
	beego.Router("/api/custom/manage/trace_report/detail", &custom.ApiCustomManageTraceReportDetailController{})
	beego.Router("/api/custom/manage/trace_report/list", &custom.ApiCustomManageTraceReportListController{})
	beego.Router("/api/custom/defender/report/submit", &custom.ApiCustomDefenderReportSubmitController{})
	beego.Router("/api/custom/defender/report/detail", &custom.ApiCustomDefenderReportDetailController{})
	beego.Router("/api/custom/defender/report/list", &custom.ApiCustomDefenderReportListController{})
	beego.Router("/api/custom/defender/assets/list", &custom.ApiCustomDefenderAssetsListController{})
	beego.Router("/api/custom/manage/attack/list", &custom.ApiCustomManageAttackListController{})
	beego.Router("/api/custom/all/boardcast/list", &custom.ApiCustomAllBoardcastListController{})
	beego.Router("/api/custom/manage/boardcast/publish", &custom.ApiCustomManageBoardcastPublishController{})
	beego.Router("/api/custom/manage/assets/list", &custom.ApiCustomManageAssetsListController{})
	beego.Router("/api/custom/manage/report/detail", &custom.ApiCustomManageReportDetailController{})
	beego.Router("/api/custom/admin/gamer/list", &custom.ApiCustomAdminGamerListController{})
	beego.Router("/api/custom/admin/game/update", &custom.ApiCustomAdminGameUpdateController{})
	beego.Router("/api/custom/all/game/detail", &custom.ApiCustomAllGameDetailController{})
	beego.Router("/api/custom/attacker/report/commit", &custom.ApiCustomAttackerReportCommitController{})
	beego.Router("/api/custom/defender/statistics", &custom.ApiCustomDefenderStatisticsController{})
	beego.Router("/api/custom/attacker/statistics", &custom.ApiCustomAttackerStatisticsController{})
	beego.Router("/api/custom/all/game/list", &custom.ApiCustomAllGameListController{})
	beego.Router("/api/custom/manage/rank/blue_team", &custom.ApiCustomManageRankBlueTeamController{})
	beego.Router("/api/custom/manage/rank/red_team", &custom.ApiCustomManageRankRedTeamController{})
	beego.Router("api/custom/admin/game/template/judge", &custom.ApiCustomAdminGameTemplateJudgeController{})
	beego.Router("/api/custom/admin/game/list", &custom.ApiCustomAdminGameListController{})
	beego.Router("/api/custom/attacker/report/appeal", &custom.ApiCustomAttackerReportAppealController{})
	beego.Router("/api/custom/attacker/report/delete", &custom.ApiCustomAttackerReportDeleteController{})
	beego.Router("/api/custom/attacker/report/detail", &custom.ApiCustomAttackerReportDetailController{})
	beego.Router("/api/custom/attacker/report/list", &custom.ApiCustomAttackerReportListController{})
	beego.Router("/api/custom/attacker/attack/detail", &custom.ApiCustomAttackerAttackDetailController{})
	beego.Router("/api/custom/attacker/attack/list", &custom.ApiCustomAttackerAttackListController{})
	beego.Router("/api/custom/attacker/attack/apply", &custom.ApiCustomAttackerAttackApplyController{})
	beego.Router("/api/custom/attacker/attack/section", &custom.ApiCustomAttackerAttackSectionController{})
	beego.Router("/api/custom/attacker/report_section", &custom.ApiCustomAttackerReportSectionController{})
	beego.Router("/api/custom/attacker/list_defender", &custom.ApiCustomAttackerListDefenderController{})
	beego.Router("/api/custom/manage/attack/reject", &custom.ApiCustomManageAttackRejectController{})
	beego.Router("/api/custom/manage/attack/accept", &custom.ApiCustomManageAttackAcceptController{})
	beego.Router("/api/custom/manage/report/comment", &custom.ApiCustomManageReportCommentController{})
	beego.Router("/api/custom/manage/report/reject", &custom.ApiCustomManageReportRejectController{})
	beego.Router("/api/custom/manage/report/accept", &custom.ApiCustomManageReportAcceptController{})
	beego.Router("/api/custom/manage/report/list", &custom.ApiCustomManageReportListController{})
	beego.Router("/api/custom/admin/game/import/part_a", &custom.ApiCustomAdminGameImportPartAController{})
	beego.Router("/api/custom/admin/game/template/part_a", &custom.ApiCustomAdminGameTemplatePartAController{})
	beego.Router("/api/custom/admin/game/import/leader", &custom.ApiCustomAdminGameImportLeaderController{})
	beego.Router("/api/custom/admin/game/template/leader", &custom.ApiCustomAdminGameTemplateLeaderController{})
	beego.Router("/api/custom/admin/game/import/judge", &custom.ApiCustomAdminGameImportJudgeController{})
	beego.Router("/api/custom/admin/game/template/judge", &custom.ApiCustomAdminGameTemplateJudgeController{})
	beego.Router("/api/custom/admin/game/import/blue_team", &custom.ApiCustomAdminGameImportBlueTeamController{})
	beego.Router("/api/custom/admin/game/template/blue_team", &custom.ApiCustomAdminGameTemplateBlueTeamController{})
	beego.Router("/api/custom/admin/game/template/defender", &custom.ApiCustomAdminGameTemplateDefenderController{})
	beego.Router("/api/custom/admin/game/import/red_team", &custom.ApiCustomAdminGameImportRedTeamController{})
	beego.Router("/api/custom/admin/game/import/defender", &custom.ApiCustomAdminGameImportDefenderController{})
	beego.Router("/api/custom/admin/game/template/red_team", &custom.ApiCustomAdminGameTemplateRedTeamController{})
	beego.Router("/api/custom/admin/game/create", &custom.ApiCustomAdminGameCreateController{})
    /* @MT-TPL-ROUTE-CUSTOM-END */
}
