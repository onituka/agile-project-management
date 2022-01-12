package projectusecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockqueryservice/mockprojectqueryservice"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

func Test_fetchProjectsUsecase_FetchProjects(t *testing.T) {
	type fields struct {
		projectQueryService *mockprojectqueryservice.MockProjectQueryService
	}
	type args struct {
		ctx context.Context
		in  *projectinput.FetchProjectsInput
	}
	trashedAt := time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *projectoutput.FetchProjectsOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}
				limit := uint32(2)
				offset := uint32(0)
				totalCount := 4

				projectDtos := []*projectoutput.ProjectOutput{
					{
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
					},
					{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "BBB",
						Name:              "管理ツール2",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
						TrashedAt:         &trashedAt,
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
				}

				f.projectQueryService.EXPECT().CountProjectsByProductID(ctx, productIDVo).Return(totalCount, nil)
				f.projectQueryService.EXPECT().FetchProjects(ctx, productIDVo, limit, offset).Return(projectDtos, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(1),
					Limit:     uint32(2),
				},
			},
			want: &projectoutput.FetchProjectsOutput{
				TotalCount: 4,
				Projects: []*projectoutput.ProjectOutput{
					{
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
					},
					{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "BBB",
						Name:              "管理ツール2",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
						TrashedAt:         &trashedAt,
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常(プロジェクトが存在しない場合)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				totalCount := 0

				f.projectQueryService.EXPECT().CountProjectsByProductID(ctx, productIDVo).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(1),
					Limit:     uint32(2),
				},
			},
			want: &projectoutput.FetchProjectsOutput{
				TotalCount: 0,
				Projects:   make([]*projectoutput.ProjectOutput, 0),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "test",
					Page:      uint32(1),
					Limit:     uint32(2),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "page値の不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(0),
					Limit:     uint32(2),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "limit値の不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(1),
					Limit:     uint32(0),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DB不正(CountProjectsByProductID)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVO, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.projectQueryService.EXPECT().CountProjectsByProductID(ctx, productIDVO).Return(0, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(1),
					Limit:     uint32(2),
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
		{
			name: "DB不正(FetchProjects)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVO, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}
				limit := uint32(2)
				offset := uint32(0)
				totalcount := 3

				apperr := apperrors.InternalServerError

				f.projectQueryService.EXPECT().CountProjectsByProductID(ctx, productIDVO).Return(totalcount, nil)
				f.projectQueryService.EXPECT().FetchProjects(ctx, productIDVO, limit, offset).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.FetchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(1),
					Limit:     uint32(2),
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
				projectQueryService: mockprojectqueryservice.NewMockProjectQueryService(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewFetchProjectsUsecase(f.projectQueryService)

			got, err := u.FetchProjects(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
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
