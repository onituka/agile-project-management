package productnoteusecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/mockqueryservice/mockproductnotequeryservice"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
)

func TestSearchProductNotesUsecaseSearchProductNotes(t *testing.T) {
	type fields struct {
		productnoteService *mockproductnotequeryservice.MockProductNoteQueryService
		productRepository  *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *productnoteinput.SearchProductNotesInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productnoteoutput.SearchProductNotesOutput
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

				titleVo, err := productnotedm.NewTitle("ノ")
				if err != nil {
					return err
				}

				totalCount := uint32(2)

				f.productnoteService.EXPECT().CountProductNotesByTitle(ctx, productIDVo, titleVo).Return(totalCount, nil)

				limit := uint32(10)
				offset := uint32(0)

				productNotesDto := []*productnoteoutput.SearchProductNoteOutput{
					{
						ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "ノート",
						Content:   "note",
						CreatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "62dfc0d0-748e-41ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "ノート2",
						Content:   "test",
						CreatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productnoteService.EXPECT().SearchProductNotes(ctx, productIDVo, titleVo, limit, offset).Return(productNotesDto, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				},
			},
			want: &productnoteoutput.SearchProductNotesOutput{
				TotalCount: 2,
				ProductNotes: []*productnoteoutput.SearchProductNoteOutput{
					{
						ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "ノート",
						Content:   "note",
						CreatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "62dfc0d0-748e-41ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
						Title:     "ノート2",
						Content:   "test",
						CreatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-44ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常(プロダクトノートが存在しない場合)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				titleVo, err := productnotedm.NewTitle("ノ")
				if err != nil {
					return err
				}

				totalCount := uint32(0)

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productnoteService.EXPECT().CountProductNotesByTitle(ctx, productIDVo, titleVo).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				},
			},
			want: &productnoteoutput.SearchProductNotesOutput{
				TotalCount:   0,
				ProductNotes: make([]*productnoteoutput.SearchProductNoteOutput, 0),
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

				titleVo, err := productnotedm.NewTitle("ノxx")
				if err != nil {
					return err
				}

				totalCount := uint32(0)

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productnoteService.EXPECT().CountProductNotesByTitle(ctx, productIDVo, titleVo).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノxx",
					Page:      1,
					Limit:     10,
				},
			},
			want: &productnoteoutput.SearchProductNotesOutput{
				TotalCount:   0,
				ProductNotes: make([]*productnoteoutput.SearchProductNoteOutput, 0),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "invalid product id",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
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
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
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
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
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
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      uint32(1),
					Limit:     uint32(0),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DBエラー(CountProductNotesByTitleメソッド実行時エラー)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				titleVo, err := productnotedm.NewTitle("ノ")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productnoteService.EXPECT().CountProductNotesByTitle(ctx, productIDVo, titleVo).Return(uint32(0), apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
		{
			name: "DBエラー(SearchProductNotes)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				titleVo, err := productnotedm.NewTitle("ノ")
				if err != nil {
					return err
				}

				totalCount := uint32(3)
				limit := uint32(10)
				offset := uint32(0)

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productnoteService.EXPECT().CountProductNotesByTitle(ctx, productIDVo, titleVo).Return(totalCount, nil)
				f.productnoteService.EXPECT().SearchProductNotes(ctx, productIDVo, titleVo, limit, offset).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
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
				productnoteService: mockproductnotequeryservice.NewMockProductNoteQueryService(gmctrl),
				productRepository:  mockproductrepository.NewMockProductRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewSearchProductNotesUsecase(f.productnoteService, f.productRepository)

			got, err := u.SearchProductNotes(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("SearchProductNotes() error = %v, wantErr %v", err, tt.wantErr)
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
