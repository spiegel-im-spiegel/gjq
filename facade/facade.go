package facade

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gjq/errs"
	"github.com/spiegel-im-spiegel/gjq/query"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	//Name is applicatin name
	Name = "gjq"
	//Version is version for applicatin
	Version = "dev-version"
)

var (
	usage = []string{ //output message of version
		Name + " " + Version,
	}
	versionFlag bool //version flag
	debugFlag   bool //debug flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: Name + " [flags] <filter string>",
		RunE: func(cmd *cobra.Command, args []string) error {
			//parse options
			if versionFlag {
				return ui.OutputErrln(strings.Join(usage, "\n"))
			}

			//interactive mode
			iFlag, err := cmd.Flags().GetBool("interactive")
			if err != nil {
				return errs.Wrap(err, "error in --interactive option")
			}

			//indent size
			i, err := cmd.Flags().GetInt("indent")
			if err != nil {
				return errs.Wrap(err, "error in --indent option")
			}
			t, err := cmd.Flags().GetBool("tab")
			if err != nil {
				return errs.Wrap(err, "error in --tab option")
			}

			//get JSON data
			var data []byte
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return errs.Wrap(err, "error in --file option")
			}
			url, err := cmd.Flags().GetString("url")
			if err != nil {
				return errs.Wrap(err, "error in --url option")
			}

			if len(file) > 0 {
				data, err = getJSONFile(file)
			} else if len(url) > 0 {
				data, err = getJSONURL(url)
			} else {
				data, err = getJSONData(ui.Reader())
			}
			if err != nil {
				if debugFlag {
					_ = ui.OutputErrln(fmt.Sprintf("%+v", err))
				}
				return err
			}
			op := query.New(
				data,
				query.WithIndent(i),
				query.WithTab(t),
			)

			if iFlag {
				if err := interactiveMode(ui, op); err != nil {
					return err
				}
				return nil
			}

			filter := "."
			if len(args) > 0 {
				filter = strings.Join(args, " ")
			}
			res, err := op.Query(filter)
			if err != nil {
				if debugFlag {
					_ = ui.OutputErrln(fmt.Sprintf("%+v", err))
				}
				return err
			}
			return ui.Outputln(string(res))
		},
	}
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "output version of "+Name)
	rootCmd.Flags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.Flags().BoolP("tab", "t", false, "use tabs for indentation")
	rootCmd.Flags().BoolP("interactive", "I", false, "interactive mode")
	rootCmd.Flags().IntP("indent", "i", 2, "indent size for formatted JSON string")
	rootCmd.Flags().StringP("file", "f", "", "JSON data (file path)")
	rootCmd.Flags().StringP("url", "u", "", "JSON data (URL)")

	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	return rootCmd
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
