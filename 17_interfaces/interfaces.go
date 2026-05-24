package main 

import "fmt"

// define a contract - for any payment gateway must have a pay() method
type paymentGateway interface {
	pay(amount float32)
}

// payment doesn't care about which payment gateway is used 
type payment struct {
	gateway paymentGateway // accepts anything that satisfies the contract or you can say a slot for ANY payment gateway
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount) // just calls the contract method
} 

// struct 1
type razorpay struct {}

func (r razorpay) pay(amount float32) {
	// logic to make payment using razorpay
	fmt.Println("paying using razorpay", amount)
}

// struct 2
type stripe struct {}

func (s stripe) pay(amount float32) {
	fmt.Println("paying using stripe", amount)
}

// struct 3
type paypal struct {}

func (p paypal) pay(amount float32) {
	fmt.Println("paying using paypal", amount)
}

func main() {
	p1 := payment {gateway: razorpay{}}   // plug razorpay into the gateway slot
    p1.makePayment(1999)     //paying using razorpay 1999

	p2 := payment {gateway: stripe{}}   // plug stripe into the gateway slot
	p2.makePayment(2999)     //paying using stripe 2999

	p3 := payment {gateway: paypal{}} // plug paypal into the gateway slot
	p3.makePayment(3999)
}