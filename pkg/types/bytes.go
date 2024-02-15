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

	"github.com/actgardner/gogen-avro/v10/util"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type Bytes []byte

func (b *Bytes) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*b = util.DecodeByteString(s)
	return nil
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b)), nil
}

type BytesWrapper struct {
	Target *Bytes
}

func (b BytesWrapper) SetBoolean(v bool) {
	panic("Unable to assign bytes to bytes field")
}

func (b BytesWrapper) SetInt(v int32) {
	panic("Unable to assign int to bytes field")
}

func (b BytesWrapper) SetLong(v int64) {
	panic("Unable to assign long to bytes field")
}

func (b BytesWrapper) SetFloat(v float32) {
	panic("Unable to assign float to bytes field")
}

func (b BytesWrapper) SetDouble(v float64) {
	panic("Unable to assign double to bytes field")
}

func (b BytesWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to bytes field")
}

func (b BytesWrapper) SetBytes(v []byte) {
	*(b.Target) = v
}

func (b BytesWrapper) SetString(v string) {
	*(b.Target) = []byte(v)
}

func (b BytesWrapper) Get(i int) types.Field {
	panic("Unable to get field from bytes field")
}

func (b BytesWrapper) SetDefault(i int) {
	panic("Unable to set default on bytes field")
}

func (b BytesWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from bytes field")
}

func (b BytesWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from bytes field")
}

func (b BytesWrapper) NullField(int) {
	panic("Unable to null field in bytes field")
}

func (b BytesWrapper) HintSize(int) {
	panic("Unable to hint size in bytes field")
}

func (b BytesWrapper) Finalize() {}