package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kibo13/todo-app/internal/entity"
	"net/http"
	"strconv"
)

// @Summary Create a new todo item
// @Security ApiKeyAuth
// @Tags Items
// @Description Create a new todo item in a specific list
// @ID create-item
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Param input body entity.TodoItem true "Todo item info"
// @Success 200 {integer} integer "Item ID"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input entity.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all items in a todo list
// @Security ApiKeyAuth
// @Tags Items
// @Description Get all todo items for a specific list
// @ID get-all-items
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {array} entity.TodoItem "List of Todo items"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "List not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get todo item by ID
// @Security ApiKeyAuth
// @Tags Items
// @Description Get a specific todo item by its ID
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} entity.TodoItem "Todo item"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update a todo item
// @Security ApiKeyAuth
// @Tags Items
// @Description Update a specific todo item
// @ID update-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param input body entity.UpdateItemInput true "Updated item info"
// @Success 200 {object} statusResponse "Status OK"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

// @Summary Delete a todo item
// @Security ApiKeyAuth
// @Tags Items
// @Description Delete a specific todo item by its ID
// @ID delete-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse "Status OK"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Failure default {object} errorResponse "Unexpected error"
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
