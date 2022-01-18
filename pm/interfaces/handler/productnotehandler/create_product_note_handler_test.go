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

func TestCreateProductNoteHandlerCreateProductNote(t *testing.T) {
	type fields struct {
		createProductNoteUsecase *mockproductnoteusecase.MockCreateProductNoteUsecase
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
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.CreateProductNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				out := &productnoteoutput.CreateProductNoteOutput{
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

				f.createProductNoteUsecase.EXPECT().CreateProductNote(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:        "jsonデコード失敗",
			fileSuffix:  "400-1",
			prepareMock: nil,
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "プロダクトID不正(リクエスト値不正)",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6xxx",
				}).Context()

				in := &productnoteinput.CreateProductNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6xxx",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InvalidParameter

				f.createProductNoteUsecase.EXPECT().CreateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6xxx",
				})
			},
		},
		{
			name:       "グループID不正(Body値不正)",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.CreateProductNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002xxx",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InvalidParameter

				f.createProductNoteUsecase.EXPECT().CreateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				}).Context()

				in := &productnoteinput.CreateProductNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.NotFound

				f.createProductNoteUsecase.EXPECT().CreateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267ax",
				})
			},
		},
		{
			name:       "Title重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.CreateProductNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.Conflict

				f.createProductNoteUsecase.EXPECT().CreateProductNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.CreateProductNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
					Title:     "ノート",
					Content:   "note",
					CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
				}

				err := apperrors.InternalServerError

				f.createProductNoteUsecase.EXPECT().CreateProductNote(ctx, in).Return(nil, err)
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
				createProductNoteUsecase: mockproductnoteusecase.NewMockCreateProductNoteUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewCreateProductNoteHandler(f.createProductNoteUsecase)

			r := httptest.NewRequest(http.MethodPost, "/products/{productID}/productnotes", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.CreateProductNote(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
