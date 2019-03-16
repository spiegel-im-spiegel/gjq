# [gjq] - Another Implementation of [jq] by golang

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/gjq.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/gjq)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gjq/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/gjq.svg)](https://github.com/spiegel-im-spiegel/gjq/releases/latest)

- Use [savaki/jq] package

## Declare [gjq] module

See [go.mod](https://github.com/spiegel-im-spiegel/gjq/blob/master/go.mod) file. 

## Command-Line Interface

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/gjq/releases/latest).

### Usage

```
$ gjq -h
Usage:
  gjq [flags] <filter string>

Flags:
      --debug         for debug
  -f, --file string   JSON data (file path)
  -h, --help          help for gjq
  -i, --indent int    indent size for formatted JSON string (default 2)
  -I, --interactive   interactive mode
  -t, --tab           use tabs for indentation
  -u, --url string    JSON data (URL)
  -v, --version       output version of gjq
```

### Syntax of Filter

See [savaki/jq].

| syntax   | meaning                                         |
| -------- | ----------------------------------------------- |
| .        | unchanged input                                 |
| .foo     | value at key                                    |
| .foo.bar | value at nested key                             |
| .[0]     | value at specified element of array             |
| .[0:1]   | array of specified elements of array, inclusive |
| .foo.[0] | nested value                                    |

### Filtering JSON data from Stdin

```
$ cat testdata/test.json
{
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
}

$ type testdata/test.json | gjq .object.array
[
  1,
  2,
  3
]
```

### Filtering JSON data from file

```
$ gjq -f testdata/test.json .object.array
[
  1,
  2,
  3
]
```

### Filtering JSON data from WWW

```
$ gjq -u https://text.baldanders.info/index.json .entry.[0]
{
  "title": "猫を殺すに猫を以ってせよ",
  "section": "remark",
  "description": "分かるかな。分っかんねーだろうな（反語）",
  "author": "Spiegel",
  "license": "http://creativecommons.org/licenses/by-sa/4.0/",
  "url": "https://text.baldanders.info/remark/2019/03/no-cat-no-life/",
  "published": "2019-03-11T13:51:41+00:00",
  "update": "2019-03-11T13:53:40+00:00"
}
```

## Interactive Mode

```
$ gjq -u https://text.baldanders.info/index.json -I
Press Ctrl+C to stop
Filter> .entry.[0]
{
  "title": "猫を殺すに猫を以ってせよ",
  "section": "remark",
  "description": "分かるかな。分っかんねーだろうな（反語）",
  "author": "Spiegel",
  "license": "http://creativecommons.org/licenses/by-sa/4.0/",
  "url": "https://text.baldanders.info/remark/2019/03/no-cat-no-life/",
  "published": "2019-03-11T13:51:41+00:00",
  "update": "2019-03-11T13:53:40+00:00"
}
Filter> 
```

[gjq]: https://github.com/spiegel-im-spiegel/gjq
[jq]: https://github.com/stedolan/jq "stedolan/jq: Command-line JSON processor"
[savaki/jq]: https://github.com/savaki/jq/ "savaki/jq: A high performance Golang implementation of the incredibly useful jq command line tool."
