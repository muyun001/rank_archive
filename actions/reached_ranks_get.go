package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank-archive/services"
	"strconv"
)

// ReachedRanksGet 获取某日达标的数据
func ReachedRanksGet(c *gin.Context) {
	date := c.Param("date")
	goalRank, err := strconv.Atoi(c.Param("goal-rank"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "读取目标排名失败",
		})
		return
	}

	rankResultsResponse, err := services.ReachedRanks(date, goalRank)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "达标排名获取失败",
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rankResultsResponse)
}
