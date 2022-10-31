package handler

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary AddArticle
// @Security ApiKeyAuth
// @Tags article
// @Description adds article
// @ID addArticle
// @Accept json
// @Produce json
// @Param input body entities.Article true "ArticleId"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/articles/ [post]
func (h *Handler) AddArticle(c *gin.Context) {
	var input entities.Article

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Articles.CreateArticle(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary GetArticle
// @Security ApiKeyAuth
// @Tags article
// @Description gets article
// @ID getArticle
// @Produce json
// @Accept json
// @Success 200 {object} entities.Article
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Param id path int true "ArticleId"
// @Router /api/articles/{id} [get]
func (h *Handler) GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	article, err := h.services.Articles.GetArticleById(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, article)
}

type getAllArticlesResponse struct {
	Data []entities.Article `json:"data"`
}

// @Summary GetAllArticles
// @Security ApiKeyAuth
// @Tags article
// @Description gets articles
// @ID getAllArticles
// @Produce json
// @Accept json
// @Success 200 {object} getAllArticlesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Param id path int true "ArticleId"
// @Router /api/articles/ [get]
func (h *Handler) GetAllArticles(c *gin.Context) {
	articles, err := h.services.Articles.GetAllArticles()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllArticlesResponse{
		Data: articles,
	})
}
