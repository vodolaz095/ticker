package models

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var client *sdk.RestClient

// Connect connects to tinkoff api
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

// GetPositions gets positions
func GetPositions(ctx context.Context, accountID string) (positions []sdk.PositionBalance, err error) {
	log.Println("Getting positions...")
	positions, err = client.PositionsPortfolio(ctx, accountID)
	if err != nil {
		return
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].Name < positions[j].Name
	})
	return
}

// GetCurrencies gets currencies
func GetCurrencies(ctx context.Context, accountID string) (currenciesInPortfolio []sdk.CurrencyBalance, err error) {
	log.Println("Getting currencies...")
	return client.CurrenciesPortfolio(ctx, accountID)
}
