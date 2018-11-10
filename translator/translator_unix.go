// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2017, Stanislav N. aka pztrn.
// All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// +build !windows

package translator

import (
	// stdlib
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Detect language on Unices.
func (t *Translator) detectLanguage() {
	// Use LC_ALL first.
	t.Language = os.Getenv("LC_ALL")
	// If LC_ALL is empty - use LANG.
	if t.Language == "" {
		t.Language = os.Getenv("LANG")
	}

	// If still nothing - force "en_US" as default locale. Otherwise
	// split language string by "." and take first part.
	// Note: en_US is a default thing, so you will not found anything
	// in "translations" directory!
	if t.Language == "" {
		fmt.Println("No locale data for current user found, using default (en_US)...")
		t.Language = "en_US"
	} else {
		t.Language = strings.Split(t.Language, ".")[0]
	}
}

func (t *Translator) detectTranslationsDirectory() error {
	// Try to use directory near binary.
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// ..which can be overriden by URTRATOR_BINDIR environment variable.
	// Useful for developers.
	envdir := os.Getenv("URTRATOR_BINDIR")
	if envdir != "" {
		dir = envdir
	}

	translations_dir := filepath.Join(dir, "translations")
	_, err := os.Stat(translations_dir)
	if err != nil {
		fmt.Println("Translations wasn't found near binary!")
		// As we're using JSON translation storage, it should be
		// put in /usr/share/urtrator/translations by package
		// maintainers in distros.
		fmt.Println("Trying /usr/share/urtrator/translations...")
		_, err := os.Stat("/usr/share/urtrator/translations")
		if err != nil {
			t.Language = "en_US"
			fmt.Println("Translations unavailable, forcing en_US language code")
			return errors.New("No translations directory was detected!")
		} else {
			t.translationsPath = "/usr/share/urtrator/translations"
		}
	} else {
		t.translationsPath = translations_dir
	}

	return nil
}
