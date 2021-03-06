# TMB epic shitty API

Simple go API that uses [gin-gonic](https://github.com/gin-gonic/gin) to make async requests to the TMB API.

## Run

### `.env` file

Few things to notice:

- The `TMB_ID` and `TMB_KEY` correspond to the app id and app key given by the [TMB developer console](https://developer.tmb.cat/).
- The `GIN_MODE` variable can be `debug` or `release`, depending on wheter you are developing or deploying to production.
- The `PORT` variable is the port used by gin to run the server. This should be `8080` (hardcoded in the `Dockerfile`). YOu can change it but make sure that it makes sense overall.

The file should be in the root directory and look like this

```bash
TMB_ID=${TMB_API_ID}
TMB_KEY=${TMB_API_KEY}
GIN_MODE=${debug|release}
PORT=${PORT}
```

### Build and run

#### Non-docker

`go build -o tmb .` and then `./tmb`.

or just

`go run .`


#### With docker

1. `docker build -t tmb .`
2. `docker run -p 8080:8080 --env-file .env tmb:latest` 


Then you can begin making requests to `localhost:${PORT}`.

## Documentation

### `GET /[pool|routines]`

There are two options:

- `/pool`
Creates a pool of async requests using [requests](https://github.com/jochasinga/requests).

- `/routines`
Create a goroutine for every pair and makes an individual async or sync request (using the same library [requests](https://github.com/jochasinga/requests)) for each one, in their own routine.

Both take the same query params and the responses are identical.

#### Query params

#### `data`

String of comma separated integers. The format of a pair is `bus_line,stop_id`.

#### `sync` (only used with `/routines`)

Boolean (`true` or `false`). Default `false`.

Defines if every request to the TMB api is made via `requests.GetAsync` or `requests.Get`. 

#### Example

```bash
GET /routines?data=D40,1554,H8,1554,27,1554 HTTP/1.1

{
   "data": [
      {
         "time": 427,
         "meta": {
            "line": "D40",
            "stop": "1554"
         },
         "elapsed": 88
      },
      {
         "time": 10,
         "meta": {
            "line": "H8",
            "stop": "1554"
         },
         "elapsed": 447
      },
      {
         "time": 30,
         "meta": {
            "line": "27",
            "stop": "1554"
         },
         "elapsed": 80
      }
   ],
   "elapsed": 447,
   "message": "OK"
}
```

*The time (`data.*.time`) parameter of the response JSON is given in seconds.*

*The elapsed times (`data.*.elapsed` and `data.elapsed`) parametere of the response JSON is given in miliseconds.*

### TODO

- [x] Docker
- [ ] Automatic deployment
- [ ] Tests
- [ ] CI/CD

### Deploy

***Note**: Now is being manually deployed to Google Cloud Run via the Google container registry.*

## Tests

***Suerte en la vida**, xd*

