package api

type LocationType struct {
	Name       string  `json:"name"`
	Region     string  `json:"region"`
	Country    string  `json:"country"`
	Latitude   float32 `json:"lat"`
	Longitude  float32 `json:"lon"`
	TimeZoneId string  `json:"tz_id"`
	LocalTime  string  `json:"localtime"`
}

type ConditionType struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type CurrentType struct {
	LastUpdated                 string        `json:"last_updated"`
	TemperatureCelsius          float32       `json:"temp_c"`
	Condition                   ConditionType `json:"condition"`
	WindKilometrePerHour        float32       `json:"wind_kph"`
	WindDegree                  int           `json:"wind_degree"`
	WindDirection               string        `json:"wind_dir"`
	PressureInMillibar          float32       `json:"pressure_mb"`
	PrecipitationInMillimetres  float32       `json:"precip_mm"`
	Humidity                    int           `json:"humidity"`
	Cloud                       int           `json:"cloud"`
	VisibilityInKilometres      float32       `json:"vis_km"`
	UV                          float32       `json:"uv"`
	GustSpeedInKiloMeterPerHour float32       `json:"gust_kph"`
}

type WeatherResponse struct {
	Location LocationType `json:"location"`
	Current  CurrentType  `json:"current"`
}
