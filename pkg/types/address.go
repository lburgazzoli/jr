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
 *     sysloglogs.avsc
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

type Address struct {
	City string `json:"city"`

	State string `json:"state"`

	Zipcode int64 `json:"zipcode"`
}

const AddressAvroCRC64Fingerprint = "\xc0$\x9c\xb4\xf4\x10ܪ"

func NewAddress() Address {
	r := Address{}
	return r
}

func DeserializeAddress(r io.Reader) (Address, error) {
	t := NewAddress()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAddressFromSchema(r io.Reader, schema string) (Address, error) {
	t := NewAddress()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAddress(r Address, w io.Writer) error {
	var err error
	err = vm.WriteString(r.City, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.State, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Zipcode, w)
	if err != nil {
		return err
	}
	return err
}

func (r Address) Serialize(w io.Writer) error {
	return writeAddress(r, w)
}

func (r Address) Schema() string {
	return "{\"fields\":[{\"name\":\"city\",\"type\":{\"arg.properties\":{\"regex\":\"City_[1-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"state\",\"type\":{\"arg.properties\":{\"regex\":\"State_[1-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"zipcode\",\"type\":{\"arg.properties\":{\"range\":{\"max\":99999,\"min\":10000}},\"type\":\"long\"}}],\"name\":\"ksql.address\",\"type\":\"record\"}"
}

func (r Address) SchemaName() string {
	return "ksql.address"
}

func (_ Address) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Address) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Address) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Address) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Address) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Address) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Address) SetString(v string)   { panic("Unsupported operation") }
func (_ Address) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Address) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.City}

		return w

	case 1:
		w := types.String{Target: &r.State}

		return w

	case 2:
		w := types.Long{Target: &r.Zipcode}

		return w

	}
	panic("Unknown field index")
}

func (r *Address) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Address) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Address) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Address) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Address) HintSize(int)                     { panic("Unsupported operation") }
func (_ Address) Finalize()                        {}

func (_ Address) AvroCRC64Fingerprint() []byte {
	return []byte(AddressAvroCRC64Fingerprint)
}

func (r Address) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["city"], err = json.Marshal(r.City)
	if err != nil {
		return nil, err
	}
	output["state"], err = json.Marshal(r.State)
	if err != nil {
		return nil, err
	}
	output["zipcode"], err = json.Marshal(r.Zipcode)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Address) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["city"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.City); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for city")
	}
	val = func() json.RawMessage {
		if v, ok := fields["state"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.State); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for state")
	}
	val = func() json.RawMessage {
		if v, ok := fields["zipcode"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Zipcode); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for zipcode")
	}
	return nil
}
