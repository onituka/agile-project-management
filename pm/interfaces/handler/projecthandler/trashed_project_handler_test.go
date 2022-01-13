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
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

func TestTrashedProjectHandlerTrashedProject(t *testing.T) {
	type fields struct {
		trashedProjectUsecase *mockprojectusecase.MockTrashedProjectUsecase
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

				in := &projectinput.TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}

				trashedAt := time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC)

				out := &projectoutput.TrashedProjectOutPut{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
					TrashedAt:         &trashedAt,
					CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
				}

				f.trashedProjectUsecase.EXPECT().TrashedProject(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "400-1-プロジェクトID不正",
			fileSuffix: "400-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac1800023",
				}).Context()

				in := &projectinput.TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac1800023",
				}
				err := apperrors.InvalidParameter

				f.trashedProjectUsecase.EXPECT().TrashedProject(ctx, in).Return(nil, err)
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

				in := &projectinput.TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.NotFound

				f.trashedProjectUsecase.EXPECT().TrashedProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "409-IDがすでにゴミ箱にある",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.Conflict

				f.trashedProjectUsecase.EXPECT().TrashedProject(ctx, in).Return(nil, err)
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

				in := &projectinput.TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.InternalServerError

				f.trashedProjectUsecase.EXPECT().TrashedProject(ctx, in).Return(nil, err)
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
				trashedProjectUsecase: mockprojectusecase.NewMockTrashedProjectUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewTrashedProjectHandler(f.trashedProjectUsecase)

			r := httptest.NewRequest(http.MethodPut, "/projects/024d71d6-1d03-11ec-a478-0242ac180002/trash-box", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.TrashedProject(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
