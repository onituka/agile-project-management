package projectusecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
)

func TestFetchProjectsUsecaseFetchProjects(t *testing.T) {
	type fields struct {
		projectRepository *mockprojectrepository.MockProjectRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields)
		args        args
		want        FetchProjectsOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) {
				ctx := context.Background()
				projectDms := make([]*projectdm.Project, 2)
				projectDms[0], _ = projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180002",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"AAA",
					"管理ツール1",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)
				projectDms[1], _ = projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180003",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"BBB",
					"管理ツール2",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)

				f.projectRepository.EXPECT().FetchProjects(ctx).Return(projectDms, nil)
			},
			args: args{
				ctx: context.Background(),
			},
			want: FetchProjectsOutput{
				&Project{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
					CreatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					UpdatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				},
				&Project{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180003",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "BBB",
					Name:              "管理ツール2",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
					CreatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					UpdatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "DBエラー",
			prepareMock: func(f *fields) {
				ctx := context.Background()
				projectDms := make([]*projectdm.Project, 2)
				projectDms[0], _ = projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180002",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"AAA",
					"管理ツール1",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)
				projectDms[1], _ = projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180003",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"BBB",
					"管理ツール2",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)
				err := apperrors.InternalServerError

				f.projectRepository.EXPECT().FetchProjects(ctx).Return(nil, err)
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)
			f := fields{
				projectRepository: mockprojectrepository.NewMockProjectRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			u := NewFetchProjectsUsecase(f.projectRepository)

			got, err := u.FetchProjects(tt.args.ctx)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("FetchProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
