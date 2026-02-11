package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kacperfilipiuk/zguba-gov/internal/model"
	"github.com/kacperfilipiuk/zguba-gov/internal/repository"
)

type APIHandler struct {
	repo *repository.FoundItemRepo
}

func NewAPIHandler(repo *repository.FoundItemRepo) *APIHandler {
	return &APIHandler{repo: repo}
}

func (h *APIHandler) ListItems(c *gin.Context) {
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if limit < 1 {
		limit = 1
	}
	if limit > 100 {
		limit = 100
	}
	if skip < 0 {
		skip = 0
	}

	params := model.ListParams{
		Skip:         skip,
		Limit:        limit,
		Category:     c.Query("category"),
		Municipality: c.Query("municipality"),
		Status:       c.Query("status"),
		Search:       c.Query("search"),
	}

	items, err := h.repo.List(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]model.FoundItemResponse, 0, len(items))
	for _, item := range items {
		resp = append(resp, item.ToResponse())
	}

	c.JSON(http.StatusOK, resp)
}

func (h *APIHandler) CreateItem(c *gin.Context) {
	var create model.FoundItemCreate
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.repo.Create(create)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := item.ToResponse()
	c.JSON(http.StatusCreated, resp)
}

func (h *APIHandler) GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item.ToResponse())
}

func (h *APIHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var update model.FoundItemUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.repo.Update(id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item.ToResponse())
}

func (h *APIHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.Delete(id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *APIHandler) CategoriesList(c *gin.Context) {
	cats, err := h.repo.Categories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cats == nil {
		cats = []map[string]string{}
	}
	c.JSON(http.StatusOK, cats)
}

func (h *APIHandler) Stats(c *gin.Context) {
	stats, err := h.repo.Stats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *APIHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
