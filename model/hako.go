package model

type HakoInfo struct {
	Address     string `json:"address"`
	Balance     int64  `json:"balance"`
	TotalSupply int64  `json:"total_supply"`
	Credit      int64  `json:"credit"`
	Debt        int64  `json:"debt"`
	MemberCount int64  `json:"member_count"`
	UpperLimit  int64  `json:"upper_limit"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    int64  `json:"decimals"`
}
