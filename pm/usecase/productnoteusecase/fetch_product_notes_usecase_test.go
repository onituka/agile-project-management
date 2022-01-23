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
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/mockqueryservice/mockproductnotequeryservice"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
)

func TestFetchProductNotesUsecaseFetchProductNotes(t *testing.T) {
	type fields struct {
		productNoteQueryService *mockproductnotequeryservice.MockProductNoteQueryService
		productRepository       *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *productnoteinput.FetchProductNotesInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productnoteoutput.FetchProductNotesOutput
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

				totalCount := uint32(2)

				f.productNoteQueryService.EXPECT().CountProductNotesByProductID(ctx, productIDVo).Return(totalCount, nil)

				limit := uint32(10)
				offset := uint32(0)

				productNoteDtos := []*productnoteoutput.ProductNoteOutput{
					{
						ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Title:     "ノート",
						Content:   "note",
						CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "62dfc0d0-748e-11ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Title:     "test",
						Content:   "test",
						CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteQueryService.EXPECT().FetchProductNotes(ctx, productIDVo, limit, offset).Return(productNoteDtos, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      1,
					Limit:     10,
				},
			},
			want: &productnoteoutput.FetchProductNotesOutput{
				TotalCount: 2,
				ProductNotes: []*productnoteoutput.ProductNoteOutput{
					{
						ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Title:     "ノート",
						Content:   "note",
						CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "62dfc0d0-748e-11ec-88fd-acde48001122",
						ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Title:     "test",
						Content:   "test",
						CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
						UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
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

				totalCount := uint32(0)

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteQueryService.EXPECT().CountProductNotesByProductID(ctx, productIDVo).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      1,
					Limit:     10,
				},
			},
			want: &productnoteoutput.FetchProductNotesOutput{
				TotalCount:   0,
				ProductNotes: make([]*productnoteoutput.ProductNoteOutput, 0),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
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
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(0),
					Limit:     uint32(10),
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
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      uint32(1),
					Limit:     uint32(0),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DBエラー(CountProductNotesByProductID)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteQueryService.EXPECT().CountProductNotesByProductID(ctx, productIDVo).Return(uint32(0), apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Page:      1,
					Limit:     10,
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
		{
			name: "DBエラー(FetchProductNotes)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				totalCount := uint32(2)

				f.productNoteQueryService.EXPECT().CountProductNotesByProductID(ctx, productIDVo).Return(totalCount, nil)

				limit := uint32(10)
				offset := uint32(0)

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteQueryService.EXPECT().FetchProductNotes(ctx, productIDVo, limit, offset).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.FetchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
				productNoteQueryService: mockproductnotequeryservice.NewMockProductNoteQueryService(gmctrl),
				productRepository:       mockproductrepository.NewMockProductRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewFetchProductNotesUsecase(f.productNoteQueryService, f.productRepository)

			got, err := u.FetchProductNotes(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("FetchProductNotes() error = %v, wantErr %v", err, tt.wantErr)
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
