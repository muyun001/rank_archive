package services

import (
	"rank-archive/common/stringse"
	"rank-archive/databases"
	"rank-archive/structs/models"
)

// ReachedRanksDiffWords 获取2日达标关键词的差集
func ReachedRanksDiffWords(goalRank int, firstDay string, secondDay string, engine string, checkMatch string) ([]string, error) {
	firstDayKeywords, err := QueryKeywords(goalRank, firstDay, engine, checkMatch)
	if err != nil {
		return []string{}, err
	}

	secondDayKeywords, err := QueryKeywords(goalRank, secondDay, engine, checkMatch)
	if err != nil {
		return []string{}, err
	}

	differenceKeywordsSet := stringse.Diff(firstDayKeywords, secondDayKeywords)
	return differenceKeywordsSet, nil
}

// queryKeywords 查询关键词
func QueryKeywords(goalRank int, day string, engine string, checkMatch string) ([]string, error) {
	var keywords []string

	err := databases.Db.Model(models.Keyword{}).
		Joins("left join have_history_ranks on keywords.id = have_history_ranks.keyword_id").
		Where(models.HaveHistoryRank{
			Date: day,
		}).
		Where(models.Keyword{
			Engine:     engine,
			CheckMatch: checkMatch,
		}).
		Where("top_rank <= ?", goalRank).
		Pluck("word", &keywords).
		Error

	if err != nil {
		return []string{}, err
	}

	return keywords, nil
}
