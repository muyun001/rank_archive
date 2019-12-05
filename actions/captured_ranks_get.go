package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank-archive/databases"
	"rank-archive/structs/models"
	"rank-archive/structs/responses"
)

func CapturedRanksGet(c *gin.Context) {
	date := c.Query("date")
	word := c.Query("word")
	engine := c.Query("engine")
	checkMatch := c.Query("check_match")

	keywordId := 0
	if word != "" || engine != "" {
		keyword := models.Keyword{}
		keywordFound := databases.Db.Where(&models.Keyword{
			Word:       word,
			Engine:     engine,
			CheckMatch: checkMatch,
		}).First(&keyword).RecordNotFound() == false
		if keywordFound {
			keywordId = keyword.ID
		}
	}

	capturedRanks := make([]models.HaveHistoryRank, 0)
	databases.Db.Model(&models.HaveHistoryRank{}).
		Where("capture_url <> ''").
		Preload("Keyword").
		Where(&models.HaveHistoryRank{
			KeywordId: keywordId,
			Date:      date,
		}).Find(&capturedRanks)

	historyRanks := make([]responses.HistoryRank, 0)
	for _, capturedRank := range capturedRanks {
		historyRank := responses.HistoryRank{
			Keyword:    capturedRank.Keyword.Word,
			Engine:     capturedRank.Keyword.Engine,
			CheckMatch: capturedRank.Keyword.CheckMatch,
			TopRank:    capturedRank.TopRank,
			Ranks:      capturedRank.Ranks,
			Date:       capturedRank.Date,
			CaptureUrl: capturedRank.CaptureUrl,
		}
		historyRanks = append(historyRanks, historyRank)
	}

	c.JSON(http.StatusOK, historyRanks)
}
