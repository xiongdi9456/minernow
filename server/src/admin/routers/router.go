package routers

import (
	"admin/controllers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(before),
		beego.NSNamespace("/admin",
			beego.NSRouter("/login", &controllers.LoginController{}, "post:HandleLogin"),
			//beego.NSRouter("/relogin", &controllers.LoginController{}, "post:HandleReLogin"),
			beego.NSRouter("/logout", &controllers.LoginController{}, "post:HandleLogout"),
			beego.NSRouter("/password", &controllers.LoginController{}, "post:HandlePassword"),

			beego.NSRouter("/operations", &controllers.OperationLogController{}, "get:HandleGetAdminOperationLog"),
			beego.NSRouter("/permissions", &controllers.PermissionController{}, "get:HandleGetPermissions;post:HandleCreatePermissions;put:HandleUpdatePermissions;delete:HandleDeletePermissions"),

			beego.NSRouter("/admins", &controllers.AdminUserController{}, "get:HandleGetAdmins;post:HandleCreateAdmins;put:HandleUpdateAdmins;delete:HandleDelAdmins"),
			beego.NSNamespace("/admins",
				beego.NSRouter("/roles", &controllers.AdminUserController{}, "get:HandleGetAdminsRoles;put:HandleAddAdminsRoles;delete:HandleDelAdminsRoles"),
				beego.NSRouter("/suspend", &controllers.AdminUserController{}, "post:HandleSuspendAdmins"),
				beego.NSRouter("/restore", &controllers.AdminUserController{}, "post:HandleRestoreAdmins"),
				beego.NSRouter("/unbind", &controllers.AdminUserController{}, "post:UnBind"),
			),
			beego.NSNamespace("/game",
				beego.NSRouter("/chargeinfo", &controllers.GameInfoOptController{}, "get:QueryChargeRecord"),
				beego.NSRouter("/manaulwrisk", &controllers.GameInfoOptController{}, "get:ManaulWriteRisk"),
				beego.NSRouter("/genstatistic/:action", &controllers.GameInfoOptController{}, "get:GenStatisticInfo"),
				beego.NSRouter("/statistic/:action", &controllers.GameInfoOptController{}, "get:GetStatisticInfo"),
				beego.NSRouter("/report/:action", &controllers.GameInfoOptController{}, "get:GetReportInfo"),
				beego.NSRouter("/manualdebug/statistic", &controllers.GameInfoOptController{}, "get:DebugStatisticGen"),
				beego.NSRouter("/order/risk", &controllers.GameInfoOptController{}, "get:ManualOrderRisk"),
			),

			beego.NSNamespace("/roles",
				beego.NSRouter("", &controllers.AdminRoleController{}, "get:HandleGetRoles;post:HandleCreateRole;put:HandleUpdateRole;delete:HandleDeleteRole"),
				beego.NSRouter("/permissions", &controllers.AdminRoleController{}, "get:HandleGetRolePermissions;put:HandleAddRolePermissions;delete:HandleDeleteRolePermissions"),
				beego.NSRouter("/members", &controllers.AdminRoleController{}, "get:HandleGetRoleMembers"),
			),

			beego.NSRouter("/commissionrates", &controllers.CommissionRateController{}, "get:GetCommRates;post:EditCommRates"),
			beego.NSRouter("/smstemplates", &controllers.SmsController{}, "get:GetSmsTemplates;post:AddSmsTemplates;put:UpdateSmsTemplates;delete:DelSmsTemplates"),
			beego.NSRouter("/smscodes", &controllers.SmsController{}, "get:GetSmsCodes"),

			beego.NSNamespace("/notifications",
				beego.NSRouter("", &controllers.SysNotificationController{}, "get:Get;post:Create;put:Update;delete:Del"),
				beego.NSRouter("/publish", &controllers.SysNotificationController{}, "post:Publish"),
				beego.NSRouter("/unpublish", &controllers.SysNotificationController{}, "post:UnPublish"),
			),

			beego.NSRouter("/systemmessages", &controllers.SystemMessageMethodControllers{}, "get:GetAllSystemMessage;post:AddNewSystemMessage"),
			beego.NSRouter("/systemmessages/:id", &controllers.SystemMessageMethodControllers{}, "get:GetSystemMessageById;put:UpdateSystemMessageById;delete:DelSystemMessageById"),
			beego.NSNamespace("/orders",
				beego.NSRouter(":order_id/messages", &controllers.AdminMessageMethodController{}, "post:GetMessage"),
			),
			beego.NSRouter(fmt.Sprintf("/%s/version", controllers.KEY_SYSTEM_INPUT), &controllers.AppVersionController{}, "get:GetLastAppVersion"),
			beego.NSRouter("/versions", &controllers.AppVersionController{}, "get:GetAppVersions;post:CreateAppVersion"),
			beego.NSRouter(fmt.Sprintf("/versions/%s", controllers.KEY_ID_INPUT), &controllers.AppVersionController{}, "put:UpdateAppVersion;delete:DelAppVersion"),
			beego.NSRouter(fmt.Sprintf("/versions/%s/publish", controllers.KEY_ID_INPUT), &controllers.AppVersionController{}, "post:PublishAppVersion"),
			beego.NSRouter(fmt.Sprintf("/versions/%s/unpublish", controllers.KEY_ID_INPUT), &controllers.AppVersionController{}, "post:UnPublishAppVersion"),

			beego.NSRouter(fmt.Sprintf("/config/%s", controllers.KEY_KEY_ACTION), &controllers.ConfigController{}, "get:List;post:Create;put:Update;delete:Delete"),
			beego.NSRouter(fmt.Sprintf("/config/%s/refresh", controllers.KEY_KEY_ACTION), &controllers.ConfigController{}, "put:Flush"),
			beego.NSRouter(fmt.Sprintf("/config/%s/clean", controllers.KEY_KEY_ACTION), &controllers.ConfigController{}, "delete:Clean"),
			beego.NSRouter("/configwarning", &controllers.ConfigController{}, "get:ListWarning;post:CreateWarning;put:UpdateWarning;delete:DelWarning"),

			beego.NSRouter("/users-bulk", &controllers.OtcUserController{}, "delete:DelOctUsers"),
			beego.NSNamespace("/users",
				beego.NSRouter("", &controllers.OtcUserController{}, "get:GetOtcUser;put:UpdateOtcUser;delete:DelOctUser"),
				beego.NSRouter("/syncusdtbyuid", &controllers.OtcUserController{}, "post:SyncTransactionByUid"),
				//beego.NSRouter("/syncusdtbymobile", &controllers.OtcUserController{}, "post:SyncTransactionByMobile"),
				beego.NSRouter("/eusdrecharge", &controllers.OtcUserController{}, "post:EusdRecharge"),
				beego.NSRouter("/payment-methods", &controllers.OtcUserController{}, "get:GetPayment"),
			),
			beego.NSRouter("/userloginlogs", &controllers.OtcUserLoginLogController{}, "get:Get"),

			beego.NSNamespace("/exchangersverify",
				beego.NSRouter("", &controllers.OtcExchangerVerifyController{}, "get:Get"),
				beego.NSRouter("/approve", &controllers.OtcExchangerVerifyController{}, "post:Approve"),
				beego.NSRouter("/reject", &controllers.OtcExchangerVerifyController{}, "post:Reject"),
			),
			beego.NSNamespace("/exchangers",
				beego.NSRouter("", &controllers.OtcExchangerController{}, "get:GetExchanger;put:UpdateExchanger;post:CreateExchanger"),
				beego.NSRouter("/order", &controllers.OtcExchangerController{}, "Get:GetOtcOrder"),
				beego.NSRouter("/otc_wealth", &controllers.OtcExchangerController{}, "Get:GetExchangerWealth"),
			),

			beego.NSNamespace("/apps",
				beego.NSRouter("", &controllers.AppController{}, "get:GetApp;post:CreateApp;put:UpdateApp;delete:DelApp"),
				beego.NSRouter("/feature", &controllers.AppController{}, "post:FeatureApp"),
				beego.NSRouter("/unfeature", &controllers.AppController{}, "post:UnFeatureApp"),
				beego.NSRouter("/publish", &controllers.AppController{}, "post:PublishApp"),
				beego.NSRouter("/unpublish", &controllers.AppController{}, "post:UnPublishApp"),
			),
			beego.NSRouter("/appwhite", &controllers.AppWhiteController{}, "get:GetWhiteApp;post:CreateWhiteApp;delete:DelWhiteApp"),
			beego.NSRouter("/apptypes", &controllers.AppTypeController{}, "get:Get;post:Create;put:Update;delete:Del"),
			beego.NSRouter("/appchannels", &controllers.AppChannelController{}, "get:Get;post:Create;put:Update;delete:Del"),
			beego.NSNamespace("/banners",
				beego.NSRouter("", &controllers.BannerController{}, "get:Get;post:Create;put:Update;delete:Del"),
				beego.NSRouter("/publish", &controllers.BannerController{}, "post:Publish"),
				beego.NSRouter("/unpublish", &controllers.BannerController{}, "post:UnPublish"),
			),

			beego.NSNamespace("/commissionstat",
				beego.NSRouter("", &controllers.CommissionStatController{}, "get:Get"),
				beego.NSRouter("/distribute", &controllers.CommissionStatController{}, "post:Distribute"),
				beego.NSRouter("/calccommission", &controllers.CommissionStatController{}, "post:CalcCommission"),
			),

			beego.NSRouter("/topagent", &controllers.TopAgentController{}, "get:Get;post:Create;put:Update;delete:Del"),
			beego.NSNamespace("/agentwhitelists",
				beego.NSRouter("", &controllers.AgentWhiteListController{}, "get:GetAgentWhiteList;post:CreateAgentWhiteList;put:UpdateAgentWhiteList;delete:DelAgentWhiteList"),
				beego.NSRouter("/agents", &controllers.AgentWhiteListController{}, "get:GetAgent;put:AddAgent;delete:DelAgent"),
			),
			beego.NSRouter("/eosrecords", &controllers.AdminEosMethods{}, "get:GetEosTransactions"),
			beego.NSRouter("/eoswallets", &controllers.AdminEosMethods{}, "get:GetEosWealth"),
			beego.NSRouter("/eosaddress", &controllers.AdminEosMethods{}, "get:GetEosAccount"),
			beego.NSNamespace("/eos",
				beego.NSRouter("/frozen", &controllers.AdminEosMethods{}, "post:EosFrozen"),
				beego.NSRouter("/unfrozen", &controllers.AdminEosMethods{}, "post:EosUnFrozen"),
			),

			beego.NSNamespace("/usdt",
				beego.NSRouter("/wallets", &controllers.UsdtController{}, "get:GetWallet"),
				beego.NSRouter("/lock", &controllers.UsdtController{}, "post:Lock"),
				beego.NSRouter("/unlock", &controllers.UsdtController{}, "post:Unlock"),
				beego.NSRouter("/transrecords", &controllers.UsdtController{}, "get:GetTransRecords"),
				beego.NSRouter("/turnrecords", &controllers.UsdtController{}, "get:GetTurnRecords"),
				beego.NSRouter("/cashrecords", &controllers.UsdtController{}, "get:GetCashRecords"),
				beego.NSNamespace("/cashrecords",
					beego.NSRouter("/approve", &controllers.UsdtController{}, "post:ApproveTransferOut"),
					beego.NSRouter("/reject", &controllers.UsdtController{}, "post:RejectTransferOut"),
				),
			),

			beego.NSRouter("/oss", &controllers.OssController{}, "get:Get;post:Post"),

			beego.NSNamespace("/appealservice",
				beego.NSRouter("", &controllers.AppealServiceController{}, "get:Get;post:Create;put:Update;delete:Del"),
				beego.NSRouter("/suspend", &controllers.AppealServiceController{}, "post:Suspend"),
				beego.NSRouter("/restore", &controllers.AppealServiceController{}, "post:Restore"),
				beego.NSRouter("/record", &controllers.AppealServiceController{}, "get:Record"),
			),
			beego.NSRouter("/appeals", &controllers.AppealController{}, "get:Get;put:Update"),
			beego.NSRouter("/prices", &controllers.PriceController{}, "get:Get"),
			beego.NSRouter("/curprice", &controllers.PriceController{}, "get:GetCurPrice"),
			beego.NSNamespace("/orders",
				beego.NSRouter("", &controllers.OrdersController{}, "get:List"),
				beego.NSRouter("/:order_id", &controllers.OrdersController{}, "get:Info"),
				beego.NSRouter("/cancel", &controllers.OrdersController{}, "post:CancelOrder"),
				beego.NSRouter("/pay", &controllers.OrdersController{}, "post:ConfirmOrderPay"),
				beego.NSRouter("/confirm", &controllers.OrdersController{}, "post:ConfirmOrder"),
				beego.NSRouter("/resolve", &controllers.OrdersController{}, "post:ResolveOrder"),
			),
			beego.NSNamespace("/eusd",
				beego.NSRouter("transfer", &controllers.EusdRetireController{}, "post:TransferToTokenAccount"),
				beego.NSRouter("retire", &controllers.EusdRetireController{}, "post:Retire"),
				beego.NSRouter("balance", &controllers.EusdRetireController{}, "post:GetBalanceByUid"),
			),
			beego.NSRouter("/endpoint", &controllers.EndPointController{}, "get:Get;post:Create;put:Update;delete:Del"),
			beego.NSNamespace("/platformuser",
				beego.NSRouter("", &controllers.PlatformUserController{}, "get:List"),
				beego.NSRouter("/add", &controllers.PlatformUserController{}, "post:Add"),
				beego.NSRouter("/status", &controllers.PlatformUserController{}, "post:Status"),
			),
			beego.NSNamespace("/platformusercate",
				beego.NSRouter("", &controllers.PlatformUserCateController{}, "get:List"),
				beego.NSRouter("/add", &controllers.PlatformUserCateController{}, "post:Add"),
			),
			beego.NSRouter("/announcements", &controllers.AnnouncementController{}, "get:Get;post:Create;put:Update;delete:Del"),
			beego.NSNamespace("/serverstop",
				beego.NSRouter("", &controllers.ServerStopController{}, "get:Get"),
				beego.NSRouter("/set", &controllers.ServerStopController{}, "post:Set"),
				beego.NSRouter("/log", &controllers.ServerStopController{}, "get:Log"),
			),
			beego.NSNamespace("/otcstat",
				beego.NSRouter("", &controllers.OtcStatController{}, "get:Get"),
				beego.NSRouter("/user", &controllers.OtcStatController{}, "get:User"),
			),
			beego.NSNamespace("/monthdividendconf",
				beego.NSRouter("/conf", &controllers.MonthDividendConfController{}, "get:GetMonthDividendConf;post:EditMonthDividendConf;delete:DelMonthDividendConf"),
				beego.NSRouter("/positionconf", &controllers.MonthDividendConfController{}, "get:GetMonthDividendPosition"),
				//beego.NSRouter("/agents", &controllers.MonthDividendConfController{}, "get:GetAgents;post:AddAgent;delete:DelAgent"),
			),
			beego.NSNamespace("/monthdividendwhitelist",
				beego.NSRouter("/add", &controllers.MonthDividendWhiteListController{}, "post:CreateMonthDividendWhiteList"),
				beego.NSRouter("/del", &controllers.MonthDividendWhiteListController{}, "delete:DelMonthDividendWhiteList"),
				beego.NSRouter("/update", &controllers.MonthDividendWhiteListController{}, "put:UpdateMonthDividendWhiteList"),
				beego.NSRouter("/get", &controllers.MonthDividendWhiteListController{}, "get:GetMonthDividendWhiteList"),

				beego.NSRouter("/addagent", &controllers.MonthDividendWhiteListController{}, "post:AddAgent"),
				beego.NSRouter("/delagent", &controllers.MonthDividendWhiteListController{}, "delete:DelAgent"),
				beego.NSRouter("/getagent", &controllers.MonthDividendWhiteListController{}, "get:GetAgent"),
				//beego.NSRouter("/agents", &controllers.MonthDividendConfController{}, "get:GetAgents;post:AddAgent;delete:DelAgent"),
			),
			beego.NSNamespace("/general",
				beego.NSRouter("/frozenuser", &controllers.GeneralOperationController{}, "post:FrozenUser"),
				beego.NSRouter("/unfrozenuser", &controllers.GeneralOperationController{}, "post:UnFrozenUser"),
				beego.NSRouter("/lockuser", &controllers.GeneralOperationController{}, "post:LockUser"),
				beego.NSRouter("/unlockuser", &controllers.GeneralOperationController{}, "post:UnlockUser"),
			),
			beego.NSNamespace("/task",
				beego.NSRouter("/query", &controllers.TaskController{}, "get:Query"),
				beego.NSRouter("/querysingle", &controllers.TaskController{}, "get:QuerySingle"),
				beego.NSRouter("/result", &controllers.TaskController{}, "get:QueryResult"),
				beego.NSRouter("/add", &controllers.TaskController{}, "post:Add"),
				beego.NSRouter("/edit", &controllers.TaskController{}, "post:Edit"),
				beego.NSRouter("/del", &controllers.TaskController{}, "post:Delete"),
				beego.NSRouter("/distribute", &controllers.TaskController{}, "post:Distribute"),
				beego.NSRouter("/distributesingle", &controllers.TaskController{}, "post:DistributeSingle"),
				beego.NSRouter("/switch", &controllers.TaskController{}, "post:Switch"),
				beego.NSRouter("/switchsingle", &controllers.TaskController{}, "post:SwitchSingle"),
				beego.NSRouter("/run", &controllers.TaskController{}, "post:Run"),
			),
			beego.NSNamespace("/servernode",
				beego.NSRouter("/query", &controllers.ServerNodeController{}, "get:Query"),
			),
			beego.NSRouter("/userfuzzysearch", &controllers.NickFuzzySearchController{}, "get:GetUsers"),
			beego.NSRouter("/personalreport", &controllers.ReportController{}, "get:GetPersonalReport"),

			beego.NSNamespace("/menu",
				beego.NSRouter("/add", &controllers.MenuController{}, "post:Add"),
				beego.NSRouter("/getlist", &controllers.MenuController{}, "get:GetAllMenu"),
				beego.NSRouter("/updatemenu", &controllers.MenuController{}, "put:UpdateMenu"),
				beego.NSRouter("/del", &controllers.MenuController{}, "delete:Del"),
				beego.NSRouter("/updatemenuaccess", &controllers.MenuController{}, "post:UpdateMenuAccess"),
				beego.NSRouter("/getaccesslist", &controllers.MenuController{}, "get:GetAccessMenus"),
				beego.NSRouter("/getallaccesslist", &controllers.MenuController{}, "get:GetAllAccessMenus"),
			),
			beego.NSNamespace("/threshold",
				beego.NSRouter("/add", &controllers.ProfitThresholdController{}, "post:Add"),
				beego.NSRouter("/del", &controllers.ProfitThresholdController{}, "delete:Del"),
				beego.NSRouter("/update", &controllers.ProfitThresholdController{}, "put:Update"),
				beego.NSRouter("/get", &controllers.ProfitThresholdController{}, "get:Find"),
				beego.NSRouter("/findcheckingrecord", &controllers.ProfitThresholdController{}, "get:FindCheckingRecords"),
				beego.NSRouter("/rejectcheckingrecord", &controllers.ProfitThresholdController{}, "post:RejectCheckingRecord"),
				beego.NSRouter("/agreecheckingrecord", &controllers.ProfitThresholdController{}, "post:AgreeCheckingRecord"),
			),

			beego.NSNamespace("/report",
				beego.NSNamespace("/team",
					beego.NSRouter("/list", &controllers.ReportTeamController{}, "get:List"),
					beego.NSRouter("/sum", &controllers.ReportTeamController{}, "get:Sum"),
					beego.NSRouter("/personal", &controllers.ReportTeamController{}, "get:Personal"),
				),
			),
		),
	)
	beego.AddNamespace(ns)

	return
}

func before(ctx *context.Context) {
	//set output Content-Type to be json
	ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
}
