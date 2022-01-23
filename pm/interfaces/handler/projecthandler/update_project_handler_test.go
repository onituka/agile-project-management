package projecthandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestUpdateProjectHandlerUpdateProject(t *testing.T) {
	type fields struct {
		updateProjectUsecase *mockprojectusecase.MockUpdateProjectUsecase
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

				in := &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
				}

				out := &projectoutput.UpdateProjectOutput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
					CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
				}

				f.updateProjectUsecase.EXPECT().UpdateProject(ctx, in).Return(out, nil)
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
			fileSuffix: "400-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-x242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-x242ac180002",
					ProductID:         "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
				}
				err := apperrors.InvalidParameter

				f.updateProjectUsecase.EXPECT().UpdateProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-x242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:        "jsonデコード失敗",
			fileSuffix:  "400-2",
			prepareMock: nil,
		},
		{
			name:       "400-3-キー名不正",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "1AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
				}
				err := apperrors.InvalidParameter

				f.updateProjectUsecase.EXPECT().UpdateProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
				}
				err := apperrors.NotFound

				f.updateProjectUsecase.EXPECT().UpdateProject(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "キー名重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
					"productID": "024d78d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
				}
				err := apperrors.Conflict

				f.updateProjectUsecase.EXPECT().UpdateProject(ctx, in).Return(nil, err)
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

				in := &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
				}
				err := apperrors.InternalServerError

				f.updateProjectUsecase.EXPECT().UpdateProject(ctx, in).Return(nil, err)
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
				updateProjectUsecase: mockprojectusecase.NewMockUpdateProjectUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewUpdateProjectHandler(f.updateProjectUsecase)

			r := httptest.NewRequest(http.MethodPut, "/projects/024d71d6-1d03-11ec-a478-0242ac180002", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.UpdateProject(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
