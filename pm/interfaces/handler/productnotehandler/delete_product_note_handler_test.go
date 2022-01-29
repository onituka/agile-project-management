package productnotehandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/productnotehandler/mockproductnoteusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
)

func TestDeleteProductNoteHandlerDeleteProductNote(t *testing.T) {
	type fields struct {
		deleteProductNoteUsecase *mockproductnoteusecase.MockDeleteProductNoteUsecase
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

				in := &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				f.deleteProductNoteUsecase.EXPECT().DeleteProductNote(ctx, in).Return(nil)
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
			fileSuffix: "400",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}).Context()

				in := &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}

				err := apperrors.InvalidParameter

				f.deleteProductNoteUsecase.EXPECT().DeleteProductNote(ctx, in).Return(err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				})
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
				}).Context()

				in := &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
				}

				err := apperrors.NotFound

				f.deleteProductNoteUsecase.EXPECT().DeleteProductNote(ctx, in).Return(err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-41ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
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

				in := &productnoteinput.DeleteProductNoteInput{
					ID:        "52dfc0d0-748e-41ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}

				err := apperrors.InternalServerError

				f.deleteProductNoteUsecase.EXPECT().DeleteProductNote(ctx, in).Return(err)
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
				deleteProductNoteUsecase: mockproductnoteusecase.NewMockDeleteProductNoteUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewDeleteProductNoteHandler(f.deleteProductNoteUsecase)

			r := httptest.NewRequest(http.MethodDelete, "/products/{productID:[a-z0-9-]{36}}/notes/{productNoteID:[a-z0-9-]{36}}", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.DeleteProductNote(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
