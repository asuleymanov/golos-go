package client

import (
	"github.com/asuleymanov/golos-go/apis/network_broadcast"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/types"
	"time"
)

//SendTrx generates and sends an array of transactions to GOLOS.
func (client *Client) SendTrx(username string, strx []types.Operation) (*BResp, error) {
	var bresp BResp

	// Получение необходимых параметров
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	// Получаем необходимый для подписи ключ
	privKeys, err := client.SigningKeys(strx[0])
	if err != nil {
		return nil, err
	}

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, client.Chain); err != nil {
		return nil, err
	}

	// Отправка транзакции
	var resp *network_broadcast.BroadcastResponse
	if client.AsyncProtocol {
		err = client.NetworkBroadcast.BroadcastTransaction(tx.Transaction)
	} else {
		resp, err = client.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)
	}

	bresp.JSONTrx = JSONTrxString(tx)

	if err != nil {
		return &bresp, err
	}
	if resp != nil {
		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}

	return &bresp, nil
}
