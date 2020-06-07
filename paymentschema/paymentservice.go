package paymentschema

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"time"
)


type PaymentService struct {
	transactionTable dynamo.Table
}

type RefundService struct {
	refundTable dynamo.Table
}


func NewPaymentService(itemTableName string) (*PaymentService, error) {
	dynamoTable, err := newDynamoTable(itemTableName, "")
	if err != nil {
		return nil, err
	}
	return &PaymentService{
		transactionTable: dynamoTable,
	}, nil
}

func NewRefundService(itemTableName string) (*RefundService, error) {
	dynamoTable, err := newDynamoTable(itemTableName, "")
	if err != nil {
		return nil, err
	}
	return &RefundService{
		refundTable: dynamoTable,
	}, nil
}


func newDynamoTable(tableName ,endpoint string) (dynamo.Table, error) {
	if tableName == "" {
		return dynamo.Table{}, fmt.Errorf("you must supply a table name")
	}
	cfg := aws.Config{}
	cfg.Region = aws.String("eu-west-2")
	if endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
	}

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &cfg)
	table := db.Table(tableName)
	return table, nil
}

func (p *PaymentService) CreateTransaction(item *Transaction) error {
	now := time.Now()
	item.CreatedOn = now
	item.UpdatedOn = now
	return p.transactionTable.Put(item).Run()
}

func (p *PaymentService) GetTransaction(transaction *Transaction) error {
	now := time.Now()
	transaction.CreatedOn = now
	transaction.UpdatedOn = now
	return p.transactionTable.Get("ID", transaction.ID).One(transaction)
}


func (r *RefundService) RefundTransaction(refund *Refund) error {
	now := time.Now()
	refund.CreatedOn = now
	refund.UpdatedOn = now
	return r.refundTable.Put(refund).Run()
}

func (r *RefundService) GetRefund(refund *Refund) error {
	now := time.Now()
	refund.CreatedOn = now
	refund.UpdatedOn = now
	return r.refundTable.Get("ID", refund.ID).One(refund)
}


