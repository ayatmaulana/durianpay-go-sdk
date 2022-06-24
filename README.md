# Durianpay SDK for Golang (Unofficial)


## 🧐 Project Philosophy

> Durianpay is a Payment Gateway that can be used in Indonesia with the support of many payment methods such as VA, Bank Transfer, EWallet, Bank Transfer, Credit Card, and of course PayLater. Not only that, to get to Go Live the process is very easy.
> 
> But it's a shame, besides the API that can be used but there is no Official SDK that has been created and supported by Durianpay, this greatly adds to the workload when you want to use Golang as a language that will be integrated with Durianpay

<!-- **Read more about the Well app on [the project homepage](https://projects.colegaw.in/well-app?utm_source=GitHub&utm_medium=readme&utm_campaign=well_app_readme).** -->

## 📒 Official Documentation

All of API Documentation is on the Official Website of Durian Pay [https://durianpay.id/docs/api/](https://durianpay.id/docs/api/).

*⚠️ API documentation is subject to change without notice, so this SDK may not work properly.*
## 🚀 Installation

```bash
$ go get github.com/ayatmaulana/durianpay-go-sdk
```

## Usage

```go
package main

import (
    "github.com/ayatmaulana/durianpay-go-sdk"
    "github.com/ayatmaulana/durianpay-go-sdk/common"
)

... 
func main() {
  dp = durianpay.NewClient(&common.ClientConfig{
    Mode: common.SANDBOX, // SANBOX or LIVE
    ApiKey: "xxx",
    EnableLogging: true,
  })
}
```

# 👨‍💻 Library Usage

Here's a brief high-level overview of the Library uses:

- This project uses pure `"net/http"` Golang package for requesting Durianpay API
- For Logging purpose we use Logrus - [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
- We use Testify - [https://github.com/stretchr/testify](https://github.com/stretchr/testify) for test our code
  - And also using httpmock - [https://github.com/jarcoal/httpmock](https://github.com/jarcoal/httpmock) to mocking http request in test


## 🌴 Example
### Create Order
```go
import (
    "time"
    "github.com/ayatmaulana/durianpay-go-sdk"
    "github.com/ayatmaulana/durianpay-go-sdk/common"
    "github.com/ayatmaulana/durianpay-go-sdk/modules/order"
)

...

  payload := &order.CreateOrderRequestPayload{
    Amount: "100000",
    Currency: "IDR",
    ExpiryDate: time.Now().Add(time.Hour + time.Duration(24)),
    OrderRefID: "INV/ORDER/xxxxx1",
    PaymentOption: "full_payment",
    Metadata: order.Metadata{},
    Items: []order.Items{
      {
        Name: "Soap",
        Qty: 1,
        Price: "100000",
        Logo: "",
      },
    },
    Customer: order.Customer{
      CustomerRefID: "CUS/xx1" ,
      GivenName: "Ayat Maulana",
      Email: "myemail@example.com",
      Mobile: "0812345678",
      Address: order.Address{
        ReceiverName: "Hehe",
        ReceiverPhone: "0812345678",
      },
    },
  }

  res, err := dp.Order.CreateOrder(context.TODO(), payload)
  if err != nil {
    // handle error
    log.Println(err)
  }
```

### Create Payment

```go
  ctx := context.TODO()
  res, err := dp.Payment.ChargePaymentVA(ctx, &payment.ChargePaymentVARequestPayload{
    OrderID: orderId,
    Amount: "100000",
    BankCode: payment.BCA,
    Name: "Ayat Maulana",
  }, &payment.SandboxOption{
    ForceFail: true,
  })

```

### Use with Context
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 10))
defer cancel()

res, err := dp.Disbursement.FetchBankList(ctx)
if err != nil {
  log.Println(err)
}
log.Println(res)
```
## 📐 Todo 
- [ ] API Integration
  - [x] Order
  - [x] Payment
  - [x] Promo
  - [x] Disbursement
  - [x] Settlement
  - [x] Refund
  - [x] Subscription
  - [ ] EWallet Account
  - [ ] Customer
  - [ ] Static VA
- [ ] Testing
- [ ] Documentation per module

## ✍️ Contributing

Interested in contributing to this project? Thanks so much for your interest! We are always looking for improvements to the project and contributions from open-source developers are greatly appreciated.

If you have a contribution in mind, just fork this repository, follow code guide and make pull request in main repository.

<!-- ## 🌟 Spread the word! -->
<!---->
<!-- If you want to say thank you and/or support active development of the Well app: -->
<!---->
<!-- - Add a GitHub Star to the project! -->
<!-- - Tweet about the project on your Twitter! -->
<!--   - Tag [@colegawin_](https://twitter.com/colegawin_) and/or `#thewellapp` -->
<!-- - Leave us a review [on the iOS App Store](https://apps.apple.com/us/app/well-reboot-your-mindset/id1573357406)! -->
<!---->
<!-- Thanks so much for your interest in growing the reach of the Well app! -->
<!---->
<!-- _**PS:** consider sponsoring me ([Cole Gawin](https://colegaw.in)) to continue the development of this project on [BuyMeACoffee](https://buymeacoffee.com/colegawin) :)_ -->

## 📃 License

The Well app is free and open-source software licensed under the GNU General Public License v3.0.


