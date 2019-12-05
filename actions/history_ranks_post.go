package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank-archive/services"
	"rank-archive/structs/requests"
)

// 批量接收历史排名数据接口
func HistoryRanksPost(c *gin.Context) {
	historyRanks := make([]requests.HistoryRank, 0)
	err := c.BindJSON(&historyRanks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传入格式错误",
		})
		return
	}
	err = services.HistoryRanksSave(historyRanks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "保存出错",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "保存成功",
	})
}
