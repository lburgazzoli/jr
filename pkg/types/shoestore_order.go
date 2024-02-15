// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     csv_product.avsc
 *     csv_user.avsc
 *     finance_stock_trade.avsc
 *     fleetmgmt_description.avsc
 *     fleetmgmt_location.avsc
 *     fleetmgmt_sensor.avsc
 *     gaming_game.avsc
 *     gaming_player.avsc
 *     gaming_player_activity.avsc
 *     genericstore_purchase.avsc
 *     insurance_customer.avsc
 *     insurance_customer_activity.avsc
 *     insurance_offer.avsc
 *     inventorymgmt_inventory.avsc
 *     inventorymgmt_product.avsc
 *     iot_device_information.avsc
 *     map_dumb_schema.avsc
 *     marketing_campaign_finance.avsc
 *     net_device.avsc
 *     payment_credit_card.avsc
 *     payment_transaction.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizzastore_order.avsc
 *     pizzastore_order_cancelled.avsc
 *     pizzastore_order_completed.avsc
 *     shoestore_clickstream.avsc
 *     shoestore_customer.avsc
 *     shoestore_order.avsc
 *     shoestore_shoe.avsc
 *     shopping_order.avsc
 *     shopping_rating.avsc
 *     siem_log.avsc
 *     store.avsc
 *     syslog_log.avsc
 *     user.avsc
 *     users.avsc
 *     users_array_map.avsc
 *     webanalytics_clickstream.avsc
 *     webanalytics_code.avsc
 *     webanalytics_page_view.avsc
 *     webanalytics_user.avsc
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

type ShoestoreOrder struct {
	Order_id int32 `json:"order_id"`

	Product_id string `json:"product_id"`

	Customer_id string `json:"customer_id"`

	Ts int64 `json:"ts"`
}

const ShoestoreOrderAvroCRC64Fingerprint = "\xe7_\xf3\xf3r\xa1:)"

func NewShoestoreOrder() ShoestoreOrder {
	r := ShoestoreOrder{}
	return r
}

func DeserializeShoestoreOrder(r io.Reader) (ShoestoreOrder, error) {
	t := NewShoestoreOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoestoreOrderFromSchema(r io.Reader, schema string) (ShoestoreOrder, error) {
	t := NewShoestoreOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoestoreOrder(r ShoestoreOrder, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Order_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Product_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Customer_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	return err
}

func (r ShoestoreOrder) Serialize(w io.Writer) error {
	return writeShoestoreOrder(r, w)
}

func (r ShoestoreOrder) Schema() string {
	return "{\"fields\":[{\"name\":\"order_id\",\"type\":\"int\"},{\"name\":\"product_id\",\"type\":\"string\"},{\"name\":\"customer_id\",\"type\":\"string\"},{\"name\":\"ts\",\"type\":\"long\"}],\"name\":\"shoes.ShoestoreOrder\",\"type\":\"record\"}"
}

func (r ShoestoreOrder) SchemaName() string {
	return "shoes.ShoestoreOrder"
}

func (_ ShoestoreOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ ShoestoreOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ShoestoreOrder) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Order_id}

		return w

	case 1:
		w := types.String{Target: &r.Product_id}

		return w

	case 2:
		w := types.String{Target: &r.Customer_id}

		return w

	case 3:
		w := types.Long{Target: &r.Ts}

		return w

	}
	panic("Unknown field index")
}

func (r *ShoestoreOrder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ShoestoreOrder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ShoestoreOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ShoestoreOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ShoestoreOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ ShoestoreOrder) Finalize()                        {}

func (_ ShoestoreOrder) AvroCRC64Fingerprint() []byte {
	return []byte(ShoestoreOrderAvroCRC64Fingerprint)
}

func (r ShoestoreOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["order_id"], err = json.Marshal(r.Order_id)
	if err != nil {
		return nil, err
	}
	output["product_id"], err = json.Marshal(r.Product_id)
	if err != nil {
		return nil, err
	}
	output["customer_id"], err = json.Marshal(r.Customer_id)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ShoestoreOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["order_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Order_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for order_id")
	}
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