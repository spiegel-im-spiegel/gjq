package query

import (
	"github.com/savaki/jq"
	"github.com/spiegel-im-spiegel/gjq/errs"
)

//Op is class for JSON query
type Op struct {
	data []byte
}

//New returns new Op instance
func New(j []byte) *Op {
	return &Op{data: j}
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
	return res, errs.Wrapf(err, "cannot parse JSON data (filter is \"%v\")", filter)
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
