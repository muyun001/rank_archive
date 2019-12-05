package services

import (
	"rank-archive/databases"
	"rank-archive/structs/models"
)

// ReachedWordsNum 获得某日达标关键词的数量
func ReachedWordsNum(goalRank int, date string, engine string) (int, error) {
	var reachedWordsNum int
	err := databases.Db.Model(models.HaveHistoryRank{}).
		Select("COUNT(DISTINCT (keyword_id))").
		Joins("left join keywords on have_history_ranks.keyword_id=keywords.id").
		Where(models.HaveHistoryRank{
			Date: date,
		}).
		Where(models.Keyword{
			Engine: engine,
		}).
		Where("have_history_ranks.top_rank <= ?", goalRank).
		Count(&reachedWordsNum).
		Error
	if err != nil {
		return 0, err
	}

	return reachedWordsNum, nil
}
