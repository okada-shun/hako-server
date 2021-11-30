package database

import (
	"hako-server/model"
)

// Create

func (d *GormDatabase) CreateTransferTokenHistory(transferTokenHistory model.TransferTokenHistory) error {
	/*
		INSERT INTO `transfer_token_histories` (`transfer_from`,`transfer_to`,`value`,`tx_hash`,`block_n`)
		VALUES ('0xE31c9fF6a8A1b952098CfeaF60c521cf68435503','0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9',
		30,'0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',77)
	*/
	return d.DB.Create(&transferTokenHistory).Error
}

func (d *GormDatabase) CreateTransferCreditHistory(transferCreditHistory model.TransferCreditHistory) error {
	/*
		INSERT INTO `transfer_credit_histories` (`transfer_from`,`transfer_to`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9','0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8',
		20,'0xbfa5afa9ccfabe79b0b21f2a9d778cf773b963e00307dd68c96abc3a30c5d3be',83)
	*/
	return d.DB.Create(&transferCreditHistory).Error
}

func (d *GormDatabase) CreateJoinHakoHistory(joinHakoHistory model.JoinHakoHistory) error {
	/*
		INSERT INTO `join_hako_histories` (`new_member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9',
		20,'0xb05efb51c3d5ae80070863c9742ead34baa1bb22d07005358c5eacc29dd99c77',80)
	*/
	return d.DB.Create(&joinHakoHistory).Error
}

func (d *GormDatabase) CreateLeaveHakoHistory(leaveHakoHistory model.LeaveHakoHistory) error {
	/*
		INSERT INTO `leave_hako_histories` (`member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x1b96b91E74c57302Cc420e8104Ee0A45811599d8',
		25,'0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0',89)
	*/
	return d.DB.Create(&leaveHakoHistory).Error
}

func (d *GormDatabase) CreateDepositTokenHistory(depositTokenHistory model.DepositTokenHistory) error {
	/*
		INSERT INTO `deposit_token_histories` (`member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9',
		10,'0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0',87)
	*/
	return d.DB.Create(&depositTokenHistory).Error
}

func (d *GormDatabase) CreateWithdrawTokenHistory(withdrawTokenHistory model.WithdrawTokenHistory) error {
	/*
		INSERT INTO `withdraw_token_histories` (`member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8',
		5,'0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0',88)
	*/
	return d.DB.Create(&withdrawTokenHistory).Error
}

func (d *GormDatabase) CreateRegisterBorrowingHistory(registerBorrowingHistory model.RegisterBorrowingHistory) error {
	/*

	 */
	return d.DB.Create(&registerBorrowingHistory).Error
}

func (d *GormDatabase) CreateLendCreditHistory(lendCreditHistory model.LendCreditHistory) error {
	/*

	 */
	return d.DB.Create(&lendCreditHistory).Error
}

func (d *GormDatabase) CreateCollectDebtFromHistory(collectDebtFromHistory model.CollectDebtFromHistory) error {
	/*

	 */
	return d.DB.Create(&collectDebtFromHistory).Error
}

func (d *GormDatabase) CreateReturnDebtToHistory(returnDebtToHistory model.ReturnDebtToHistory) error {
	/*

	 */
	return d.DB.Create(&returnDebtToHistory).Error
}

func (d *GormDatabase) CreateCreateCreditHistory(createCreditHistory model.CreateCreditHistory) error {
	/*

	 */
	return d.DB.Create(&createCreditHistory).Error
}

func (d *GormDatabase) CreateReduceDebtHistory(reduceDebtHistory model.ReduceDebtHistory) error {
	/*

	 */
	return d.DB.Create(&reduceDebtHistory).Error
}

func (d *GormDatabase) CreateChangeHakoOwnerHistory(changeHakoOwnerHistory model.ChangeHakoOwnerHistory) error {
	/*

	 */
	return d.DB.Create(&changeHakoOwnerHistory).Error
}

func (d *GormDatabase) CreateChangeUpperLimitHistory(changeUpperLimitHistory model.ChangeUpperLimitHistory) error {
	/*

	 */
	return d.DB.Create(&changeUpperLimitHistory).Error
}

func (d *GormDatabase) CreateGetRewardHistory(getRewardHistory model.GetRewardHistory) error {
	/*

	 */
	return d.DB.Create(&getRewardHistory).Error
}

// Get

func (d *GormDatabase) GetTransferTokenFromHistory(address string) ([]model.TransferTokenHistory, error) {
	var transferTokenHistories []model.TransferTokenHistory
	// SELECT * FROM `transfer_token_histories` WHERE transfer_from = '0xE31c9fF6a8A1b952098CfeaF60c521cf68435503'
	err := d.DB.Where("transfer_from = ?", address).Find(&transferTokenHistories).Error
	return transferTokenHistories, err
}

func (d *GormDatabase) GetTransferTokenToHistory(address string) ([]model.TransferTokenHistory, error) {
	var transferTokenHistories []model.TransferTokenHistory
	// SELECT * FROM `transfer_token_histories` WHERE transfer_to = '0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9'
	err := d.DB.Where("transfer_to = ?", address).Find(&transferTokenHistories).Error
	return transferTokenHistories, err
}

func (d *GormDatabase) GetTransferCreditFromHistory(address string) ([]model.TransferCreditHistory, error) {
	var transferCreditHistories []model.TransferCreditHistory
	// SELECT * FROM `transfer_credit_histories` WHERE transfer_from = '0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9'
	err := d.DB.Where("transfer_from = ?", address).Find(&transferCreditHistories).Error
	return transferCreditHistories, err
}

