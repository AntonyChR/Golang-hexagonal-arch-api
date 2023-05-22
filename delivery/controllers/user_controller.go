package controllers

import (
	"chat/domain/entities"
	"chat/domain/services"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserController(service *services.UserService) *UserController {
	return &UserController{userService: *service}
}

type UserController struct {
	userService services.UserService
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUserBydId(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "user not found")
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.userService.GetUserBydId(id)
	errorMsg := "Error deleting user with id: " + id

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	err = c.userService.DeleteUserById(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	ctx.JSON(http.StatusOK, "User deleted")
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	data, _ := io.ReadAll(ctx.Request.Body)
	var user entities.User

	json.Unmarshal(data, &user)

	newUser, err := c.userService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusCreated, newUser)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	data, err := c.userService.GetAllUsers()
	clientId, _ := ctx.Get("userId")
	newData := removeUser(data, clientId.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Data not found")
		return
	}

	ctx.JSON(http.StatusOK, newData)
}

func removeUser(users []*entities.User, id string) []*entities.User{
	index := -1
	for i, user := range users {
		if user.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		users = append(users[:index], users[index+1:]...)
	}

	return users
}
