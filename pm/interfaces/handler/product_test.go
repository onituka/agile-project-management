package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/mockusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

func TestProductHandlerCreateProduct(t *testing.T) {
	type fields struct {
		productUsecase *mockusecase.MockProductUsecase
	}
	tests := []struct {
		name        string
		fileSuffix  string
		prepareMock func(f *fields)
	}{
		{
			name:       "200-正常",
			fileSuffix: "200",
			prepareMock: func(f *fields) {
				ctX := context.Background()

				in := &productusecase.CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				out := &productusecase.CreateProductOutput{
					ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:      "プロジェクト管理ツール",
					LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
					CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				}

				f.productUsecase.EXPECT().CreateProduct(ctX, in).Return(out, nil)

			},
		},
		{
			name:        "400-1-jsonデコード失敗",
			fileSuffix:  "400-1",
			prepareMock: nil,
		},
		{
			name:       "400-2-グループID不正",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &productusecase.CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac1800020",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}
				err := apperrors.InvalidParameter

				f.productUsecase.EXPECT().CreateProduct(ctx, in).Return(nil, err)
			},
		},
		{
			name:       "409-プロダクト名重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &productusecase.CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.Conflict

				f.productUsecase.EXPECT().CreateProduct(ctx, in).Return(nil, err)
			},
		},
		{
			name:       "500-DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := context.Background()

				in := &productusecase.CreateProductInput{
					GroupID:  "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InternalServerError

				f.productUsecase.EXPECT().CreateProduct(ctx, in).Return(nil, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				productUsecase: mockusecase.NewMockProductUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewProductHandler(f.productUsecase)

			r := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			h.CreateProduct(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
