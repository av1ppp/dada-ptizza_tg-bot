package yoomoney

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

////////////////////////// operation-history

type OperationHistoryRequest struct {
	// Перечень типов операций, которые требуется отобразить. Возможные значения:
	//		deposition — пополнение счета (приход);
	//		payment — платежи со счета (расход);
	//		incoming-transfers-unaccepted — непринятые входящие P2P-переводы любого типа.
	//
	// Типы операций перечисляются через пробел. Если параметр отсутствует, выводятся все операции.
	Type string

	// Отбор платежей по значению метки. Выбираются платежи, у которых указано заданное значение параметра label
	// вызова request-payment.
	Label string

	From string // Вывести операции от момента времени
	Till string // Вывести операции до момента времени

	// Если параметр присутствует, то будут отображены операции, начиная с номера start_record.
	StartRecord string

	// Количество запрашиваемых записей истории операций. Допустимые значения: от 1
	// до 100, по умолчанию — 30.
	Records int

	// Показывать подробные детали операции. По умолчанию false.
	// Для отображения деталей операции требуется наличие права operation-details.
	Details bool
}

type OperationHistoryResponse struct {
	Error      string `json:"error"`       // Код ошибки. Присутствует при ошибке выполнения запроса.
	NextRecord string `json:"next_record"` // Порядковый номер первой записи на следующей странице истории операций.
	Operations []struct {
		OperationID string `json:"operation_id"` // Идентификатор операции.

		// Статус платежа (перевода). Может принимать следующие значения:
		//		- success - платеж успешно завершен;
		//		- refused - платеж отвергнут получателем или отменен отправителем;
		//		- in_progess - платеж не завершен, перевод не принят получателем
		//			или ожидает ввода кода протекции.
		Status string `json:"status"`

		Datetime  string `json:"datetime"`   // Дата и время совершения операции.
		Title     string `json:"title"`      // Краткое описание операции (название магазина или источник пополнения).
		PatternID string `json:"pattern_id"` // Идентификатор шаблона, по которому совершен платеж.

		// Направление движения средств. Может принимать значения:
		//		- in (приход);
		//		- out (расход);
		Direction string `json:"direction"`

		Amount float32 `json:"amount"` // Сумма операции
		Label  string  `json:"label"`  // Метка платежа

		// Тип операции. Возможные значения:
		//		- payment-shop - исходящий платеж в магазин;
		//		- outgoing-transfer - исходящий P2P-перевод любого типа;
		//		- deposition - зачисление;
		//		- incoming-transfer - входящий перевод;
		//		- incoming-transfer-protected - входящий перевод с кодом протекции.
		Type string `json:"type"`
	} `json:"operations"`
}

func (client *Client) CallOperationHistory(req *OperationHistoryRequest) (*OperationHistoryResponse, error) {
	body := make(map[string]string)
	if req.Type != "" {
		body["type"] = req.Type
	}
	if req.Label != "" {
		body["label"] = req.Label
	}
	if req.From != "" {
		body["from"] = req.From
	}
	if req.Till != "" {
		body["till"] = req.Till
	}
	if req.StartRecord != "" {
		body["start_record"] = req.StartRecord
	}
	if req.Records != 0 {
		body["records"] = fmt.Sprint(req.Records)
	}
	if req.Details {
		body["details"] = "true"
	}

	resp, err := client.sendRequest("operation-history", body)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var operationHistoryResponse OperationHistoryResponse

	if err := json.Unmarshal(data, &operationHistoryResponse); err != nil {
		return nil, err
	}

	if operationHistoryResponse.Error != "" {
		return nil, errors.New(operationHistoryResponse.Error)
	}

	return &operationHistoryResponse, nil
}
