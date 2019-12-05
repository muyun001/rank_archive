package routers

import (
	"github.com/gin-gonic/gin"
	"rank-archive/actions"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.POST("/history-ranks", actions.HistoryRanksPost)
	r.GET("/ranks", actions.RanksGet)
	r.GET("/captured-ranks", actions.CapturedRanksGet)
	r.GET("/reached-num/:goal-rank/:date", actions.ReachedWordsNumGet)
	r.GET("/reached-diff-words/:goal-rank/:first-day/:second-day", actions.DiffWordsGet)
	r.GET("/reached-ranks/:goal-rank/:date", actions.ReachedRanksGet)

	return r
}

func init() {
	r = gin.Default()
}