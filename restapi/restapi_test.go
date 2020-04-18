package restapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

			in: httptest.NewRequest("POST", "/games", strings.NewReader(`{"rows":5, "columns":5}`)),
			out: httptest.NewRecorder(),
			expectedStatus: http.StatusCreated,
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
