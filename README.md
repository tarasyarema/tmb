# TMB epic shitty API

Simple go API that uses [gin-gonic](https://github.com/gin-gonic/gin) to make async requests to the TMB API.

## Run

### `.env` file

Few things to notice:

- The `TMB_ID` and `TMB_KEY` correspond to the app id and app key given by the [TMB developer console](https://developer.tmb.cat/).
- The `GIN_MODE` variable can be `debug` or `release`, depending on wheter you are developing or deploying to production.

The file should be in the root directory and look like this

```bash
TMB_ID=xxx
TMB_KEY=xxx
GIN_MODE=debug|release
```

### Build and run

Two options:

1. `go build` and then run the executable `./tmb.exe`
2. `go run .`

By default the port is `8000` or `8080`, depending on the os. Just look at the fucking terminal.

### Deploy

...

### TODO

Setup docker, *bla, bla, ...*

## Documentation

### `GET /[pool|routines]`

There are two options:

- `/pool`
Creates a pool of async requests using [requests](https://github.com/jochasinga/requests).

- `/routines`
Create a goroutine for every pair and makes an individual async request (using the same library [requests](https://github.com/jochasinga/requests)) for each one, in their own routine.

Both take the same query params and the responses are identical.

#### Query params

#### `data`

String of comma separated integers. The format of a pair is `bus_line,stop_id`.

#### `sync` (only used with `/routines`)

Boolean (`true` or `false`). Default `false`.

Defines if every request to the TMB api is made via `requests.GetAsync` or `requests.Get`. 

#### Example

##### Request

```bash
GET /routines?data=54,208,13,37&sync=true
```

##### Response

```bash
HTTP/1.1 200 OK

{
  "data": [
    {
      "Time": 23,
      "Meta": {
        "Line": 54,
        "Stop": 208
      }
    },
    {
      "Time": 420,
      "Meta": {
        "Line": 13,
        "Stop": 37
      }
    }
  ],
  "message": "OK"
}
```

*The time (`data.*.Time`) is given in seconds.*

## Tests

***Suerte en la vida**, xd*