func (d *GormDatabase) GetTransferCreditToHistory(address string) ([]model.TransferCreditHistory, error) {
	var transferCreditHistories []model.TransferCreditHistory
	// SELECT * FROM `transfer_credit_histories` WHERE transfer_to = '0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8'
	err := d.DB.Where("transfer_to = ?", address).Find(&transferCreditHistories).Error
	return transferCreditHistories, err
}

func (d *GormDatabase) GetJoinHakoHistory(address string) ([]model.JoinHakoHistory, error) {
	var joinHakoHistories []model.JoinHakoHistory
	// SELECT * FROM `join_hako_histories` WHERE new_member = '0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8'
	err := d.DB.Where("new_member = ?", address).Find(&joinHakoHistories).Error
	return joinHakoHistories, err
}

func (d *GormDatabase) GetLeaveHakoHistory(address string) ([]model.LeaveHakoHistory, error) {
	var leaveHakoHistories []model.LeaveHakoHistory
	// SELECT * FROM `leave_hako_histories` WHERE member = '0x1b96b91E74c57302Cc420e8104Ee0A45811599d8'
	err := d.DB.Where("member = ?", address).Find(&leaveHakoHistories).Error
	return leaveHakoHistories, err
}

func (d *GormDatabase) GetDepositTokenHistory(address string) ([]model.DepositTokenHistory, error) {
	var depositTokenHistories []model.DepositTokenHistory
	// SELECT * FROM `deposit_token_histories` WHERE member = '0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9'
	err := d.DB.Where("member = ?", address).Find(&depositTokenHistories).Error
	return depositTokenHistories, err
}

func (d *GormDatabase) GetWithdrawTokenHistory(address string) ([]model.WithdrawTokenHistory, error) {
	var withdrawTokenHistories []model.WithdrawTokenHistory
	// SELECT * FROM `withdraw_token_histories` WHERE member = '0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8'
	err := d.DB.Where("member = ?", address).Find(&withdrawTokenHistories).Error
	return withdrawTokenHistories, err
}

func (d *GormDatabase) GetRegisterBorrowingHistory(address string) ([]model.RegisterBorrowingHistory, error) {
	var registerBorrowingHistories []model.RegisterBorrowingHistory
	//
	err := d.DB.Where("member = ?", address).Find(&registerBorrowingHistories).Error
	return registerBorrowingHistories, err
}

func (d *GormDatabase) GetLendCreditFromHistory(address string) ([]model.LendCreditHistory, error) {
	var lendCreditHistories []model.LendCreditHistory
	//
	err := d.DB.Where("lend_from = ?", address).Find(&lendCreditHistories).Error
	return lendCreditHistories, err
}

func (d *GormDatabase) GetLendCreditToHistory(address string) ([]model.LendCreditHistory, error) {
	var lendCreditHistories []model.LendCreditHistory
	//
	err := d.DB.Where("lend_to = ?", address).Find(&lendCreditHistories).Error
	return lendCreditHistories, err
}

func (d *GormDatabase) GetCollectDebtFromCreditorHistory(address string) ([]model.CollectDebtFromHistory, error) {
	var collectDebtFromHistories []model.CollectDebtFromHistory
	//
	err := d.DB.Where("creditor = ?", address).Find(&collectDebtFromHistories).Error
	return collectDebtFromHistories, err
}

func (d *GormDatabase) GetCollectDebtFromDebtorHistory(address string) ([]model.CollectDebtFromHistory, error) {
	var collectDebtFromHistories []model.CollectDebtFromHistory
	//
	err := d.DB.Where("debtor = ?", address).Find(&collectDebtFromHistories).Error
	return collectDebtFromHistories, err
}

func (d *GormDatabase) GetReturnDebtToCreditorHistory(address string) ([]model.ReturnDebtToHistory, error) {
	var returnDebtToHistories []model.ReturnDebtToHistory
	//
	err := d.DB.Where("creditor = ?", address).Find(&returnDebtToHistories).Error
	return returnDebtToHistories, err
}

func (d *GormDatabase) GetReturnDebtToDebtorHistory(address string) ([]model.ReturnDebtToHistory, error) {
	var returnDebtToHistories []model.ReturnDebtToHistory
	//
	err := d.DB.Where("debtor = ?", address).Find(&returnDebtToHistories).Error
	return returnDebtToHistories, err
}

func (d *GormDatabase) GetCreateCreditHistory(address string) ([]model.CreateCreditHistory, error) {
	var createCreditHistories []model.CreateCreditHistory
	//
	err := d.DB.Where("member = ?", address).Find(&createCreditHistories).Error
	return createCreditHistories, err
}

func (d *GormDatabase) GetReduceDebtHistory(address string) ([]model.ReduceDebtHistory, error) {
	var reduceDebtHistories []model.ReduceDebtHistory
	//
	err := d.DB.Where("member = ?", address).Find(&reduceDebtHistories).Error
	return reduceDebtHistories, err
}

func (d *GormDatabase) GetChangeHakoOwnerHistory() ([]model.ChangeHakoOwnerHistory, error) {
	var changeHakoOwnerHistories []model.ChangeHakoOwnerHistory
	//
	err := d.DB.Find(&changeHakoOwnerHistories).Error
	return changeHakoOwnerHistories, err
}

func (d *GormDatabase) GetChangeUpperLimitHistory() ([]model.ChangeUpperLimitHistory, error) {
	var changeUpperLimitHistories []model.ChangeUpperLimitHistory
	//
	err := d.DB.Find(&changeUpperLimitHistories).Error
	return changeUpperLimitHistories, err
}

func (d *GormDatabase) GetGetRewardHistory() ([]model.GetRewardHistory, error) {
	var getRewardHistories []model.GetRewardHistory
	//
	err := d.DB.Find(&getRewardHistories).Error
	return getRewardHistories, err
}
