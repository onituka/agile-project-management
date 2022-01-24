package productnoteusecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/mockrepository/mockproductnoterepository"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
)

func TestDeleteProductNoteUsecaseDeleteProductNote(t *testing.T) {
	type fields struct {
		productNoteRepository *mockproductnoterepository.MockProductNoteRepository
		productRepository     *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *productnoteinput.DeleteProductNoteInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productnoteoutput.DeleteProductNoteMsg
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

				productNoteIDVo, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteIDVo, productIDVo).Return(nil, nil)
				f.productNoteRepository.EXPECT().DeleteProductNote(ctx, productNoteIDVo, productIDVo).Return(nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want: &productnoteoutput.DeleteProductNoteMsg{
				Message: "プロダクトノートを削除しました。",
			},
			wantErr: nil,
		},
		{
			name:        "プロダクトIDの不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "invalid product id",
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
				in: &productnoteinput.DeleteProductNoteInput{
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
				in: &productnoteinput.DeleteProductNoteInput{
					ID:        "invalid product note id",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "指定したIDでプロダクトノートが存在しない",
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
				in: &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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

				productIDVo, err := productdm.NewProductID("4495c574-34c2-4fb3-9ca4-3a7c79c267a6")
				if err != nil {
					return err
				}

				productNoteIDVo, err := productnotedm.NewProductNoteID("52dfc0d0-748e-11ec-88fd-acde48001122")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByIDForUpdate(ctx, productIDVo).Return(nil, err)
				f.productNoteRepository.EXPECT().FetchProductNoteByIDForUpdate(ctx, productNoteIDVo, productIDVo).Return(nil, nil)
				f.productNoteRepository.EXPECT().DeleteProductNote(ctx, productNoteIDVo, productIDVo).Return(apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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

			u := NewDeleteProductNoteUsecase(f.productNoteRepository, f.productRepository)

			got, err := u.DeleteProductNote(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("DeleteProductNote() error = %v, wantErr %v", err, tt.wantErr)
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
