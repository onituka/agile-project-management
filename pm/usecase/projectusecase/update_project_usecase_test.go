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
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

func TestUpdateProjectUsecaseUpdateProject(t *testing.T) {
	type fields struct {
		projectRepository *mockprojectrepository.MockProjectRepository
		productRepository *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *projectinput.UpdateProjectInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *projectoutput.UpdateProjectOutput
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
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				oldProjectDm, err := projectdm.Reconstruct(
					"024d71d6-1d03-11ec-a478-0242ac180002",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"BBB",
					"管理ツール2",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-9242ac180002",
					nil,
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().UpdateProject(ctx, gomock.Any()).Return(nil)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(projectDm, nil)
				f.projectRepository.EXPECT().FetchProjectByID(ctx, projectIDVo, productIDVo).Return(oldProjectDm, nil)
				f.projectRepository.EXPECT().FetchProjectByGroupIDAndKeyName(ctx, groupIDVo, keyNameVo).Return(nil, apperr)
				f.projectRepository.EXPECT().FetchProjectByGroupIDAndName(ctx, groupIDVo, nameVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyName:           "AAA",
					Name:              "管理ツール1",
					LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
					DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				},
			},
			want: &projectoutput.UpdateProjectOutput{
				ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
				ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
				KeyName:           "AAA",
				Name:              "管理ツール1",
				LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
				DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac180002",
				TrashedAt:         nil,
				CreatedAt:         time.Date(2021, 11, 20, 0, 0, 0, 0, time.UTC),
				UpdatedAt:         time.Now().UTC(),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "invalid product id",
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
			name: "指定したIDでプロダクトが存在しない",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, apperr)

				return err
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name: "プロジェクトID不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "invalid project id",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
			name: "指定したIDでプロジェクトが存在しない",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-11ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, apperr)

				return err
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:        "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name: "keyName不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(projectDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
			name: "プロジェクト名不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(projectDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
			name: "leaderID不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(projectDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
			name: "defaultAssigneeID不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(projectDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(projectDm, nil)
				f.projectRepository.EXPECT().FetchProjectByID(ctx, projectIDVo, productIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.UpdateProjectInput{
					ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
					ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
				productRepository: mockproductrepository.NewMockProductRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewUpdateProjectUsecase(f.projectRepository, f.productRepository)

			got, err := u.UpdateProject(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("Updateproject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(projectoutput.UpdateProjectOutput{}, "ID", "UpdatedAt")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
