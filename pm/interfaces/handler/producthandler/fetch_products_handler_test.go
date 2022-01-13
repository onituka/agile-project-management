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

func TestFetchProductsHandlerFetchProducts(t *testing.T) {
	type fields struct {
		fetchProductsUsecase *mockproductusecase.MockFetchProductsUsecase
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

				in := &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    1,
					Limit:   10,
				}

				out := &productoutput.FetchProductsOutput{
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
				}

				f.fetchProductsUsecase.EXPECT().FetchProducts(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "正常(プロダクトが存在しない場合)",
			fileSuffix: "200-2",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    1,
					Limit:   10,
				}

				out := &productoutput.FetchProductsOutput{
					TotalCount: 0,
					Products:   make([]*productoutput.ProductOutput, 0),
				}

				f.fetchProductsUsecase.EXPECT().FetchProducts(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
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

				in := &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    0,
					Limit:   10,
				}

				err := apperrors.InvalidParameter

				f.fetchProductsUsecase.EXPECT().FetchProducts(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("page", "0")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "usecaseでの500エラー(DBエラー)",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := context.TODO()

				in := &productinput.FetchProductsInput{
					GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
					Page:    1,
					Limit:   10,
				}

				err := apperrors.InternalServerError

				f.fetchProductsUsecase.EXPECT().FetchProducts(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				q := r.URL.Query()
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
				fetchProductsUsecase: mockproductusecase.NewMockFetchProductsUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewFetchProductsHandler(f.fetchProductsUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.FetchProducts(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
