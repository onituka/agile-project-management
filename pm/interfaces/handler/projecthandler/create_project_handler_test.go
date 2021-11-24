package projecthandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projecthandler/mockprojectusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

func TestCreateProjectHandlerCreateProject(t *testing.T) {
	type fields struct {
		createProjectUsecase *mockprojectusecase.MockCreateProjectUsecase
	}
	tests := []struct {
		name        string
		fileSuffix  string
		prepareMock func(f *fields)
	}{
		{
			name:       "200-正常",
			fileSuffix: "200",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &projectusecase.CreateProjectInput{
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				}

				out := &projectusecase.CreateProjectOutput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
					CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
				}

				f.createProjectUsecase.EXPECT().CreateProject(ctx, in).Return(out, nil)
			},
		},
		{
			name:        "400-1-jsonデコード失敗",
			fileSuffix:  "400-1",
			prepareMock: nil,
		},
		{
			name:       "400-2-キー名不正",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &projectusecase.CreateProjectInput{
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "1AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				}

				err := apperrors.InvalidParameter

				f.createProjectUsecase.EXPECT().CreateProject(ctx, in).Return(nil, err)
			},
		},
		{
			name:       "409-キー名重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &projectusecase.CreateProjectInput{
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				}

				err := apperrors.Conflict

				f.createProjectUsecase.EXPECT().CreateProject(ctx, in).Return(nil, err)
			},
		},
		{
			name:       "500-DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &projectusecase.CreateProjectInput{
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				}

				err := apperrors.InternalServerError

				f.createProjectUsecase.EXPECT().CreateProject(ctx, in).Return(nil, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				createProjectUsecase: mockprojectusecase.NewMockCreateProjectUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewCreateProjectHandler(f.createProjectUsecase)

			r := httptest.NewRequest(http.MethodPost, "/projects", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			h.CreateProject(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
