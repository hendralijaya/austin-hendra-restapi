package controller

import (
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var u web.UserLoginRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)

	if ok {
		return
	}

	user, err := c.authService.VerifyCredential(u)
	ok = helper.ValidationError(ctx, err)

	if ok {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *authController) Register(ctx *gin.Context) {
	var u web.UserRegisterRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	user, err := c.authService.Create(u)
	ok = helper.ValidationError(ctx, err)
	if ok {
		return
	}
	userIdString := strconv.FormatUint(user.Id, 10)
	token, err := service.JWTService.GenerateToken(c.jwtService,userIdString,user.Username)
	ok = helper.InternalServerError(ctx, err)
	if ok {
		return
	}
	mainLink := helper.GetMainLink()
	helper.SendMail(`<a href="`+ mainLink+`/verify_register_token/`+token + `</a>`, "Verification Email",user.Email, "", "")
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

func (c *authController) Logout(ctx *gin.Context) {
	helper.ClearSession(ctx)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *authController) ForgotPassword(ctx *gin.Context) {

}
