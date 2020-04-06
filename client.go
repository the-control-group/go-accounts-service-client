package accountsservice

// stdlib
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var ACCOUNTS_SERVICE_TIMEOUT = os.Getenv("ACCOUNTS_SERVICE_TIMEOUT")
var ACCOUNTS_SERVICE_API_URL = os.Getenv("ACCOUNTS_SERVICE_API_URL")
var AUTHORIZATION_HEADER     = os.Getenv("AUTHORIZATION_HEADER")

var client *http.Client

func init() {
	if ACCOUNTS_SERVICE_TIMEOUT == "" {
		ACCOUNTS_SERVICE_TIMEOUT = "30s"
		log.Println("Using default accounts service timeout", ACCOUNTS_SERVICE_TIMEOUT)
	}
	timeout, err := time.ParseDuration(ACCOUNTS_SERVICE_TIMEOUT)
	if err != nil {
	}
	client = &http.Client{
		Timeout: timeout,
	}
	if ACCOUNTS_SERVICE_API_URL == "" {
		ACCOUNTS_SERVICE_API_URL = "http://localhost:8000"
	}
}

type Error struct {
	error string `json:"error"`
	message string `json:"message"`
	failures []map[string]interface{} `json:"failures,omitempty"`
}

func (e *Error) Error() string {
	return e.message
}

type Filter struct {
	Filters [][]interface{}
}

func (f *Filter) Add(field, operator, value string) *Filter {
	f.Filters = append(f.Filters, []interface{}{field, operator, value})
	return f
}

func (f *Filter) String() (str string) {
	for _, filter := range f.Filters {
		str += fmt.Sprintf("filter[%s][%s]=%s",filter...)
	}
	return
}

func GetCustomer(customerId int) (customer *Customer, err error) {
	var getCustomerRequest *http.Request

	getCustomerRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/customers/%d", ACCOUNTS_SERVICE_API_URL, customerId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getCustomerRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getCustomerRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&customer)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetCustomerOrders(customerId int) (orders []Order, err error) {
	var getCustomerOrdersRequest *http.Request

	getCustomerOrdersRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/customers/%d/orders", ACCOUNTS_SERVICE_API_URL, customerId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getCustomerOrdersRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getCustomerOrdersRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&orders)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetCustomerPaymentOptions(customerId int) (paymentOptions []PaymentOption, err error) {
	var getCustomerPaymentOptionsRequest *http.Request

	getCustomerPaymentOptionsRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/customers/%d/payment_options", ACCOUNTS_SERVICE_API_URL, customerId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getCustomerPaymentOptionsRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getCustomerPaymentOptionsRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&paymentOptions)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetCustomerTransactions(customerId int) (transactions []Transaction, err error) {
	var getCustomerTransactionsRequest *http.Request

	getCustomerTransactionsRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/transactions?filter[customer_id][eq]=%d", ACCOUNTS_SERVICE_API_URL, customerId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getCustomerTransactionsRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getCustomerTransactionsRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&transactions)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetSubscription(subscriptionId int) (subscription *Subscription, err error) {
	var getSubscriptionRequest *http.Request

	getSubscriptionRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/subscriptions/%d", ACCOUNTS_SERVICE_API_URL, subscriptionId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getSubscriptionRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getSubscriptionRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&subscription)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetSubscriptionByOrderPlan(orderId int, planSku string) (subscription *Subscription, err error) {
	var getSubscriptionByOrderPlanRequest *http.Request

	getSubscriptionByOrderPlanRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/subscriptions/?filter[order_id][eq]=%d&filter[plan_sku][eq]=%s", ACCOUNTS_SERVICE_API_URL, orderId, planSku), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getSubscriptionByOrderPlanRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getSubscriptionByOrderPlanRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	var subscriptions []*Subscription

	err = json.NewDecoder(resp.Body).Decode(&subscriptions)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	subscription = subscriptions[0]

	return
}

