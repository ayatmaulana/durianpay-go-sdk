package payment

import "github.com/ayatmaulana/durianpay-go-sdk/common"

const (
  ROUTE_GROUP = "/payments"
)


type Payment struct {
  *common.Agent
}


func (p *Payment) FetchPayment() {

}

func (p *Payment) FetchPaymentById() {

}

func (p *Payment) ChargePayment() {

}

func (p *Payment) CheckPaymentStatus() {

}

func (p *Payment) VerifyPayment(){

}

func (p *Payment) CancelPayment(){

}

func (p *Payment) MDRFeeCalculation() {

}
