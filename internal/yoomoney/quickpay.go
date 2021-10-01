package yoomoney

import (
	"fmt"
	"net/url"
)

var baseFormUrl = baseUrl + "/quickpay/confirm.xml"

type CreateFormOptions struct {
	Receiver     string  // Номер кошелька ЮMoney, на который нужно зачислять деньги отправителей.
	QuickpayForm string  // Возможные значения: shop (для унив. формы), small, donate
	Targets      string  // Назначение платежа.
	PaymentType  string  // Способ оплаты (PC - ЮMoney, AC - банк. карта, MC - баланс мобильного)
	Sum          float32 // Сумма отправителя

	AmountDue float32 // Сумма получения
	A         float32 // Коэффициент комиссии

	// Необязательные параметры
	FormComment string // Название перевода в истории отправителя
	ShortDest   string // Название перевода на странице подтверждения.
	Label       string // Метка, которую сайт или приложение присваивает конкретному переводу.
	Comment     string // Поле, в котором можно передать комментарий отправителя перевода.
	SuccessURL  string // URL-адрес для редиректа после совершения перевода.
	NeedFio     bool   // Нужны ФИО отправителя.
	NeedEmail   bool   // Нужна электронная почты отправителя.
	NeedPhone   bool   // Нужен телефон отправителя.
	NeedAddress bool   // Нужен адрес отправителя.
}

type CreateFormResponse struct {
	OrigURL *url.URL
	TempURL *url.URL
}

func (client *Client) CreateFormURL(opts CreateFormOptions) (*CreateFormResponse, error) {
	values := url.Values{}
	values.Add("receiver", opts.Receiver)
	values.Add("targets", opts.Targets)
	values.Add("payment-type", opts.PaymentType)
	values.Add("sum", fmt.Sprint(opts.Sum))
	if opts.AmountDue != 0 && opts.A != 0 {
		values.Add("amount_due", fmt.Sprint(opts.AmountDue))
		values.Add("a", fmt.Sprint(opts.A))
	}
	values.Add("quickpay-form", opts.QuickpayForm)

	if opts.FormComment != "" {
		values.Add("form-comment", opts.FormComment)
	}
	if opts.ShortDest != "" {
		values.Add("short-dest", opts.ShortDest)
	}
	if opts.Label != "" {
		values.Add("label", opts.Label)
	}
	if opts.Comment != "" {
		values.Add("comment", opts.Comment)
	}
	if opts.Comment != "" {
		values.Add("success-url", opts.SuccessURL)
	}

	values.Add("need-fio", fmt.Sprint(opts.NeedFio))
	values.Add("need-email", fmt.Sprint(opts.NeedEmail))
	values.Add("need-phone", fmt.Sprint(opts.NeedPhone))
	values.Add("need-address", fmt.Sprint(opts.NeedAddress))

	origU, _ := url.Parse(baseFormUrl)
	origU.RawQuery = values.Encode()

	// get temp url
	resp, err := client.httpClient.Get(origU.String())
	if err != nil {
		return nil, err
	}

	tempU, err := url.Parse(resp.Header.Get("Location"))
	if err != nil {
		return nil, err
	}

	// return u
	return &CreateFormResponse{
		OrigURL: origU,
		TempURL: tempU,
	}, nil
}
