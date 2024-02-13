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

type Product struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Price float64 `json:"price"`
}

const ProductAvroCRC64Fingerprint = "\x11\x95\xf9C\xff4\xa8\x9c"

func NewProduct() Product {
	r := Product{}
	return r
}

func DeserializeProduct(r io.Reader) (Product, error) {
	t := NewProduct()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeProductFromSchema(r io.Reader, schema string) (Product, error) {
	t := NewProduct()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeProduct(r Product, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Price, w)
	if err != nil {
		return err
	}
	return err
}

func (r Product) Serialize(w io.Writer) error {
	return writeProduct(r, w)
}

func (r Product) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"long\"}},{\"name\":\"name\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"string\"}},{\"name\":\"description\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"string\"}},{\"name\":\"price\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"double\"}}],\"name\":\"ksql.product\",\"type\":\"record\"}"
}

func (r Product) SchemaName() string {
	return "ksql.product"
}

func (_ Product) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Product) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Product) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Product) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Product) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Product) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Product) SetString(v string)   { panic("Unsupported operation") }
func (_ Product) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Product) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Name}

		return w

	case 2:
		w := types.String{Target: &r.Description}

		return w

	case 3:
		w := types.Double{Target: &r.Price}

		return w

	}
	panic("Unknown field index")
}

func (r *Product) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Product) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Product) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Product) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Product) HintSize(int)                     { panic("Unsupported operation") }
func (_ Product) Finalize()                        {}

func (_ Product) AvroCRC64Fingerprint() []byte {
	return []byte(ProductAvroCRC64Fingerprint)
}

func (r Product) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	output["price"], err = json.Marshal(r.Price)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Product) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["description"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Description); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for description")
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
	return nil
}
