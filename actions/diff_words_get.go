package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank-archive/services"
	"strconv"
)

// DiffWordsGet 获取2日达标差集
func DiffWordsGet(c *gin.Context) {
	firstDay := c.Param("first-day")
	secondDay := c.Param("second-day")
	engine := c.Query("engine")
	checkMatch := c.Query("check-match")
	goalRank, err := strconv.Atoi(c.Param("goal-rank"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "读取目标排名失败",
		})
		return
	}

	keywords, err := services.ReachedRanksDiffWords(goalRank, firstDay, secondDay, engine, checkMatch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "排名差集获取失败",
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, keywords)
}
