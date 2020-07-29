package corona

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJP(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	tests := map[string]struct {
		want Country
	}{
		"success": {
			want: Country{
				Name:        "Japan",
				Cases:       1,
				TodayCases:  2,
				Deaths:      3,
				TodayDeaths: 4,
			},
		},
	}

	for name, test := range tests {
		corona := &Corona{endpoint: srv.URL}

		t.Run(name, func(t *testing.T) {
			country, err := corona.JP()
			if err != nil {
				t.Fatal(err)
			}

			if *country != test.want {
				t.Errorf(`want="%v" err="%v"`, test.want, *country)
			}
		})
	}
}

func TestWorldRanking(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	tests := map[string]struct {
		want []Country
	}{
		"success": {
			want: []Country{
				{
					Name:        "Japan",
					Cases:       1,
					TodayCases:  2,
					Deaths:      3,
					TodayDeaths: 4,
				},
				{
					Name:        "USA",
					Cases:       5,
					TodayCases:  6,
					Deaths:      7,
					TodayDeaths: 8,
				},
			},
		},
	}

	for name, test := range tests {
		corona := &Corona{endpoint: srv.URL}

		t.Run(name, func(t *testing.T) {
			countries, err := corona.WorldRanking()
			if err != nil {
				t.Fatal(err)
			}

			for i, country := range countries {
				if *country != test.want[i] {
					t.Errorf(`want="%v" err="%v"`, test.want[i], *country)
				}
			}
		})
	}
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v3/covid-19/countries/JP", jpMock)
	handler.HandleFunc("/v3/covid-19/countries", rankingMock)

	srv := httptest.NewServer(handler)

	return srv
}

func jpMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`
{"country": "Japan",
"cases": 1,
"todayCases": 2,
"deaths": 3,
"todayDeaths": 4}
	`))
}

func rankingMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`
[{"country": "Japan",
"cases": 1,
"todayCases": 2,
"deaths": 3,
"todayDeaths": 4},
{"country": "USA",
"cases": 5,
"todayCases": 6,
"deaths": 7,
"todayDeaths": 8}]
	`))
}
