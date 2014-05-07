package assets

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const fakeHandlerReturnValue string = "RSVP"

type FakeHandler struct {
	isCalled bool
}

func (handler *FakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.isCalled = true
	w.Write([]byte(fakeHandlerReturnValue))
}

type FakeAssetsRetriever struct {
	isCalled bool
	path     string
}

func (retriever *FakeAssetsRetriever) Retrieve(assetsPath string, w http.ResponseWriter) {
	retriever.isCalled = true
	retriever.path = assetsPath
}

var handlerFunc http.Handler
var fakeHandler FakeHandler
var fakeAssetsRetriever FakeAssetsRetriever
var assetHandler AssetsHandler
var responseRecorder *httptest.ResponseRecorder

func initAssetsTest() {
	fakeHandler = FakeHandler{}
	fakeAssetsRetriever = FakeAssetsRetriever{}
	assetHandler = AssetsHandler{&fakeAssetsRetriever}
	handlerFunc = assetHandler.HandleAssets(&fakeHandler)
	responseRecorder = httptest.NewRecorder()
}

func TestWhenRequestUriIsRequestingAnAsset(t *testing.T) {
	initAssetsTest()
	var request = &http.Request{RequestURI: "/assets/path/to/some/file.js"}

	handlerFunc.ServeHTTP(responseRecorder, request)

	if fakeHandler.isCalled {
		t.Error("should not have been called fake handler")
	}
	if !fakeAssetsRetriever.isCalled {
		t.Errorf("should have tried to get assets from retriever")
	}
	if fakeAssetsRetriever.path != "/path/to/some/file.js" {
		t.Errorf("should have passed %v as path but got %v",
			"/path/to/some/file.js", fakeAssetsRetriever.path)
	}
}

func TestWhenRequestUriIsNotRequestingAnAsset(t *testing.T) {
	initAssetsTest()
	var request = &http.Request{RequestURI: "/"}

	handlerFunc.ServeHTTP(responseRecorder, request)

	if !fakeHandler.isCalled {
		t.Error("should have called fake handler")
	}
	if responseRecorder.Body.String() != fakeHandlerReturnValue {
		t.Errorf("should have written %v in response but found %v", fakeHandlerReturnValue, responseRecorder.Body.String())
	}
}
