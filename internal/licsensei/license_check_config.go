// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Geza Corp authors
package licsensei

import (
	"fmt"
	"regexp"
	"strings"
)

type LicenseCheckConfig struct {
	Copyrights   []string
	Authors      []string
	LicenseTypes []string
	LicenseTexts []string
}

func isLicenseCheckConfigMatch(lcc LicenseCheckConfig, pcg parsedCommentGroup) (bool, error) {
	if ok, err := pcg.isCopyrightMatch(lcc); err != nil {
		return false, err
	} else if !ok {
		return false, nil
	}

	if match := pcg.isLicenseMatch(lcc); !match {
		return false, nil
	}

	return true, nil
}

func (c LicenseCheckConfig) IsLicenseMatch(g parsedCommentGroup) bool {
	if len(c.LicenseTypes) == 0 && len(c.LicenseTexts) == 0 {
		return true
	}

	re := regexp.MustCompile(`\W+`)
	g.comment = re.ReplaceAllString(g.comment, "")

	for _, licenseType := range c.LicenseTypes {
		// either match spdx license identifier
		for _, spdx := range g.spdxs {
			spdx = strings.ToLower(spdx)
			spdx = strings.TrimSpace(strings.TrimPrefix(spdx, "spdx-license-identifier:"))

			if strings.ToLower(licenseType) == spdx {
				return true
			}
		}

		// or match actual license header
		if t, ok := knownLicenses[licenseType]; ok && re.ReplaceAllString(t, "") == g.comment {
			return true
		}
	}

	for _, lt := range c.LicenseTexts {
		if lt != "" && g.comment == re.ReplaceAllString(lt, "") {
			return true
		}
	}

	return false
}

// Original implementation of the copyright match regex is from the Licensei project (https://github.com/goph/licensei).
func (c LicenseCheckConfig) IsCopyrightMatch(copyright string) (bool, error) {
	if len(c.Copyrights) == 0 {
		return true, nil
	}

	for _, cr := range c.Copyrights {
		cr = regexp.QuoteMeta(cr)

		// add word boundaries
		cr = strings.ReplaceAll(cr, ":YEAR:", "([0-9]{4}[,-]?[\\s]?(\\band \\b)?)+")
		if len(c.Authors) > 0 {
			authors := fmt.Sprintf(`(\b%s\b)`, strings.Join(c.Authors, `\b|\b`))
			// fix end . word boundary mishap
			authors = strings.ReplaceAll(authors, `.\b`, `\b.`)
			cr = strings.ReplaceAll(cr, ":AUTHOR:", authors)
		}

		matched, err := regexp.MatchString(cr, copyright)
		if err != nil {
			return false, err
		}

		if matched {
			return matched, nil
		}
	}

	return false, nil
}
