package controllers

type LoginController struct {
	BaseController
}

func (l *LoginController) Login() {

	l.TplName = "login/login.html"
}
