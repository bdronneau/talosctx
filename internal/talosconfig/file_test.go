package talosconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/bdronneau/talosctx/internal/env"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestGetTalosConfigEnv(t *testing.T) {
	expectedPath := "../../tests/assets/.talos/config"

	t.Setenv(env.TALOS_ENV_CONFIG_NAME, expectedPath)

	expectedData, _ := os.ReadFile(expectedPath)

	content, path, err := GetTalosConfigContent()
	assert.Nil(t, err)
	assert.Equal(t, expectedData, content)
	assert.Equal(t, expectedPath, path)

	t.Setenv(env.TALOS_ENV_CONFIG_NAME, "no/path")

	_, _, err = GetTalosConfigContent()
	assert.ErrorContains(t, err, "no such file or directory")
}

func TestGetTalosConfigHome(t *testing.T) {
	expectedHome := "../../tests/assets/"
	expectedPath := filepath.Join("../../tests/assets/", ".talos/config")

	expectedData, _ := os.ReadFile(expectedPath)

	getHomeDir = func() (string, error) {
		return expectedHome, nil
	}

	content, path, err := GetTalosConfigContent()
	assert.Nil(t, err)
	assert.Equal(t, expectedData, content)
	assert.Equal(t, expectedPath, path)

	getHomeDir = func() (string, error) {
		return expectedHome, fmt.Errorf("Oh no")
	}

	_, _, err = GetTalosConfigContent()
	assert.Error(t, fmt.Errorf("Oh no"), err)

	getHomeDir = func() (string, error) {
		return "no/path", nil
	}

	_, _, err = GetTalosConfigContent()
	assert.ErrorContains(t, err, "no such file or directory")
}

func TestWritetalosConfig(t *testing.T) {
	guid := xid.New()
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("%s_%s", "talosctx_test", guid.String()))
	if err != nil {
		t.Fatalf("can not create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	expectedData := []byte("test data")
	expecteFilePath := tmpDir + "/testfile"
	expecteFileMode := os.FileMode(0o600)

	err = WritetalosConfig(expecteFilePath, expectedData)
	assert.ErrorContains(t, err, "no such file or directory")

	err = os.WriteFile(expecteFilePath, expectedData, expecteFileMode)
	assert.NoError(t, err)

	actualFile, err := os.Stat(expecteFilePath)
	assert.NoError(t, err)
	assert.Equal(t, expecteFileMode, actualFile.Mode())

	actualData, err := os.ReadFile(expecteFilePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, actualData)
}
