package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank-archive/services"
	"strconv"
)

// ReachedWordsNumGet 获得某日达标关键词的数量
func ReachedWordsNumGet(c *gin.Context) {
	date := c.Param("date")
	engine := c.Query("engine")
	goalRank, err := strconv.Atoi(c.Param("goal-rank"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "读取目标排名失败",
		})
		return
	}

	reachedWordsNum, err := services.ReachedWordsNum(goalRank, date, engine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "排名达标数获取失败",
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reachedWordsNum)
}
