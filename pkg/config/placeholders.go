package config

import (
	"strings"
)

const (
	// placeholderBasePath used inside linters to evaluate relative paths.
	placeholderBasePath = "${base-path}"

	// placeholderConfigPath used inside linters to paths relative to configuration.
	placeholderConfigPath = "${config-path}"
)

func NewPlaceholderReplacer(cfg *Config) *strings.Replacer {
	// Linters that compile globs with '/' as the separator (notably depguard)
	// match slash-normalized paths. On Windows, GetBasePath/GetConfigDir use
	// native backslashes, which gobwas/glob treats as escape characters, so
	// patterns like "${base-path}/pkg/a.go" never match.
	//
	// filepath.ToSlash is a no-op for '\' on non-Windows GOOS, so replace
	// backslashes explicitly for placeholder expansion only. Callers that need
	// OS paths still use GetBasePath/GetConfigDir directly.
	return strings.NewReplacer(
		placeholderBasePath, slashPath(cfg.GetBasePath()),
		placeholderConfigPath, slashPath(cfg.GetConfigDir()),
	)
}

func slashPath(p string) string {
	return strings.ReplaceAll(p, `\`, `/`)
}
