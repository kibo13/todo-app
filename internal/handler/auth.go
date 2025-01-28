package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kibo13/todo-app/internal/entity"
	"net/http"
)

// @Summary Registration
// @Tags Auth
// @Description User registration
// @ID registration
// @Accept json
// @Produce json
// @Param input body entity.User true "Account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Login
// @Tags Auth
// @Description User login
// @ID login
// @Accept json
// @Produce json
// @Param input body signInInput true "User credentials"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": token})
}
