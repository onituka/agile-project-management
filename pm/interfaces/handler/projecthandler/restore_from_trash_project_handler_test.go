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

func TestRestoreFromTrashProjectHandlerRestoreFromTrashProject(t *testing.T) {
	type fields struct {
		restoreFromTrashProjectUsecase *mockprojectusecase.MockRestoreFromTrashProjectUsecase
	}

	tests := []struct {
		name           string
		fileSuffix     string
		prepareMock    func(f *fields)
		prepareRequest func(r *http.Request)
	}{
		{
			name:       "正常",
			fileSuffix: "200",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.RestoreFromTrashProjectIDInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "024d78d6-1d03-11ec-a478-0242ac180002",
				}

				trashedAt := time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC)

				out := &projectoutput.RestoreFromTrashProjectIDOutPut{
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

				f.restoreFromTrashProjectUsecase.EXPECT().RestoreFromTrashProject(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトID不正",
			fileSuffix: "400",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002xxx",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.RestoreFromTrashProjectIDInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002xxx",
					ProductID: "024d78d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.InvalidParameter

				f.restoreFromTrashProjectUsecase.EXPECT().RestoreFromTrashProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002xxx",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "IDが存在しない",
			fileSuffix: "404-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.RestoreFromTrashProjectIDInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "024d78d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.NotFound

				f.restoreFromTrashProjectUsecase.EXPECT().RestoreFromTrashProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.RestoreFromTrashProjectIDInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "024d78d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.NotFound

				f.restoreFromTrashProjectUsecase.EXPECT().RestoreFromTrashProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトがすでに復元されている",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.RestoreFromTrashProjectIDInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "024d78d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.Conflict

				f.restoreFromTrashProjectUsecase.EXPECT().RestoreFromTrashProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.RestoreFromTrashProjectIDInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "024d78d6-1d03-11ec-a478-0242ac180002",
				}
				err := apperrors.InternalServerError

				f.restoreFromTrashProjectUsecase.EXPECT().RestoreFromTrashProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				restoreFromTrashProjectUsecase: mockprojectusecase.NewMockRestoreFromTrashProjectUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewRestoreFromTrashProjectHandler(f.restoreFromTrashProjectUsecase)

			r := httptest.NewRequest(http.MethodPut, "/products/{productID}/projects/{projectID}/restore/trash-box", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.RestoreFromTrashProject(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
