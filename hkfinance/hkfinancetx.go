package hkfinance

import (
	"io/ioutil"
	"math/big"

	"hako-server/config"
	"hako-server/model"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type HkfinanceTx struct {
	Config    *config.Config
	Ethclient *ethclient.Client
	Hkfinance *Hkfinance
}

// return Hkfinance instance
func getHkfinanceInstance(config *config.Config) (*Hkfinance, error) {
	client, err := ethclient.Dial(config.Ethereum.NetworkURL)
	if err != nil {
		return nil, err
	}
	contractAddressBytes, err := ioutil.ReadFile(config.Ethereum.ContractAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(string(contractAddressBytes))
	instance, err := NewHkfinance(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// return eth client
func GetEthclient(config *config.Config) (*ethclient.Client, error) {
	client, err := ethclient.Dial(config.Ethereum.NetworkURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// create new HkfinanceTx instance
func NewHkfinanceTx(config *config.Config) (*HkfinanceTx, error) {
	hkfinanceInstance, err := getHkfinanceInstance(config)
	if err != nil {
		return nil, err
	}
	ethclient, err := GetEthclient(config)
	if err != nil {
		return nil, err
	}
	return &HkfinanceTx{
		Config:    config,
		Hkfinance: hkfinanceInstance,
		Ethclient: ethclient,
	}, nil
}

func (h *HkfinanceTx) GetHakoInfo() (*model.HakoInfo, error) {
	address, err := ioutil.ReadFile(h.Config.Ethereum.ContractAddress)
	if err != nil {
		return nil, err
	}
	balance, err := h.Hkfinance.BalanceOfHako(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	totalSupply, err := h.Hkfinance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	credit, err := h.Hkfinance.CreditOfHako(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	debt, err := h.Hkfinance.DebtOfHako(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	memberCount, err := h.Hkfinance.MemberCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	upperLimit, err := h.Hkfinance.UpperLimit(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	name, err := h.Hkfinance.Name(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	symbol, err := h.Hkfinance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	decimals, err := h.Hkfinance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &model.HakoInfo{
		Address:     string(address),
		Balance:     balance.Int64(),
		TotalSupply: totalSupply.Int64(),
		Credit:      credit.Int64(),
		Debt:        debt.Int64(),
		MemberCount: memberCount.Int64(),
		UpperLimit:  upperLimit.Int64(),
		Name:        name,
		Symbol:      symbol,
		Decimals:    decimals.Int64(),
	}, nil
}

func (h *HkfinanceTx) GetUserInfo(hexaddress string) (*model.UserInfo, error) {
	address := common.HexToAddress(hexaddress)
	balance, err := h.Hkfinance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	member, err := h.Hkfinance.MemberCheckOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	credit, err := h.Hkfinance.CreditToHakoOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	debt, err := h.Hkfinance.DebtToHakoOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	lending, err := h.Hkfinance.CreditToMemberOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	borrowing, err := h.Hkfinance.DebtToMemberOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}

	netAssets := big.NewInt(0)
	netAssets = netAssets.Add(netAssets, balance).Add(netAssets, credit).Add(netAssets, lending)
	netAssets = netAssets.Sub(netAssets, debt).Sub(netAssets, borrowing)

	valueDuration, err := h.Hkfinance.BorrowValueDurationOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	value := valueDuration[0]
	duration := valueDuration[1]

	return &model.UserInfo{
		Address:   hexaddress,
		Balance:   balance.Int64(),
		Member:    int(member.Int64()),
		Credit:    credit.Int64(),
		Debt:      debt.Int64(),
		Lending:   lending.Int64(),
		Borrowing: borrowing.Int64(),
		NetAssets: netAssets.Int64(),
		Value:     value.Int64(),
		Duration:  duration.Int64(),
	}, nil
}

func (h *HkfinanceTx) GetOwnerInfo() (*model.OwnerInfo, error) {
	hakoAddress, err := ioutil.ReadFile(h.Config.Ethereum.ContractAddress)
	if err != nil {
		return nil, err
	}
	ownerAddress, err := h.Hkfinance.HakoOwner(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	hakoBalance, err := h.Hkfinance.BalanceOfHako(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	totalSupply, err := h.Hkfinance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	hakoCredit, err := h.Hkfinance.CreditOfHako(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	hakoDebt, err := h.Hkfinance.DebtOfHako(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ownerBalance, err := h.Hkfinance.BalanceOfHakoOwner(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	lendingCount, err := h.Hkfinance.LendCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	memberCount, err := h.Hkfinance.MemberCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	upperLimit, err := h.Hkfinance.UpperLimit(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &model.OwnerInfo{
		HakoAddress:  string(hakoAddress),
		OwnerAddress: ownerAddress.String(),
		HakoBalance:  hakoBalance.Int64(),
		TotalSupply:  totalSupply.Int64(),
		HakoCredit:   hakoCredit.Int64(),
		HakoDebt:     hakoDebt.Int64(),
		OwnerBalance: ownerBalance.Int64(),
		LendingCount: lendingCount.Int64(),
		MemberCount:  memberCount.Int64(),
		UpperLimit:   upperLimit.Int64(),
	}, nil
}
