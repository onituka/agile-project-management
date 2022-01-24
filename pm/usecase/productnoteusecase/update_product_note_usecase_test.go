package productnoteusecase

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
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/mockrepository/mockproductnoterepository"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
)

func TestUpdateProductNoteUsecaseUpdateProductNote(t *testing.T) {
	type fields struct {
		productNoteRepository *mockproductnoterepository.MockProductNoteRepository
		productRepository     *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *productnoteinput.UpdateProductNoteInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productnoteoutput.UpdateProductNoteOutput
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

				productNoteID, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				productNoteDm, err := productnotedm.Reconstruct(
					"52dfc0d0-748e-11ec-88fd-acde48001122",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"ノート",
					"note",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				oldProductNoteDm, err := productnotedm.Reconstruct(
					"52dfc0d0-748e-11ec-88fd-acde48001122",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"ノート2",
					"note",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				titleVo, err := productnotedm.NewTitle("ノート")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().UpdateProductNote(ctx, gomock.Any()).Return(nil)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)

				gomock.InOrder(
					f.productNoteRepository.EXPECT().FetchProductNoteByID(ctx, productNoteID, productIDVo).Return(productNoteDm, nil),
					f.productNoteRepository.EXPECT().FetchProductNoteByID(ctx, productNoteID, productIDVo).Return(oldProductNoteDm, nil),
				)

				f.productNoteRepository.EXPECT().FetchProductNoteByProductIDAndTitle(ctx, productIDVo, titleVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					UserID:    "024d78d6-1d03-11ec-a478-0242ac184402",
				},
			},
			want: &productnoteoutput.UpdateProductNoteOutput{
				ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
				ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
				Title:     "ノート",
				Content:   "note",
				CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Now().UTC(),
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "invalid product id",
					Title:     "ノート",
					Content:   "note",
					UserID:    "024d78d6-1d03-11ec-a478-0242ac184402",
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

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
		},
		{
			name: "プロダクトノートIDの不正",
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
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "invalid product note id",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					UserID:    "024d78d6-1d03-11ec-a478-0242ac184402",
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

				productNoteID, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteID, productIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want:    nil,
			wantErr: apperrors.NotFound,
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

				productNoteID, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				productNoteDm, err := productnotedm.Reconstruct(
					"52dfc0d0-748e-11ec-88fd-acde48001122",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"ノート",
					"note",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)
				f.productNoteRepository.EXPECT().FetchProductNoteByID(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "",
					Content:   "note",
					UserID:    "024d78d6-1d03-11ec-a478-0242ac184402",
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

				productNoteID, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				productNoteDm, err := productnotedm.Reconstruct(
					"52dfc0d0-748e-11ec-88fd-acde48001122",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"ノート",
					"note",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)
				f.productNoteRepository.EXPECT().FetchProductNoteByID(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "",
					UserID:    "024d78d6-1d03-11ec-a478-0242ac184402",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "createdBy不正",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				productNoteID, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				productNoteDm, err := productnotedm.Reconstruct(
					"52dfc0d0-748e-11ec-88fd-acde48001122",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"ノート",
					"note",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)
				f.productNoteRepository.EXPECT().FetchProductNoteByID(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					UserID:    "invalid created By",
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

				productNoteID, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				productNoteDm, err := productnotedm.Reconstruct(
					"52dfc0d0-748e-11ec-88fd-acde48001122",
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"ノート",
					"note",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteID, productIDVo).Return(productNoteDm, nil)
				f.productNoteRepository.EXPECT().FetchProductNoteByID(ctx, productNoteID, productIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					UserID:    "024d78d6-1d03-11ec-a478-0242ac184402",
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
				productNoteRepository: mockproductnoterepository.NewMockProductNoteRepository(gmctrl),
				productRepository:     mockproductrepository.NewMockProductRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewUpdateProductNoteUsecase(f.productNoteRepository, f.productRepository)

			got, err := u.UpdateProductNote(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("UpdateProductNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(productnoteoutput.UpdateProductNoteOutput{}, "ID", "CreatedAt", "UpdatedAt")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
