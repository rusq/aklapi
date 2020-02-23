# Auckland Rubbish and Recycling days webservice

### Usage

This simple app exposes 2 endpoints so far, both accepting `addr` parameter.

* /rr/ - rubbish and recycling, returns the JSON of the following format:

      {
          "rubbish": "2020-02-25",
          "recycle": "2020-02-25",
          "address": "Britomart, CBD"
      }

* /rrext/ - extended rubbish and recycling.  Returns the JSON in the following format:

      {
          "Collections": [
              {
                  "Day": "Monday 24 January",
                  "Date": "2020-02-24T00:00:00+13:00",
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
$ curl --location --request GET 'https://<server>/api/v1/rr/?addr=500%20Queen%20Street'
{"rubbish":"2020-02-24","recycle":"2020-02-24","address":"500 Queen Street, Auckland Central"}
```

### Integrating with Home Assistant

Add the following to your `configuration.yaml`:

```yaml
sensor:
  - platform: rest
    resource: https://your_server/api/v1/rr?addr=500%20Mystreet
    name: Recycle
    value_template: '{{ value_json.recycle }}'
```