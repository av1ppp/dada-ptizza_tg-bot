package yoomoney

import (
	"encoding/json"
	"io"
)

////////////////////////// request-payment
// TODO

type RequestPaymentRequest struct {
	PatternID string `schema:"pattern_id"` // 	Фиксированное значение: p2p

	PhoneNumber string // Номер телефона в формате ITU-T E.164, полный номер, начиная с 7

	// Идентификатор получателя перевода (номер счета, номер телефона или email).
	To string `schema:"to"`

	Amount    float32 `schema:"amount"`     // Сумма к оплате (столько заплатит отправитель).
	AmountDue float32 `schema:"amount_due"` // Сумма к получению (придет на счет получателя после оплаты).
	Comment   string  `schema:"comment"`    // Комментарий к переводу, отображается в истории отправителя.
	Message   string  `schema:"message"`    // Комментарий к переводу, отображается получателю.
	Label     string  `schema:"label"`      // Метка платежа. Необязательный параметр.

	// Значение параметра true - признак того, что перевод защищен
	// кодом протекции. По умолчанию параметр отсутствует (обычный перевод).
	Codepro bool `schema:"codepro"`

	// Число дней, в течении которых получатель перевода может ввести
	// код протекции и получить перевод на свой счет.
	ExpirePeriod int `schema:"expire_period"`
}

// type RequestPaymentResponseMoneySource struct {
// 	// TODO
// }

type RequestPaymentResponse struct {
	Status                 string  `json:"status"`
	Error                  string  `json:"error"`
	MoneySource            string  `json:"money_source"`
	RequestID              string  `json:"request_id"`
	ContractAmount         float32 `json:"contract_amount"`
	Balance                float32 `json:"balance"`
	RecipientAccountStatus string  `json:"recipient_account_status"`
	RecipientAccountType   string  `json:"recipient_account_type"`
	ProtectionCode         string  `json:"protection_code"`
	AccountUnblockUri      string  `json:"account_unblock_uri"`
	ExtActionUri           string  `json:"ext_action_uri"`
}

func (client *Client) CallRequestPayment(req *RequestPaymentRequest) (*RequestPaymentResponse, error) {
	resp, err := client.sendRequest("request-payment", req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var requestPaymentResponse RequestPaymentResponse

	if err := json.Unmarshal(data, &requestPaymentResponse); err != nil {
		return nil, err
	}

	return &requestPaymentResponse, nil
}
