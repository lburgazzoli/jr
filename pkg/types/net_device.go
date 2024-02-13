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

type NetDevice struct {
	VLAN string `json:"VLAN"`

	IPV4_SRC_ADDR string `json:"IPV4_SRC_ADDR"`

	IPV4_DST_ADDR string `json:"IPV4_DST_ADDR"`

	IN_BYTES int32 `json:"IN_BYTES"`

	FIRST_SWITCHED int64 `json:"FIRST_SWITCHED"`

	LAST_SWITCHED int64 `json:"LAST_SWITCHED"`

	L4_SRC_PORT int32 `json:"L4_SRC_PORT"`

	L4_DST_PORT int32 `json:"L4_DST_PORT"`

	TCP_FLAGS int32 `json:"TCP_FLAGS"`

	PROTOCOL int32 `json:"PROTOCOL"`

	SRC_TOS int32 `json:"SRC_TOS"`

	SRC_AS int32 `json:"SRC_AS"`

	DST_AS int32 `json:"DST_AS"`

	L7_PROTO int32 `json:"L7_PROTO"`

	L7_PROTO_NAME string `json:"L7_PROTO_NAME"`

	L7_PROTO_CATEGORY string `json:"L7_PROTO_CATEGORY"`
}

const NetDeviceAvroCRC64Fingerprint = "\x17tJ7\xff\x9b\x8bs"

func NewNetDevice() NetDevice {
	r := NetDevice{}
	return r
}

func DeserializeNetDevice(r io.Reader) (NetDevice, error) {
	t := NewNetDevice()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNetDeviceFromSchema(r io.Reader, schema string) (NetDevice, error) {
	t := NewNetDevice()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNetDevice(r NetDevice, w io.Writer) error {
	var err error
	err = vm.WriteString(r.VLAN, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IPV4_SRC_ADDR, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IPV4_DST_ADDR, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IN_BYTES, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FIRST_SWITCHED, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.LAST_SWITCHED, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.L4_SRC_PORT, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.L4_DST_PORT, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TCP_FLAGS, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.PROTOCOL, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.SRC_TOS, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.SRC_AS, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.DST_AS, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.L7_PROTO, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.L7_PROTO_NAME, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.L7_PROTO_CATEGORY, w)
	if err != nil {
		return err
	}
	return err
}

func (r NetDevice) Serialize(w io.Writer) error {
	return writeNetDevice(r, w)
}

func (r NetDevice) Schema() string {
	return "{\"fields\":[{\"name\":\"VLAN\",\"type\":\"string\"},{\"name\":\"IPV4_SRC_ADDR\",\"type\":\"string\"},{\"name\":\"IPV4_DST_ADDR\",\"type\":\"string\"},{\"name\":\"IN_BYTES\",\"type\":\"int\"},{\"name\":\"FIRST_SWITCHED\",\"type\":\"long\"},{\"name\":\"LAST_SWITCHED\",\"type\":\"long\"},{\"name\":\"L4_SRC_PORT\",\"type\":\"int\"},{\"name\":\"L4_DST_PORT\",\"type\":\"int\"},{\"name\":\"TCP_FLAGS\",\"type\":\"int\"},{\"name\":\"PROTOCOL\",\"type\":\"int\"},{\"name\":\"SRC_TOS\",\"type\":\"int\"},{\"name\":\"SRC_AS\",\"type\":\"int\"},{\"name\":\"DST_AS\",\"type\":\"int\"},{\"name\":\"L7_PROTO\",\"type\":\"int\"},{\"name\":\"L7_PROTO_NAME\",\"type\":\"string\"},{\"name\":\"L7_PROTO_CATEGORY\",\"type\":\"string\"}],\"name\":\"jr.NetDevice\",\"type\":\"record\"}"
}

func (r NetDevice) SchemaName() string {
	return "jr.NetDevice"
}

func (_ NetDevice) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NetDevice) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NetDevice) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NetDevice) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NetDevice) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NetDevice) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NetDevice) SetString(v string)   { panic("Unsupported operation") }
func (_ NetDevice) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NetDevice) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.VLAN}

		return w

	case 1:
		w := types.String{Target: &r.IPV4_SRC_ADDR}

		return w

	case 2:
		w := types.String{Target: &r.IPV4_DST_ADDR}

		return w

	case 3:
		w := types.Int{Target: &r.IN_BYTES}

		return w

	case 4:
		w := types.Long{Target: &r.FIRST_SWITCHED}

		return w

	case 5:
		w := types.Long{Target: &r.LAST_SWITCHED}

		return w

	case 6:
		w := types.Int{Target: &r.L4_SRC_PORT}

		return w

	case 7:
		w := types.Int{Target: &r.L4_DST_PORT}

		return w

	case 8:
		w := types.Int{Target: &r.TCP_FLAGS}

		return w

	case 9:
		w := types.Int{Target: &r.PROTOCOL}

		return w

	case 10:
		w := types.Int{Target: &r.SRC_TOS}

		return w

	case 11:
		w := types.Int{Target: &r.SRC_AS}

		return w

	case 12:
		w := types.Int{Target: &r.DST_AS}

		return w

	case 13:
		w := types.Int{Target: &r.L7_PROTO}

		return w

	case 14:
		w := types.String{Target: &r.L7_PROTO_NAME}

		return w

	case 15:
		w := types.String{Target: &r.L7_PROTO_CATEGORY}

		return w

	}
	panic("Unknown field index")
}

