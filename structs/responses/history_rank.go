package responses

// 用于输出排名数据的格式
type HistoryRank struct {
	Keyword    string `json:"keyword"`
	Engine     string `json:"engine"`
	CheckMatch string `json:"check_match"`
	TopRank    int    `json:"top_rank"`
	Ranks      string `json:"ranks"`
	Date       string `json:"date"`
	CaptureUrl string `json:"capture_url"`
}

// 用于输出达标数据的格式
type RankResult struct {
	Word string `json:"word"`
	Rank int    `json:"rank"`
}

// 获取达标排名结果的回复
type RankResultsResponse []RankResult
