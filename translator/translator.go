// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.

// This is translator - package which translates everything in languages
// other than en_US.
// Available only on unixes for now, Windows version in ToDo.
package translator

import (
	// stdlib
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Translator struct {
	// Accepted languages.
	AcceptedLanguages map[string]string
	// Currently active language.
	Language string
	// Translations.
	translations map[string]map[string]string
	// Path to translations files.
	translationsPath string
}

// Formats string from passed map.
// We expect replaceable strings to be named as:
//
//      {{ VAR }}
//
// We will change whole "{{ VAR }}" to params["VAR"] value. E.g.:
//
//      data = "Version {{ version }}"
//      map = map[string]string{"version": "0.1"}
//
// will be formatted as:
//
//      result = "Version 0.1"
//
// Also note that we will replace ALL occurences of "{{ VAR }}" within string!
// All untranslated variables will not be touched at all.
func (t *Translator) formatFromMap(data string, params map[string]string) string {
	new_data := data
	for k, v := range params {
		new_data = strings.Replace(new_data, "{{ "+k+" }}", v, -1)
	}
	return new_data
}

// Translator initialization.
func (t *Translator) Initialize() {
	fmt.Println("Initializing translations...")

	t.AcceptedLanguages = map[string]string{
		"System's default language": "default",
		"English":                   "en_US",
		"French":                    "fr_FR",
		"Russian":                   "ru_RU",
	}

	// Initialize storages.
	t.translations = make(map[string]map[string]string)
	t.translationsPath = ""

	// Getting locale name from environment.
	// ToDo: Windows compatability. Possible reference:
	// https://github.com/cloudfoundry-attic/jibber_jabber/blob/master/jibber_jabber_windows.go
	t.detectLanguage()

	fmt.Println("Using translations for '" + t.Language + "'")

	err := t.detectTranslationsDirectory()
	if err == nil {
		t.loadTranslations()
	} else {
		fmt.Println("Skipping translations loading due to missing translations directory.")
	}
}

// Load translations into memory.
func (t *Translator) loadTranslations() {
	fmt.Println("Loading translations for language " + t.Language)
	fmt.Println("Translations directory: " + t.translationsPath)

	t.translations[t.Language] = make(map[string]string)

	if t.translationsPath != "" {
		// Check if language was selected in options dialog. In that
		// case it will overwrite autodetected language.
		var translationsDir string = ""
		if cfg.Cfg["/general/language"] != "" {
			translationsDir = filepath.Join(t.translationsPath, cfg.Cfg["/general/language"])
			t.Language = cfg.Cfg["/general/language"]
		} else {
			translationsDir = filepath.Join(t.translationsPath, t.Language)
		}
		files_list, _ := ioutil.ReadDir(translationsDir)
		if len(files_list) > 0 {
			for i := range files_list {
				// Read file.
				file_path := filepath.Join(translationsDir, files_list[i].Name())
				file_data, _ := ioutil.ReadFile(file_path)
				var trans map[string]string
				json.Unmarshal(file_data, &trans)
				// Assign parsed translations to language code.
				t.translations[t.Language] = trans
			}
		}
	}
}

// Actual translation function.
// Parameters:
//   * str - raw string from source which we will try to translate.
//   * params - map[string]string with parameters to replace in string.
//
// If string wasn't found in loaded translations for current language
// we will use English as fallback.
// Translates passed data from loaded translations file.
// Returns passed data without changes if translation wasn't found.
func (t *Translator) Translate(data string, params map[string]string) string {
	val, ok := t.translations[t.Language][data]
	if !ok {
		if params != nil && len(params) > 0 {
			return t.formatFromMap(data, params)
		} else {
			return data
		}
	}

	if params != nil && len(params) > 0 {
		return t.formatFromMap(val, params)
	}

	return val
}
