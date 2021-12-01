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
	"github.com/onituka/agile-project-management/project-management/usecase/mocktime"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/mockrepository/mockproductrepository"
)

func TestCreateProductUsecaseCreateProduct(t *testing.T) {
	type fields struct {
		productRepository *mockproductrepository.MockProductRepository
		timeManager       *mocktime.MockTimeManager
	}
	type args struct {
		ctx context.Context
		in  *CreateProductInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields)
		args        args
		want        *CreateProductOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) {
				ctx := context.TODO()
				now := time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC)
				groupIDVo := groupdm.GroupID("024d78d6-1d03-11ec-a478-0242ac180002")
				nameVo := productdm.Name("プロジェクト管理ツール")
				err := apperrors.NotFound

				f.timeManager.EXPECT().Now().Return(now)
				f.productRepository.EXPECT().FetchProductByGroupIDAndName(ctx, groupIDVo, nameVo).Return(nil, err)
				f.productRepository.EXPECT().CreateProduct(ctx, gomock.Any()).Return(nil)
			},
			args: args{
				ctx: context.TODO(),
				in: &CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				},
			},
			want: &CreateProductOutput{
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
			name:        "グループID不正",
			prepareMock: nil,
			args: args{
				ctx: context.TODO(),
				in: &CreateProductInput{
					GroupID:  "invalid group id",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
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
				in: &CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "A",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
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
				in: &CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "invalid leader id",
				},
			},
			want:    nil,
			wantErr: apperrors.InvalidParameter,
		},
		{
			name: "DBエラー",
			prepareMock: func(f *fields) {
				ctx := context.TODO()
				now := time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC)
				groupIDVo := groupdm.GroupID("024d78d6-1d03-11ec-a478-0242ac180002")
				nameVo := productdm.Name("プロジェクト管理ツール")
				err := apperrors.InternalServerError

				f.timeManager.EXPECT().Now().Return(now)
				f.productRepository.EXPECT().FetchProductByGroupIDAndName(ctx, groupIDVo, nameVo).Return(nil, err)
			},
			args: args{
				ctx: context.TODO(),
				in: &CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
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
				timeManager:       mocktime.NewMockTimeManager(gmctrl),
			}
			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			u := NewCreateProductUsecase(f.productRepository, f.timeManager)

			got, err := u.CreateProduct(tt.args.ctx, tt.args.in)
			if hasErr, expectErr := err != nil, tt.wantErr != nil; hasErr != expectErr {
				t.Errorf("FetchProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(CreateProductOutput{}, "ID")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("differs: (-wantErr +gotErr)\n- %v\n+ %v", tt.wantErr, err)
			}
		})
	}
}
