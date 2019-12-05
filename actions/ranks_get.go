package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank-archive/databases"
	"rank-archive/structs/models"
	"rank-archive/structs/responses"
)

func RanksGet(c *gin.Context) {
	date := c.Query("date")
	word := c.Query("word")
	engine := c.Query("engine")
	checkMatch := c.Query("check_match")

	keywordId := 0
	if word != "" || engine != "" {
		if word == "" || engine == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请同时指定关键词和引擎",
			})
			return
		}
		keyword := models.Keyword{}
		if databases.Db.Where(&models.Keyword{
			Word:       word,
			Engine:     engine,
			CheckMatch: checkMatch,
		}).First(&keyword).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "关键词未找到排名记录",
			})
			return
		}
		keywordId = keyword.ID
	}

	ranks := make([]models.HaveHistoryRank, 0)
	databases.Db.Model(&models.HaveHistoryRank{}).
		Preload("Keyword").
		Where(&models.HaveHistoryRank{
			KeywordId: keywordId,
			Date:      date,
		}).Find(&ranks)

	historyRanks := make([]responses.HistoryRank, 0)
	for _, rank := range ranks {
		historyRank := responses.HistoryRank{
			Keyword:    rank.Keyword.Word,
			Engine:     rank.Keyword.Engine,
			CheckMatch: rank.Keyword.CheckMatch,
			TopRank:    rank.TopRank,
			Ranks:      rank.Ranks,
			Date:       rank.Date,
			CaptureUrl: rank.CaptureUrl,
		}
		historyRanks = append(historyRanks, historyRank)
	}

	c.JSON(http.StatusOK, historyRanks)
}
