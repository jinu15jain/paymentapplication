package paymentschema

import (
	"time"
)

type Transaction struct  {
	ID string `dynamo:"ID,hash"`
	MID string `dynamo:"MID"`
	Money string `dynamo:"Money"`
	Status string `dynamo:"Status"`
	CreatedOn time.Time `dynamo:"CreatedOn"`
	UpdatedOn time.Time `dynamo:"UpdatedOn"`
}


type Refund struct  {
	ID string `dynamo:"ID,hash"`
	MID string `dynamo:"MID"`
	TransactionId string `dynamo:"TransactionId"`
	Money string `dynamo:"MID"`
	Status string `dynamo:"Status"`
	CreatedOn time.Time `dynamo:"CreatedOn"`
	UpdatedOn time.Time `dynamo:"UpdatedOn"`
}