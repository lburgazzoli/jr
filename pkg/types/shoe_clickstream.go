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

type ShoeClickstream struct {
	Product_id string `json:"product_id"`

	User_id string `json:"user_id"`

	View_time int32 `json:"view_time"`

	Page_url string `json:"page_url"`

	Ip string `json:"ip"`

	Ts int64 `json:"ts"`
}

const ShoeClickstreamAvroCRC64Fingerprint = "\x1e\xd3E\xf0\f\xc2\b\xf6"

func NewShoeClickstream() ShoeClickstream {
	r := ShoeClickstream{}
	return r
}

func DeserializeShoeClickstream(r io.Reader) (ShoeClickstream, error) {
	t := NewShoeClickstream()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoeClickstreamFromSchema(r io.Reader, schema string) (ShoeClickstream, error) {
	t := NewShoeClickstream()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoeClickstream(r ShoeClickstream, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Product_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.User_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.View_time, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Page_url, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ip, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	return err
}

func (r ShoeClickstream) Serialize(w io.Writer) error {
	return writeShoeClickstream(r, w)
}

func (r ShoeClickstream) Schema() string {
	return "{\"fields\":[{\"name\":\"product_id\",\"type\":\"string\"},{\"name\":\"user_id\",\"type\":\"string\"},{\"name\":\"view_time\",\"type\":\"int\"},{\"name\":\"page_url\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"ts\",\"type\":\"long\"}],\"name\":\"shoes.ShoeClickstream\",\"type\":\"record\"}"
}

func (r ShoeClickstream) SchemaName() string {
	return "shoes.ShoeClickstream"
}

func (_ ShoeClickstream) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ShoeClickstream) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ShoeClickstream) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ShoeClickstream) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ShoeClickstream) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ShoeClickstream) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ShoeClickstream) SetString(v string)   { panic("Unsupported operation") }
func (_ ShoeClickstream) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ShoeClickstream) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Product_id}

		return w

	case 1:
		w := types.String{Target: &r.User_id}

		return w

	case 2:
		w := types.Int{Target: &r.View_time}

		return w

	case 3:
		w := types.String{Target: &r.Page_url}

		return w

	case 4:
		w := types.String{Target: &r.Ip}

		return w

	case 5:
		w := types.Long{Target: &r.Ts}

		return w

	}
	panic("Unknown field index")
}

func (r *ShoeClickstream) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ShoeClickstream) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ShoeClickstream) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ShoeClickstream) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ShoeClickstream) HintSize(int)                     { panic("Unsupported operation") }
func (_ ShoeClickstream) Finalize()                        {}

func (_ ShoeClickstream) AvroCRC64Fingerprint() []byte {
	return []byte(ShoeClickstreamAvroCRC64Fingerprint)
}

func (r ShoeClickstream) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["product_id"], err = json.Marshal(r.Product_id)
	if err != nil {
		return nil, err
	}
	output["user_id"], err = json.Marshal(r.User_id)
	if err != nil {
		return nil, err
	}
	output["view_time"], err = json.Marshal(r.View_time)
	if err != nil {
		return nil, err
	}
	output["page_url"], err = json.Marshal(r.Page_url)
	if err != nil {
		return nil, err
	}
	output["ip"], err = json.Marshal(r.Ip)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ShoeClickstream) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["product_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Product_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for product_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["user_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for user_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["view_time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.View_time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for view_time")
	}
	val = func() json.RawMessage {
		if v, ok := fields["page_url"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Page_url); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for page_url")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ip"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ip); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ip")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ts")
	}
	return nil
}
