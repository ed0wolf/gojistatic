package ping

import (
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestWhenGettingPingItShouldReturn200(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	Ping(web.C{}, responseRecorder, nil)

	if responseRecorder.Code != 200 {
		t.Errorf("It should return a 200 status code but returned %d", responseRecorder.Code)
	}
}
