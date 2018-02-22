package gowu

type ForecastResponse struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Forecast int `json:"forecast"`
		} `json:"features"`
	} `json:"response"`
	Forecast Forecast `json: "forecast"`
}

type Forecast struct {
	TxtForecast struct {
		Date        string `json:"date"`
		Forecastday []struct {
			Period        int    `json:"period"`
			Icon          string `json:"icon"`
			IconURL       string `json:"icon_url"`
			Title         string `json:"title"`
			Fcttext       string `json:"fcttext"`
			FcttextMetric string `json:"fcttext_metric"`
			Pop           string `json:"pop"`
		} `json:"forecastday"`
	} `json:"txt_forecast"`
	Simpleforecast struct {
		Forecastday []struct {
			Date struct {
				Epoch          string `json:"epoch"`
				Pretty         string `json:"pretty"`
				Day            int    `json:"day"`
				Month          int    `json:"month"`
				Year           int    `json:"year"`
				Yday           int    `json:"yday"`
				Hour           int    `json:"hour"`
				Min            string `json:"min"`
				Sec            int    `json:"sec"`
				Isdst          string `json:"isdst"`
				Monthname      string `json:"monthname"`
				MonthnameShort string `json:"monthname_short"`
				WeekdayShort   string `json:"weekday_short"`
				Weekday        string `json:"weekday"`
				Ampm           string `json:"ampm"`
				TzShort        string `json:"tz_short"`
				TzLong         string `json:"tz_long"`
			} `json:"date"`
			Period int `json:"period"`
			High   struct {
				Fahrenheit string `json:"fahrenheit"`
				Celsius    string `json:"celsius"`
			} `json:"high"`
			Low struct {
				Fahrenheit string `json:"fahrenheit"`
				Celsius    string `json:"celsius"`
			} `json:"low"`
			Conditions string `json:"conditions"`
			Icon       string `json:"icon"`
			IconURL    string `json:"icon_url"`
			Skyicon    string `json:"skyicon"`
			Pop        int    `json:"pop"`
			QpfAllday  struct {
				In float64 `json:"in"`
				Mm int     `json:"mm"`
			} `json:"qpf_allday"`
			QpfDay struct {
				In interface{} `json:"in"`
				Mm interface{} `json:"mm"`
			} `json:"qpf_day"`
			QpfNight struct {
				In float64 `json:"in"`
				Mm int     `json:"mm"`
			} `json:"qpf_night"`
			SnowAllday struct {
				In float64 `json:"in"`
				Cm float64 `json:"cm"`
			} `json:"snow_allday"`
			SnowDay struct {
				In interface{} `json:"in"`
				Cm interface{} `json:"cm"`
			} `json:"snow_day"`
			SnowNight struct {
				In float64 `json:"in"`
				Cm float64 `json:"cm"`
			} `json:"snow_night"`
			Maxwind struct {
				Mph     int    `json:"mph"`
				Kph     int    `json:"kph"`
				Dir     string `json:"dir"`
				Degrees int    `json:"degrees"`
			} `json:"maxwind"`
			Avewind struct {
				Mph     int    `json:"mph"`
				Kph     int    `json:"kph"`
				Dir     string `json:"dir"`
				Degrees int    `json:"degrees"`
			} `json:"avewind"`
			Avehumidity int `json:"avehumidity"`
			Maxhumidity int `json:"maxhumidity"`
			Minhumidity int `json:"minhumidity"`
		} `json:"forecastday"`
	} `json:"simpleforecast"`
}
