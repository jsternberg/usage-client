# Enterprise Client

## Registering a Server

Here is the registration workflow for InfluxDB:

When an app wakes up (influxdb, chronograf, etc…) it posts the following JSON to `/api/v1/servers`:


```json
# POST /api/v1/servers:
{
  "cluster_id" : "abc123",
  "server_id" : "abc123",
  "host" : "111.222.333.444:1234",
  "product" : "influxdb",
  "version" : "1.2.3"
}
```

This should happen every time the application starts.

If there is a token stored locally (more on that in a minute). It should be sent along with __ALL__ requests to Enterprise. The token can either be on the query string (e.g. `/api/v1/servers?token=123456`) or using the "X-Authorization" header on the request.

If there is no token stored locally (it is up to each app to store this token), when the application starts up they user should be prompted to register with Enterprise.

The user should be sent to:

```bash
https://enterprise.influxdata.com/start?cluster_id=abc123&product=influxdb&redirect_url=http://some.host:port
```

The host (defaults to `enterprise.influxdata.com`) should be configurable in each of the apps.

### Parameters:

* `cluster_id`: should be the id of the cluster you want to register
* `product`: should be the name of the product; chronograf, influxdb, etc...
* `redirect_url` (optional): should be the url to redirect the user to after they complete their register or sign in if they already have an account.

Once registration is complete Enterprise will redirect to the `redirect_url` with the query param "token" set. The app should save this token for making future requests to Enterprise. If there is no `redirect_url` then the user will be taken to their profile page where they'll be shown their token.

Does that all make sense? I hope so. :) Let me know if there are questions.

## Posting Product Stats:

Ideally when posting product stats data to Enterprise you will pass the "token" either using the `token` query param or the "X-Authorization" header. If no token is set then the data will __NOT__ be associated with any organization.

#### Example Body

```json
# POST /api/v1/stats/:product
{
  "cluster_id": "abc123",
  "server_id": "abc123",
  "stats": [{
    "name": "engine",
    "tags": {
      "path": "/home/philip/.influxdb/data/_internal/monitor/1",
      "version": "bz1"
    },
    "values": {
      "blks_write": 39,
      "blks_write_bytes": 2421,
      "blks_write_bytes_c": 2202,
      "points_write": 39,
      "points_write_dedupe": 39
    }
  }]
}
```

The `:product` param in the API URL can be one of the following:

* influxdb
* kapicator
* telegraf
* chronograf

## Posting Usage Stats:

Usage stats are anonymous stats sent to Enterprise every 12 hours.

#### Example Body

```json
# POST /api/v1/usage/:product
[{
  "tags": {
    "version": "0.9.5",
    "arch": "amd64",
    "os": "linux"
  },
  "values": {
    "cluster_id": "23423",
    "server_id": "1",
    "num_databases": 3,
    "num_measurements": 2342,
    "num_series": 87232
  }
}]
```

The `:product` param in the API URL can be one of the following:

* influxdb
* kapicator
* telegraf
* chronograf

## API Errors

It's possible that if you don't send the Enterprise application required or valid data you might get some errors.

In the case that you do get errors back from Enterprise they will come with a none `success` code, most likely a `422` or `500`. You will also receive a JSON payload that looks something like the following:

```json
{
  "errors": {
    "cluster_id": [
      "ClusterID can not be blank."
    ],
    "host": [
      "Host can not be blank."
    ],
    "product": [
      "Product can not be blank."
    ],
    "server_id": [
      "ServerID can not be blank."
    ],
    "version": [
      "Version can not be blank."
    ]
  }
}
```

The exact errors you might receive will depend on the end-point you hit, but you get the idea.

### 500 Errors

If you get a 500 error it will most likely look like the following:

```json
{"error":"json: cannot unmarshal number into Go value of type string"}
``
