// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2018, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
// URTrator contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
