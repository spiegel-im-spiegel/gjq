package facade

import (
	"github.com/atotto/clipboard"
	"github.com/spiegel-im-spiegel/gjq/errs"
	"github.com/spiegel-im-spiegel/gjq/query"
	"github.com/spiegel-im-spiegel/gocli/prompt"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func interactiveMode(ui *rwi.RWI, op *query.Op) error {
	p := prompt.New(
		rwi.New(
			rwi.WithReader(ui.Reader()),
			rwi.WithWriter(ui.Writer()),
		),
		func(s string) (string, error) {
			res, err := op.Query(s)
			if err != nil {
				return err.Error(), nil
			}
			return string(res), errs.Wrap(clipboard.WriteAll(string(res)), "error when output result")
		},
		prompt.WithPromptString("Filter> "),
		prompt.WithHeaderMessage("Press Ctrl+C to stop"),
	)
	if !p.IsTerminal() {
		return errs.Wrap(prompt.ErrNotTerminal, "error in interactive mode")
	}
	return errs.Wrap(p.Run(), "error in interactive mode")
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
