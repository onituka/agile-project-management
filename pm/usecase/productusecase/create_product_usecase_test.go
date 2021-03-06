package productusecase

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
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

func TestCreateProductUsecaseCreateProduct(t *testing.T) {
	type fields struct {
		productRepository *mockproductrepository.MockProductRepository
	}
	type args struct {
		ctx context.Context
		in  *productinput.CreateProductInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productoutput.CreateProductOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()
				var err error

				groupIDVo, err := groupdm.NewGroupID("024d78d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				nameVo, err := productdm.NewName("プロジェクト管理ツール")
				if err != nil {
					return err
				}

				apperr := apperrors.NotFound

				f.productRepository.EXPECT().FetchProductByGroupIDAndName(ctx, groupIDVo, nameVo).Return(nil, apperr)
				f.productRepository.EXPECT().CreateProduct(ctx, gomock.Any()).Return(nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.CreateProductInput{
					GroupID:  "024d78d6-1d03-41ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-44ec-a478-0242ac184402",
				},
			},
			want: &productoutput.CreateProductOutput{
				ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
				Name:      "プロジェクト管理ツール",
				LeaderID:  "024d78d6-1d03-44ec-a478-0242ac184402",
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
			wantErr: nil,
		},
		{
			name:        "グループID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.CreateProductInput{
					GroupID:  "invalid group id",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-44ec-a478-0242ac184402",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "プロダクト名不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.CreateProductInput{
					GroupID:  "024d78d6-1d03-41ec-a478-0242ac180002",
					Name:     "A",
					LeaderID: "024d78d6-1d03-44ec-a478-0242ac184402",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "leaderID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.CreateProductInput{
					GroupID:  "024d78d6-1d03-41ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "invalid leader id",
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

				groupIDVo, err := groupdm.NewGroupID("024d78d6-1d03-41ec-a478-0242ac180002")
				if err != nil {
					return err
				}

				nameVo, err := productdm.NewName("プロジェクト管理ツール")
				if err != nil {
					return err
				}

				apperr := apperrors.InternalServerError

				f.productRepository.EXPECT().FetchProductByGroupIDAndName(ctx, groupIDVo, nameVo).Return(nil, apperr)

				return err
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.CreateProductInput{
					GroupID:  "024d78d6-1d03-41ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-44ec-a478-0242ac184402",
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

			u := NewCreateProductUsecase(f.productRepository)

			got, err := u.CreateProduct(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opt := cmpopts.IgnoreFields(productoutput.CreateProductOutput{}, "ID", "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(tt.want, got, opt); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
