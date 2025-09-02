package handlers

import (
	"net/http"
	"rest-api/src/apps/api/handlers/dto"
	"rest-api/src/core/domain"
	"rest-api/src/core/interfaces/primary"
	"strconv"

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

func (u *UserHandler) GetUserByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	user, err := u.userService.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	userResponse := dto.FromDomain(user)
	return c.JSON(http.StatusOK, userResponse)

}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	err = u.userService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "user deleted successfully"})
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	userResponses := dto.FromDomainList(users)
	return c.JSON(http.StatusOK, userResponses)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req dto.CreateUserDTO
	err = c.Bind(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	err = req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	existingUser, err := u.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	existingUser.SetName(req.Name)
	existingUser.SetEmail(req.Email)
	existingUser.SetAge(req.Age)

	updatedUser, err := u.userService.Update(existingUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	userResponse := dto.FromDomain(updatedUser)
	return c.JSON(http.StatusOK, userResponse)
}
