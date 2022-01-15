package productusecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockqueryservice/mockproductqueryservice"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

func TestFetchProductsUsecaseFetchProducts(t *testing.T) {
	type fields struct {
		productqueryservice *mockproductqueryservice.MockProductQueryService
	}
	type args struct {
		ctx context.Context
		in  *productinput.FetchProductsInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *productoutput.FetchProductsOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				groupID := "024d78d6-1d03-11ec-a478-0242ac180002"
				totalCount := uint32(2)

				f.productqueryservice.EXPECT().CountProductsByGroupID(ctx, groupID).Return(totalCount, nil)

				limit := uint32(10)
				offset := uint32(0)

				productsDto := []*productoutput.ProductOutput{
					{
						ID:        "4487c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Name:      "test",
						LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Name:      "プロジェクト管理ツール",
						LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					},
				}

				f.productqueryservice.EXPECT().FetchProducts(ctx, groupID, limit, offset).Return(productsDto, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(1),
					Limit:   uint32(10),
				},
			},
			want: &productoutput.FetchProductsOutput{
				TotalCount: 2,
				Products: []*productoutput.ProductOutput{
					{
						ID:        "4487c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Name:      "test",
						LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
						GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
						Name:      "プロジェクト管理ツール",
						LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
						CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常(プロダクトが存在しない場合)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				groupID := "024d78d6-1d03-11ec-a478-0242ac180002"
				totalCount := uint32(0)

				f.productqueryservice.EXPECT().CountProductsByGroupID(ctx, groupID).Return(totalCount, nil)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(1),
					Limit:   uint32(10),
				},
			},
			want: &productoutput.FetchProductsOutput{
				TotalCount: 0,
				Products:   make([]*productoutput.ProductOutput, 0),
			},
			wantErr: nil,
		},
		{
			name:        "グループID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "invalid group id",
					Page:    uint32(1),
					Limit:   uint32(10),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "page値不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(0),
					Limit:   uint32(10),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "limit値不正（下限値）",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(1),
					Limit:   uint32(0),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name:        "limit値不正（上限値）",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(1),
					Limit:   uint32(51),
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DBエラー(CountProductsメソッド実行時エラー)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				groupID := "024d78d6-1d03-11ec-a478-0242ac180002"
				apperr := apperrors.InternalServerError

				f.productqueryservice.EXPECT().CountProductsByGroupID(ctx, groupID).Return(uint32(0), apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(1),
					Limit:   uint32(10),
				},
			},
			want:    nil,
			wantErr: apperrors.InternalServerError,
		},
		{
			name: "DBエラー(FetchProductsメソッド実行時エラー)",
			prepareMock: func(f *fields) error {
				ctx := context.TODO()

				groupID := "024d78d6-1d03-11ec-a478-0242ac180002"
				totalCount := uint32(2)

				f.productqueryservice.EXPECT().CountProductsByGroupID(ctx, groupID).Return(totalCount, nil)

				limit := uint32(10)
				offset := uint32(0)

				apperr := apperrors.InternalServerError

				f.productqueryservice.EXPECT().FetchProducts(ctx, groupID, limit, offset).Return(nil, apperr)

				return nil
			},
			args: args{
				ctx: context.TODO(),
				in: &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    uint32(1),
					Limit:   uint32(10),
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
				productqueryservice: mockproductqueryservice.NewMockProductQueryService(gmctrl),
			}
			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewFetchProductsUsecase(f.productqueryservice)

			got, err := u.FetchProducts(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("FetchProducts() error = %v, wantErr %v", err, tt.wantErr)
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
