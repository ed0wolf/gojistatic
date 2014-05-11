package ping

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Ping(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1.0.0")
}
