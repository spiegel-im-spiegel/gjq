package query

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/savaki/jq"
	"github.com/spiegel-im-spiegel/gjq/errs"
)

//Op is class for JSON query
type Op struct {
	data   []byte
	indent int
	tab    bool
}

//OptFunc is self-referential function for functional options pattern
type OptFunc func(*Op)

//New returns new Op instance
func New(j []byte, opts ...OptFunc) *Op {
	o := &Op{data: j, indent: 2, tab: false}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

//WithIndent returns function for setting Reader
func WithIndent(i int) OptFunc {
	return func(o *Op) {
		if i < 0 {
			o.indent = 0
		} else {
			o.indent = i
		}
	}
}

//WithTab returns function for setting Reader
func WithTab(t bool) OptFunc {
	return func(o *Op) {
		o.tab = t
	}
}

//Query returns result of query
func (o *Op) Query(filter string) ([]byte, error) {
	if o == nil {
		return nil, errs.Wrap(errs.ErrNullPointer, "cannot parse JSON data")
	}
	op, err := jq.Parse(filter)
	if err != nil {
		return nil, errs.Wrapf(err, "cannot parse JSON data (filter is \"%v\")", filter)
	}
	res, err := op.Apply(o.data)
	if err != nil {
		return nil, errs.Wrapf(err, "cannot parse JSON data (filter is \"%v\")", filter)
	}

	//format
	if o.tab {
		return formatJSON(res, "\t")
	}
	if o.indent > 0 {
		return formatJSON(res, strings.Repeat(" ", o.indent))
	}
	return compactJSON(res)
}

func formatJSON(j []byte, indent string) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.Indent(buf, j, "", indent)
	return buf.Bytes(), errs.Wrap(err, "format error in JSON string")
}

func compactJSON(j []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, j)
	return buf.Bytes(), errs.Wrap(err, "format error in JSON string")
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
