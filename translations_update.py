#!/usr/bin/env python3

# URTrator - Urban Terror server browser and game launcher, written in
# Go.
#
# Copyright (c) 2016-2020, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
# URTrator contributors.
#
# Permission is hereby granted, free of charge, to any person obtaining
# a copy of this software and associated documentation files (the
# "Software"), to deal in the Software without restriction, including
# without limitation the rights to use, copy, modify, merge, publish,
# distribute, sublicense, and/or sell copies of the Software, and to
# permit persons to whom the Software is furnished to do so, subject
# to the following conditions:
#
# The above copyright notice and this permission notice shall be
# included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
# IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
# CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
# TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
# OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


# This is a translations updating tool. It will parse thru all
# source, compile JSONs and put them into "translations/en_US"
# directory for later use.

import glob
import json
import os
import sys

FILES_TO_TRANSLATE = []
TRANSLATABLE_JSON = {}
SOURCE_PATH = ""

def get_files_for_translation():
    print("Getting files list for translation...")
    SOURCE_PATH = os.path.dirname(os.path.realpath(__file__))
    print("Directory:", SOURCE_PATH)
    for file in glob.iglob(os.path.join(SOURCE_PATH, "**", "*.go")):
        print(file)
        FILES_TO_TRANSLATE.append(file)

    for file in glob.iglob(os.path.join(SOURCE_PATH, "ui", "**", "*.go")):
        print(file)
        FILES_TO_TRANSLATE.append(file)

    # Append main file also.
    FILES_TO_TRANSLATE.append("urtrator.go")

    print("Found " + str(len(FILES_TO_TRANSLATE)) + " files.")

def generate_default_translation():
    print("Generating default translations...")
    for file in FILES_TO_TRANSLATE:
        translations_found = False

        file_data = open(os.path.join(SOURCE_PATH, file), "r").read().split("\n")
        for line in file_data:
            if "Translator.Translate(" in line:
                translations_found = True
                translatable_string = line.split("Translator.Translate(")[1].split("\", ")[0]
                # Skip variables translation.
                if translatable_string[1] != "\"":
                    continue
                translatable_string = translatable_string[1:]
                TRANSLATABLE_JSON[translatable_string] = translatable_string

    # Just a stat :)
    print("Got " + str(len(TRANSLATABLE_JSON)) + " translations.")

def save_sections():
    print("Saving translations...")
    file_to_write = open(os.path.join(SOURCE_PATH, "translations", "en_US", "strings.json"), "w")
    file_to_write.write(json.dumps(TRANSLATABLE_JSON, indent = 4))
    file_to_write.close()

    print("Saving empty translation...")
    for item in TRANSLATABLE_JSON:
        TRANSLATABLE_JSON[item] = ""

    file_to_write = open(os.path.join(SOURCE_PATH, "translations", "empty", "strings.json"), "w")
    file_to_write.write(json.dumps(TRANSLATABLE_JSON, indent = 4))
    file_to_write.close()

get_files_for_translation()
generate_default_translation()
save_sections()
