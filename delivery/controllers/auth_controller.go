package controllers

import (
	"chat/domain/entities"
	services "chat/domain/services"
	"chat/infrastructure/security"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthController(service *services.UserService) *AuthController {
	return &AuthController{userService: *service}
}

type AuthController struct {
	userService services.UserService
}

func (c *AuthController) Login(ctx *gin.Context) {
	data, _ := io.ReadAll(ctx.Request.Body)
	var userSearched entities.User
	json.Unmarshal(data, &userSearched)

	user, err := c.userService.GetBy("email", userSearched.Email)

	errorMsg := "Error: wrong credentials or user does not exist"

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	isValidPass := security.CheckPasswordHash(userSearched.Password, user.Password)

	if !isValidPass {
		ctx.JSON(http.StatusOK, errorMsg)
		return
	}

	jwt, err := security.GenerateJWT(*user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusOK, gin.H{"user": user, "jwt": jwt})

}
func (c *AuthController) RevalidateJWT(ctx *gin.Context) {

	token, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	userId, err := security.ValidateJWT("id", token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	user, err := c.userService.GetUserBydId(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "user not found")
		return
	}

	user.Password = ""

	newJWT, _ := security.GenerateJWT(*user)

	ctx.JSON(http.StatusOK, gin.H{"jwt": newJWT, "user": user})

}
