package analytics

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAnalyticsEvents(t *testing.T) {
	// Set up a test router
	router := gin.New()
	router.GET("/analytics/events", GetAnalyticsEvents)

	// Create a test server to handle the request
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Set BASE_URL to the test server URL
	os.Setenv("BASE_URL", ts.URL)

	// Create a test request to the /analytics/events endpoint
	req, err := http.NewRequest("GET", ts.URL+"/analytics/events", nil)
	assert.NoError(t, err)

	// Create a test request with query parameters
	reqWithoutType, err := http.NewRequest("GET", ts.URL+"/analytics/events?room_id=65797aca2230de2e7bd21539", nil)
	assert.NoError(t, err)
	reqWithoutRoomId, err := http.NewRequest("GET", ts.URL+"/analytics/events?type=track.add.success", nil)
	assert.NoError(t, err)
	reqWithRequiredParams, err := http.NewRequest("GET", ts.URL+"/analytics/events?type=track.add.success&room_id=65797aca2230de2e7bd21539", nil)
	assert.NoError(t, err)

	tests := []struct {
		name         string
		request      *http.Request
		expectedCode int
	}{
		{
			name:         "Get analytics events without parameters",
			request:      req,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get analytics events without type param",
			request:      reqWithoutType,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get analytics events without room_id param",
			request:      reqWithoutRoomId,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Get analytics events with parameters",
			request:      reqWithRequiredParams,
			expectedCode: http.StatusOK,
		},
	}

	for index, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Perform the request
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, test.request)

			// without room_id, response contains and empty events list
			if index == 2 {
				var response map[string]interface{}
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Nil(t, err)

				if events, ok := response["events"].([]interface{}); ok {
					assert.Equal(t, len(events), 0)
				} else {
					t.Error("Expected 'events' key in the response")
				}

			}
			// Check the response code
			assert.Equal(t, test.expectedCode, resp.Code)

		})
	}
}
