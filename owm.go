package owm

import (
	"fmt"
	"net/http"
)

type API interface {
	Current(city string) (interface{}, error)
}

type Doer func(http.Request) (*http.Response, error)

func New(do Doer, apiKey string) API {
	return api{
		key: apiKey,
		do:  do,
	}
}

const OWMAPI string = "api.openweathermap.org"

type api struct {
	key string
	do  Doer
}

func (a api) Current(city string) (interface{}, error) {
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s/data/2.5/weather", OWMAPI), nil)
	q := r.URL.Query()
	q.Add("appid", a.key)
	r.URL.RawQuery = q.Encode()

	return a.do(*r)
}
