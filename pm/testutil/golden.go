package testutil

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/xerrors"
)

const (
	baseFixtureDir = "testdata"

	// folder name for where the response fixtures are stored
	responseFixtureDir = "response"

	// golden file name info
	goldenFileNamePrefix = "response"
	goldenFileNameSuffix = ".json.golden"

	// permissions
	filePerms os.FileMode = 0644
	dirPerms  os.FileMode = 0755
)

var (
	update = flag.Bool("update", false, "update golden test files")
	clean  = flag.Bool("clean", false, "clean old golden test files")
)

func AssertResponseBody(t *testing.T, res *http.Response, fileNameOpts ...string) {
	t.Helper()

	goldenFile := newGoldenFilePath(t, fileNameOpts...)

	gotJSON, err := newGotJson(res)
	if err != nil {
		t.Fatalf("unexpected error by newGotJson '%v'", err)
	}

	if *update {
		if err = updateGoldenFile(goldenFile, gotJSON.Bytes()); err != nil {
			t.Fatalf("unexpected error by updateGoldenFile '%v'", err)
		}
	}

	wantJSON, err := newWantJson(goldenFile)
	if err != nil {
		t.Fatalf("unexpected error by newGotJson '%v'", err)
	}

	if diff := cmp.Diff(wantJSON.String(), gotJSON.String()); len(diff) != 0 {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func newGoldenFilePath(t *testing.T, fileNameOpts ...string) string {
	var goldenFilePath string
	testFuncName := strings.Split(t.Name(), "/")[0]

	if len(fileNameOpts) == 0 {
		// file path example: ./testdata/response/{testFuncName}/response.json.golden
		goldenFilePath = filepath.Join(
			baseFixtureDir,
			responseFixtureDir,
			testFuncName,
			fmt.Sprintf("%s%s", goldenFileNamePrefix, goldenFileNameSuffix),
		)
	} else {
		optsStr := strings.Join(fileNameOpts, "-")

		// file path example: ./testdata/response/{testFuncName}/response-{fileNameOpts}.json.golden
		goldenFilePath = filepath.Join(
			baseFixtureDir,
			responseFixtureDir,
			testFuncName,
			fmt.Sprintf("%s-%s%s", goldenFileNamePrefix, optsStr, goldenFileNameSuffix),
		)
	}

	return goldenFilePath
}

func newGotJson(res *http.Response) (*bytes.Buffer, error) {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var compactJson bytes.Buffer
	if err = json.Compact(&compactJson, b); err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, compactJson.Bytes(), "", "  "); err != nil {
		return nil, err
	}

	return &prettyJSON, nil
}

func newWantJson(goldenFile string) (*bytes.Buffer, error) {
	b, err := os.ReadFile(goldenFile)
	if err != nil {
		return nil, err
	}

	var compactJson bytes.Buffer
	if err = json.Compact(&compactJson, b); err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, compactJson.Bytes(), "", "  "); err != nil {
		return nil, err
	}

	return &prettyJSON, nil
}

func updateGoldenFile(goldenFile string, actualData []byte) error {
	goldenFileDir := filepath.Dir(goldenFile)

	if err := ensureDir(goldenFileDir); err != nil {
		return err
	}

	// 最終行に改行挿入
	actualData = append(actualData, '\n')

	if err := os.WriteFile(goldenFile, actualData, filePerms); err != nil {
		return err
	}

	now := time.Now()
	if err := os.Chtimes(goldenFileDir, now, now); err != nil {
		return err
	}

	return nil
}

func ensureDir(goldenFileDir string) error {
	s, err := os.Stat(goldenFileDir)
	defer func() {
		*clean = false
	}()

	switch {
	case err != nil && os.IsNotExist(err):
		return os.MkdirAll(goldenFileDir, dirPerms)

	case err == nil && s.IsDir() && *clean:
		// ./testdata/response folder削除
		if err = os.RemoveAll(filepath.Join(baseFixtureDir, responseFixtureDir)); err != nil {
			return err
		}
		return os.MkdirAll(goldenFileDir, dirPerms)

	case err == nil && !s.IsDir():
		return xerrors.Errorf("fixture folder is a file: %s", goldenFileDir)
	}

	return err
}
