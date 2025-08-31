package handlers

import (
	"net/http"
	"rest-api/src/apps/api/handlers/dto"
	"rest-api/src/core/domain"
	"rest-api/src/core/interfaces/primary"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService primary.UserServicePort
}

func NewUserHandler(userService primary.UserServicePort) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var req dto.CreateUserDTO

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	err = req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var user domain.User
	user.SetName(req.Name)
	user.SetEmail(req.Email)
	user.SetAge(req.Age)

	createdUser, err := u.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	userResponse := dto.FromDomain(createdUser)

	return c.JSON(http.StatusCreated, userResponse)
}

func (u *UserHandler) GetAllUsers(c echo.Context) error  {
	users, err := u.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	usersDTO := make([]dto.UserResponseDTO, 0, len(users))
	for _, userDomain := range users {
		usersDTO = append(usersDTO, dto.FromDomain(userDomain))
	}

	return c.JSON(http.StatusOK, usersDTO)
}

