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
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/mockprojectnotequeryservice"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockprojectrepository"
)

func Test_searchProjectNotesUsecase_SearchProjectNotes(t *testing.T) {
	type fields struct {
		projectNoteQueryService *mockprojectnotequeryservice.MockProjectNoteQueryService
		productRepository       *mockproductrepository.MockProductRepository
		projectRepository       *mockprojectrepository.MockProjectRepository
	}
	type args struct {
		ctx context.Context
		in  *projectnoteinput.SearchProjectNotesInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *projectnoteoutput.SearchProjectNotesOutput
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
				limit := uint32(50)
				offset := uint32(0)
				totalCount := uint32(2)
				title := "プロジェクト"

				projectNotesDto := []*projectnoteoutput.SearchProjectNoteOutPut{
					{
						ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "プロジェクトまとめ",
						Content:   "プロジェクトに関する資料まとめ",
						CreatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						UpdatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "プロジェクトまとめ1",
						Content:   "プロジェクトに関する資料まとめ",
						CreatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						UpdatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteQueryService.EXPECT().CountProjectNotesByTitle(ctx, productIDVo, projectIDVo, title).Return(totalCount, nil)
				f.projectNoteQueryService.EXPECT().SearchProjectNotes(ctx, productIDVo, projectIDVo, title, limit, offset).Return(projectNotesDto, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
				},
			},
			want: &projectnoteoutput.SearchProjectNotesOutput{
				TotalCount: 2,
				ProjectNotes: []*projectnoteoutput.SearchProjectNoteOutPut{
					{
						ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "プロジェクトまとめ",
						Content:   "プロジェクトに関する資料まとめ",
						CreatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						UpdatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "プロジェクトまとめ1",
						Content:   "プロジェクトに関する資料まとめ",
						CreatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						UpdatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常(プロジェクトノートが存在しない)",
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

				totalCount := uint32(0)
				title := "プロジェクト"

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteQueryService.EXPECT().CountProjectNotesByTitle(ctx, productIDVo, projectIDVo, title).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
				},
			},
			want: &projectnoteoutput.SearchProjectNotesOutput{
				TotalCount:   0,
				ProjectNotes: make([]*projectnoteoutput.SearchProjectNoteOutPut, 0),
			},
			wantErr: nil,
		},
		{
			name:        "page値の不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      0,
					Limit:     50,
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
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
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
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
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-x478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
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
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},

		{
			name: "DBエラー(CountProjectNotesByTitle)",
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

				title := "プロジェクト"

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteQueryService.EXPECT().CountProjectNotesByTitle(ctx, productIDVo, projectIDVo, title).Return(uint32(0), apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
		{
			name: "DBエラー(SearchProjectNotes)",
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

				totalCount := uint32(2)
				limit := uint32(50)
				offset := uint32(0)
				title := "プロジェクト"

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectRepository.EXPECT().FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo).Return(nil, err)
				f.projectNoteQueryService.EXPECT().CountProjectNotesByTitle(ctx, productIDVo, projectIDVo, title).Return(totalCount, nil)
				f.projectNoteQueryService.EXPECT().SearchProjectNotes(ctx, productIDVo, projectIDVo, title, limit, offset).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectnoteinput.SearchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクト",
					Page:      1,
					Limit:     50,
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
				projectNoteQueryService: mockprojectnotequeryservice.NewMockProjectNoteQueryService(gmctrl),
				productRepository:       mockproductrepository.NewMockProductRepository(gmctrl),
				projectRepository:       mockprojectrepository.NewMockProjectRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewSearchProjectNotesUsecase(f.projectNoteQueryService, f.productRepository, f.projectRepository)

			got, err := u.SearchProjectNotes(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("SearchprojectNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(projectnoteoutput.SearchProjectNoteOutPut{}, "ID", "CreatedAt", "UpdatedAt")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
