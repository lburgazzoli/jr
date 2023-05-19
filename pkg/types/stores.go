// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     NetDevice.avsc
 *     User.avsc
 *     campaign_finance.avsc
 *     clickstream.avsc
 *     clickstream_codes.avsc
 *     clickstream_users.avsc
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
 *     orders.avsc
 *     pageviews.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe_clickstream.avsc
 *     shoe_customers.avsc
 *     shoe_orders.avsc
 *     shoes_product.avsc
 *     siem_logs.avsc
 *     stockTrades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
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

type Stores struct {
	Store_id int32 `json:"store_id"`

	City string `json:"city"`

	State string `json:"state"`
}

const StoresAvroCRC64Fingerprint = "\xc8@p2i:\xbc\xee"

func NewStores() Stores {
	r := Stores{}
	return r
}

func DeserializeStores(r io.Reader) (Stores, error) {
	t := NewStores()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStoresFromSchema(r io.Reader, schema string) (Stores, error) {
	t := NewStores()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStores(r Stores, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Store_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.City, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.State, w)
	if err != nil {
		return err
	}
	return err
}

func (r Stores) Serialize(w io.Writer) error {
	return writeStores(r, w)
}

func (r Stores) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"city\":\"Raleigh\",\"state\":\"NC\",\"store_id\":1},{\"city\":\"Chicago\",\"state\":\"IL\",\"store_id\":2},{\"city\":\"Sacramento\",\"state\":\"CA\",\"store_id\":3},{\"city\":\"Austin\",\"state\":\"TX\",\"store_id\":4},{\"city\":\"Boston\",\"state\":\"MA\",\"store_id\":5},{\"city\":\"Atlanta\",\"state\":\"GA\",\"store_id\":6},{\"city\":\"Lexington\",\"state\":\"SC\",\"store_id\":7}]},\"fields\":[{\"name\":\"store_id\",\"type\":\"int\"},{\"name\":\"city\",\"type\":\"string\"},{\"name\":\"state\",\"type\":\"string\"}],\"name\":\"datagen.example.stores\",\"type\":\"record\"}"
}

func (r Stores) SchemaName() string {
	return "datagen.example.stores"
}

func (_ Stores) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Stores) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Stores) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Stores) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Stores) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Stores) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Stores) SetString(v string)   { panic("Unsupported operation") }
func (_ Stores) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Stores) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Store_id}

		return w

	case 1:
		w := types.String{Target: &r.City}

		return w

	case 2:
		w := types.String{Target: &r.State}

		return w

	}
	panic("Unknown field index")
}

func (r *Stores) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Stores) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Stores) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Stores) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Stores) HintSize(int)                     { panic("Unsupported operation") }
func (_ Stores) Finalize()                        {}

func (_ Stores) AvroCRC64Fingerprint() []byte {
	return []byte(StoresAvroCRC64Fingerprint)
}

func (r Stores) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["store_id"], err = json.Marshal(r.Store_id)
	if err != nil {
		return nil, err
	}
	output["city"], err = json.Marshal(r.City)
	if err != nil {
		return nil, err
	}
	output["state"], err = json.Marshal(r.State)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Stores) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["store_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_id")
	}
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
	return nil
}