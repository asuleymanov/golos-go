package golos

import (
	"time"

	"github.com/asuleymanov/golos-go/api"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/types"
)

//SendTrx generates and sends an array of transactions to VIZ.
func (client *Client) SendTrx(username string, strx []types.Operation) (*types.OperationResponse, error) {
	var bresp types.OperationResponse

	// Getting the necessary parameters
	props, err := client.API.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Creating a Transaction
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Adding Operations to a Transaction
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	// Obtain the key required for signing
	privKeys, err := client.SigningKeys(strx[0])
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	if err := tx.Sign(privKeys, client.Config.ChainID); err != nil {
		return nil, err
	}

	// Sending a transaction
	var resp *api.BroadcastResponse
	if client.asyncProtocol {
		err = client.API.BroadcastTransaction(tx.Transaction)
	} else {
		resp, err = client.API.BroadcastTransactionSynchronous(tx.Transaction)
	}

	bresp.JSONTrx, _ = JSONTrxString(tx)

	if err != nil {
		return &bresp, err
	}
	if resp != nil && !client.asyncProtocol {
		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}

	return &bresp, nil
}

func (client *Client) GetTrx(strx []types.Operation) (*types.Transaction, error) {
	// Getting the necessary parameters
	props, err := client.API.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Creating a Transaction
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	}

	// Adding Operations to a Transaction
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	return tx, nil
}