func (r *NetDevice) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NetDevice) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NetDevice) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NetDevice) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NetDevice) HintSize(int)                     { panic("Unsupported operation") }
func (_ NetDevice) Finalize()                        {}

func (_ NetDevice) AvroCRC64Fingerprint() []byte {
	return []byte(NetDeviceAvroCRC64Fingerprint)
}

func (r NetDevice) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["VLAN"], err = json.Marshal(r.VLAN)
	if err != nil {
		return nil, err
	}
	output["IPV4_SRC_ADDR"], err = json.Marshal(r.IPV4_SRC_ADDR)
	if err != nil {
		return nil, err
	}
	output["IPV4_DST_ADDR"], err = json.Marshal(r.IPV4_DST_ADDR)
	if err != nil {
		return nil, err
	}
	output["IN_BYTES"], err = json.Marshal(r.IN_BYTES)
	if err != nil {
		return nil, err
	}
	output["FIRST_SWITCHED"], err = json.Marshal(r.FIRST_SWITCHED)
	if err != nil {
		return nil, err
	}
	output["LAST_SWITCHED"], err = json.Marshal(r.LAST_SWITCHED)
	if err != nil {
		return nil, err
	}
	output["L4_SRC_PORT"], err = json.Marshal(r.L4_SRC_PORT)
	if err != nil {
		return nil, err
	}
	output["L4_DST_PORT"], err = json.Marshal(r.L4_DST_PORT)
	if err != nil {
		return nil, err
	}
	output["TCP_FLAGS"], err = json.Marshal(r.TCP_FLAGS)
	if err != nil {
		return nil, err
	}
	output["PROTOCOL"], err = json.Marshal(r.PROTOCOL)
	if err != nil {
		return nil, err
	}
	output["SRC_TOS"], err = json.Marshal(r.SRC_TOS)
	if err != nil {
		return nil, err
	}
	output["SRC_AS"], err = json.Marshal(r.SRC_AS)
	if err != nil {
		return nil, err
	}
	output["DST_AS"], err = json.Marshal(r.DST_AS)
	if err != nil {
		return nil, err
	}
	output["L7_PROTO"], err = json.Marshal(r.L7_PROTO)
	if err != nil {
		return nil, err
	}
	output["L7_PROTO_NAME"], err = json.Marshal(r.L7_PROTO_NAME)
	if err != nil {
		return nil, err
	}
	output["L7_PROTO_CATEGORY"], err = json.Marshal(r.L7_PROTO_CATEGORY)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NetDevice) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["VLAN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VLAN); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VLAN")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IPV4_SRC_ADDR"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IPV4_SRC_ADDR); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IPV4_SRC_ADDR")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IPV4_DST_ADDR"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IPV4_DST_ADDR); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IPV4_DST_ADDR")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IN_BYTES"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IN_BYTES); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IN_BYTES")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FIRST_SWITCHED"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FIRST_SWITCHED); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FIRST_SWITCHED")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LAST_SWITCHED"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LAST_SWITCHED); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LAST_SWITCHED")
	}
	val = func() json.RawMessage {
		if v, ok := fields["L4_SRC_PORT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.L4_SRC_PORT); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for L4_SRC_PORT")
	}
	val = func() json.RawMessage {
		if v, ok := fields["L4_DST_PORT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.L4_DST_PORT); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for L4_DST_PORT")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TCP_FLAGS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TCP_FLAGS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TCP_FLAGS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PROTOCOL"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PROTOCOL); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PROTOCOL")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SRC_TOS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SRC_TOS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SRC_TOS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SRC_AS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SRC_AS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SRC_AS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DST_AS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DST_AS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DST_AS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["L7_PROTO"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.L7_PROTO); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for L7_PROTO")
	}
	val = func() json.RawMessage {
		if v, ok := fields["L7_PROTO_NAME"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.L7_PROTO_NAME); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for L7_PROTO_NAME")
	}
	val = func() json.RawMessage {
		if v, ok := fields["L7_PROTO_CATEGORY"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.L7_PROTO_CATEGORY); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for L7_PROTO_CATEGORY")
	}
	return nil
}
