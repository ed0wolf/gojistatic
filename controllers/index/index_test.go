package index

import (
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestWhenIndexFileDoesNotExist(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	Index("/some/path/that/doesnt/exist")(web.C{}, responseRecorder, nil)

	if responseRecorder.Code != 404 {
		t.Errorf("It should return a 200 status code but returned %d", responseRecorder.Code)
	}
}

/*func TestWhenIndexPageDoesExist(t *testing.T) {

}*/
