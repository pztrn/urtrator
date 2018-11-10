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

// +build windows

package translator

import (
	// stdlib
	"fmt"
	"os"
	"path/filepath"
)

// Detect language on Windows.
func (t *Translator) detectLanguage() {
	fmt.Println("ToDo! Forcing en_US for now!")
	t.Language = "en_US"
}

func (t *Translator) detectTranslationsDirectory() error {
	// Translations MUST reside in directory neear binary!
	// ToDo: more checks.
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	t.translationsPath = filepath.Join(dir, "translations")

	return nil
}
