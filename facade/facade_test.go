package facade

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
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

var res1 = "[1,2,3]\n"

func TestLoadByStdin(t *testing.T) {
	inData := bytes.NewBufferString(imputStr)
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(
		rwi.WithReader(inData),
		rwi.WithWriter(outBuf),
		rwi.WithErrorWriter(outErrBuf),
	)
	args := []string{"-i", "0", "--debug", ".object.array"}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(stdin) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(stdin) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	if str != res1 {
		t.Errorf("Execute(stdin) = \"%v\", want \"%v\".", str, res1)
	}
}

func TestLoadByStdinError(t *testing.T) {
	inData := bytes.NewBufferString(imputStr)
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(
		rwi.WithReader(inData),
		rwi.WithWriter(outBuf),
		rwi.WithErrorWriter(outErrBuf),
	)
	args := []string{"-i", "0", "--debug", ".[0]"}

	exit := Execute(ui, args)
	if exit != exitcode.Abnormal {
		t.Errorf("Execute(stdin) = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
		fmt.Printf("Info: %+v\n", outBuf.String())
	} else {
		fmt.Printf("Info: %+v\n", outErrBuf.String())
	}
}

func TestLoadByFile(t *testing.T) {
	inData := bytes.NewBufferString(imputStr)
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(
		rwi.WithReader(inData),
		rwi.WithWriter(outBuf),
		rwi.WithErrorWriter(outErrBuf),
	)
	args := []string{"-i", "0", "--debug", "-f", "../testdata/test.json", ".object.array"}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(file) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(file) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	if str != res1 {
		t.Errorf("Execute(file) = \"%v\", want \"%v\".", str, res1)
	}
}

func TestLoadByFileError(t *testing.T) {
	inData := bytes.NewBufferString(imputStr)
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(
		rwi.WithReader(inData),
		rwi.WithWriter(outBuf),
		rwi.WithErrorWriter(outErrBuf),
	)
	args := []string{"-i", "0", "--debug", "-f", "noexist.txt", ".object.array"}

	exit := Execute(ui, args)
	if exit != exitcode.Abnormal {
		t.Errorf("Execute(file) = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
		fmt.Printf("Info: %+v\n", outBuf.String())
	} else {
		fmt.Printf("Info: %+v\n", outErrBuf.String())
	}
}

func TestLoadByURL(t *testing.T) {
	inData := bytes.NewBufferString(imputStr)
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(
		rwi.WithReader(inData),
		rwi.WithWriter(outBuf),
		rwi.WithErrorWriter(outErrBuf),
	)
	args := []string{"-i", "0", "--debug", "-u", "https://text.baldanders.info/index.json", ".title"}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(url) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(url) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "\"text.Baldanders.info\"\n"
	if str != res {
		t.Errorf("Execute(url) = \"%v\", want \"%v\".", str, res)
	}
}

func TestLoadByURLError(t *testing.T) {
	inData := bytes.NewBufferString(imputStr)
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(
		rwi.WithReader(inData),
		rwi.WithWriter(outBuf),
		rwi.WithErrorWriter(outErrBuf),
	)
	args := []string{"-i", "0", "--debug", "-u", "http://foo.bar/json.json", ".object.array"}

	exit := Execute(ui, args)
	if exit != exitcode.Abnormal {
		t.Errorf("Execute(url) = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
		fmt.Printf("Info: %+v\n", outBuf.String())
	} else {
		fmt.Printf("Info: %+v\n", outErrBuf.String())
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
