// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siem_logs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type InsuranceCustomerActivity struct {
	Activity_id int32 `json:"activity_id"`

	Customer_id int32 `json:"customer_id"`

	Activity_type string `json:"activity_type"`

	Propensity_to_churn float64 `json:"propensity_to_churn"`

	Ip_address string `json:"ip_address"`
}

const InsuranceCustomerActivityAvroCRC64Fingerprint = "\xd2\xe8Vu\x9a\xbe{\xe9"

func NewInsuranceCustomerActivity() InsuranceCustomerActivity {
	r := InsuranceCustomerActivity{}
	return r
}

func DeserializeInsuranceCustomerActivity(r io.Reader) (InsuranceCustomerActivity, error) {
	t := NewInsuranceCustomerActivity()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInsuranceCustomerActivityFromSchema(r io.Reader, schema string) (InsuranceCustomerActivity, error) {
	t := NewInsuranceCustomerActivity()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInsuranceCustomerActivity(r InsuranceCustomerActivity, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Activity_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Customer_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Activity_type, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Propensity_to_churn, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ip_address, w)
	if err != nil {
		return err
	}
	return err
}

func (r InsuranceCustomerActivity) Serialize(w io.Writer) error {
	return writeInsuranceCustomerActivity(r, w)
}

func (r InsuranceCustomerActivity) Schema() string {
	return "{\"fields\":[{\"name\":\"activity_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1,\"step\":1}},\"type\":\"int\"}},{\"name\":\"customer_id\",\"type\":{\"arg.properties\":{\"options\":[415,871,597,210,833,371,387,313,882,854,895,198,468,860,383,226,630,546,949,276,595,406,565,129,973,10,579,167,273,548,175,957,45,582,66,472,578,825,572,322,864,272,910,507,942,368,643,421,92,433,75,305,173,863,315,344,526,264,225,40,866,473,662,611,401,993,158,997,43,36,843,814,965,967,470,55,615,30,894,333,763,929,994,771,601,634,955,310,463,887,770,386,188,62,41,465,932,832,819,103,377,216,958,39,699,557,195,966,382,796,223,58,232,486,462,441,731,388,395,606,307,137,56,719,104,493,114,392,343,640,102,268,294,169,57,35,956,49,752,522,845,980,204,581,886,694,229,753,190,206]},\"type\":\"int\"}},{\"name\":\"activity_type\",\"type\":{\"arg.properties\":{\"options\":[\"web_open\",\"mobile_open\",\"new_account\"]},\"type\":\"string\"}},{\"name\":\"propensity_to_churn\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1,\"min\":0.01}},\"type\":\"double\"}},{\"name\":\"ip_address\",\"type\":{\"arg.properties\":{\"options\":[\"223.135.0.204\",\"36.149.179.191\",\"60.194.249.18\",\"136.186.114.150\",\"74.124.70.237\",\"69.212.203.214\",\"70.1.63.102\",\"37.130.125.130\",\"193.33.103.229\",\"32.61.200.245\",\"51.196.238.21\",\"111.187.166.26\",\"217.237.199.242\",\"202.194.220.196\",\"205.196.213.241\",\"137.209.182.50\",\"198.43.175.27\",\"124.166.107.190\",\"192.61.188.203\",\"57.125.102.149\",\"74.163.164.66\",\"67.66.179.220\",\"201.63.8.144\",\"50.57.65.143\",\"135.146.254.138\",\"225.103.12.27\",\"147.187.228.215\",\"94.87.184.82\",\"88.145.28.140\",\"208.216.1.101\",\"175.36.181.120\",\"194.7.156.81\",\"30.209.200.179\",\"72.112.128.155\",\"77.82.56.87\",\"76.241.80.115\",\"177.13.65.63\",\"92.160.162.102\",\"216.212.0.222\",\"76.119.252.185\",\"134.186.192.26\",\"85.228.219.76\",\"195.82.242.168\",\"109.187.240.70\",\"125.191.141.103\",\"176.58.219.253\",\"98.53.186.67\",\"21.127.2.182\",\"58.73.132.227\",\"251.119.237.69\"]},\"type\":\"string\"}}],\"name\":\"insurance.InsuranceCustomerActivity\",\"type\":\"record\"}"
}

func (r InsuranceCustomerActivity) SchemaName() string {
	return "insurance.InsuranceCustomerActivity"
}

func (_ InsuranceCustomerActivity) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetString(v string)   { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InsuranceCustomerActivity) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Activity_id}

		return w

	case 1:
		w := types.Int{Target: &r.Customer_id}

		return w

	case 2:
		w := types.String{Target: &r.Activity_type}

		return w

	case 3:
		w := types.Double{Target: &r.Propensity_to_churn}

		return w

	case 4:
		w := types.String{Target: &r.Ip_address}

		return w

	}
	panic("Unknown field index")
}

func (r *InsuranceCustomerActivity) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *InsuranceCustomerActivity) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ InsuranceCustomerActivity) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) HintSize(int)                     { panic("Unsupported operation") }
func (_ InsuranceCustomerActivity) Finalize()                        {}

func (_ InsuranceCustomerActivity) AvroCRC64Fingerprint() []byte {
	return []byte(InsuranceCustomerActivityAvroCRC64Fingerprint)
}

func (r InsuranceCustomerActivity) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["activity_id"], err = json.Marshal(r.Activity_id)
	if err != nil {
		return nil, err
	}
	output["customer_id"], err = json.Marshal(r.Customer_id)
	if err != nil {
		return nil, err
	}
	output["activity_type"], err = json.Marshal(r.Activity_type)
	if err != nil {
		return nil, err
	}
	output["propensity_to_churn"], err = json.Marshal(r.Propensity_to_churn)
	if err != nil {
		return nil, err
	}
	output["ip_address"], err = json.Marshal(r.Ip_address)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InsuranceCustomerActivity) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["activity_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Activity_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for activity_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["customer_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Customer_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for customer_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["activity_type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Activity_type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for activity_type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["propensity_to_churn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propensity_to_churn); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propensity_to_churn")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ip_address"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ip_address); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ip_address")
	}
	return nil
}