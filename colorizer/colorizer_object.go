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
