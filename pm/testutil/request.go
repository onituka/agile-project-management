package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	// folder name for where the request fixtures are stored
	requestFixtureDir = "request"

	// request file name info
	requestFileNamePrefix = "request"
	requestFileNameSuffix = ".json"
)

func GetRequestJsonFromTestData(t *testing.T, fileNameOpts ...string) string {
	t.Helper()

	requestFilePath := newRequestFilePath(t, fileNameOpts...)

	byteData, err := os.ReadFile(requestFilePath)
	if err != nil {
		t.Fatalf("unexpected error by os.ReadFile '%v'", err)
	}

	return string(byteData)
}

func newRequestFilePath(t *testing.T, fileNameOpts ...string) string {
	var requestFilePath string
	testFuncName := strings.Split(t.Name(), "/")[0]

	if len(fileNameOpts) == 0 {
		// file path example: ./testdata/request/{testFuncName}/request.json
		requestFilePath = filepath.Join(
			baseFixtureDir,
			requestFixtureDir,
			testFuncName,
			fmt.Sprintf("%s%s", requestFileNamePrefix, requestFileNameSuffix),
		)
	} else {
		optsStr := strings.Join(fileNameOpts, "-")

		// file path example: ./testdata/request/{testFuncName}/request-{fileNameOpts}.json
		requestFilePath = filepath.Join(
			baseFixtureDir,
			requestFixtureDir,
			testFuncName,
			fmt.Sprintf("%s-%s%s", requestFileNamePrefix, optsStr, requestFileNameSuffix),
		)
	}

	return requestFilePath
}
