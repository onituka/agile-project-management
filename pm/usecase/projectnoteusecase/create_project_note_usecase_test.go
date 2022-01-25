package projectnoteusecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/mockprojectnoterepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
)

func TestCreateProjectNoteUsecaseCreateProjectNote(t *testing.T) {
	type fields struct {
		projectNoteRepository *mockprojectnoterepository.MockProjectNoteRepository
		productRepository     *mockproductrepository.MockProductRepository
		projectRepository     *mockprojectrepository.MockProjectRepository
	}
	type args struct {
		ctx context.Context
		in  *projectnoteinput.CreateProjectNoteInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *projectnoteoutput.CreateProjectNoteOutput
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

				titleVo, err := projectnotedm.NewTitle("プロジェクトまとめ")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().FetchProjectNoteByProjectIDAndTitle(ctx, projectIDVo, titleVo).Return(nil, apperr)
				f.projectNoteRepository.EXPECT().CreateProjectNote(ctx, gomock.Any()).Return(nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want: &projectnoteoutput.CreateProjectNoteOutput{
				ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
				ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
				GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
				Title:     "プロジェクトまとめ",
				Content:   "プロジェクトに関する資料まとめ",
				CreatedBy: "777d78d6-1d03-11ec-a478-0242ac184402",
				UpdatedBy: "777d78d6-1d03-11ec-a478-0242ac184402",
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want:    nil,
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
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want:    nil,
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
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-x478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want:    nil,
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

				projectIDVo, err := projectdm.NewProjectID("024d71d6-1d03-11ec-a478-0242ac180002")
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
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name: "グループID不正",
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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-x478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "タイトル不正",
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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "",
					Content:   "プロジェクトに関する資料まとめ",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "コンテンツ不正",
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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "UserID不正",
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

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-x478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
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

				titleVo, err := projectnotedm.NewTitle("プロジェクトまとめ")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteRepository.EXPECT().FetchProjectNoteByProjectIDAndTitle(ctx, projectIDVo, titleVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
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
				projectNoteRepository: mockprojectnoterepository.NewMockProjectNoteRepository(gmctrl),
				productRepository:     mockproductrepository.NewMockProductRepository(gmctrl),
				projectRepository:     mockprojectrepository.NewMockProjectRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewCreateProjectNoteUsecase(f.projectNoteRepository, f.productRepository, f.projectRepository)

			got, err := u.CreateProjectNote(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("CreateProjectNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(projectnoteoutput.CreateProjectNoteOutput{}, "ID", "CreatedAt", "UpdatedAt")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
