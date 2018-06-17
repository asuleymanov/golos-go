package client

import (
	"github.com/asuleymanov/golos-go/apis/network_broadcast"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/types"
	"time"
)

//SendTrx generates and sends an array of transactions to GOLOS.
func (client *Client) SendTrx(username string, strx []types.Operation, opts ...map[string]interface{}) (*BResp, error) {
	var options map[string]interface{}
	if len(opts) > 0 {
		options = opts[0]
	}

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

	// Получаем необходимый для подписи ключ
	privKeys, err := client.SigningKeys(strx[0])
	if err != nil {
		return nil, err
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, client.Chain); err != nil {
		return nil, err
	}

	// Отправка транзакции
	var resp *network_broadcast.BroadcastResponse
	if options != nil && options["async"].(bool) {
		err = client.NetworkBroadcast.BroadcastTransaction(tx.Transaction)
	} else {
		resp, err = client.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)
	}

	if err != nil {
		return nil, err
	}
	if resp != nil {
		var bresp BResp

		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}

	return nil, nil
}
