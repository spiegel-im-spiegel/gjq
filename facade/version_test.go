package facade

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestVersionMin(t *testing.T) {
	result := "gjq dev-version\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{"-v"}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outBuf.String()
	if str != "" {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, "")
	}
	str = outErrBuf.String()
	if str != result {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, result)
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
