package query

import (
	"fmt"
	"testing"
)

var imputStr = `{
  "string": "a",
  "number": 1.23,
  "simple": ["a", "b", "c"],
  "mixed": [
    "a",
    1,
    {"hello":"world"}
  ],
  "object": {
    "first": "joe",
    "array": [1,2,3]
  }
}`

func TestQuery(t *testing.T) {
	testCases := []struct {
		indent int
		tab    bool
		filter string
		res    string
	}{
		{indent: -1, tab: false, filter: ".object.array", res: "[1,2,3]"},
		{indent: 0, tab: false, filter: ".object.array", res: "[1,2,3]"},
		{indent: 0, tab: true, filter: ".object.array", res: "[\n\t1,\n\t2,\n\t3\n]"},
		{indent: 1, tab: true, filter: ".object.array", res: "[\n\t1,\n\t2,\n\t3\n]"},
		{indent: 1, tab: false, filter: ".object.array", res: "[\n 1,\n 2,\n 3\n]"},
	}

	for _, tc := range testCases {
		op := New(
			[]byte(imputStr),
			WithIndent(tc.indent),
			WithTab(tc.tab),
		)
		res, err := op.Query(tc.filter)
		if err != nil {
			t.Errorf("Op.Query(\"%v\")  = %+v, want nil.", tc.filter, err)
			continue
		} else if string(res) != tc.res {
			t.Errorf("Op.Query(\"%v\")  = \"%v\", want \"%v\".", tc.filter, string(res), tc.res)
		}
	}
}

func TestParseError(t *testing.T) {
	op := New(
		[]byte(imputStr),
		WithIndent(0),
		WithTab(false),
	)
	res, err := op.Query(".[0]")
	if err == nil {
		t.Error("Op.Query(\"string\")  = nil, want not nil.")
		fmt.Printf("Info: %v\n", string(res))
	} else {
		fmt.Printf("Info: %+v\n", err)
	}
}

func TestQueryNil(t *testing.T) {
	op := (*Op)(nil)
	_, err := op.Query(".")
	if err == nil {
		t.Error("Op.Query(\".\")  = nil, want not nil.")
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
