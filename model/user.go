package model

type UserInfo struct {
	Address   string `json:"address"`
	Balance   int64  `json:"balance"`
	Member    int    `json:"member"`
	Credit    int64  `json:"credit"`
	Debt      int64  `json:"debt"`
	Lending   int64  `json:"lending"`
	Borrowing int64  `json:"borrowing"`
	NetAssets int64  `json:"net_assets"`
	Value     int64  `json:"value"`
	Duration  int64  `json:"duration"`
}
