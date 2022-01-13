package producthandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/producthandler/mockproductusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

func TestUpdateProductHandlerUpdateProduct(t *testing.T) {
	type fields struct {
		updateProductUsecase *mockproductusecase.MockUpdateProductUsecase
	}
	tests := []struct {
		name           string
		fileSuffix     string
		prepareMock    func(f *fields)
		prepareRequest func(r *http.Request)
	}{
		{
			name:       "200-正常",
			fileSuffix: "200",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productinput.UpdateProductInput{
					ID:       "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				out := &productoutput.UpdateProductOutput{
					ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:      "プロジェクト管理ツール",
					LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
					CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				}

				f.updateProductUsecase.EXPECT().UpdateProduct(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:        "400-1-jsonデコード失敗",
			fileSuffix:  "400-1",
			prepareMock: nil,
		},
		{
			name:       "400-2-プロダクトID不正",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6xxxxxxxx",
				}).Context()

				in := &productinput.UpdateProductInput{
					ID:       "4495c574-34c2-4fb3-9ca4-3a7c79c267a6xxxxxxxx",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}
				err := apperrors.InvalidParameter

				f.updateProductUsecase.EXPECT().UpdateProduct(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6xxxxxxxx",
				})
			},
		},
		{
			name:       "404-IDが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productinput.UpdateProductInput{
					ID:       "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}
				err := apperrors.NotFound

				f.updateProductUsecase.EXPECT().UpdateProduct(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "409-プロダクト名重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productinput.UpdateProductInput{
					ID:       "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}
				err := apperrors.Conflict

				f.updateProductUsecase.EXPECT().UpdateProduct(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "500-DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productinput.UpdateProductInput{
					ID:       "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Name:     "プロジェクト管理ツール",
					LeaderID: "024d78d6-1d03-11ec-a478-0242ac184402",
				}
				err := apperrors.InternalServerError

				f.updateProductUsecase.EXPECT().UpdateProduct(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				updateProductUsecase: mockproductusecase.NewMockUpdateProductUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewUpdateProductHandler(f.updateProductUsecase)

			r := httptest.NewRequest(http.MethodPut, "/products/4495c574-34c2-4fb3-9ca4-3a7c79c267a6", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.UpdateProduct(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
