package API_test

//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//// endpoints_test.go
//func TestMetricsHandler(t *testing.T) {
//	testingStruct := []struct{
//		routeVariable string
//		shouldPass bool
//	}{
//		{"goroutines", true},
//		{"heap", true},
//		{"counters", true},
//		{"queries", true},
//		{"adhadaeqm3k", false},
//	}
//
//	for _, testingVar := range testingStruct {
//		path := fmt.Sprintf("/metrics/%s", testingVar.routeVariable)
//		req, err := http.NewRequest("GET", path, nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		rr := httptest.NewRecorder()
//
//		// Need to create a router that we can pass the request through so that the vars will be added to the context
//		router := mux.NewRouter()
//
//
//		router.HandleFunc("/metrics/{type}", HealthCheckHandler)
//		router.ServeHTTP(rr, req)
//
//		// In this case, our MetricsHandler returns a non-200 response
//		// for a route variable it doesn't know about.
//		if rr.Code == http.StatusOK && !testingVar.shouldPass {
//			t.Errorf("handler should have failed on routeVariable %s: got %v want %v",
//				testingVar.routeVariable, rr.Code, http.StatusOK)
//		}
//	}
//}
//func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
//	param := mux.Vars(r)
//	fmt.Println("params: ", param["type"])
//
//	// A very simple health check.
//	w.Header().Set("Content-Type", "application/json")
//	if param["type"] == "adhadaeqm3k"{
//		w.WriteHeader(http.StatusBadRequest)
//		io.WriteString(w, `{"alive": false}`)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//
//	// In the future we could report back on the status of our DB, or our cache
//	// (e.g. Redis) by performing a simple PING, and include them in the response.
//	io.WriteString(w, `{"alive": true}`)
//}


//func TestHealthCheckHandler(t *testing.T) {
//	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//	// pass 'nil' as the third parameter.
//	req, err := http.NewRequest("GET", "/intents/createagent/Users%2Fvee%2FDesktop%2FCreateAgents-1ea24b69d927.json", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(GetIntentsHandler)
//
//	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//	// directly and pass in our Request and ResponseRecorder.
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code is what we expect.
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	// Check the response body is what we expect.
//	expected := `{"alive": true}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
//	}
//}