package services

import (
	"rank-archive/databases"
	"rank-archive/structs/models"
	"rank-archive/structs/responses"
)

// ReachedRanks 获取某日达标的数据
func ReachedRanks(date string, goalRank int) (responses.RankResultsResponse, error) {
	rankResultsResponse := make(responses.RankResultsResponse, 0)
	err := databases.Db.Model(models.Keyword{}).
		Joins("left join have_history_ranks on keywords.id = have_history_ranks.keyword_id").
		Select("keywords.word, min(have_history_ranks.top_rank) as rank").
		Where(models.HaveHistoryRank{
			Date: date,
		}).
		Where("top_rank <= ?", goalRank).
		Group("keywords.id").
		Scan(&rankResultsResponse).
		Error

	if err != nil {
		return rankResultsResponse, err
	}

	return rankResultsResponse, nil
}
