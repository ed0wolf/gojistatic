package index

import (
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Index(indexPagePath string) func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile(indexPagePath)

		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.Write(content)
	}
}
