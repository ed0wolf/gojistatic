package assets

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
)

const expectedFileBody string = "BODY"

var rootDir, assetPath, expectedContentType string
var writer *httptest.ResponseRecorder
var retriever *FileAssetsRetriever

func initRetrieverTest() {
	rootDir = os.TempDir()
	expectedContentType = "text/javascript"
	assetPath = fmt.Sprintf("/%s.js", uuid.New())
	writer = httptest.NewRecorder()
	retriever = &FileAssetsRetriever{rootDir}
}

func TestWhenTheAssetDoesNotExist(t *testing.T) {
	initRetrieverTest()

	retriever.Retrieve(assetPath, writer)

	if writer.Code != 404 {
		t.Errorf("should have returned 404 but instead returned: %s", writer.Code)
	}
}

func TestWhenTheReqeustedAssetExists(t *testing.T) {
	initRetrieverTest()
	createAsset(rootDir+assetPath, expectedFileBody, t)

	retriever.Retrieve(assetPath, writer)

	body := writer.Body.String()
	if writer.Code != 200 {
		t.Errorf("should have returned 200 but instead returned: %s", writer.Code)
	}
	if body != expectedFileBody {
		t.Errorf("should have written %s to responseWriter but wrote: %s", expectedFileBody, body)
	}

	actualContentType := writer.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("should have set content-type header to %s but was instead %s", expectedContentType, actualContentType)
	}

	deleteAsset(rootDir + assetPath)
}

func TestWhenGettingMimeForJsAssetPath(t *testing.T) {
	mime, err := getAssetMime("jquery.js")

	if err != nil {
		t.Error("should not have returned an error")
	}

	if mime != "text/javascript" {
		t.Errorf("should have returned 'text/javascript' but returned: %s", mime)
	}
}

func TestWhenGettingMimeForCssAssetPath(t *testing.T) {
	mime, err := getAssetMime("main.css")

	if err != nil {
		t.Error("should not have returned an error")
	}

	if mime != "text/css" {
		t.Errorf("should have returned 'text/css' but returned: %s", mime)
	}
}

func TestWhenGettingMimeForXmlAssetPath(t *testing.T) {
	mime, err := getAssetMime("interesting.xml")

	if err != nil {
		t.Error("should not have returned an error")
	}

	if mime != "text/xml" {
		t.Errorf("should have returned 'text/xml' but returned: %s", mime)
	}
}

func TestWhenGettingMimeForHtmlAssetPath(t *testing.T) {
	mime, err := getAssetMime("interesting.html")

	if err != nil {
		t.Error("should not have returned an error")
	}

	if mime != "text/html" {
		t.Errorf("should have returned 'text/html' but returned: %s", mime)
	}
}

func TestWhenGettingMimeForHtmAssetPath(t *testing.T) {
	mime, err := getAssetMime("interesting.htm")

	if err != nil {
		t.Error("should not have returned an error")
	}

	if mime != "text/html" {
		t.Errorf("should have returned 'text/html' but returned: %s", mime)
	}
}

func TestWhenGettingMimeForJsonAssetPath(t *testing.T) {
	mime, err := getAssetMime("interesting.json")

	if err != nil {
		t.Error("should not have returned an error")
	}

	if mime != "application/json" {
		t.Errorf("should have returned 'application/json' but returned: %s", mime)
	}
}

func TestWhenGettingUnknownMimeType(t *testing.T) {
	mime, err := getAssetMime("mime.fyjhfgjhfgjh")

	if mime != "" {
		t.Errorf("should have returned '' but returned: %s", mime)
	}

	if err == nil {
		t.Error("should have returned an error but returned nil")
	}
}

func createAsset(path, content string, t *testing.T) {
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	} else {
		file.WriteString(content)
	}
}

func deleteAsset(path string) {
	os.Remove(path)
}
