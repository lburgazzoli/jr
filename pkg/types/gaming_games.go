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
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payrollbonus.avsc
 *     payrollemployee.avsc
 *     payrollemployeelocation.avsc
 *     pizzaorders.avsc
 *     pizzaorderscancelled.avsc
 *     pizzaorderscompleted.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
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

type GamingGames struct {
	Id int32 `json:"id"`

	Room_name string `json:"room_name"`

	Created_date int64 `json:"created_date"`
}

const GamingGamesAvroCRC64Fingerprint = "}\xc4\x1d2\x85\x16\r]"

func NewGamingGames() GamingGames {
	r := GamingGames{}
	return r
}

func DeserializeGamingGames(r io.Reader) (GamingGames, error) {
	t := NewGamingGames()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeGamingGamesFromSchema(r io.Reader, schema string) (GamingGames, error) {
	t := NewGamingGames()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeGamingGames(r GamingGames, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Room_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Created_date, w)
	if err != nil {
		return err
	}
	return err
}

func (r GamingGames) Serialize(w io.Writer) error {
	return writeGamingGames(r, w)
}

func (r GamingGames) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5000,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"room_name\",\"type\":{\"arg.properties\":{\"options\":[\"Arcade -- Rookie\",\"Arcade -- Skilled\",\"Arcade -- Expert\",\"Survival -- Rookie\",\"Survival -- Skilled\",\"Survival -- Expert\",\"Classic -- Rookie\",\"Classic -- Skilled\",\"Classic -- Expert\"]},\"type\":\"string\"}},{\"name\":\"created_date\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1609459200000,\"step\":100000}},\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"gaming.GamingGames\",\"type\":\"record\"}"
}

func (r GamingGames) SchemaName() string {
	return "gaming.GamingGames"
}

func (_ GamingGames) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ GamingGames) SetInt(v int32)       { panic("Unsupported operation") }
func (_ GamingGames) SetLong(v int64)      { panic("Unsupported operation") }
func (_ GamingGames) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ GamingGames) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ GamingGames) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ GamingGames) SetString(v string)   { panic("Unsupported operation") }
func (_ GamingGames) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *GamingGames) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Room_name}

		return w

	case 2:
		w := types.Long{Target: &r.Created_date}

		return w

	}
	panic("Unknown field index")
}

func (r *GamingGames) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *GamingGames) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ GamingGames) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ GamingGames) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ GamingGames) HintSize(int)                     { panic("Unsupported operation") }
func (_ GamingGames) Finalize()                        {}

func (_ GamingGames) AvroCRC64Fingerprint() []byte {
	return []byte(GamingGamesAvroCRC64Fingerprint)
}

func (r GamingGames) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["room_name"], err = json.Marshal(r.Room_name)
	if err != nil {
		return nil, err
	}
	output["created_date"], err = json.Marshal(r.Created_date)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *GamingGames) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["room_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Room_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for room_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["created_date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Created_date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for created_date")
	}
	return nil
}
