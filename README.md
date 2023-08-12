# Consuming weather API

In this project we consume an API from, [weatherapi](https://www.weatherapi.com/). Sign up and get your API key. Create a `.env` file as `sample.env`.

## Run

`make run` or `go run cmd/main.go`

## Resource

- [how-to-make-an-http-server-in-go](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)
- [how-to-make-http-requests-in-go](https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go)

## Request and Response

```sh
# request
curl --location 'localhost:3000/?city=Accra'
```

```json
{
  "location": {
    "name": "Accra",
    "region": "Greater Accra",
    "country": "Ghana",
    "lat": 5.55,
    "lon": -0.22,
    "tz_id": "Africa/Accra",
    "localtime_epoch": 1691876472,
    "localtime": "2023-08-12 21:41"
  },
  "current": {
    "last_updated_epoch": 1691875800,
    "last_updated": "2023-08-12 21:30",
    "temp_c": 25.0,
    "temp_f": 77.0,
    "is_day": 0,
    "condition": {
      "text": "Partly cloudy",
      "icon": "//cdn.weatherapi.com/weather/64x64/night/116.png",
      "code": 1003
    },
    "wind_mph": 13.6,
    "wind_kph": 22.0,
    "wind_degree": 220,
    "wind_dir": "SW",
    "pressure_mb": 1013.0,
    "pressure_in": 29.91,
    "precip_mm": 0.0,
    "precip_in": 0.0,
    "humidity": 89,
    "cloud": 50,
    "feelslike_c": 27.3,
    "feelslike_f": 81.1,
    "vis_km": 10.0,
    "vis_miles": 6.0,
    "uv": 1.0,
    "gust_mph": 18.1,
    "gust_kph": 29.2
  }
}
```

```sh
curl --location 'localhost:3000/format?city=Accra'

```

```json
{
  "location": {
    "name": "Accra",
    "region": "Greater Accra",
    "country": "Ghana",
    "lat": 5.55,
    "lon": -0.22,
    "tz_id": "Africa/Accra",
    "localtime": "2023-08-12 21:41"
  },
  "current": {
    "last_updated": "2023-08-12 21:30",
    "temp_c": 25,
    "condition": {
      "text": "Partly cloudy",
      "icon": "//cdn.weatherapi.com/weather/64x64/night/116.png",
      "code": 1003
    },
    "wind_kph": 22,
    "wind_degree": 220,
    "wind_dir": "SW",
    "pressure_mb": 1013,
    "precip_mm": 0,
    "humidity": 89,
    "cloud": 50,
    "vis_km": 10,
    "uv": 1,
    "gust_kph": 29.2
  }
}
```