func GetOrderSubscriptions(orderId int) (subscriptions []Subscription, err error) {
	var getOrderSubscriptionsRequest *http.Request

	getOrderSubscriptionsRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/subscriptions?filter[order_id][eq]=%d", ACCOUNTS_SERVICE_API_URL, orderId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getOrderSubscriptionsRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getOrderSubscriptionsRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&subscriptions)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetPlan(brandSlug, sku string) (plan *Plan, err error) {
	var getPlanRequest *http.Request

	getPlanRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/brands/%s/plans/%s", ACCOUNTS_SERVICE_API_URL, brandSlug, sku), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getPlanRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getPlanRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error

		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&plan)

	if err != nil {
		fmt.Println("GetPlan: Unable to json decode response body", err)
		return
	}

	return
}

func GetProduct(brandSlug, sku string) (product *Product, err error) {
	var getProductRequest *http.Request

	getProductRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/brands/%s/products/%s", ACCOUNTS_SERVICE_API_URL, brandSlug, sku), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getProductRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getProductRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("GetProduct: Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&product)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetSubscriptionOrders(subscriptionId int) (orders []Order, err error) {
	var getSubscriptionOrdersRequest *http.Request

	getSubscriptionOrdersRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/subscriptions/%d/orders", ACCOUNTS_SERVICE_API_URL, subscriptionId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getSubscriptionOrdersRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getSubscriptionOrdersRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&orders)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetOrder(orderId int) (order *Order, err error) {
	var getOrderRequest *http.Request

	getOrderRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/orders/%d", ACCOUNTS_SERVICE_API_URL, orderId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getOrderRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getOrderRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&order)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

// TODO
func GetOrderPlans(orderId int) (plans []Plan, err error) {
	return
}

// TODO
func GetOrderProducts(orderId int) (product []Product, err error) {
	return
}

func GetPaymentOptions(filter string) (paymentOptions []PaymentOption, err error) {
	var getPaymentOptionsRequest *http.Request

	getPaymentOptionsRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/payment_options?%s", ACCOUNTS_SERVICE_API_URL, filter), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getPaymentOptionsRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getPaymentOptionsRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&paymentOptions)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetPaymentOption(paymentOptionId int) (paymentOption *PaymentOption, err error) {
	var getPaymentOptionRequest *http.Request

	getPaymentOptionRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/payment_options/%d", ACCOUNTS_SERVICE_API_URL, paymentOptionId), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getPaymentOptionRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getPaymentOptionRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountsError Error
		err = json.NewDecoder(resp.Body).Decode(&accountsError)

		if err != nil {
			fmt.Println("Unable to json decode error response body", err)
		}

		err = &accountsError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&paymentOption)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetOrdersAggregate(filter *Filter, group string, aggregate []string) (agg map[string]interface{}, err error) {
	var getOrdersAggregateRequest *http.Request

	var jsonAgg []byte
	jsonAgg, err = json.Marshal(aggregate)

	if err != nil {
		return
	}

	getOrdersAggregateRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/transactions?%s&group=%s&aggregate=%s", ACCOUNTS_SERVICE_API_URL, filter, group, jsonAgg), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getOrdersAggregateRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getOrdersAggregateRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountServiceError Error
		err = json.NewDecoder(resp.Body).Decode(&accountServiceError)

		if err != nil {
			fmt.Println("Unable to json decode response body", err)
			return
		}

		err = &accountServiceError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&agg)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}

func GetTransactions(filter string) (transactions []Transaction, err error) {
	var getTransactionsRequest *http.Request

	getTransactionsRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/transactions?%s", ACCOUNTS_SERVICE_API_URL, filter), nil)
	
	if err != nil {
		return
	}

	if AUTHORIZATION_HEADER != "" {
		getTransactionsRequest.Header.Add("Authorization", AUTHORIZATION_HEADER)
	}

	var resp *http.Response

	resp, err = client.Do(getTransactionsRequest)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		var accountServiceError Error
		err = json.NewDecoder(resp.Body).Decode(&accountServiceError)

		if err != nil {
			fmt.Println("Unable to json decode response body", err)
			return
		}

		err = &accountServiceError

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&transactions)

	if err != nil {
		fmt.Println("Unable to json decode response body", err)
		return
	}

	return
}
