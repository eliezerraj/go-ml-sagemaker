package core

import (
	"time"

)

type SageMakerEndpoint struct {
	FraudEndpoint		string		`json:"fraud_endpoint,omitempty"`
	CustomerEndpoint	string		`json:"customer_endpoint,omitempty"`
}

type PaymentFraud struct {
	AccountID		string		`json:"account_id,omitempty"`
	CardNumber		string		`json:"card_number,omitempty"`
	TerminalName	string		`json:"terminal_name,omitempty"`
	CoordX			int32		`json:"coord_x,omitempty"`
	CoordY			int32		`json:"coord_y,omitempty"`
	CardType		string  	`json:"card_type,omitempty"`
	CardModel		string  	`json:"card_model,omitempty"`
	PaymentAt		time.Time	`json:"payment_at,omitempty"`
	MCC				string  	`json:"mcc,omitempty"`
	Status			string  	`json:"status,omitempty"`
	Currency		string  	`json:"currency,omitempty"`
	Amount			float64 	`json:"amount,omitempty"`
	Tx1Day			float64 	`json:"tx_1d,omitempty"`
	Avg1Day			float64 	`json:"avg_1d,omitempty"`
	Tx7Day			float64 	`json:"tx_7d,omitempty"`
	Avg7Day			float64 	`json:"avg_7d,omitempty"`
	Tx30Day			float64 	`json:"tx_30d,omitempty"`
	Avg30Day		float64 	`json:"avg_30d,omitempty"`
	TimeBtwTx		int32 		`json:"time_btw_cc_tx,omitempty"`
	Fraud			float64	  	`json:"fraud,omitempty"`
}

type CustomerClassification struct {
	Dim_1		float64		`json:"dim_1,omitempty"`
	Dim_2		float64		`json:"dim_2,omitempty"`
	Dim_3		float64		`json:"dim_3,omitempty"`
	Dim_4		float64		`json:"dim_4,omitempty"`
	Cluster		string		`json:"cluster,omitempty"`
}