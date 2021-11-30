package eventlogs

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"hako-server/config"
	"hako-server/database"
	"hako-server/hkfinance"
	"hako-server/model"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EventDatabase struct {
	DB *database.GormDatabase
}

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type LogTransferCredit struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type LogJoinHako struct {
	NewMember common.Address
	Value     *big.Int
}

type LogLeaveHako struct {
	Member common.Address
	Value  *big.Int
}

type LogDepositToken struct {
	Member common.Address
	Value  *big.Int
}

type LogWithdrawToken struct {
	Member common.Address
	Value  *big.Int
}

type LogRegisterBorrowing struct {
	Member   common.Address
	Value    *big.Int
	Duration *big.Int
}

type LogLendCredit struct {
	From     common.Address
	To       common.Address
	Value    *big.Int
	Duration *big.Int
	ID       *big.Int
	Time     *big.Int
}

type LogCollectDebtFrom struct {
	Creditor common.Address
	Debtor   common.Address
	ID       *big.Int
}

type LogReturnDebtTo struct {
	Debtor   common.Address
	Creditor common.Address
	ID       *big.Int
}

type LogCreateCredit struct {
	Member common.Address
	Value  *big.Int
}

type LogReduceDebt struct {
	Member common.Address
	Value  *big.Int
}

type LogChangeHakoOwner struct {
	OldHakoOwner common.Address
	NewHakoOwner common.Address
}

type LogChangeUpperLimit struct {
	HakoOwner     common.Address
	NewUpperLimit *big.Int
}

type LogGetReward struct {
	HakoOwner   common.Address
	RewardValue *big.Int
}

