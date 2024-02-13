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

type GamingPlayers struct {
	Player_id int32 `json:"player_id"`

	Player_name string `json:"player_name"`

	Ip string `json:"ip"`
}

const GamingPlayersAvroCRC64Fingerprint = ";+۟7\r\xe9W"

func NewGamingPlayers() GamingPlayers {
	r := GamingPlayers{}
	return r
}

func DeserializeGamingPlayers(r io.Reader) (GamingPlayers, error) {
	t := NewGamingPlayers()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeGamingPlayersFromSchema(r io.Reader, schema string) (GamingPlayers, error) {
	t := NewGamingPlayers()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeGamingPlayers(r GamingPlayers, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Player_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Player_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ip, w)
	if err != nil {
		return err
	}
	return err
}

func (r GamingPlayers) Serialize(w io.Writer) error {
	return writeGamingPlayers(r, w)
}

func (r GamingPlayers) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"ip\":\"117.119.34.13\",\"player_id\":1001,\"player_name\":\"Frieda Baldi\"},{\"ip\":\"253.165.97.77\",\"player_id\":1002,\"player_name\":\"Cherrita Gallaccio\"},{\"ip\":\"98.173.165.60\",\"player_id\":1003,\"player_name\":\"Matt Cleugh\"},{\"ip\":\"213.77.228.11\",\"player_id\":1004,\"player_name\":\"Dulciana Murfill\"},{\"ip\":\"40.62.99.184\",\"player_id\":1005,\"player_name\":\"Germayne Streetley\"},{\"ip\":\"46.152.206.98\",\"player_id\":1006,\"player_name\":\"Brenna Woolfall\"},{\"ip\":\"241.18.67.160\",\"player_id\":1007,\"player_name\":\"Gerhardt Tenbrug\"},{\"ip\":\"60.105.214.197\",\"player_id\":1008,\"player_name\":\"Hayley Tuma\"},{\"ip\":\"68.145.84.22\",\"player_id\":1009,\"player_name\":\"Winny Cadigan\"},{\"ip\":\"189.253.70.170\",\"player_id\":1010,\"player_name\":\"Bonnibelle Macek\"},{\"ip\":\"183.153.131.220\",\"player_id\":1011,\"player_name\":\"Lionel Byneth\"},{\"ip\":\"190.205.93.27\",\"player_id\":1012,\"player_name\":\"Trev Roper\"},{\"ip\":\"122.144.88.117\",\"player_id\":1013,\"player_name\":\"Lena MacFadzean\"},{\"ip\":\"233.122.238.221\",\"player_id\":1014,\"player_name\":\"Benton Allcorn\"},{\"ip\":\"119.248.57.152\",\"player_id\":1015,\"player_name\":\"Avis Moyler\"},{\"ip\":\"221.76.79.166\",\"player_id\":1016,\"player_name\":\"Marchall Rochewell\"},{\"ip\":\"214.194.20.254\",\"player_id\":1017,\"player_name\":\"Adele Bohl\"},{\"ip\":\"48.183.248.225\",\"player_id\":1018,\"player_name\":\"Barnett Mcall\"},{\"ip\":\"109.187.240.70\",\"player_id\":1019,\"player_name\":\"Frieda Pirrone\"},{\"ip\":\"66.106.114.58\",\"player_id\":1020,\"player_name\":\"Pattin Eringey\"},{\"ip\":\"200.181.48.97\",\"player_id\":1021,\"player_name\":\"Kalila Fewings\"},{\"ip\":\"105.180.133.32\",\"player_id\":1022,\"player_name\":\"Giacobo Beuscher\"},{\"ip\":\"195.82.242.168\",\"player_id\":1023,\"player_name\":\"Rozalin Hair\"},{\"ip\":\"193.33.103.229\",\"player_id\":1024,\"player_name\":\"Egon Beagan\"},{\"ip\":\"151.58.194.8\",\"player_id\":1025,\"player_name\":\"Owen Strotton\"},{\"ip\":\"151.18.78.73\",\"player_id\":1026,\"player_name\":\"Fernando Rosensaft\"},{\"ip\":\"245.93.20.235\",\"player_id\":1027,\"player_name\":\"Carleton Gwyther\"},{\"ip\":\"71.191.233.253\",\"player_id\":1028,\"player_name\":\"Kata Coll\"},{\"ip\":\"151.228.5.139\",\"player_id\":1029,\"player_name\":\"Rossie Hobben\"},{\"ip\":\"195.209.115.240\",\"player_id\":1030,\"player_name\":\"Stephanie Gookey\"},{\"ip\":\"137.210.104.16\",\"player_id\":1031,\"player_name\":\"Robyn Milazzo\"},{\"ip\":\"59.26.115.135\",\"player_id\":1032,\"player_name\":\"Tilda O'Lunney\"},{\"ip\":\"242.91.60.243\",\"player_id\":1033,\"player_name\":\"Nolan Kidney\"},{\"ip\":\"44.125.117.30\",\"player_id\":1034,\"player_name\":\"Jori Ottiwill\"},{\"ip\":\"58.22.186.13\",\"player_id\":1035,\"player_name\":\"Benito Graveson\"},{\"ip\":\"11.107.127.127\",\"player_id\":1036,\"player_name\":\"Zechariah Wrate\"},{\"ip\":\"145.156.130.86\",\"player_id\":1037,\"player_name\":\"Chelsae Napton\"},{\"ip\":\"128.91.240.205\",\"player_id\":1038,\"player_name\":\"Jeremy Heffernon\"},{\"ip\":\"175.36.181.120\",\"player_id\":1039,\"player_name\":\"Derk McAviy\"},{\"ip\":\"71.176.200.104\",\"player_id\":1040,\"player_name\":\"Constantin Mears\"},{\"ip\":\"61.160.45.31\",\"player_id\":1041,\"player_name\":\"Fitz Ballin\"},{\"ip\":\"174.111.209.131\",\"player_id\":1042,\"player_name\":\"Essy Bettles\"},{\"ip\":\"28.50.5.234\",\"player_id\":1043,\"player_name\":\"Gene Klemt\"},{\"ip\":\"56.180.162.177\",\"player_id\":1044,\"player_name\":\"Nikolai Arnopp\"},{\"ip\":\"206.250.61.47\",\"player_id\":1045,\"player_name\":\"Gustave Westhofer\"},{\"ip\":\"29.250.237.107\",\"player_id\":1046,\"player_name\":\"Simona Mayhow\"},{\"ip\":\"26.45.199.135\",\"player_id\":1047,\"player_name\":\"Cort Bainbridge\"},{\"ip\":\"31.46.35.158\",\"player_id\":1048,\"player_name\":\"Sibyl Vockins\"},{\"ip\":\"199.22.148.96\",\"player_id\":1049,\"player_name\":\"Andriette Gaze\"},{\"ip\":\"9.130.149.231\",\"player_id\":1050,\"player_name\":\"Shaughn De Simoni\"},{\"ip\":\"206.228.92.173\",\"player_id\":1051,\"player_name\":\"Nathaniel Hallowell\"},{\"ip\":\"236.38.31.50\",\"player_id\":1052,\"player_name\":\"Charley Dudill\"},{\"ip\":\"37.136.192.70\",\"player_id\":1053,\"player_name\":\"Cirstoforo Joblin\"},{\"ip\":\"107.27.178.139\",\"player_id\":1054,\"player_name\":\"Hyacinthia Kinastan\"},{\"ip\":\"29.183.6.183\",\"player_id\":1055,\"player_name\":\"Dur Lasselle\"},{\"ip\":\"118.222.29.232\",\"player_id\":1056,\"player_name\":\"Gay Chadburn\"},{\"ip\":\"232.200.180.97\",\"player_id\":1057,\"player_name\":\"Livvie Hawyes\"},{\"ip\":\"198.1.226.227\",\"player_id\":1058,\"player_name\":\"Aldrich MacVay\"},{\"ip\":\"64.39.185.31\",\"player_id\":1059,\"player_name\":\"Riva Rossant\"},{\"ip\":\"243.28.102.230\",\"player_id\":1060,\"player_name\":\"Johanna Reichartz\"},{\"ip\":\"252.49.141.204\",\"player_id\":1061,\"player_name\":\"Trent Gantlett\"},{\"ip\":\"215.235.104.14\",\"player_id\":1062,\"player_name\":\"Aryn Haskell\"},{\"ip\":\"61.59.180.203\",\"player_id\":1063,\"player_name\":\"Byrann Barock\"},{\"ip\":\"159.45.141.92\",\"player_id\":1064,\"player_name\":\"Gerda Cleugher\"},{\"ip\":\"212.156.139.166\",\"player_id\":1065,\"player_name\":\"Sonnie Guildford\"},{\"ip\":\"72.176.106.166\",\"player_id\":1066,\"player_name\":\"Vergil Borge\"},{\"ip\":\"199.23.37.30\",\"player_id\":1067,\"player_name\":\"Lurline Rocco\"},{\"ip\":\"204.124.16.39\",\"player_id\":1068,\"player_name\":\"Geoff Eddy\"},{\"ip\":\"144.103.183.21\",\"player_id\":1069,\"player_name\":\"Zea Leighton\"},{\"ip\":\"46.192.83.190\",\"player_id\":1070,\"player_name\":\"Leif Baden\"},{\"ip\":\"107.164.34.84\",\"player_id\":1071,\"player_name\":\"Quint Bidgod\"},{\"ip\":\"104.16.237.57\",\"player_id\":1072,\"player_name\":\"Talbot Cashell\"},{\"ip\":\"106.79.164.228\",\"player_id\":1073,\"player_name\":\"Sheridan Foulsham\"},{\"ip\":\"243.10.150.38\",\"player_id\":1074,\"player_name\":\"Camile Shrimplin\"},{\"ip\":\"194.5.112.254\",\"player_id\":1075,\"player_name\":\"Marcel Nayshe\"},{\"ip\":\"34.10.81.209\",\"player_id\":1076,\"player_name\":\"Lea Murrish\"},{\"ip\":\"136.17.217.222\",\"player_id\":1077,\"player_name\":\"Lucais Midson\"},{\"ip\":\"174.7.59.85\",\"player_id\":1078,\"player_name\":\"Zeb Rylatt\"},{\"ip\":\"219.151.0.93\",\"player_id\":1079,\"player_name\":\"Nertie Zuker\"},{\"ip\":\"183.255.62.100\",\"player_id\":1080,\"player_name\":\"Babara Henderson\"},{\"ip\":\"131.82.63.116\",\"player_id\":1081,\"player_name\":\"Electra Ridgley\"},{\"ip\":\"88.103.168.162\",\"player_id\":1082,\"player_name\":\"Jere Standingford\"},{\"ip\":\"255.154.216.143\",\"player_id\":1083,\"player_name\":\"Cyril Yellowlea\"},{\"ip\":\"25.97.116.94\",\"player_id\":1084,\"player_name\":\"Isadora Peegrem\"},{\"ip\":\"160.99.254.200\",\"player_id\":1085,\"player_name\":\"Caria Smewings\"},{\"ip\":\"113.47.135.39\",\"player_id\":1086,\"player_name\":\"Karena Kauffman\"},{\"ip\":\"187.131.135.163\",\"player_id\":1087,\"player_name\":\"Haywood Snowball\"},{\"ip\":\"143.25.215.3\",\"player_id\":1088,\"player_name\":\"Winslow Starcks\"},{\"ip\":\"78.107.115.65\",\"player_id\":1089,\"player_name\":\"Alis Ponton\"},{\"ip\":\"101.231.38.202\",\"player_id\":1090,\"player_name\":\"Marietta Lezemere\"},{\"ip\":\"22.30.43.226\",\"player_id\":1091,\"player_name\":\"Emilee Broadbridge\"},{\"ip\":\"77.208.184.143\",\"player_id\":1092,\"player_name\":\"Faye Beaument\"},{\"ip\":\"72.88.181.74\",\"player_id\":1093,\"player_name\":\"Shannah Beatson\"},{\"ip\":\"232.20.218.30\",\"player_id\":1094,\"player_name\":\"West Doy\"},{\"ip\":\"141.46.127.99\",\"player_id\":1095,\"player_name\":\"Chryste Wren\"},{\"ip\":\"120.63.46.231\",\"player_id\":1096,\"player_name\":\"Trumann Labba\"},{\"ip\":\"82.97.69.85\",\"player_id\":1097,\"player_name\":\"Anatollo Beckwith\"},{\"ip\":\"213.241.29.166\",\"player_id\":1098,\"player_name\":\"Konstanze Dunsford\"},{\"ip\":\"183.237.217.217\",\"player_id\":1099,\"player_name\":\"Raychel Roset\"},{\"ip\":\"165.19.12.241\",\"player_id\":1100,\"player_name\":\"Heindrick Ravenscroft\"}]},\"fields\":[{\"name\":\"player_id\",\"type\":\"int\"},{\"name\":\"player_name\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"gaming.GamingPlayers\",\"type\":\"record\"}"
}

