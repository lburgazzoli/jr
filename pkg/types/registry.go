// Autogenerated code. DO NOT EDIT.
package types

var address Address

var array_orderline_wrapper ArrayOrderlineWrapper

var bytes Bytes

var campaignfinance Campaignfinance

var clickstream Clickstream

var clickstreamusers Clickstreamusers

var codes Codes

var creditcards Creditcards

var destination Destination

var deviceinformation Deviceinformation

var fleetmgmtdescription Fleetmgmtdescription

var fleetmgmtlocation Fleetmgmtlocation

var fleetmgmtsensors Fleetmgmtsensors

var gaminggames Gaminggames

var gamingplayeractivity Gamingplayeractivity

var gamingplayers Gamingplayers

var insurancecustomeractivity Insurancecustomeractivity

var insurancecustomers Insurancecustomers

var insuranceoffers Insuranceoffers

var inventory Inventory

var location Location

var map_bool_wrapper MapBoolWrapper

var map_bytes_wrapper MapBytesWrapper

var map_float_wrapper MapFloatWrapper

var map_int_wrapper MapIntWrapper

var map_long_wrapper MapLongWrapper

var map_string_wrapper MapStringWrapper

var my_map_test_record MyMapTestRecord

var net_device NetDevice

var orderline Orderline

var orders Orders

var pageviews Pageviews

var payrollbonus Payrollbonus

var payrollemployee Payrollemployee

var payrollemployeelocation Payrollemployeelocation

var pizzaorders Pizzaorders

var pizzaorderscancelled Pizzaorderscancelled

var pizzaorderscompleted Pizzaorderscompleted

var product Product

var purchase Purchase

var ratings Ratings

var shoe Shoe

var shoe_clickstream ShoeClickstream

var shoe_customer ShoeCustomer

var shoe_order ShoeOrder

var siemlogs Siemlogs

var source Source

var stock_trade StockTrade

var stores Stores

var sysloglogs Sysloglogs

var transactions Transactions

var user User

var users Users

//gocyclo:ignore
func GetType(templateType string) interface{} {

	switch templateType {

	case "address":
		return &address

	case "array_orderline_wrapper":
		return &array_orderline_wrapper

	case "bytes":
		return &bytes

	case "campaignfinance":
		return &campaignfinance

	case "clickstream":
		return &clickstream

	case "clickstreamusers":
		return &clickstreamusers

	case "codes":
		return &codes

	case "creditcards":
		return &creditcards

	case "destination":
		return &destination

	case "deviceinformation":
		return &deviceinformation

	case "fleetmgmtdescription":
		return &fleetmgmtdescription

	case "fleetmgmtlocation":
		return &fleetmgmtlocation

	case "fleetmgmtsensors":
		return &fleetmgmtsensors

	case "gaminggames":
		return &gaminggames

	case "gamingplayeractivity":
		return &gamingplayeractivity

	case "gamingplayers":
		return &gamingplayers

	case "insurancecustomeractivity":
		return &insurancecustomeractivity

	case "insurancecustomers":
		return &insurancecustomers

	case "insuranceoffers":
		return &insuranceoffers

	case "inventory":
		return &inventory

	case "location":
		return &location

	case "map_bool_wrapper":
		return &map_bool_wrapper

	case "map_bytes_wrapper":
		return &map_bytes_wrapper

	case "map_float_wrapper":
		return &map_float_wrapper

	case "map_int_wrapper":
		return &map_int_wrapper

	case "map_long_wrapper":
		return &map_long_wrapper

	case "map_string_wrapper":
		return &map_string_wrapper

	case "my_map_test_record":
		return &my_map_test_record

	case "net_device":
		return &net_device

	case "orderline":
		return &orderline

	case "orders":
		return &orders

	case "pageviews":
		return &pageviews

	case "payrollbonus":
		return &payrollbonus

	case "payrollemployee":
		return &payrollemployee

	case "payrollemployeelocation":
		return &payrollemployeelocation

	case "pizzaorders":
		return &pizzaorders

	case "pizzaorderscancelled":
		return &pizzaorderscancelled

	case "pizzaorderscompleted":
		return &pizzaorderscompleted

	case "product":
		return &product

	case "purchase":
		return &purchase

	case "ratings":
		return &ratings

	case "shoe":
		return &shoe

	case "shoe_clickstream":
		return &shoe_clickstream

	case "shoe_customer":
		return &shoe_customer

	case "shoe_order":
		return &shoe_order

	case "siemlogs":
		return &siemlogs

	case "source":
		return &source

	case "stock_trade":
		return &stock_trade

	case "stores":
		return &stores

	case "sysloglogs":
		return &sysloglogs

	case "transactions":
		return &transactions

	case "user":
		return &user

	case "users":
		return &users
	}
	return nil
}
