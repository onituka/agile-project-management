package projectusecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
)

func TestTrashedProjectUsecaseTrashedProject(t *testing.T) {
	type fields struct {
		projectRepository *mockprojectrepository.MockProjectRepository
	}
	type args struct {
		ctx context.Context
		in  *TrashedProjectIDInput
	}
	trashedAt := time.Now().UTC()
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *TrashedProjectOutPut
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-11ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				projectDm, err := projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180002",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"AAA",
					"管理ツール1",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					nil,
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Now().UTC(),
				)
				if err != nil {
					return err
				}

				f.projectRepository.EXPECT().UpdateProject(ctx, gomock.Any()).Return(nil)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo).Return(projectDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				},
			},

			want: &TrashedProjectOutPut{
				ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
				ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
				KeyName:           "AAA",
				Name:              "管理ツール1",
				LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
				DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				TrashedAt:         &trashedAt,
				CreatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				UpdatedAt:         time.Now().UTC(),
			},
			wantErr: nil,
		},
		{
			name:        "projectID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &TrashedProjectIDInput{
					ID: "invalid project id",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "指定したIDでプロジェクトが存在しない",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-11ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo).Return(nil, apperr)

				return err
			},
			args: args{
				ctx: context.TODO(),
				in: &TrashedProjectIDInput{
					ID: "024d71d6-1d03-11ec-a478-0242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name: "DBエラー",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-11ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				projectDm, err := projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180002",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"AAA",
					"管理ツール1",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					nil,
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Now().UTC(),
				)
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo).Return(projectDm, nil)
				f.projectRepository.EXPECT().UpdateProject(ctx, projectDm).Return(apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &TrashedProjectIDInput{
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
				projectRepository: mockprojectrepository.NewMockProjectRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewTrashedProjectUsecase(f.projectRepository)

			got, err := u.TrashedProject(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("Trashedproject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(TrashedProjectOutPut{}, "ID", "TrashedAt", "UpdatedAt")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
