package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"github.com/s2ar/unstable/internal/application"
	"github.com/s2ar/unstable/internal/service/opendota"
)

func TestHandler_teamGetInfo(t *testing.T) {
	type mockBehavior func(r *opendota.MockOpendota)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "408 Error",
			mockBehavior:         func(r *opendota.MockOpendota) {},
			expectedStatusCode:   408,
			expectedResponseBody: `{"message":"Request Timeout"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			app := application.NewMockApplication(c)

			repo := opendota.NewMockOpendota(c)
			test.mockBehavior(repo)

			r := mux.NewRouter()
			generalMiddleware := []mux.MiddlewareFunc{
				application.WithApp(app),
			}
			api := r.PathPrefix("/api/").Subrouter()
			api.Use(generalMiddleware...)
			GenRouting(api)

			//r.Path("/team/top").Methods(http.MethodGet).HandlerFunc(TeamGetInfo)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/team/top/", nil)

			// Make Request
			api.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)

			/*
				services := &service.Service{Authorization: repo}
				handler := Handler{services}

				// Init Endpoint
				r := gin.New()
				r.POST("/sign-up", handler.signUp)

				// Create Request
				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/sign-up",
					bytes.NewBufferString(test.inputBody))

				// Make Request
				r.ServeHTTP(w, req)

				// Assert
				assert.Equal(t, w.Code, test.expectedStatusCode)
				assert.Equal(t, w.Body.String(), test.expectedResponseBody)
			*/
		})
	}

}

/*
// TestTeamGetInfo tests handler.handleTeamGetInfo
func TestTeamGetInfo(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := application.NewMockApplication(ctrl)

	serviceOpendota := app.ServiceOpendota()

	serviceOpendota.GetTopTeam()

	objectService := NewMockObjectService(ctrl)
	testerService := NewMockTesterService(ctrl)

	callbackHandler := New(objectService, testerService)

	ts := httptest.NewServer(http.HandlerFunc(callbackHandler.Post))
	defer ts.Close()

	var testObjectID uint = 1

		tests := []struct {
			name    string
			payload string
			expect  func()
			assert  func()
		}{
			{
				name:    "object offline",
				payload: fmt.Sprintf(`{"object_ids": [%d]}`, testObjectID),
				expect: func() {
					testerService.EXPECT().GetObject(ctx, testObjectID).Return(model.TesterObject{
						ID:     testObjectID,
						Online: false,
					}, nil)
				},
			},
			{
				name:    "object online",
				payload: `{"object_ids": [1]}`,
				expect: func() {
					testerService.EXPECT().GetObject(ctx, testObjectID).Return(model.TesterObject{
						ID:     testObjectID,
						Online: true,
					}, nil)
					objectService.EXPECT().UpdateObject(ctx, &model.DBObject{
						ID: testObjectID,
					}).Return(nil)
				},
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				if tc.expect != nil {
					tc.expect()
				}

				if tc.assert != nil {
					tc.assert()
				}
			})

			res, err := http.Post(ts.URL, "application/json", strings.NewReader(tc.payload))
			if err != nil {
				t.Error(err)
			}
			defer res.Body.Close()

			content, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, content, []byte("ok"))
		}
}*/
