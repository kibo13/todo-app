package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kibo13/todo-app/internal/entity"
	"net/http"
	"strconv"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags Lists
// @Description Create a new todo list for the user
// @ID create-list
// @Accept json
// @Produce json
// @Param input body entity.TodoList true "List info"
// @Success 200 {integer} integer "List ID"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input entity.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []entity.TodoList `json:"data"`
}

// @Summary Get all lists
// @Security ApiKeyAuth
// @Tags Lists
// @Description Get all todo lists for a specific user
// @ID get-all-lists
// @Accept json
// @Produce json
// @Success 200 {array} entity.TodoList "List of Todo lists"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get todo list by ID
// @Security ApiKeyAuth
// @Tags Lists
// @Description Get a specific todo list by its ID
// @ID get-list-by-id
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} entity.TodoList "Todo list"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update todo list
// @Security ApiKeyAuth
// @Tags Lists
// @Description Update an existing todo list
// @ID update-list
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Param input body entity.UpdateListInput true "Updated list information"
// @Success 200 {object} statusResponse "Status response"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete todo list
// @Security ApiKeyAuth
// @Tags Lists
// @Description Delete a todo list by ID
// @ID delete-list
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} statusResponse "Status response"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
