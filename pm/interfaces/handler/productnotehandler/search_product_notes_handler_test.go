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

func TestSearchProductNotesHandlerSearchProductNotes(t *testing.T) {
	type fields struct {
		searchProductNotesUsecase *mockproductnoteusecase.MockSearchProductNotesUsecase
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
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				}

				out := &productnoteoutput.SearchProductNotesOutput{
					TotalCount: 2,
					ProductNotes: []*productnoteoutput.SearchProductNoteOutput{
						{
							ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
							ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
							Title:     "ノート",
							Content:   "note",
							CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
							UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
							CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						},
						{
							ID:        "62dfc0d0-748e-11ec-88fd-acde48001122",
							ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
							Title:     "ノート2",
							Content:   "test",
							CreatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
							UpdatedBy: "024d78d6-1d03-11ec-a478-0242ac184402",
							CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						},
					},
				}

				f.searchProductNotesUsecase.EXPECT().SearchProductNotes(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("title", "ノ")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "正常(プロダクトノートが存在しない場合)",
			fileSuffix: "200-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				}

				out := &productnoteoutput.SearchProductNotesOutput{
					TotalCount:   0,
					ProductNotes: make([]*productnoteoutput.SearchProductNoteOutput, 0),
				}

				f.searchProductNotesUsecase.EXPECT().SearchProductNotes(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("title", "ノ")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "正常(検索にヒットしなかった場合)",
			fileSuffix: "200-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノxxx",
					Page:      1,
					Limit:     10,
				}

				out := &productnoteoutput.SearchProductNotesOutput{
					TotalCount:   0,
					ProductNotes: make([]*productnoteoutput.SearchProductNoteOutput, 0),
				}

				f.searchProductNotesUsecase.EXPECT().SearchProductNotes(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("title", "ノxxx")
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
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("title", "ノ")
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
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("title", "ノ")
				q.Set("page", "1")
				q.Set("limit", "test")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "usecaseでの400エラー(page値不正)",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      0,
					Limit:     10,
				}

				err := apperrors.InvalidParameter

				f.searchProductNotesUsecase.EXPECT().SearchProductNotes(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("title", "ノ")
				q.Set("page", "0")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロダクトノートが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
				}).Context()

				in := &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				}

				err := apperrors.NotFound

				f.searchProductNotesUsecase.EXPECT().SearchProductNotes(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a9",
				})
				q := r.URL.Query()
				q.Set("title", "ノ")
				q.Set("page", "1")
				q.Set("limit", "10")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "usecaseでの500エラー(DBエラー)",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &productnoteinput.SearchProductNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					Title:     "ノ",
					Page:      1,
					Limit:     10,
				}

				err := apperrors.InternalServerError

				f.searchProductNotesUsecase.EXPECT().SearchProductNotes(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("title", "ノ")
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
				searchProductNotesUsecase: mockproductnoteusecase.NewMockSearchProductNotesUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewSearchProductNotesHandler(f.searchProductNotesUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products/{productID:[a-z0-9-]{36}}/notes/search", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.SearchProductNotes(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
