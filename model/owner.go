package model

type OwnerInfo struct {
	HakoAddress  string `json:"hako_address"`
	OwnerAddress string `json:"owner_address"`
	HakoBalance  int64  `json:"hako_balance"`
	TotalSupply  int64  `json:"total_supply"`
	HakoCredit   int64  `json:"hako_credit"`
	HakoDebt     int64  `json:"hako_debt"`
	OwnerBalance int64  `json:"owner_balance"`
	LendingCount int64  `json:"lending_count"`
	MemberCount  int64  `json:"member_count"`
	UpperLimit   int64  `json:"upper_limit"`
}
