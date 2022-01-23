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
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/mockqueryservice/mockprojectqueryservice"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

func TestSearchProjectsUsecaseSearchProjects(t *testing.T) {
	type fields struct {
		projectQueryService *mockprojectqueryservice.MockProjectQueryService
		productRepository   *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *projectinput.SearchProjectsInput
	}
	trashedAt := time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *projectoutput.SearchProjectsOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				keyword := "A"
				limit := uint32(2)
				offset := uint32(0)
				totalCount := uint32(2)

				projectsDto := []*projectoutput.SearchProjectOutput{
					{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "A",
						Name:              "管理ツール1",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
						TrashedAt:         &trashedAt,
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "BBB",
						Name:              "管理ツールA",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
						TrashedAt:         &trashedAt,
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectQueryService.EXPECT().CountProjectsByKeyNameAndName(ctx, productIDVo, keyword).Return(totalCount, nil)
				f.projectQueryService.EXPECT().SearchProjects(ctx, productIDVo, keyword, limit, offset).Return(projectsDto, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				},
			},
			want: &projectoutput.SearchProjectsOutput{
				TotalCount: 2,
				Projects: []*projectoutput.SearchProjectOutput{
					{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "A",
						Name:              "管理ツール1",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
						TrashedAt:         &trashedAt,
						CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
						ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
						KeyName:           "BBB",
						Name:              "管理ツールA",
						LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
						DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
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

				keyword := "A"
				totalCount := uint32(0)

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectQueryService.EXPECT().CountProjectsByKeyNameAndName(ctx, productIDVo, keyword).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				},
			},
			want: &projectoutput.SearchProjectsOutput{
				TotalCount: 0,
				Projects:   make([]*projectoutput.SearchProjectOutput, 0),
			},
			wantErr: nil,
		},
		{
			name: "正常(検索にヒットしなかった場合)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				keyword := "A"
				totalCount := uint32(0)

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectQueryService.EXPECT().CountProjectsByKeyNameAndName(ctx, productIDVo, keyword).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				},
			},
			want: &projectoutput.SearchProjectsOutput{
				TotalCount: 0,
				Projects:   make([]*projectoutput.SearchProjectOutput, 0),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
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
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name:        "page値の不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
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
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      uint32(1),
					Limit:     uint32(0),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DBエラー(CountProjectsByKeyNameAndName)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				keyword := "A"

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectQueryService.EXPECT().CountProjectsByKeyNameAndName(ctx, productIDVo, keyword).Return(uint32(0), apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
		{
			name: "DBエラー(SearchProjects)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				keyword := "A"
				limit := uint32(2)
				offset := uint32(0)
				totalCount := uint32(3)

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.projectQueryService.EXPECT().CountProjectsByKeyNameAndName(ctx, productIDVo, keyword).Return(totalCount, nil)
				f.projectQueryService.EXPECT().SearchProjects(ctx, productIDVo, keyword, limit, offset).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &projectinput.SearchProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
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
				productRepository:   mockproductrepository.NewMockProductRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewSearchProjectsUsecase(f.projectQueryService, f.productRepository)

			got, err := u.SearchProjects(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("SearchProjects() error = %v, wantErr %v", err, tt.wantErr)
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
