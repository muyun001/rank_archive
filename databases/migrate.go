package databases

import "rank-archive/structs/models"

func AutoMigrate() {
	Db.AutoMigrate(&models.Keyword{}, &models.HaveHistoryRank{}, &models.NoHistoryRank{})
}