func (d *EventDatabase) ReadEventLogs(config *config.Config) error {
	client, err := hkfinance.GetEthclient(config)
	if err != nil {
		return err
	}
	contractAddressBytes, err := ioutil.ReadFile(config.Ethereum.ContractAddress)
	if err != nil {
		return err
	}
	contractAddress := common.HexToAddress(string(contractAddressBytes))
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(200),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(hkfinance.HkfinanceABI)))
	if err != nil {
		return err
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	logTransferCreditSig := []byte("TransferCredit(address,address,uint256)")
	logTransferCreditSigHash := crypto.Keccak256Hash(logTransferCreditSig)

	logJoinHakoSig := []byte("JoinHako(address,uint256)")
	logJoinHakoSigHash := crypto.Keccak256Hash(logJoinHakoSig)

	logLeaveHakoSig := []byte("LeaveHako(address,uint256)")
	logLeaveHakoSigHash := crypto.Keccak256Hash(logLeaveHakoSig)

	logDepositTokenSig := []byte("DepositToken(address,uint256)")
	logDepositTokenSigHash := crypto.Keccak256Hash(logDepositTokenSig)

	logWithdrawTokenSig := []byte("WithdrawToken(address,uint256)")
	logWithdrawTokenSigHash := crypto.Keccak256Hash(logWithdrawTokenSig)

	logRegisterBorrowingSig := []byte("RegisterBorrowing(address,uint256,uint256)")
	logRegisterBorrowingSigHash := crypto.Keccak256Hash(logRegisterBorrowingSig)

	logLendCreditSig := []byte("LendCredit(address,address,uint256,uint256,uint256,uint256)")
	logLendCreditSigHash := crypto.Keccak256Hash(logLendCreditSig)

	logCollectDebtFromSig := []byte("CollectDebtFrom(address,address,uint256)")
	logCollectDebtFromSigHash := crypto.Keccak256Hash(logCollectDebtFromSig)

	logReturnDebtToSig := []byte("ReturnDebtTo(address,address,uint256)")
	logReturnDebtToSigHash := crypto.Keccak256Hash(logReturnDebtToSig)

	logCreateCreditSig := []byte("CreateCredit(address,uint256)")
	logCreateCreditSigHash := crypto.Keccak256Hash(logCreateCreditSig)

	logReduceDebtSig := []byte("ReduceDebt(address,uint256)")
	logReduceDebtSigHash := crypto.Keccak256Hash(logReduceDebtSig)

	logChangeHakoOwnerSig := []byte("ChangeHakoOwner(address,address)")
	logChangeHakoOwnerSigHash := crypto.Keccak256Hash(logChangeHakoOwnerSig)

	logChangeUpperLimitSig := []byte("ChangeUpperLimit(address,uint256)")
	logChangeUpperLimitSigHash := crypto.Keccak256Hash(logChangeUpperLimitSig)

	logGetRewardSig := []byte("GetReward(address,uint256)")
	logGetRewardSigHash := crypto.Keccak256Hash(logGetRewardSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {

		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				return err
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			d.DB.CreateTransferTokenHistory(model.TransferTokenHistory{
				TransferFrom: transferEvent.From.Hex(),
				TransferTo:   transferEvent.To.Hex(),
				Value:        transferEvent.Value.Int64(),
				TxHash:       logTransferSigHash.Hex(),
				BlockN:       vLog.BlockNumber,
			})

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Value: %s\n", transferEvent.Value.String())

		case logTransferCreditSigHash.Hex():
			fmt.Printf("Log Name: Transfer Credit\n")

			var transferCreditEvent LogTransferCredit

			err := contractAbi.UnpackIntoInterface(&transferCreditEvent, "TransferCredit", vLog.Data)
			if err != nil {
				return err
			}

			transferCreditEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferCreditEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			d.DB.CreateTransferCreditHistory(model.TransferCreditHistory{
				TransferFrom: transferCreditEvent.From.Hex(),
				TransferTo:   transferCreditEvent.To.Hex(),
				Value:        transferCreditEvent.Value.Int64(),
				TxHash:       logTransferCreditSigHash.Hex(),
				BlockN:       vLog.BlockNumber,
			})

			fmt.Printf("From: %s\n", transferCreditEvent.From.Hex())
			fmt.Printf("To: %s\n", transferCreditEvent.To.Hex())
			fmt.Printf("Value: %s\n", transferCreditEvent.Value.String())

		case logJoinHakoSigHash.Hex():
			fmt.Printf("Log Name: Join Hako\n")

			var joinHakoEvent LogJoinHako

			err := contractAbi.UnpackIntoInterface(&joinHakoEvent, "JoinHako", vLog.Data)
			if err != nil {
				return err
			}

			joinHakoEvent.NewMember = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateJoinHakoHistory(model.JoinHakoHistory{
				NewMember: joinHakoEvent.NewMember.Hex(),
				Value:     joinHakoEvent.Value.Int64(),
				TxHash:    logJoinHakoSigHash.Hex(),
				BlockN:    vLog.BlockNumber,
			})

			fmt.Printf("New Member: %s\n", joinHakoEvent.NewMember.Hex())
			fmt.Printf("Value: %s\n", joinHakoEvent.Value.String())

		case logLeaveHakoSigHash.Hex():
			fmt.Printf("Log Name: Leave Hako\n")

			var leaveHakoEvent LogLeaveHako

			err := contractAbi.UnpackIntoInterface(&leaveHakoEvent, "LeaveHako", vLog.Data)
			if err != nil {
				return err
			}

			leaveHakoEvent.Member = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateLeaveHakoHistory(model.LeaveHakoHistory{
				Member: leaveHakoEvent.Member.Hex(),
				Value:  leaveHakoEvent.Value.Int64(),
				TxHash: logLeaveHakoSigHash.Hex(),
				BlockN: vLog.BlockNumber,
			})

			fmt.Printf("Member: %s\n", leaveHakoEvent.Member.Hex())
			fmt.Printf("Value: %s\n", leaveHakoEvent.Value.String())

		case logDepositTokenSigHash.Hex():
			fmt.Printf("Log Name: Deposit Token\n")

			var depositTokenEvent LogDepositToken

			err := contractAbi.UnpackIntoInterface(&depositTokenEvent, "DepositToken", vLog.Data)
			if err != nil {
				return err
			}

			depositTokenEvent.Member = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateDepositTokenHistory(model.DepositTokenHistory{
				Member: depositTokenEvent.Member.Hex(),
				Value:  depositTokenEvent.Value.Int64(),
				TxHash: logDepositTokenSigHash.Hex(),
				BlockN: vLog.BlockNumber,
			})

			fmt.Printf("Member: %s\n", depositTokenEvent.Member.Hex())
			fmt.Printf("Value: %s\n", depositTokenEvent.Value.String())

		case logWithdrawTokenSigHash.Hex():
			fmt.Printf("Log Name: Withdraw Token\n")

			var withdrawTokenEvent LogWithdrawToken

			err := contractAbi.UnpackIntoInterface(&withdrawTokenEvent, "WithdrawToken", vLog.Data)
			if err != nil {
				return err
			}

			withdrawTokenEvent.Member = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateWithdrawTokenHistory(model.WithdrawTokenHistory{
				Member: withdrawTokenEvent.Member.Hex(),
				Value:  withdrawTokenEvent.Value.Int64(),
				TxHash: logWithdrawTokenSigHash.Hex(),
				BlockN: vLog.BlockNumber,
			})

			fmt.Printf("Member: %s\n", withdrawTokenEvent.Member.Hex())
			fmt.Printf("Value: %s\n", withdrawTokenEvent.Value.String())

		case logRegisterBorrowingSigHash.Hex():
			fmt.Printf("Log Name: Register Borrowing\n")

			var registerBorrowingEvent LogRegisterBorrowing

			err := contractAbi.UnpackIntoInterface(&registerBorrowingEvent, "RegisterBorrowing", vLog.Data)
			if err != nil {
				return err
			}

			registerBorrowingEvent.Member = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateRegisterBorrowingHistory(model.RegisterBorrowingHistory{
				Member:   registerBorrowingEvent.Member.Hex(),
				Value:    registerBorrowingEvent.Value.Int64(),
				Duration: registerBorrowingEvent.Duration.Int64(),
				TxHash:   logRegisterBorrowingSigHash.Hex(),
				BlockN:   vLog.BlockNumber,
			})

			fmt.Printf("Member: %s\n", registerBorrowingEvent.Member.Hex())
			fmt.Printf("Value: %s\n", registerBorrowingEvent.Value.String())
			fmt.Printf("Duration: %s\n", registerBorrowingEvent.Duration.String())

		case logLendCreditSigHash.Hex():
			fmt.Printf("Log Name: Lend Credit\n")

			var lendCreditEvent LogLendCredit

			err := contractAbi.UnpackIntoInterface(&lendCreditEvent, "LendCredit", vLog.Data)
			if err != nil {
				return err
			}

			lendCreditEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			lendCreditEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			d.DB.CreateLendCreditHistory(model.LendCreditHistory{
				LendFrom:  lendCreditEvent.From.Hex(),
				LendTo:    lendCreditEvent.To.Hex(),
				Value:     lendCreditEvent.Value.Int64(),
				Duration:  lendCreditEvent.Duration.Int64(),
				LendingID: lendCreditEvent.ID.Int64(),
				Time:      lendCreditEvent.Time.Int64(),
				TxHash:    logLendCreditSigHash.Hex(),
				BlockN:    vLog.BlockNumber,
			})

			fmt.Printf("From: %s\n", lendCreditEvent.From.Hex())
			fmt.Printf("To: %s\n", lendCreditEvent.To.Hex())
			fmt.Printf("Value: %s\n", lendCreditEvent.Value.String())
			fmt.Printf("Duration: %s\n", lendCreditEvent.Duration.String())
			fmt.Printf("ID: %s\n", lendCreditEvent.ID.String())
			fmt.Printf("Time: %s\n", lendCreditEvent.Time.String())

		case logCollectDebtFromSigHash.Hex():
			fmt.Printf("Log Name: Collect Debt From\n")

			var collectDebtFromEvent LogCollectDebtFrom

			err := contractAbi.UnpackIntoInterface(&collectDebtFromEvent, "CollectDebtFrom", vLog.Data)
			if err != nil {
				return err
			}

			collectDebtFromEvent.Creditor = common.HexToAddress(vLog.Topics[1].Hex())
			collectDebtFromEvent.Debtor = common.HexToAddress(vLog.Topics[2].Hex())

			d.DB.CreateCollectDebtFromHistory(model.CollectDebtFromHistory{
				Creditor:  collectDebtFromEvent.Creditor.Hex(),
				Debtor:    collectDebtFromEvent.Debtor.Hex(),
				LendingID: collectDebtFromEvent.ID.Int64(),
				TxHash:    logCollectDebtFromSigHash.Hex(),
				BlockN:    vLog.BlockNumber,
			})

			fmt.Printf("Creditor: %s\n", collectDebtFromEvent.Creditor.Hex())
			fmt.Printf("Debtor: %s\n", collectDebtFromEvent.Debtor.Hex())
			fmt.Printf("ID: %s\n", collectDebtFromEvent.ID.String())

		case logReturnDebtToSigHash.Hex():
			fmt.Printf("Log Name: Return Debt To\n")

			var returnDebtToEvent LogReturnDebtTo

			err := contractAbi.UnpackIntoInterface(&returnDebtToEvent, "ReturnDebtTo", vLog.Data)
			if err != nil {
				return err
			}

			returnDebtToEvent.Debtor = common.HexToAddress(vLog.Topics[1].Hex())
			returnDebtToEvent.Creditor = common.HexToAddress(vLog.Topics[2].Hex())

			d.DB.CreateReturnDebtToHistory(model.ReturnDebtToHistory{
				Debtor:    returnDebtToEvent.Debtor.Hex(),
				Creditor:  returnDebtToEvent.Creditor.Hex(),
				LendingID: returnDebtToEvent.ID.Int64(),
				TxHash:    logReturnDebtToSigHash.Hex(),
				BlockN:    vLog.BlockNumber,
			})

			fmt.Printf("Debtor: %s\n", returnDebtToEvent.Debtor.Hex())
			fmt.Printf("Creditor: %s\n", returnDebtToEvent.Creditor.Hex())
			fmt.Printf("ID: %s\n", returnDebtToEvent.ID.String())

		case logCreateCreditSigHash.Hex():
			fmt.Printf("Log Name: Create Credit\n")

			var createCreditEvent LogCreateCredit

			err := contractAbi.UnpackIntoInterface(&createCreditEvent, "CreateCredit", vLog.Data)
			if err != nil {
				return err
			}

			createCreditEvent.Member = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateCreateCreditHistory(model.CreateCreditHistory{
				Member: createCreditEvent.Member.Hex(),
				Value:  createCreditEvent.Value.Int64(),
				TxHash: logCreateCreditSigHash.Hex(),
				BlockN: vLog.BlockNumber,
			})

			fmt.Printf("Member: %s\n", createCreditEvent.Member.Hex())
			fmt.Printf("Value: %s\n", createCreditEvent.Value.String())

		case logReduceDebtSigHash.Hex():
			fmt.Printf("Log Name: Reduce Debt\n")

			var reduceDebtEvent LogReduceDebt

			err := contractAbi.UnpackIntoInterface(&reduceDebtEvent, "ReduceDebt", vLog.Data)
			if err != nil {
				return err
			}

			reduceDebtEvent.Member = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateReduceDebtHistory(model.ReduceDebtHistory{
				Member: reduceDebtEvent.Member.Hex(),
				Value:  reduceDebtEvent.Value.Int64(),
				TxHash: logReduceDebtSigHash.Hex(),
				BlockN: vLog.BlockNumber,
			})

			fmt.Printf("Member: %s\n", reduceDebtEvent.Member.Hex())
			fmt.Printf("Value: %s\n", reduceDebtEvent.Value.String())

		case logChangeHakoOwnerSigHash.Hex():
			fmt.Printf("Log Name: Change Hako Owner\n")

			var changeHakoOwnerEvent LogChangeHakoOwner

			err := contractAbi.UnpackIntoInterface(&changeHakoOwnerEvent, "ChangeHakoOwner", vLog.Data)
			if err != nil {
				return err
			}

			changeHakoOwnerEvent.OldHakoOwner = common.HexToAddress(vLog.Topics[1].Hex())
			changeHakoOwnerEvent.NewHakoOwner = common.HexToAddress(vLog.Topics[2].Hex())

			d.DB.CreateChangeHakoOwnerHistory(model.ChangeHakoOwnerHistory{
				OldHakoOwner: changeHakoOwnerEvent.OldHakoOwner.Hex(),
				NewHakoOwner: changeHakoOwnerEvent.NewHakoOwner.Hex(),
				TxHash:       logChangeHakoOwnerSigHash.Hex(),
				BlockN:       vLog.BlockNumber,
			})

			fmt.Printf("OldHakoOwner: %s\n", changeHakoOwnerEvent.OldHakoOwner.Hex())
			fmt.Printf("NewHakoOwner: %s\n", changeHakoOwnerEvent.NewHakoOwner.Hex())

		case logChangeUpperLimitSigHash.Hex():
			fmt.Printf("Log Name: Change Upper Limit\n")

			var changeUpperLimitEvent LogChangeUpperLimit

			err := contractAbi.UnpackIntoInterface(&changeUpperLimitEvent, "ChangeUpperLimit", vLog.Data)
			if err != nil {
				return err
			}

			changeUpperLimitEvent.HakoOwner = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateChangeUpperLimitHistory(model.ChangeUpperLimitHistory{
				HakoOwner:     changeUpperLimitEvent.HakoOwner.Hex(),
				NewUpperLimit: changeUpperLimitEvent.NewUpperLimit.Int64(),
				TxHash:        logChangeUpperLimitSigHash.Hex(),
				BlockN:        vLog.BlockNumber,
			})

			fmt.Printf("HakoOwner: %s\n", changeUpperLimitEvent.HakoOwner.Hex())
			fmt.Printf("NewUpperLimit: %s\n", changeUpperLimitEvent.NewUpperLimit.String())

		case logGetRewardSigHash.Hex():
			fmt.Printf("Log Name: Get Reward\n")

			var getRewardEvent LogGetReward

			err := contractAbi.UnpackIntoInterface(&getRewardEvent, "GetReward", vLog.Data)
			if err != nil {
				return err
			}

			getRewardEvent.HakoOwner = common.HexToAddress(vLog.Topics[1].Hex())

			d.DB.CreateGetRewardHistory(model.GetRewardHistory{
				HakoOwner:   getRewardEvent.HakoOwner.Hex(),
				RewardValue: getRewardEvent.RewardValue.Int64(),
				TxHash:      logGetRewardSigHash.Hex(),
				BlockN:      vLog.BlockNumber,
			})

			fmt.Printf("HakoOwner: %s\n", getRewardEvent.HakoOwner.Hex())
			fmt.Printf("RewardValue: %s\n", getRewardEvent.RewardValue.String())

		}

	}
	return nil
}
