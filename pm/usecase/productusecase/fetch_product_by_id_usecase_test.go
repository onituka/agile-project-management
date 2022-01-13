package productusecase

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
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

func TestFetchProductByIDUsecaseFetchProductByID(t *testing.T) {
	type fields struct {
		productRepository *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *productinput.FetchProductByIDInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productoutput.FetchProductByIDOutput
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

				productDm, err := productdm.Reconstruct(
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"プロジェクト管理ツール",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				f.productRepository.EXPECT().FetchProductByID(ctx, productIDVo).Return(productDm, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				},
			},
			want: &productoutput.FetchProductByIDOutput{
				ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
				Name:      "プロジェクト管理ツール",
				LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
				CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:        "productID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductByIDInput{
					ID: "invalid product id",
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

				f.productRepository.EXPECT().FetchProductByID(ctx, productIDVo).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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

				productDm, err := productdm.Reconstruct(
					"4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"024d78d6-1d03-11ec-a478-0242ac180002",
					"プロジェクト管理ツール",
					"024d78d6-1d03-11ec-a478-0242ac184402",
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				)
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByID(ctx, productIDVo).Return(productDm, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
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
				productRepository: mockproductrepository.NewMockProductRepository(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewFetchProductByIDUsecase(f.productRepository)

			got, err := u.FetchProductByID(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("FetchProductByID() error = %v, wantErr %v", err, tt.wantErr)
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
