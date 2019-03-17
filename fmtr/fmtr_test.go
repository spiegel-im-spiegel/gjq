package fmtr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFormatter(t *testing.T) {
	testCases := []struct {
		indent int
		tab    bool
		color  bool
		json   []byte
		res    []byte
	}{
		{indent: 0, tab: false, color: false, json: []byte(`{"foo":"bar"}`), res: []byte(`{"foo":"bar"}`)},
		{indent: 1, tab: false, color: false, json: []byte(`{"foo":"bar"}`), res: []byte("{\n \"foo\": \"bar\"\n}")},
		{indent: 0, tab: true, color: false, json: []byte(`{"foo":"bar"}`), res: []byte("{\n\t\"foo\": \"bar\"\n}")},
		{indent: 0, tab: false, color: true, json: []byte(`{"foo":"bar"}`), res: []byte{0x1b, 0x5b, 0x6d, 0x7b, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x36, 0x6d, 0x22, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x36, 0x6d, 0x66, 0x6f, 0x6f, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x36, 0x6d, 0x22, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x6d, 0x3a, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x36, 0x3b, 0x31, 0x6d, 0x22, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x36, 0x3b, 0x31, 0x6d, 0x62, 0x61, 0x72, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x36, 0x3b, 0x31, 0x6d, 0x22, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x6d, 0x7d, 0x1b, 0x5b, 0x30, 0x6d}},
	}

	for _, tc := range testCases {
		res, err := New(tc.indent, tc.tab, tc.color).Format(tc.json)
		if err != nil {
			t.Errorf("Fmtr.Format(\"%v\") = %+v, want nil.", string(tc.json), err)
			continue
		} else if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("Fmtr.Format(\"%v\") = \"%v\", want \"%v\".", string(tc.json), res, tc.res)
		}
	}
}

func TestFormatRawNil(t *testing.T) {
	_, err := (*Fmtr)(nil).FormatRaw([]byte(`{"foo":"bar"}`))
	if err == nil {
		t.Error("Fmtr.FormatRaw() = nil, not want nil.")
	} else {
		fmt.Printf("Info: %+v\n", err)
	}
}

func TestFormatNil(t *testing.T) {
	_, err := (*Fmtr)(nil).Format([]byte(`{"foo":"bar"}`))
	if err == nil {
		t.Error("Fmtr.Format() = nil, not want nil.")
	} else {
		fmt.Printf("Info: %+v\n", err)
	}
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
