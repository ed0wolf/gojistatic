package assets

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type AssetsRetriever interface {
	Retrieve(assetPath string, w http.ResponseWriter)
}

type FileAssetsRetriever struct {
	RootDir string
}

func (assetsRetriever *FileAssetsRetriever) Retrieve(assetPath string, w http.ResponseWriter) {
	content, err := ioutil.ReadFile(assetsRetriever.RootDir + "/" + assetPath)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	mime, mimeErr := getAssetMime(assetPath)
	if mimeErr == nil {
		w.Header().Set("Content-Type", mime)
	}

	w.Write(content)
}

func getAssetMime(assetPath string) (string, error) {
	if strings.HasSuffix(assetPath, ".js") {
		return "text/javascript", nil
	} else if strings.HasSuffix(assetPath, ".json") {
		return "application/json", nil
	} else if strings.HasSuffix(assetPath, ".css") {
		return "text/css", nil
	} else if strings.HasSuffix(assetPath, ".xml") {
		return "text/xml", nil
	} else if strings.HasSuffix(assetPath, ".html") || strings.HasSuffix(assetPath, ".htm") {
		return "text/html", nil
	}
	return "", errors.New("Couldn't find MIME for" + assetPath)
}