func (r GamingPlayers) SchemaName() string {
	return "gaming.GamingPlayers"
}

func (_ GamingPlayers) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ GamingPlayers) SetInt(v int32)       { panic("Unsupported operation") }
func (_ GamingPlayers) SetLong(v int64)      { panic("Unsupported operation") }
func (_ GamingPlayers) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ GamingPlayers) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ GamingPlayers) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ GamingPlayers) SetString(v string)   { panic("Unsupported operation") }
func (_ GamingPlayers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *GamingPlayers) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Player_id}

		return w

	case 1:
		w := types.String{Target: &r.Player_name}

		return w

	case 2:
		w := types.String{Target: &r.Ip}

		return w

	}
	panic("Unknown field index")
}

func (r *GamingPlayers) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *GamingPlayers) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ GamingPlayers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ GamingPlayers) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ GamingPlayers) HintSize(int)                     { panic("Unsupported operation") }
func (_ GamingPlayers) Finalize()                        {}

func (_ GamingPlayers) AvroCRC64Fingerprint() []byte {
	return []byte(GamingPlayersAvroCRC64Fingerprint)
}

func (r GamingPlayers) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["player_id"], err = json.Marshal(r.Player_id)
	if err != nil {
		return nil, err
	}
	output["player_name"], err = json.Marshal(r.Player_name)
	if err != nil {
		return nil, err
	}
	output["ip"], err = json.Marshal(r.Ip)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *GamingPlayers) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["player_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Player_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for player_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["player_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Player_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for player_name")
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
	return nil
}
