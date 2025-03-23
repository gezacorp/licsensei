// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Geza Corp authors
package licsensei

import (
	"go/ast"
	"strings"
)

type parsedCommentGroup struct {
	spdxs      []string
	copyrights []string
	comment    string
}

func parseCommentGroup(comments *ast.CommentGroup) parsedCommentGroup {
	pcg := parsedCommentGroup{
		spdxs:      []string{},
		copyrights: []string{},
	}

	for line := range strings.SplitSeq(comments.Text(), "\n") {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "* ")

		if strings.HasPrefix(line, "+") {
			continue
		}

		if strings.HasPrefix(strings.ToLower(line), "spdx-license-identifier") {
			pcg.spdxs = append(pcg.spdxs, line)

			continue
		}

		if strings.HasPrefix(strings.ToLower(line), "copyright") {
			pcg.copyrights = append(pcg.copyrights, line)

			continue
		}

		pcg.comment += line + "\n"
	}

	pcg.comment = strings.TrimSpace(pcg.comment)

	return pcg
}

func (g parsedCommentGroup) isCopyrightMatch(lcc LicenseCheckConfig) (bool, error) {
	if len(lcc.Copyrights) == 0 {
		return true, nil
	}

	for _, copyright := range g.copyrights {
		match, err := lcc.IsCopyrightMatch(copyright)
		if err != nil {
			return false, err
		}

		if match {
			return true, nil
		}
	}

	return false, nil
}

func (g parsedCommentGroup) isLicenseMatch(lcc LicenseCheckConfig) bool {
	if len(lcc.LicenseTypes) == 0 && len(lcc.LicenseTexts) == 0 {
		return true
	}

	if match := lcc.IsLicenseMatch(g); match {
		return match
	}

	return false
}
