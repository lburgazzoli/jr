{{$id:=uuid}}{{add_v_to_list "customers_id_list" $id}}{
  "id": "{{$id}}",
  "first_name": "{{name}}",
  "last_name": "{{surname}}",
  "email": "{{email}}",
  "state": "{{state}}",
  "zip_code": "{{zip}}"
  "street_address": "{{city}}, {{street}} {{building 2}}, {{zip}}",
  "phone_number": "{{land_prefix}} {{regex "[0-9]{7}"}}",
  "country": "United States",
  "country_code: "US"
}