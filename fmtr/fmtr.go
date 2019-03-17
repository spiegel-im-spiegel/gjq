package fmtr

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
	"github.com/spiegel-im-spiegel/gjq/errs"
)

var (
	regularNormal = color.New()
	regularBold   = color.New(color.Bold)
	cyanNormal    = color.New(color.FgCyan)
	cyanBold      = color.New(color.FgCyan, color.Bold)
	redBold       = color.New(color.FgRed, color.Bold)
)

//Fmtr is class of JSON formatter
type Fmtr struct {
	formatter *jsoncolor.Formatter
	colorable bool
}

//New returns new Fmtr instance
func New(indentSize int, tabs, colorable bool) *Fmtr {
	color.NoColor = false
	formatter := jsoncolor.NewFormatter()
	formatter.SpaceColor = regularNormal
	formatter.CommaColor = regularNormal
	formatter.ColonColor = regularNormal
	formatter.ObjectColor = regularNormal
	formatter.ArrayColor = regularNormal
	formatter.FieldQuoteColor = cyanNormal
	formatter.FieldColor = cyanNormal
	formatter.StringQuoteColor = cyanBold
	formatter.StringColor = cyanBold
	formatter.TrueColor = redBold
	formatter.FalseColor = redBold
	formatter.NumberColor = regularBold
	formatter.NullColor = redBold
	if tabs {
		formatter.Indent = "\t"
	} else if indentSize > 0 {
		formatter.Indent = strings.Repeat(" ", indentSize)
	} else {
		formatter.Indent = ""
	}
	formatter.Prefix = ""

	return &Fmtr{formatter: formatter, colorable: colorable}
}

//FormatRaw retur formatted JSON string
func (f *Fmtr) FormatRaw(j []byte) ([]byte, error) {
	if f == nil {
		return nil, errs.Wrap(errs.ErrNullPointer, "format error in JSON string")
	}
	buf := &bytes.Buffer{}
	if len(f.formatter.Indent) > 0 {
		err := json.Indent(buf, j, f.formatter.Prefix, f.formatter.Indent)
		return buf.Bytes(), errs.Wrap(err, "format error in JSON string")
	} else {
		err := json.Compact(buf, j)
		return buf.Bytes(), errs.Wrap(err, "format error in JSON string")
	}
}

//Format retur formatted JSON string (colorized)
func (f *Fmtr) Format(j []byte) ([]byte, error) {
	if f == nil {
		return nil, errs.Wrap(errs.ErrNullPointer, "format error in JSON string (colorized)")
	}
	if f.colorable {
		buf := &bytes.Buffer{}
		err := f.formatter.Format(buf, j)
		return buf.Bytes(), errs.Wrap(err, "format error in JSON string (colorized)")
	}
	return f.FormatRaw(j)
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
