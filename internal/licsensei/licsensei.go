// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Geza Corp authors
package licsensei

import (
	"context"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"emperror.dev/errors"
	"github.com/go-logr/logr"
)

type CheckErrors map[string]string

type CheckConfiguration struct {
	IgnorePaths         []string
	IgnoreFiles         []string
	LicenseCheckConfigs []LicenseCheckConfig
}

func Check(ctx context.Context, root string, config CheckConfiguration) (CheckErrors, error) {
	if root == "" {
		return nil, errors.New("root cannot be empty")
	}

	fileErrors := CheckErrors{}
	logger := logr.FromContextOrDiscard(ctx)

	err := filepath.Walk(root, func(path string, info os.FileInfo, fileerr error) error {
		if fileerr != nil {
			return fileerr
		}

		if info.IsDir() || filepath.Ext(path) != ".go" || config.isPathIgnored(root, path) {
			return nil
		}

		logger.V(1).Info("checking", "path", path)

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.PackageClauseOnly|parser.ParseComments)
		if err != nil {
			fileErrors[path] = err.Error()

			//nolint:nilerr
			return nil
		}

		for _, lcc := range config.LicenseCheckConfigs {
			if len(node.Comments) == 0 {
				fileErrors[path] = "missing license header"

				return nil
			}

			for _, cg := range node.Comments {
				pcg := parseCommentGroup(cg)

				match, err := isLicenseCheckConfigMatch(lcc, pcg)
				if err != nil {
					fileErrors[path] = err.Error()

					//nolint:nilerr
					return nil
				}

				if match {
					return nil
				}
			}
		}

		fileErrors[path] = "invalid license header"

		return nil
	})

	return fileErrors, err
}

// This func contains code from [Licensei](https://github.com/goph/licensei).
func (c CheckConfiguration) isPathIgnored(root string, path string) bool {
	for _, ignorePath := range c.IgnorePaths {
		if !filepath.IsAbs(ignorePath) {
			ignorePath = filepath.Join(root, ignorePath)
		}

		if strings.HasPrefix(path, ignorePath+"/") {
			return true
		}
	}

	for _, glob := range c.IgnoreFiles {
		if matched, err := filepath.Match(glob, filepath.Base(path)); err == nil && matched {
			return true
		}
	}

	return false
}
