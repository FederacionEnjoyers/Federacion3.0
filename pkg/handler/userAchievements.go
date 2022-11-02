package handler

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Entity that allows return slice of userAchievements

type getAllUserAchievementsResponse struct {
	Data []entities.UserAchievement `json:"data"`
}

// Returns User Achievement By Token
// If user is not exists return http.BadRequest

func (h *Handler) GetUserAchievementsByToken(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userAchievements, err := h.services.Achievement.GetUserAchievements(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUserAchievementsResponse{
		Data: userAchievements,
	})
}

// Returns achievements by user Id

func (h *Handler) GetUserAchievementsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	userAchievements, err := h.services.Achievement.GetUserAchievements(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUserAchievementsResponse{
		Data: userAchievements,
	})
}
