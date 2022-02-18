package projectnoteusecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/mockprojectnoterepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
)

func TestDeleteProjectNoteUsecaseDeleteProjectNote(t *testing.T) {
	type fields struct {
		projectNoteRepository *mockprojectnoterepository.MockProjectNoteRepository
		productRepository     *mockproductrepository.MockProductRepository
		projectRepository     *mockprojectrepository.MockProjectRepository
	}
	type args struct {
		ctx context.Context
		in  *projectnoteinput.DeleteProjectNoteInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
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

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				projectNoteIDVo, err := projectnotedm.NewProjectNoteID("52dfc0d0-748e-41ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().FetchProjectNoteByIDForUpdate(ctx, projectNoteIDVo, projectIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().DeleteProjectNote(ctx, projectNoteIDVo, productIDVo, projectIDVo).Return(nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},

			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "プロダクトが存在しない",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},
			wantErr: apperrors.NotFound,
		},
		{
			name: "プロジェクトIDの不正",
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
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-x478-0242ac180002",
				},
			},
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "プロジェクトが存在しない",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},
			wantErr: apperrors.NotFound,
		},
		{
			name: "プロジェクトノートIDの不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-x8fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "プロジェクトノートが存在しない",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				projectNoteIDVo, err := projectnotedm.NewProjectNoteID("52dfc0d0-748e-41ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().FetchProjectNoteByIDForUpdate(ctx, projectNoteIDVo, projectIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},
			wantErr: apperrors.NotFound,
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

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				projectNoteIDVo, err := projectnotedm.NewProjectNoteID("52dfc0d0-748e-41ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().FetchProjectNoteByIDForUpdate(ctx, projectNoteIDVo, projectIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().DeleteProjectNote(ctx, projectNoteIDVo, productIDVo, projectIDVo).Return(apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.DeleteProjectNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
				},
			},
			wantErr: apperrors.InternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				projectNoteRepository: mockprojectnoterepository.NewMockProjectNoteRepository(gmctrl),
				productRepository:     mockproductrepository.NewMockProductRepository(gmctrl),
				projectRepository:     mockprojectrepository.NewMockProjectRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewDeleteProjectNoteUsecase(f.projectNoteRepository, f.productRepository, f.projectRepository)

			err := u.DeleteProjectNote(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("DeleteProjectNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
