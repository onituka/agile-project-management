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

func TestUpdateProjectNoteHandlerUpdateProjectNote(t *testing.T) {
	type fields struct {
		updateProjectNoteUsecase *mockprojectnoteusecase.MockUpdateProjectNoteUsecase
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
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				out := &projectnoteoutput.UpdateProjectNoteOutput{
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
				}

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(out, nil)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロダクトID不正",
			fileSuffix: "400-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-xca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトID不正",
			fileSuffix: "400-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-x478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-x478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},

			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-x478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトノートID不正",
			fileSuffix: "400-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-9478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-x478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-x478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-9478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-9478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-x478-0242ac180002",
				})
			},
		},
		{
			name:       "タイトル不正",
			fileSuffix: "400-4",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InvalidParameter

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:        "jsonデコード失敗",
			fileSuffix:  "400-5",
			prepareMock: nil,
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロダクトが存在しない",
			fileSuffix: "404-1",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.NotFound

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトが存在しない",
			fileSuffix: "404-2",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.NotFound

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "プロジェクトノートが存在しない",
			fileSuffix: "404-3",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.NotFound

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "Title重複",
			fileSuffix: "409",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.Conflict

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
		{
			name:       "DBエラー",
			fileSuffix: "500",
			prepareMock: func(f *fields) {
				ctx := mux.SetURLVars(&http.Request{}, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				}).Context()

				in := &projectnoteinput.UpdateProjectNoteInput{
					ID:        "777d71d6-1d03-41ec-a478-0242ac180002",
					ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					ProjectID: "024d71d6-1d03-41ec-a478-0242ac180002",
					UserID:    "777d78d6-1d03-41ec-a478-0242ac184402",
					Title:     "プロジェクトまとめ",
					Content:   "プロジェクトに関する資料まとめ",
				}

				err := apperrors.InternalServerError

				f.updateProjectNoteUsecase.EXPECT().UpdateProjectNote(ctx, in).Return(nil, err)
			},
			prepareRequest: func(r *http.Request) {
				*r = *mux.SetURLVars(r, map[string]string{
					"productID":     "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
					"projectID":     "024d71d6-1d03-41ec-a478-0242ac180002",
					"projectNoteID": "777d71d6-1d03-41ec-a478-0242ac180002",
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				updateProjectNoteUsecase: mockprojectnoteusecase.NewMockUpdateProjectNoteUsecase(gmctrl),
			}

			if tt.prepareMock != nil {
				tt.prepareMock(&f)
			}

			h := NewUpdateProjectNoteHandler(f.updateProjectNoteUsecase)

			r := httptest.NewRequest(http.MethodPut, "/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/{projectNoteID:[a-z0-9-]{36}}", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareRequest != nil {
				tt.prepareRequest(r)
			}

			h.UpdateProjectNote(w, r)

			res := w.Result()
			defer res.Body.Close()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
