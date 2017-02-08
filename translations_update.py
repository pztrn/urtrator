#!/usr/bin/env python3

# URTator - Urban Terror server browser and game launcher, written in
# Go.
#
# Copyright (c) 2016-2017, Stanislav N. aka pztrn.
# All rights reserved.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
