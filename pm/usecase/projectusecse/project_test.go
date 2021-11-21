package projectusecse

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/mockrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/mocktime"
)

func TestProjectUsecaseFetchProjectByID(t *testing.T) {
	type fields struct {
		projectRepository *mockrepository.MockProjectRepository
		timeManager       *mocktime.MockTimeManager
	}
	type args struct {
		ctx context.Context
		in  *FetchProjectByIDInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields)
		args        args
		want        *FetchProjectByIDOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) {
				ctx := context.Background()
				projectIDVo := projectdm.ProjectID("024d71d6-1d03-11ec-a478-0242ac180002")

				projectDm, _ := projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180002",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"AAA",
					"管理ツール1",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)

				f.projectRepository.EXPECT().FetchProjectByID(ctx, projectIDVo).Return(projectDm, nil)
			},
			args: args{
				ctx: context.Background(),
				in: &FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				},
			},
			want: &FetchProjectByIDOutput{
				ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
				GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
				KeyName:           "AAA",
				Name:              "管理ツール1",
				LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
				DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				CreatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				UpdatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:        "プロジェクトID不正",
			prepareMock: nil,
			args: args{
				ctx: context.Background(),
				in: &FetchProjectByIDInput{
					ID: "invalid project id",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "指定したIDでプロジェクトが存在しない",
			prepareMock: func(f *fields) {
				ctx := context.Background()
				projectIDVo := projectdm.ProjectID("024d71d6-1d03-11ec-a478-0242ac180002")

				err := apperrors.NotFound

				f.projectRepository.EXPECT().FetchProjectByID(ctx, projectIDVo).Return(nil, err)
			},
			args: args{
				ctx: context.Background(),
				in: &FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name: "DBエラー",
			prepareMock: func(f *fields) {
				ctx := context.Background()
				projectIDVo := projectdm.ProjectID("024d71d6-1d03-11ec-a478-0242ac180002")

				err := apperrors.InternalServerError

				f.projectRepository.EXPECT().FetchProjectByID(ctx, projectIDVo).Return(nil, err)
			},
			args: args{
				ctx: context.Background(),
				in: &FetchProjectByIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				projectRepository: mockrepository.NewMockProjectRepository(gmctrl),
				timeManager:       mocktime.NewMockTimeManager(gmctrl),
			}
			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			u := NewProjectUsecase(f.projectRepository, f.timeManager)

			got, err := u.FetchProjectByID(tt.args.ctx, tt.args.in)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("FetchProjectByID() error = %v, wantErr %v", err, tt.wantErr)
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
