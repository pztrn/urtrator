// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2020, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
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

package ui

import (
	// local
	"go.dev.pztrn.name/urtrator/common"

	// other
	"github.com/mattn/go-gtk/gtk"
)

func ShowAboutDialog() {
	ad := gtk.NewAboutDialog()

	ad.SetProgramName("URTrator")
	ad.SetComments(ctx.Translator.Translate("Urban Terror servers browser and game launcher", nil))
	ad.SetVersion(common.URTRATOR_VERSION)
	ad.SetWebsite("https://gitlab.com/pztrn/urtrator")
	ad.SetLogo(logo)

	// ToDo: put it in plain text files.
	var authors []string
	authors = append(authors, "Stanislav N. aka pztrn - project creator, main developer.")
	ad.SetAuthors(authors)

	var artists []string
	artists = append(artists, "UrTConnector team, for great icons and allowing to use them.")
	ad.SetArtists(artists)

	var documenters []string
	documenters = append(documenters, "No one at this moment")
	ad.SetDocumenters(documenters)

	ad.SetCopyright("Stanislav N. aka pztrn")
	ad.SetLicense(MITLicense)

	ad.Run()
	ad.Destroy()
}

var MITLicense = `Copyright 2016-2018, Stanislav N. aka pztrn (or p0z1tr0n) and
URTrator contributors.

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.`
