package main

import (
	//"abigen/bind"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/sirupsen/logrus"
	abi "Conflux_Smart_Contract/contracts/abi"
)

func main() {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
	})
	if err != nil {
		panic(err)
	}

	err = client.AccountManager.UnlockDefault("hello")
	if err != nil {
		panic(err)
	}
	contractAddr := cfxaddress.MustNew("cfxtest:acd7apn6pnfhna7w1pa8evzhwhv3085vjjp1b8bav5")
	instance, err := abi.NewMyERC20Token(contractAddr, client)
	if err != nil {
		panic(err)
	}

	user := cfxaddress.MustNew("cfxtest:aasfup1wgjyxkzy3575cbnn87xj5tam2zud125twew")
	result, err := instance.BalanceOf(nil, user.MustGetCommonAddress())
	if err != nil {
		panic(err)
	}

	logrus.WithField("balance", result).WithField("user", user).Info("access contract") // "bar"
}