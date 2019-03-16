# [gjq] - Another Implementation of [jq] by golang

- Use [savaki/jq] package

## Declare [gjq] module

See [go.mod](https://github.com/spiegel-im-spiegel/gjq/blob/master/go.mod) file. 

## Command-Line Interface

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/gpgpdump/releases/latest).

### Usage

```
$ gjq -h
Usage:
  gjq [flags] <query string>

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
