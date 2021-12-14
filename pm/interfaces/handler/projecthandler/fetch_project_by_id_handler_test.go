package projecthandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projecthandler/mockprojectusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

func TestFetchProjectByIDHandlerFetchProjectByID(t *testing.T) {
	type fields struct {
		fetchProjectByIDUsecase *mockprojectusecase.MockFetchProjectByIDUsecase
	}
	tests := []struct {
		name           string
		fileSuffix     string
		prepareMock    func(f *fields)
		prepareRequest func(r *http.Request)
	}{
		{
			name:       "200-正常",
			fileSuffix: "200",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectusecase.FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}

				out := &projectusecase.FetchProjectByIDOutput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
					CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
				}

				f.fetchProjectByIDUsecase.EXPECT().FetchProjectByID(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "400-プロジェクトID不正",
			fileSuffix: "400",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac1800023",
				}).Context()

				in := &projectusecase.FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac1800023",
				}

				err := apperrors.InvalidParameter

				f.fetchProjectByIDUsecase.EXPECT().FetchProjectByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac1800023",
				})
			},
		},
		{
			name:       "404-IDが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectusecase.FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}

				err := apperrors.NotFound

				f.fetchProjectByIDUsecase.EXPECT().FetchProjectByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "500-DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectusecase.FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}

				err := apperrors.InternalServerError

				f.fetchProjectByIDUsecase.EXPECT().FetchProjectByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				fetchProjectByIDUsecase: mockprojectusecase.NewMockFetchProjectByIDUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewFetchProjectByIDHandler(f.fetchProjectByIDUsecase)

			r := httptest.NewRequest(http.MethodGet, "/projects/{projectID}", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.FetchProjectByID(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
