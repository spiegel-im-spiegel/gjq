package facade

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spiegel-im-spiegel/gjq/errs"
)

func getJSONURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errs.Wrapf(err, "cannot read JSON data from %v", url)
	}
	defer resp.Body.Close()

	return getJSONData(resp.Body)
}

func getJSONFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errs.Wrap(err, "cannot read JSON data file")
	}
	defer file.Close()

	return getJSONData(file)
}

func getJSONData(r io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errs.Wrap(err, "cannot read JSON data")
	}
	return b, nil
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
