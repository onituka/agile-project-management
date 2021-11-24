package projecthandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projecthandler/mockprojectusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

func TestFetchProjectsHandlerFetchProjects(t *testing.T) {
	type fields struct {
		fetchProjectsUsecase *mockprojectusecase.MockFetchProjectsUsecase
	}
	tests := []struct {
		name        string
		fileSuffix  string
		prepareMock func(f *fields)
	}{
		{
			name:       "200-1-正常",
			fileSuffix: "200-1",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				out := projectusecase.FetchProjectsOutput{
					&projectusecase.Project{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "AAA",
						Name:              "管理ツール1",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
					&projectusecase.Project{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180003",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "BBB",
						Name:              "管理ツール2",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
				}

				f.fetchProjectsUsecase.EXPECT().FetchProjects(ctx).Return(out, nil)
			},
		},
		{
			name:       "200-2-正常",
			fileSuffix: "200-2",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				out := make(projectusecase.FetchProjectsOutput, 0)

				f.fetchProjectsUsecase.EXPECT().FetchProjects(ctx).Return(out, nil)
			},
		},
		{
			name:       "500-DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				err := apperrors.InternalServerError

				f.fetchProjectsUsecase.EXPECT().FetchProjects(ctx).Return(nil, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				fetchProjectsUsecase: mockprojectusecase.NewMockFetchProjectsUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewFetchProjectsHandler(f.fetchProjectsUsecase)

			r := httptest.NewRequest(http.MethodPost, "/projects", nil)
			w := httptest.NewRecorder()

			h.FetchProjects(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
