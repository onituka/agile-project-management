package producthandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/producthandler/mockproductusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

func TestSearchProductsHandlerSearchProducts(t *testing.T) {
	type fields struct {
		searchProductsUsecase *mockproductusecase.MockSearchProductsUsecase
	}
	tests := []struct {
		name           string
		fileSuffix     string
		prepareMock    func(f *fields)
		prepareRequest func(r *http.Request)
	}{
		{
			name:       "正常",
			fileSuffix: "200-1",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.SearchProductsInput{
					GroupID:     "024d78d6-1d03-11ec-a478-0242ac180002",
					ProductName: "プ",
					Page:        1,
					Limit:       10,
				}

				out := &productoutput.SearchProductsOutput{
					TotalCount: 1,
					Products: []*productoutput.SearchProductOutput{
						{
							ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
							Name:      "プロジェクト管理ツール",
							LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
							CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
						},
					},
				}

				f.searchProductsUsecase.EXPECT().SearchProducts(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プ")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "正常(該当するプロダクト名が存在しない場合)",
			fileSuffix: "200-2",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.SearchProductsInput{
					GroupID:     "024d78d6-1d03-11ec-a478-0242ac180002",
					ProductName: "プxxx",
					Page:        1,
					Limit:       10,
				}

				out := &productoutput.SearchProductsOutput{
					TotalCount: 0,
					Products:   make([]*productoutput.SearchProductOutput, 0),
				}

				f.searchProductsUsecase.EXPECT().SearchProducts(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プxxx")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:        "クエリーパラメータ文字列によるpage不正",
			fileSuffix:  "400-1",
			prepareMock: nil,
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プ")
				q.Set("page", "test")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:        "クエリーパラメータ文字列によるlimit不正",
			fileSuffix:  "400-2",
			prepareMock: nil,
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プ")
				q.Set("page", "1")
				q.Set("limit", "test")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "usecaseでの400エラー(page値不正)",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.SearchProductsInput{
					GroupID:     "024d78d6-1d03-11ec-a478-0242ac180002",
					ProductName: "プ",
					Page:        0,
					Limit:       10,
				}

				err := apperrors.InvalidParameter

				f.searchProductsUsecase.EXPECT().SearchProducts(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プ")
				q.Set("page", "0")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.SearchProductsInput{
					GroupID:     "024d78d6-1d03-11ec-a478-0242ac180002",
					ProductName: "プ",
					Page:        1,
					Limit:       10,
				}

				err := apperrors.NotFound

				f.searchProductsUsecase.EXPECT().SearchProducts(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プ")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "usecaseでの500エラー(DBエラー)",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.SearchProductsInput{
					GroupID:     "024d78d6-1d03-11ec-a478-0242ac180002",
					ProductName: "プ",
					Page:        1,
					Limit:       10,
				}

				err := apperrors.InternalServerError

				f.searchProductsUsecase.EXPECT().SearchProducts(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("name", "プ")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				searchProductsUsecase: mockproductusecase.NewMockSearchProductsUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewSearchProductsHandler(f.searchProductsUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products/search", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.SearchProducts(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
