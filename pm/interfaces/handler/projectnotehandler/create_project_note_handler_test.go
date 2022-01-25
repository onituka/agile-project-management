package projectnotehandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateProjectNoteHandlerCreateProjectNote(t *testing.T) {
	type fields struct {
		createProjectNoteUsecase *mockprojectnoteusecase.MockCreateProjectNoteUsecase
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
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				out := &projectnoteoutput.CreateProjectNoteOutput{
					ID:        "52dfc0d0-748e-11ec-88fd-acde48001122",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
					CreatedBy: "777d78d6-1d03-11ec-a478-0242ac184402",
					UpdatedBy: "777d78d6-1d03-11ec-a478-0242ac184402",
					CreatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 11, 05, 0, 0, 0, 0, time.UTC),
				}

				f.createProjectNoteUsecase.EXPECT().CreateProjectNote(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロダクトID不正",
			fileSuffix: "400-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.createProjectNoteUsecase.EXPECT().CreateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトID不正",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-x478-0242ac180002",
				}).Context()

				in := &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-x478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.createProjectNoteUsecase.EXPECT().CreateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-x478-0242ac180002",
				})
			},
		},
		{
			name:       "タイトル不正",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.createProjectNoteUsecase.EXPECT().CreateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:        "jsonデコード失敗",
			fileSuffix:  "400-4",
			prepareMock: nil,
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "Title重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.Conflict

				f.createProjectNoteUsecase.EXPECT().CreateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.CreateProjectNoteInput{
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-11ec-a478-0242ac180002",
					GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InternalServerError

				f.createProjectNoteUsecase.EXPECT().CreateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID": "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID": "024d71d6-1d03-11ec-a478-0242ac180002",
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				createProjectNoteUsecase: mockprojectnoteusecase.NewMockCreateProjectNoteUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewCreateProjectNoteHandler(f.createProjectNoteUsecase)

			r := httptest.NewRequest(http.MethodPost, "/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.CreateProjectNote(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
