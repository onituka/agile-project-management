package projecthandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projecthandler/mockprojectusecase"
	"github.com/onituka/agile-project-management/project-management/testutil"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

func TestSearchTrashedProjectsHandlerSearchTrashedProjects(t *testing.T) {
	type fields struct {
		searchTrashedProjectsUsecase *mockprojectusecase.MockSearchTrashedProjectsUsecase
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

				in := &projectinput.SearchTrashedProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				}

				trashedAt := time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC)

				out := &projectoutput.SearchTrashedProjectsOutput{
					TotalCount: 2,
					Projects: []*projectoutput.SearchTrashedProjectOutput{
						{
							ID:                "024d71d6-1d03-11ec-a478-0242ac180002",
							ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
							KeyName:           "AAA",
							Name:              "管理ツール1",
							LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
							DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
							TrashedAt:         &trashedAt,
							CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
							UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						},
						{
							ID:                "024d71d6-1d03-11ec-a478-0242ac180003",
							ProductID:         "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
							GroupID:           "024d78d6-1d03-11ec-a478-0242ac180002",
							KeyName:           "BBB",
							Name:              "管理ツールA",
							LeaderID:          "024d78d6-1d03-11ec-a478-0242ac184402",
							DefaultAssigneeID: "024d78d6-1d03-11ec-a478-9242ac182002",
							TrashedAt:         &trashedAt,
							CreatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
							UpdatedAt:         time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC),
						},
					},
				}

				f.searchTrashedProjectsUsecase.EXPECT().SearchTrashedProjects(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "1")
				q.Set("limit", "2")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "正常(プロジェクトが存在しない場合)",
			fileSuffix: "200-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &projectinput.SearchTrashedProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				}
				out := &projectoutput.SearchTrashedProjectsOutput{
					TotalCount: 0,
					Projects:   make([]*projectoutput.SearchTrashedProjectOutput, 0),
				}

				f.searchTrashedProjectsUsecase.EXPECT().SearchTrashedProjects(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "1")
				q.Set("limit", "2")
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

				in := &projectinput.SearchTrashedProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				}
				out := &projectoutput.SearchTrashedProjectsOutput{
					TotalCount: 0,
					Projects:   make([]*projectoutput.SearchTrashedProjectOutput, 0),
				}

				f.searchTrashedProjectsUsecase.EXPECT().SearchTrashedProjects(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "1")
				q.Set("limit", "2")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:        "クエリパラメータ文字列によるpage不正",
			fileSuffix:  "400-1",
			prepareMock: nil,

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "test")
				q.Set("limit", "2")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:        "クエリパラメータ文字列によるlimit不正",
			fileSuffix:  "400-2",
			prepareMock: nil,

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("keyword", "A")
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

				in := &projectinput.SearchTrashedProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      0,
					Limit:     2,
				}

				err := apperrors.InvalidParameter

				f.searchTrashedProjectsUsecase.EXPECT().SearchTrashedProjects(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "0")
				q.Set("limit", "2")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &projectinput.SearchTrashedProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				}

				err := apperrors.NotFound

				f.searchTrashedProjectsUsecase.EXPECT().SearchTrashedProjects(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})
				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "1")
				q.Set("limit", "2")
				r.URL.RawQuery = q.Encode()
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				}).Context()

				in := &projectinput.SearchTrashedProjectsInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					KeyWord:   "A",
					Page:      1,
					Limit:     2,
				}

				err := apperrors.InternalServerError

				f.searchTrashedProjectsUsecase.EXPECT().SearchTrashedProjects(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
				})

				q := r.URL.Query()
				q.Set("keyword", "A")
				q.Set("page", "1")
				q.Set("limit", "2")
				r.URL.RawQuery = q.Encode()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				searchTrashedProjectsUsecase: mockprojectusecase.NewMockSearchTrashedProjectsUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewSearchTrashedProjectsHandler(f.searchTrashedProjectsUsecase)

			r := httptest.NewRequest(http.MethodGet, "/products/{productID}/projects/search/trash-box", nil)
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.SearchTrashedProjects(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
