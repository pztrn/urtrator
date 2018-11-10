// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package colorizer

import (
	// stdlib
	"fmt"
	"html"
	"strings"
)

type Colorizer struct {
	// RAW colors to Pango relation.
	colors map[string]string
}

func (c *Colorizer) ClearFromMarkup(data string) string {
	var result string = ""

	data = html.EscapeString(data)

	data_splitted := strings.Split(data, "&gt;")

	if len(data_splitted) > 1 {
		for item := range data_splitted {
			if len(data_splitted[item]) > 0 {
				result += strings.Split(data_splitted[item], "&lt;")[0]
			}
		}
	} else {
		result = data_splitted[0]
	}

	return result
}

func (c *Colorizer) Fix(data string) string {
	var result string = ""

	data = html.EscapeString(data)

	data_splitted := strings.Split(data, "^")
	if len(data_splitted) > 1 {
		for item := range data_splitted {
			if len(data_splitted[item]) > 0 {
				colorcode_raw := string([]rune(data_splitted[item])[0])
				colorcode, ok := c.colors[colorcode_raw]
				if !ok {
					colorcode = "#000000"
				}
				result += "<span foreground=\"" + colorcode + "\">" + string([]rune(data_splitted[item])[1:]) + "</span>"
			} else {
				result += data_splitted[item]
			}
		}
	} else {
		result = data_splitted[0]
	}
	return "<markup>" + result + "</markup>"
}

func (c *Colorizer) Initialize() {
	fmt.Println("Initializing colorizer...")
	c.initializeStorages()
}

func (c *Colorizer) initializeStorages() {
	c.colors = map[string]string{
		"1": "#cc0000",
		"2": "#00cc00",
		"3": "#eeee00",
		"4": "#1c86ee",
		"5": "#00eeee",
		"6": "#ee00ee",
		"7": "#000000",
		"8": "#000000",
	}
}
