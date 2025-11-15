# Auckland Council API (unofficial)

Full list of available endpoints, for detailed description see below.

| Name                        | Endpoint        | Parameters              | Comments                            |
|:----------------------------|:----------------|:------------------------|:------------------------------------|
| Address                     | `/api/v1/addr`  | `addr`: partial address | Address Query                       |
| Rubbish and Recycling short | `/api/v1/rr`    | `addr`: partial address | Rubbish and Recycling, short format |
| Rubbish and Recycling       | `/api/v1/rrext` | `addr`: partial address | Rubbish and Recycling               |
| Healthcheck                 | `/healthcheck`  |                         | Returns OK if alive                 |


### Address search

* `/api/v1/addr`, parameter: `addr` - 

### Rubbish and Recycling

Two endpoints so far, both accepting `addr` parameter.

* `/api/v1/rr` - rubbish and recycling, returns the JSON of the following format:

      {
          "rubbish": "2020-02-25",
          "recycle": "2020-02-25",
          "address": "Britomart, CBD"
      }

* `/api/v1/rrext` - extended rubbish and recycling.  Returns the JSON in the following format:

      {
          "Collections": [
              {
                  "Day": "Monday 24 January",
                  "Date": "2020-01-24T00:00:00+13:00",
                  "Rubbish": true,
                  "Recycle": true
              },
              {
                  "Day": "Monday 31 January",
                  "Date": "2020-01-31T00:00:00+13:00",
                  "Rubbish": true,
                  "Recycle": false
              }
          ],
          "Address": {
              "ACRateAccountKey": "12342478585",
              "Address": "500 Queen Street, Auckland Central",
              "Suggestion": "500 Queen Street, Auckland Central"
          }
      }

Example:

```sh
$ curl --location --request GET 'https://<server>/api/v1/rr?addr=500%20Queen%20Street'
{"rubbish":"2020-02-24","recycle":"2020-02-24","address":"500 Queen Street, Auckland Central"}
```

### Integrating with Home Assistant

Assuming your aklapi API server running on localhost:5010, add the following
to your `configuration.yaml`:

```yaml
sensor:
  - platform: rest
    resource: "http://localhost:5010/api/v1/rr?addr=xx"
    name: Recycle
    scan_interval: 300
    value_template: "{{ value_json.recycle }}"
    method: GET
    unique_id: recycle_date

  - platform: rest
    resource: "http://localhost:5010/api/v1/rr?addr=xx"
    name: Food Scraps
    scan_interval: 300
    value_template: "{{ value_json.foodscraps }}"
    method: GET
    unique_id: foodscraps_date

  - platform: rest
    resource: "http://localhost:5010/api/v1/rr?addr=xx"
    name: Rubbish
    scan_interval: 300
    value_template: "{{ value_json.rubbish }}"
    method: GET
    unique_id: rubbish_date

```
