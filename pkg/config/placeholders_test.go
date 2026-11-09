package config

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPlaceholderReplacer_slashNormalized(t *testing.T) {
	t.Parallel()

	cfg := NewDefault()
	cfg.basePath = `D:\a\repo\repo\repro`
	cfg.cfgDir = `D:\a\repo\repo\repro`

	replacer := NewPlaceholderReplacer(cfg)

	got := replacer.Replace("${base-path}/pkg/a.go")
	require.Equal(t, "D:/a/repo/repo/repro/pkg/a.go", got)

	got = replacer.Replace("${config-path}/pkg/a.go")
	require.Equal(t, "D:/a/repo/repo/repro/pkg/a.go", got)

	// Native getters stay unchanged for filesystem use.
	require.Equal(t, `D:\a\repo\repo\repro`, cfg.GetBasePath())
	require.Equal(t, `D:\a\repo\repo\repro`, cfg.GetConfigDir())
}

func TestNewPlaceholderReplacer_unixPathsUnchanged(t *testing.T) {
	t.Parallel()

	cfg := NewDefault()
	cfg.basePath = "/home/runner/work/repo/repro"
	cfg.cfgDir = "/home/runner/work/repo/repro"

	replacer := NewPlaceholderReplacer(cfg)
	require.Equal(t, "/home/runner/work/repo/repro/pkg/a.go", replacer.Replace("${base-path}/pkg/a.go"))
}

func TestNewPlaceholderReplacer_nativeWindowsPath(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("native Windows path construction")
	}
	t.Parallel()

	cfg := NewDefault()
	native := filepath.Join("D:", "a", "repo", "repo", "repro")
	cfg.basePath = native
	cfg.cfgDir = native

	replacer := NewPlaceholderReplacer(cfg)
	require.Equal(t, "D:/a/repo/repo/repro/pkg/a.go", replacer.Replace("${base-path}/pkg/a.go"))
}
