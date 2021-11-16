package models

import (
	"context"
	"github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
	"time"
)

var client *sdk.RestClient

func Connect(token string) (err error) {
	c := sdk.NewRestClient(token)
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	accounts, err := c.Accounts(ctx)
	if err != nil {
		return
	}
	for _, account := range accounts {
		log.Printf("Account %s %s is found!", account.ID, account.Type)
	}
	client = c
	return
}

func GetPositions(ctx context.Context, accountID string) (positions []sdk.PositionBalance, err error) {
	log.Println("Getting positions...")
	return client.PositionsPortfolio(ctx, accountID)
}

func GetCurrencies(ctx context.Context, accountID string) (currenciesInPortfolio []sdk.CurrencyBalance, err error) {
	log.Println("Getting currencies...")
	return client.CurrenciesPortfolio(ctx, accountID)
}
