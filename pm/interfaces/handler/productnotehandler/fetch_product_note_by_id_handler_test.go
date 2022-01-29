package productnotehandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/productnotehandler/mockproductnoteusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

func TestFetchProductNoteByIDHandlerFetchProductNoteByID(t *testing.T) {
	type fields struct {
		fetchProductNoteByIDUsecase *mockproductnoteusecase.MockFetchProductNoteByIDUsecase
	}
	tests := []struct {
		name           string
		fileSuffix     string
		prepareMock    func(f *fields)
		prepareRequest func(r *http.Request)
	}{
		{
			name:       "正常",
			fileSuffix: "200",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.FetchProductNoteByIDInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				out := &productnoteoutput.FetchProductNoteByIDOutput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-41ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-41ec-a478-0242ac184402",
					CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
				}

				f.fetchProductNoteByIDUsecase.EXPECT().FetchProductNoteByID(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "プロダクトID不正(リクエスト値不正)",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}).Context()

				in := &productnoteinput.FetchProductNoteByIDInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}

				err := apperrors.InvalidParameter

				f.fetchProductNoteByIDUsecase.EXPECT().FetchProductNoteByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				})
			},
		},
		{
			name:       "IDが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001129",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.FetchProductNoteByIDInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001129",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				err := apperrors.NotFound

				f.fetchProductNoteByIDUsecase.EXPECT().FetchProductNoteByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001129",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.FetchProductNoteByIDInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				err := apperrors.InternalServerError

				f.fetchProductNoteByIDUsecase.EXPECT().FetchProductNoteByID(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				fetchProductNoteByIDUsecase: mockproductnoteusecase.NewMockFetchProductNoteByIDUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewFetchProductNoteByIDHandler(f.fetchProductNoteByIDUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products/{productID:[a-z0-9-]{36}}/notes/{productNoteID:[a-z0-9-]{36}}", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.FetchProductNoteByID(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
