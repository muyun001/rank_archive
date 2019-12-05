package requests

// 用于接收排名数据的格式
type HistoryRank struct {
	Ip         string `json:"ip"`
	Keyword    string `json:"keyword"`
	Engine     string `json:"engine"`
	CheckMatch string `json:"check_match"`
	TopRank    int    `json:"top_rank"`
	Ranks      string `json:"ranks"`
	Date       string `json:"date"`
	CaptureUrl string `json:"capture_url"`
}
