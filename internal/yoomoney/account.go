package yoomoney

import (
	"encoding/json"
	"io"
)

////////////////////////// accout-info

type AccountInfoResponseBalanceDetails struct {
	Total             float32 `json:"total"`
	Available         float32 `json:"available"`
	DepositionPending float32 `json:"deposition_pending"`
	Blocked           float32 `json:"blocked"`
	Debt              float32 `json:"debt"`
	Hold              float32 `json:"hold"`
}

type AccountInfoResponseCardsLinked struct {
	PanFragment string `json:"pan_fragment"`
	Type        string `json:"type"`
}

type AccountInfoResponse struct {
	Account        string                            `json:"account"`
	Balance        float32                           `json:"balance"`
	Currency       string                            `json:"currency"`
	AccountStatus  string                            `json:"account_status"`
	AccountType    string                            `json:"account_type"`
	BalanceDetails AccountInfoResponseBalanceDetails `json:"balance_details"`
	CardsLinked    []AccountInfoResponseCardsLinked  `json:"cards_linked"`
}

func (client *Client) CallAccountInfo() (*AccountInfoResponse, error) {
	resp, err := client.sendRequest("account-info", nil)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var accountInfoResponse AccountInfoResponse

	if err := json.Unmarshal(data, &accountInfoResponse); err != nil {
		return nil, err
	}

	return &accountInfoResponse, nil
}

////////////////////////// operation-history

type OperationHistoryRequest struct {
	Type        string `schema:"type"`
	Label       string `schema:"label"`
	From        string `schema:"from"`
	Till        string `schema:"till"`
	StartRecord string `schema:"start_record"`
	Records     int    `schema:"records"`
	Details     bool   `schema:"details"`
}

type OperationHistoryResponseOperation struct {
	OperationID string  `json:"operation_id"`
	Status      string  `json:"status"`
	Datetime    string  `json:"datetime"`
	Title       string  `json:"title"`
	PatternID   string  `json:"pattern_id"`
	Direction   string  `json:"direction"`
	Amount      float32 `json:"amount"`
	Label       string  `json:"label"`
	Type        string  `json:"type"`
}

type OperationHistoryResponse struct {
	Error      string                            `json:"error"`
	NextRecord string                            `json:"next_record"`
	Operations OperationHistoryResponseOperation `json:"operations"`
}

func (client *Client) CallOperationHistory(req *OperationHistoryRequest) (*OperationHistoryResponse, error) {
	resp, err := client.sendRequest("operation-history", req)
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

	return &operationHistoryResponse, nil
}

////////////////////////// operation-details
// TODO
