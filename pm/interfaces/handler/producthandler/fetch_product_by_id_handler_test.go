package producthandler

import (
	"net/http"
	"net/http/httptest"
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

func TestFetchProductByIDHandlerFetchProductByID(t *testing.T) {
	type fields struct {
		fetchProductByIDUsecase *mockproductusecase.MockFetchProductByIDUsecase
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

				in := &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				out := &productoutput.FetchProductByIDOutput{
					ID:        "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Name:      "プロジェクト管理ツール",
					LeaderID:  "024d78d6-1d03-11ec-a478-0242ac184402",
					CreatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
				}

				f.fetchProductByIDUsecase.EXPECT().FetchProductByID(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "400-プロダクトID不正",
			fileSuffix: "400",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}).Context()

				in := &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}

				err := apperrors.InvalidParameter

				f.fetchProductByIDUsecase.EXPECT().FetchProductByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				})
			},
		},
		{
			name:       "404-IDが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a0",
				}).Context()

				in := &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a0",
				}

				err := apperrors.NotFound

				f.fetchProductByIDUsecase.EXPECT().FetchProductByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a0",
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

				in := &productinput.FetchProductByIDInput{
					ID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				err := apperrors.InternalServerError

				f.fetchProductByIDUsecase.EXPECT().FetchProductByID(ctx, in).Return(nil, err)
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
				fetchProductByIDUsecase: mockproductusecase.NewMockFetchProductByIDUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewFetchProductByIDHandler(f.fetchProductByIDUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products/{productID:[a-z0-9-]{36}}", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.FetchProductByID(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
