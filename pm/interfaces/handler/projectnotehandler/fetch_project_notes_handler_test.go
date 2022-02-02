package projectnotehandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projectnotehandler/mockprojectnoteusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

func TestFetchProjectNotesHandlerFetchProjectNotes(t *testing.T) {
	type fields struct {
		fetchProjectNotesUsecase *mockprojectnoteusecase.MockFetchProjectNotesUsecase
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
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				out := &projectnoteoutput.FetchProjectNotesOutput{
					TotalCount: 2,
					ProjectNotes: []*projectnoteoutput.ProjectNoteOutput{
						{
							ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
							ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
							GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
							Title:     "プロジェクトまとめ",
							Content:   "プロジェクトに関する資料まとめ",
							CreatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
							UpdatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
							CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						},
						{
							ID:        "888d71d6-1d03-41ec-a478-0242ac180002",
							ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
							GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
							Title:     "プロジェクトまとめ1",
							Content:   "プロジェクトに関する資料まとめ",
							CreatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
							UpdatedBy: "777d78d6-1d03-41ec-a478-0242ac184402",
							CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
						},
					},
				}

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "正常(プロジェクトノートが存在しない)",
			fileSuffix: "200-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				out := &projectnoteoutput.FetchProjectNotesOutput{
					TotalCount:   0,
					ProjectNotes: make([]*projectnoteoutput.ProjectNoteOutput, 0),
				}

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロダクトID不正",
			fileSuffix: "400-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				err := apperrors.InvalidParameter

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロジェクトID不正",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-x478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-x478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				err := apperrors.InvalidParameter

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-x478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:        "クエリパラメータ文字列によるpage不正",
			fileSuffix:  "400-3",
			prepareMock: nil,

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "test")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:        "クエリパラメータ文字列によるlimit不正",
			fileSuffix:  "400-4",
			prepareMock: nil,

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "test")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "usecaseでの400エラー(page値不正)",
			fileSuffix: "400-5",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      0,
					Limit:     50,
				}

				err := apperrors.InvalidParameter

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "0")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				err := apperrors.NotFound

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロジェクトが存在しない",
			fileSuffix: "404-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				err := apperrors.NotFound

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.FetchProjectNotesInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					Page:      1,
					Limit:     50,
				}

				err := apperrors.InternalServerError

				f.fetchProjectNotesUsecase.EXPECT().FetchProjectNotes(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-41ec-a478-0242ac180002",
				})

				q := r.URL.Query()
				q.Set("page", "1")
				q.Set("limit", "50")
				r.URL.RawQuery = q.Encode()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				fetchProjectNotesUsecase: mockprojectnoteusecase.NewMockFetchProjectNotesUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewFetchProjectNotesHandler(f.fetchProjectNotesUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.FetchProjectNotes(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
