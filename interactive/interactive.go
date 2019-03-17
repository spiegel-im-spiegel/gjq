package interactive

import (
	"context"
	"os"

	"github.com/atotto/clipboard"
	"github.com/mattn/go-colorable"
	"github.com/spiegel-im-spiegel/gjq/errs"
	"github.com/spiegel-im-spiegel/gjq/fmtr"
	"github.com/spiegel-im-spiegel/gjq/query"
	"github.com/spiegel-im-spiegel/gocli/prompt"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gocli/signal"
)

func Run(op *query.Op, formtter *fmtr.Fmtr, clip bool) error {
	errCh := make(chan error)
	defer close(errCh)

	go func() {
		child, cancelChild := context.WithCancel(
			signal.Context(context.Background(), os.Interrupt), // cancel event by SIGNAL
		)
		defer cancelChild()

		m := &mode{op: op, formtter: formtter, clip: clip}
		errCh <- m.run(child)
	}()

	if err := <-errCh; !errs.Is(err, prompt.ErrTerminate) {
		return err
	}
	return nil
}

type mode struct {
	op       *query.Op
	formtter *fmtr.Fmtr
	clip     bool
}

func (m *mode) run(ctx context.Context) error {
	lastStr := ""
	strCh := make(chan string)
	defer close(strCh)
	errCh := make(chan error)
	defer close(errCh)

	go func() {
		if err := m.proc(strCh); err != nil {
			errCh <- err
		}
	}()
	for {
		select {
		case str := <-strCh:
			lastStr = str
		case err := <-errCh:
			return err
		case <-ctx.Done(): // cancel event from context
			if m.clip && len(lastStr) > 0 {
				err := clipboard.WriteAll(lastStr)
				if err != nil {
					return errs.Wrap(err, "error when output result to clipboard")
				}
			}
			if err := ctx.Err(); err != nil {
				if !errs.Is(err, context.Canceled) {
					return errs.Wrap(err, "error in cancel event")
				}
			}
			return prompt.ErrTerminate
		}
	}
}

func (m *mode) proc(ch chan string) error {
	return prompt.New(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(colorable.NewColorableStdout()),
			rwi.WithErrorWriter(colorable.NewColorableStderr()),
		),
		func(s string) (string, error) {
			res, err := m.op.Query(s)
			if err != nil {
				return err.Error(), nil
			}
			j, _ := m.formtter.FormatRaw(res)
			ch <- string(j)
			j, err = m.formtter.Format(res)
			return string(j), err
		},
		prompt.WithPromptString("Filter> "),
		prompt.WithHeaderMessage("Press Ctrl+C to stop"),
	).Run()
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
