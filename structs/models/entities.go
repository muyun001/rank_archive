package models

import "time"

type Keyword struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Word       string    `gorm:"type:varchar(255);not null;unique_index:word_engine" json:"word"`
	Engine     string    `gorm:"type:varchar(255);not null;unique_index:word_engine" json:"engine"`
	CheckMatch string    `gorm:"type:varchar(255);not null;unique_index:word_engine" json:"check_match"`
	CreatedAt  time.Time `gorm:"type:datetime" json:"created_at"`
}

type HaveHistoryRank struct {
	ID         int       `gorm:"primary_key" json:"id"`
	KeywordId  int       `gorm:"type:int;not null;index:keyword_date" json:"keyword_id"`
	Ranks      string    `gorm:"type:varchar(64);" json:"ranks"`
	Date       string    `gorm:"type:date;not null;index:keyword_date,date_top_rank" json:"date"`
	TopRank    int       `gorm:"type:int;index:date_top_rank" json:"top_rank"`
	CaptureUrl string    `gorm:"type:varchar(255);" json:"capture_url"`
	Ip         string    `gorm:"type:varchar(32)" json:"ip"`
	CreatedAt  time.Time `gorm:"type:datetime" json:"created_at"`
	Keyword    Keyword
}

type NoHistoryRank struct {
	ID        int       `gorm:"primary_key" json:"id"`
	KeywordId int       `gorm:"type:int;not null;index:keyword_start_date" json:"keyword_id"`
	StartDate string    `gorm:"type:date;not null;index:keyword_start_date" json:"start_date"`
	EndDate   string    `gorm:"type:date;not null" json:"end_date"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}
