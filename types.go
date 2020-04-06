package accountsservice

// stdlib
import (
	"time"
)

// internal
import (
	"github.com/the-control-group/go-currency"
)

// TODO create custom time type for proper json decoding

type Customer struct {
	Id        int          `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Phone     *string       `json:"phone"`
	Created   time.Time    `json:"created"`
	Updated   time.Time    `json:"updated"`
	BrandSlug string       `json:"brand_slug"`
	Data      CustomerData `json:"data"`
}

type CustomerData struct {
	LoginHistory    []CustomerLoginHistory `json:"login.history"`
	LoginLatestIp   string                 `json:"login.latest_ip"`
	LoginLatestTime int64                  `json:"login.latest_time"`
}

type CustomerLoginHistory struct {
	Ip   string `json:"ip"`
	Time int64  `json:"time"`
}

type PaymentOption struct {
	Id                      int                                 `json:"id"`
	CustomerId              int                                 `json:"customer_id"`
	Created                 time.Time                           `json:"created"`
	Status                  string                              `json:"status"`
	Updated                 time.Time                           `json:"updated"`
	PaymentProcessor        string                              `json:"payment_processor"`
	PaymentProcessorDetails PaymentOptonPaymentProcessorDetails `json:"payment_processor_details"`
	FailureCode             *string                             `json:"failure_code,omitempty"`
	BrandSlug               string                              `json:"brand_slug"`
	FailureMessage          *string                             `json:"failure_message,omitempty"`
}

type PaymentOptonPaymentProcessorDetails struct {
	Bin         *string `json:"bin,omitempty"`
	Last4       *string `json:"last4,omitempty"`
	ExpDate     *string `json:"exp_date,omitempty"`
	ExpMonth    *string `json:"exp_month,omitempty"`
	ExpYear     *string `json:"exp_year,omitempty"`
	PaymentType *string `json:"payment_type,omitempty"`
	CardNetwork *string `json:"card_network,omitempty"`
	CardName    *string `json:"card_name,omitempty"`
	CardType    *string `json:"card_type,omitempty"`
	CardValidationResult *string `json:"card_validation_result,omitempty"`
}

type Order struct {
	Id                      int                          `json:"id"`
	PaymentOptionid int `json:"payment_option_id"`
	SubscriptionId          *int                         `json:"subsription_id"`
	CustomerId              int                          `json:"customer_id"`
	Type                    string                       `json:"type"`
	Cycle                   *int                         `json:"cycle"`
	Status                  string                       `json:"status"`
	Amount                  currency.Amount              `json:"amount"`
	Created                 time.Time                    `json:"created"`
	PaymentProcessor        string                       `json:"payment_processor"`
	PaymentProcessorDetails OrderPaymentProcessorDetails `json:"payment_processor_details"`
	BrandSlug               string                       `json:"brand_slug"`
	Updated                 time.Time                    `json:"updated"`
	Begins time.Time `json:"begins"`
	Ends time.Time `json:"ends"`
	Plans map[string]OrderQuantity `json:"plans"`
	Products map[string]OrderQuantity `json:"products"`
}

type OrderQuantity struct {
	Quantity int `json:"quantity"`
}

type OrderPaymentProcessorDetails struct {
}

type Transaction struct {
	Id                      int                                `json:"id"`
	BrandSlug               string                             `json:"brand_slug"`
	OrderId                 int                                `json:"order_id"`
	CustomerId              int                                `json:"customer_id"`
	Type                    string                             `json:"type"`
	Status                  string                             `json:"status"`
	Amount                  currency.Amount                    `json:"amount"`
	Created                 time.Time                          `json:"created"`
	PaymentProcessor        string                             `json:"payment_processor"`
	PaymentProcessorId *string `json:"payment_processor_id"`
	PaymentProcessorDetails TransactionPaymentProcessorDetails `json:"payment_processor_details"`
	Updated                 time.Time                          `json:"updated"`
	FailureCode             *string                            `json:"failure_code,omitempty"`
	FailureMessage          *string                            `json:"failure_message,omitempty"`
	PaymentOptionId int `json:"payment_option_id"`
}

type TransactionPaymentProcessorDetails struct {
}

type Subscription struct {
	Id                      int                                 `json:"id"`
	BrandSlug               string                              `json:"brand_slug"`
	OrderId                 int                                 `json:"order_id"`
	Status                  string                              `json:"status"`
	Cycle                   int                                 `json:"cycle"`
	PlanSku                 string                              `json:"plan_sku"`
	PaymentProcessor        string                              `json:"payment_processor"`
	Created                 time.Time                           `json:"created"`
	Updated                 time.Time                           `json:"updated"`
	Canceled                *time.Time                           `json:"canceled"`
	CustomerId              int                                 `json:"customer_id"`
	Next                    time.Time                           `json:"next"`
	FailureCode             *string                             `json:"failure_code,omitempty"`
	FailureMessage          *string                             `json:"failure_message,omitempty"`
	PaymentProcessorDetails SubscriptionPaymentProcessorDetails `json:"payment_processor_details"`
}

type SubscriptionPaymentProcessorDetails struct{}

type Plan struct {
	Id int `json:"id"`
	BrandSlug         string          `json:"brand_slug"`
	Sku               string          `json:"sku"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Terms             string          `json:"terms"`
	TrialPrice        currency.Amount `json:"trial_price"`
	TrialPeriod       int             `json:"trial_period"`
	TrialInterval     string          `json:"trial_interval"`
	RecurringPrice    currency.Amount `json:"recurring_price"`
	RecurringPeriod   int             `json:"recurring_period"`
	RecurringInterval string          `json:"recurring_interval"`
	RecurringCycles   *int            `json:"recurring_cycles,omitempty"`
	Created           time.Time       `json:"created"`
	Updated           time.Time       `json:"updated"`
	Status            string          `json:"status"`
	Data map[string]interface{} `json:"data"`
	Href string `json:"href"`
	Products map[string]PlanQuantity `json:"products"`
}

type PlanQuantity struct {
	RecurringQuantity int `json:"recurring_quantity"`
	TrialQuantity int `json:"trial_quantity"`
}

type Product struct {
	Id int `json:"id"`
	BrandSlug         string          `json:"brand_slug"`
	Sku               string          `json:"sku"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Price        currency.Amount `json:"price"`
	Type string `json:"type"`
}
