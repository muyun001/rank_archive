package services

import (
	"rank-archive/databases"
	"rank-archive/structs/models"
	"rank-archive/structs/requests"
	"time"
)

// 批量保存历史排名数据
func HistoryRanksSave(historyRanks []requests.HistoryRank) error {
	for _, historyRank := range historyRanks {
		rankArchiveKeyword := models.Keyword{
			Word:       historyRank.Keyword,
			Engine:     historyRank.Engine,
			CheckMatch: historyRank.CheckMatch,
		}
		err := databases.Db.Where(rankArchiveKeyword).FirstOrInit(&rankArchiveKeyword).Error
		if err != nil {
			return err
		}
		if rankArchiveKeyword.ID == 0 {
			rankArchiveKeyword.CreatedAt = time.Now()
			databases.Db.Save(&rankArchiveKeyword)
		}

		if historyRank.TopRank != 0 {
			rankArchiveRanks := models.HaveHistoryRank{
				KeywordId:  rankArchiveKeyword.ID,
				TopRank:    historyRank.TopRank,
				Ranks:      historyRank.Ranks,
				Date:       historyRank.Date,
				CaptureUrl: historyRank.CaptureUrl,
				Ip:         historyRank.Ip,
				CreatedAt:  time.Now(),
			}
			err = databases.Db.Save(&rankArchiveRanks).Error
			if err != nil {
				return err
			}
		} else {
			noHistoryRank := &models.NoHistoryRank{}
			notFound := databases.Db.
				Where("keyword_id = ?", rankArchiveKeyword.ID).
				Where("end_date <= ?", historyRank.Date).
				Order("id desc").
				First(&noHistoryRank).
				RecordNotFound()

			if notFound || !isTodayOrYesterdayOf(noHistoryRank.EndDate, historyRank.Date) {
				rankArchiveRanks := models.NoHistoryRank{
					KeywordId: rankArchiveKeyword.ID,
					StartDate: historyRank.Date,
					EndDate:   historyRank.Date,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				err = databases.Db.Save(&rankArchiveRanks).Error
				if err != nil {
					return err
				}
			} else {
				databases.Db.Model(noHistoryRank).Where("id = ?", noHistoryRank.ID).Updates(models.NoHistoryRank{EndDate: historyRank.Date, UpdatedAt: time.Now()})
			}
		}
	}

	return nil
}

func isTodayOrYesterdayOf(dateStr string, baseDateStr string) bool {
	isToday := dateStr[0:10] == baseDateStr
	baseTime, err := time.ParseInLocation("2006-01-02", baseDateStr, time.Now().Location())
	if err != nil {
		return false
	}
	yesterdayStr := baseTime.Add(-time.Hour * 24).Format("2006-01-02")
	isYesterday := dateStr[0:10] == yesterdayStr
	return isToday || isYesterday
}
