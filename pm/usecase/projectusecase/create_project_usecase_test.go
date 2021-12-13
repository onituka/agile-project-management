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
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/mocktime"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
)

func TestCreateProjectUsecaseCreateProject(t *testing.T) {
	type fields struct {
		projectRepository *mockprojectrepository.MockProjectRepository
		timeManager       *mocktime.MockTimeManager
	}
	type args struct {
		ctx context.Context
		in  *CreateProjectInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *CreateProjectOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				now := time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC)
				var err error

				groupIDVo, err := groupdm.NewGroupID("024d78d6-1d03-11ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				keyNameVo, err := projectdm.NewKeyName("AAA")
				if err != nil {
					return err
				}

				nameVo, err := projectdm.NewName("管理ツール1")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.timeManager.EXPECT().Now().Return(now)
				f.projectRepository.EXPECT().FetchProjectByGroupIDAndKeyName(ctx, groupIDVo, keyNameVo).Return(nil, apperr)
				f.projectRepository.EXPECT().FetchProjectByGroupIDAndName(ctx, groupIDVo, nameVo).Return(nil, apperr)
				f.projectRepository.EXPECT().CreateProject(ctx, gomock.Any()).Return(nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want: &CreateProjectOutput{
				ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
				ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "invalid product id",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "グループID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "invalid group id",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "keyName不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "invalid keyName",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "プロジェクト名不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "A",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "leaderID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "invalid leader id",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "defaultAssigneeID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "invalid defaultAssignee id",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DBエラー",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				now := time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC)
				var err error

				groupVo, err := groupdm.NewGroupID("024d78d6-1d03-11ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				keyNameVo, err := projectdm.NewKeyName("AAA")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.timeManager.EXPECT().Now().Return(now)
				f.projectRepository.EXPECT().FetchProjectByGroupIDAndKeyName(ctx, groupVo, keyNameVo).Return(nil, apperr)

				return err
			},
			args: args{
				ctx: context.TODO(),
				in: &CreateProjectInput{
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
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
				timeManager:       mocktime.NewMockTimeManager(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewCreateProjectUsecase(f.projectRepository, f.timeManager)

			got, err := u.CreateProject(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("CreateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(CreateProjectOutput{}, "ID")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
