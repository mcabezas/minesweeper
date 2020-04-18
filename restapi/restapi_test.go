package restapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/game"
)

func Test_CreateGameHandler(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "create",
			in:             httptest.NewRequest("POST", "/games", strings.NewReader(`{"rows":5, "columns":5}`)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusCreated,
		},
		{
			name: "create with column value 0",
			in:             httptest.NewRequest("POST", "/games", strings.NewReader(`{"rows":5, "columns":0}`)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "create with row value 0",
			in:             httptest.NewRequest("POST", "/games", strings.NewReader(`{"rows":0, "columns":5}`)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "create with row & column values 0",
			in:             httptest.NewRequest("POST", "/games", strings.NewReader(`{"rows":0, "columns":0}`)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			game := game.NewFactory()
			CreateGameHandler(game)(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}
		})
	}
}

func Test_GetGameHandler(t *testing.T) {
	f := game.NewFactory()
	createdGame, _ := f.CreateGame(10, 10)

	r := mux.NewRouter()
	r.HandleFunc("/games/{gameID}", GetGameHandler(f)).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/games/" + createdGame.ID
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusOK)
	}

}

func Test_CannotReturnFakeGames(t *testing.T) {
	f := game.NewFactory()
	createdGame, _ := f.CreateGame(10, 10)

	r := mux.NewRouter()
	r.HandleFunc("/games/{gameID}", GetGameHandler(f)).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/games/" + createdGame.ID + "fakefake"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if status := resp.StatusCode; status != http.StatusNoContent {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusNoContent)
	}

}
