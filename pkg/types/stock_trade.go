// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     clickstream.avsc
 *     clickstreamcodes.avsc
 *     clickstreamusers.avsc
 *     creditcards.avsc
 *     deviceinformation.avsc
 *     fleetmgmtdescription.avsc
 *     fleetmgmtlocation.avsc
 *     fleetmgmtsensors.avsc
 *     gaminggames.avsc
 *     gamingplayeractivity.avsc
 *     gamingplayers.avsc
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     netdevice.avsc
 *     orders.avsc
 *     pageviews.avsc
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
 *     shoeclickstream.avsc
 *     shoecustomer.avsc
 *     shoeorder.avsc
 *     siemlogs.avsc
 *     stockTrades.avsc
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

// Defines a hypothetical stock trade using some known test stock symbols.
type StockTrade struct {
	// A simulated trade side (buy or sell or short)
	Side string `json:"side"`
	// A simulated random quantity of the trade
	Quantity int32 `json:"quantity"`
	// Simulated stock symbols
	Symbol string `json:"symbol"`
	// A simulated random trade price in pennies
	Price int32 `json:"price"`
	// Simulated accounts assigned to the trade
	Account string `json:"account"`
	// The simulated user who executed the trade
	Userid string `json:"userid"`
}

const StockTradeAvroCRC64Fingerprint = "tj\b\x8cd\x01\xf0\xf6"

func NewStockTrade() StockTrade {
	r := StockTrade{}
	return r
}

func DeserializeStockTrade(r io.Reader) (StockTrade, error) {
	t := NewStockTrade()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStockTradeFromSchema(r io.Reader, schema string) (StockTrade, error) {
	t := NewStockTrade()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStockTrade(r StockTrade, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Side, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Quantity, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Symbol, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Price, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Account, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Userid, w)
	if err != nil {
		return err
	}
	return err
}

func (r StockTrade) Serialize(w io.Writer) error {
	return writeStockTrade(r, w)
}

func (r StockTrade) Schema() string {
	return "{\"doc\":\"Defines a hypothetical stock trade using some known test stock symbols.\",\"fields\":[{\"doc\":\"A simulated trade side (buy or sell or short)\",\"name\":\"side\",\"type\":{\"arg.properties\":{\"options\":[\"BUY\",\"SELL\"]},\"type\":\"string\"}},{\"doc\":\"A simulated random quantity of the trade\",\"name\":\"quantity\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5000,\"min\":1}},\"type\":\"int\"}},{\"doc\":\"Simulated stock symbols\",\"name\":\"symbol\",\"type\":{\"arg.properties\":{\"options\":[\"ZBZX\",\"ZJZZT\",\"ZTEST\",\"ZVV\",\"ZVZZT\",\"ZWZZT\",\"ZXZZT\"]},\"type\":\"string\"}},{\"doc\":\"A simulated random trade price in pennies\",\"name\":\"price\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1000,\"min\":5}},\"type\":\"int\"}},{\"doc\":\"Simulated accounts assigned to the trade\",\"name\":\"account\",\"type\":{\"arg.properties\":{\"options\":[\"ABC123\",\"LMN456\",\"XYZ789\"]},\"type\":\"string\"}},{\"doc\":\"The simulated user who executed the trade\",\"name\":\"userid\",\"type\":{\"arg.properties\":{\"regex\":\"User_[1-9]\"},\"type\":\"string\"}}],\"name\":\"ksql.StockTrade\",\"type\":\"record\"}"
}

func (r StockTrade) SchemaName() string {
	return "ksql.StockTrade"
}

func (_ StockTrade) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StockTrade) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StockTrade) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StockTrade) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StockTrade) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StockTrade) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StockTrade) SetString(v string)   { panic("Unsupported operation") }
func (_ StockTrade) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StockTrade) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Side}

		return w

	case 1:
		w := types.Int{Target: &r.Quantity}

		return w

	case 2:
		w := types.String{Target: &r.Symbol}

		return w

	case 3:
		w := types.Int{Target: &r.Price}

		return w

	case 4:
		w := types.String{Target: &r.Account}

		return w

	case 5:
		w := types.String{Target: &r.Userid}

		return w

	}
	panic("Unknown field index")
}

func (r *StockTrade) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StockTrade) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StockTrade) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StockTrade) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StockTrade) HintSize(int)                     { panic("Unsupported operation") }
func (_ StockTrade) Finalize()                        {}

func (_ StockTrade) AvroCRC64Fingerprint() []byte {
	return []byte(StockTradeAvroCRC64Fingerprint)
}

func (r StockTrade) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["side"], err = json.Marshal(r.Side)
	if err != nil {
		return nil, err
	}
	output["quantity"], err = json.Marshal(r.Quantity)
	if err != nil {
		return nil, err
	}
	output["symbol"], err = json.Marshal(r.Symbol)
	if err != nil {
		return nil, err
	}
	output["price"], err = json.Marshal(r.Price)
	if err != nil {
		return nil, err
	}
	output["account"], err = json.Marshal(r.Account)
	if err != nil {
		return nil, err
	}
	output["userid"], err = json.Marshal(r.Userid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *StockTrade) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["side"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Side); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for side")
	}
	val = func() json.RawMessage {
		if v, ok := fields["quantity"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Quantity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for quantity")
	}
	val = func() json.RawMessage {
		if v, ok := fields["symbol"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Symbol); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for symbol")
	}
	val = func() json.RawMessage {
		if v, ok := fields["price"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Price); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for price")
	}
	val = func() json.RawMessage {
		if v, ok := fields["account"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Account); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for account")
	}
	val = func() json.RawMessage {
		if v, ok := fields["userid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Userid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for userid")
	}
	return nil
}
