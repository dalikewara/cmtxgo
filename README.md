# cmtxgo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/cmtxgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/cmtxgo)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/cmtxgo)
![GitHub license](https://img.shields.io/github/license/dalikewara/cmtxgo)

**cmtxgo** generates cemtext format in Golang. Cemtext file format is a format used by banks to allow for batch
transactions. Banks adopted this format by default, but they may have different format style. This package allows you to
create a cemtext format by your own style—that means you can use it to meet your bank cemtext requirements.

> If you're working with JavaScript, you can use this module instead -> [https://github.com/dalikewara/vcmcemtex](https://github.com/dalikewara/vcmcemtex)

## Getting started

### Installation

You can use the `go install` method:

```bash
go install github.com/dalikewara/cmtxgo@latest
```

or, you can also use the `go get` method (DEPRECATED since `go1.17`):

```bash
go get github.com/dalikewara/cmtxgo
```

### Usage

A cemtext format commonly has 3 sections: `header`, `detail` and `footer`.

```text
(header) 0                 01BQL       MY NAME                   1111111004231633  230410
(detail) 1123-456157108231 530000001234S R SMITH                       TEST BATCH        062-000 12223123MY ACCOUNT      00001200
(detail) 1123-783 12312312 530000002200J K MATTHEWS                    TEST BATCH        062-000 12223123MY ACCOUNT      00000030
(detail) 1456-789   125123 530003123513P R JONES                       TEST BATCH        062-000 12223123MY ACCOUNT      00000000
(detail) 1121-232    11422 530000002300S MASLIN                        TEST BATCH        062-000 12223123MY ACCOUNT      00000000
(footer) 7999-999            000312924700031292470000000000                        000004
```

Every section detailed information about the transaction, and they can be separated by new line `\n` or just combined to
a single string. Both is valid, because it depends on the bank specification&mdash;some Indonesian bank uses single
string format, for example:

```text
0                 01BQL       MY NAME                   1111111004231633  2304101123-456157108231 530000001234S R SMITH                       TEST BATCH        062-000 12223123MY ACCOUNT      000012001123-783 12312312 530000002200J K MATTHEWS                    TEST BATCH        062-000 12223123MY ACCOUNT      000000301456-789   125123 530003123513P R JONES                       TEST BATCH        062-000 12223123MY ACCOUNT      000000001121-232    11422 530000002300S MASLIN                        TEST BATCH        062-000 12223123MY ACCOUNT      000000007999-999            000312924700031292470000000000                        000004
```

> Ref: [https://www.cemtexaba.com/aba-format](https://www.cemtexaba.com/aba-format)

This is the very basic usage of this package:

```go
cmtx := cmtxgo.NewCmtx()

// Set header
cmtx.SetHeader(&cmtxgo.Field{
    "myHeaderFieldName": &cmtxgo.Attribute{
        Order:  1,
        Type:   "lpz",
        Length: 11,
        Value:  "999",
    },
    "myHeaderFieldName2": &cmtxgo.Attribute{
        Order:  2,
        Type:   "rps",
        Length: 35,
        Value:  "My COMPANY",
    },
})

// Set detail
detailData := []map[string]interface{}{
    {
        "myDetailFieldName":  "123-456-789",
        "myDetailFieldName2": "John Doe",
        "myDetailFieldName3": "1000000",
    },
    {
        "myDetailFieldName":  "098-765-432",
        "myDetailFieldName2": "Adam Smith",
        "myDetailFieldName3": "1005500",
    },
    {
        "myDetailFieldName":  "123-098-456",
        "myDetailFieldName2": "Dali Kewara",
        "myDetailFieldName3": "204000",
    },
}
cmtx.SetDetail(&cmtxgo.Field{
    "myDetailFieldName": &cmtxgo.Attribute{
        Order:  1,
        Type:   "rps",
        Length: 15,
    },
    "myDetailFieldName2": &cmtxgo.Attribute{
        Order:  2,
        Type:   "rps",
        Length: 20,
    },
    "myDetailFieldName3": &cmtxgo.Attribute{
        Order:  3,
        Type:   "lpz",
        Length: 11,
    },
}, &detailData)

// Set footer
cmtx.SetFooter(&cmtxgo.Field{
    "myFooterFieldName": &cmtxgo.Attribute{
        Order:  1,
        Type:   "lpz",
        Length: 11,
        Value:  fmt.Sprintf("%v", len(detailData)),
    },
    "myFooterFieldName2": &cmtxgo.Attribute{
        Order:  2,
        Type:   "lpz",
        Length: 35,
        Value:  "2209500",
    },
})

// Generate the cemtext format as string
cmtxStr := cmtx.Generate()
```

The output is:

```text
00000000999My COMPANY                         123-456-789    John Doe            00001000000098-765-432    Adam Smith          00001005500123-098-456    Dali Kewara         000002040000000000000300000000000000000000000000002209500
```

If you want to save the cemtext format into a file, just use `GenerateToFile` method instead:

```go
err := cmtx.GenerateToFile("/path/to/filename.ctx")
if err != nil {
    panic(err)
}
```

### Attribute

- `Order`: sorts the generation of the fields. Example:

```go
cmtx.SetHeader(&cmtxgo.Field{
    "first": &cmtxgo.Attribute{
        Order:  2,
        Type:   "lps",
        Length: 6,
        Value:  "Second",
    },
    "second": &cmtxgo.Attribute{
        Order:  1,
        Type:   "lps",
        Length: 5,
        Value:  "First",
    },
})
// Output: "FirstSecond"
```

- `Type`: formats the field value based on padding type. Available padding type: `lps`, `lpz`, `rps`, `rpz`. 
Default is `rps`.
  - `lps`: Left Padding Space (ex: "  john doe")
  - `lpz`: Left Padding Zero (ex: "0000000123")
  - `rps`: Right Padding Space (ex: "john doe   ")
  - `rpz`: Right Padding Zero (ex: "1230000000")
- `Length`: sets max length to the field value.
- `Value`: sets field value. This only effect on `header` & `footer` section.
For `detail` section, the field value got from the matched field name in detail data.
If Value is empty, DefaultValue will be used.
- `DefaultValue`: sets the default value to the field if the Value is empty. Default is "".
- `RemoveAllChars`: removes specified characters from the field value. Example:

```go
cmtx.SetFooter(&cmtxgo.Field{
    "total_amount": &cmtxgo.Attribute{
        Order:          1,
        Type:           "lpz",
        Length:         10,
        Value:          "1.000.000.000",
        RemoveAllChars: ".",
    },
})
// Output: "1000000000"
```

- `ReplaceAllChars`: ReplaceAllChars replaces specified characters with the new one from the field value. Example:

```go
cmtx.SetFooter(&cmtxgo.Field{
    "total_amount": &cmtxgo.Attribute{
        Order:           1,
        Type:            "lpz",
        Length:          13,
        Value:           "1.000.000.000",
        ReplaceAllChars: [2]string{".", ","},
    },
})
// Output: "1,000,000,000"
```

### Option

- `MaxLengthPerSection`: sets max length on every section value.
- `AddCharPerSection`: adds specified character at the end of the value on every section. Keep in mind that this will add 
new character, so it will break the value length setting.

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/cmtxgo/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2021 [Dali Kewara](https://www.dalikewara.com)

### License

[GNU General Public License v3](https://github.com/dalikewara/cmtxgo/blob/master/LICENSE)