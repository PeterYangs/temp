package routes

import (
	"github.com/PeterYangs/superAdminCore/route"
	"superadmin/controller"
	"superadmin/controller/access"
	"superadmin/controller/admin"
	"superadmin/controller/captcha"
	"superadmin/controller/category"
	"superadmin/controller/file"
	"superadmin/controller/login"
	"superadmin/controller/menu"
	"superadmin/controller/queue"
	"superadmin/controller/role"
	"superadmin/controller/rule"
	"superadmin/controller/upload"
	"superadmin/middleware/authCheck"
	"superadmin/middleware/loginCheck"
	"superadmin/middleware/loginLimiter"
)

func Routes(_r route.Group) {

	_r.Registered(route.GET, "/index", controller.Index).Bind()

	_r.Group("/login", func(g route.Group) {

		g.Registered(route.POST, "/login", login.Login, loginLimiter.LoginLimiter).Bind()

		g.Registered(route.POST, "/logout", login.Logout).Bind()

	})

	_r.Group("/admin", func(_admin route.Group) {

		_admin.Group("/rule", func(_rule route.Group) {

			_rule.Registered(route.POST, "/update", rule.Update).Bind()
			_rule.Registered(route.GET, "/list", rule.List).Bind()
			_rule.Registered(route.GET, "/detail/:id", rule.Detail).Bind()
			_rule.Registered(route.POST, "/destroy/:id", rule.Destroy).Bind()

		})

		_admin.Group("/role", func(_role route.Group) {

			_role.Registered(route.GET, "/GetRuleList", role.GetRuleList).SetTag("skip_auth").Bind()
			_role.Registered(route.POST, "/update", role.Update).Bind()
			_role.Registered(route.GET, "/list", role.List).Bind()
			_role.Registered(route.GET, "/detail/:id", role.Detail).Bind()
			_role.Registered(route.POST, "/destroy/:id", role.Destroy).Bind()

		})

		_admin.Group("/admin", func(_admins route.Group) {

			_admins.Registered(route.GET, "/getRoleList", admin.GetRoleList).Bind()
			_admins.Registered(route.POST, "/registered", login.Registered).Bind()
			_admins.Registered(route.POST, "/list", admin.List).Bind()
			_admins.Registered(route.POST, "/detail/:id", admin.Detail).Bind()
			_admins.Registered(route.GET, "/info", admin.Info).SetTag("skip_auth").Bind()
			_admins.Registered(route.GET, "/SearchRule", admin.SearchRule).SetTag("skip_auth").Bind()
			_admins.Registered(route.POST, "/destroy/:id", admin.Destroy).Bind()
			_admins.Registered(route.POST, "/getMyMenu", admin.GetMyMenu).SetTag("skip_auth").Bind()
			_admins.Registered(route.POST, "/roleList", admin.RoleList).Bind()
			_admins.Registered(route.GET, "/getAllRule", admin.GetAllRule).SetTag("skip_auth").Bind()

		})

		_admin.Group("/menu", func(_menu route.Group) {

			_menu.Registered(route.GET, "/getFatherMenu", menu.GetFatherMenu).SetTag("skip_auth").Bind()
			_menu.Registered(route.POST, "/update", menu.Update).Bind()
			_menu.Registered(route.POST, "/list", menu.List).Bind()
			_menu.Registered(route.GET, "/detail/:id", menu.Detail).Bind()

		})

		_admin.Group("/category", func(_category route.Group) {

			_category.Registered(route.GET, "/list", category.List).Bind()
			_category.Registered(route.POST, "/update", category.Update).Bind()

		})

		_admin.Group("/upload", func(_upload route.Group) {

			_upload.Registered(route.POST, "/upload", upload.Upload).Bind()
			_upload.Registered(route.ANY, "/big_file", upload.BigFile).Bind()
		})

		_admin.Group("/queue", func(_queue route.Group) {

			_queue.Registered(route.POST, "/list", queue.List).Bind()
			_queue.Registered(route.POST, "/delay_list", queue.DelayList).Bind()

		})

		_admin.Group("/access", func(_access route.Group) {

			_access.Registered(route.POST, "/list", access.List).Bind()

		})

		_admin.Group("/file", func(_file route.Group) {

			_file.Registered(route.POST, "/update", file.Update).Bind()
			_file.Registered(route.POST, "/list", file.List).Bind()
			_file.Registered(route.POST, "/destroy/:id", file.Destroy).Bind()

		})

	}, loginCheck.LoginCheck, authCheck.AuthCheck)

	_r.Group("/captcha", func(g route.Group) {

		g.Registered(route.GET, "/captcha", captcha.Captcha).Bind()
	})

}
