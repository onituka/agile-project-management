package productnotehandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestUpdateProductNoteHandlerUpdateProductNote(t *testing.T) {
	type fields struct {
		updateProductNoteUsecase *mockproductnoteusecase.MockUpdateProductNoteUsecase
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
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				out := &productnoteoutput.UpdateProductNoteOutput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
				}

				f.updateProductNoteUsecase.EXPECT().UpdateProductNote(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:        "jsonデコード失敗",
			fileSuffix:  "400-1",
			prepareMock: nil,
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "プロダクトID不正(リクエスト値不正)",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}).Context()

				in := &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InvalidParameter

				f.updateProductNoteUsecase.EXPECT().UpdateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				})
			},
		},
		{
			name:       "タイトル不正(Body値不正)",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InvalidParameter

				f.updateProductNoteUsecase.EXPECT().UpdateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
				}).Context()

				in := &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.NotFound

				f.updateProductNoteUsecase.EXPECT().UpdateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
				})
			},
		},
		{
			name:       "Title重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.Conflict

				f.updateProductNoteUsecase.EXPECT().UpdateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.UpdateProductNoteInput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InternalServerError

				f.updateProductNoteUsecase.EXPECT().UpdateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productNoteID": "52dfc0d0-748e-11ec-88fd-acde48001122",
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				updateProductNoteUsecase: mockproductnoteusecase.NewMockUpdateProductNoteUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewUpdateProductNoteHandler(f.updateProductNoteUsecase)

			r := httptest.NewRequest(http.MethodPut, "/products/{productID:[a-z0-9-]{36}}/notes/{productNoteID:[a-z0-9-]{36}}", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.UpdateProductNote(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
