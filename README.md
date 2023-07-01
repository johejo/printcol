# printcol

Print columns like `awk '{print $N}'`

## Motivation

awk (also gawk, goawk) is a powerfull tool.

For most of my use cases, I rarely use awk as a scripting language, I use awk to output specific columns for csv or tsv.

## Install

```
go install github.com/johejo/printcol@latest
```

## Usage

```
$ printcol -h
Usage of printcol:
  -col string
        0-based column indexes (negative index is supported) for example "0", "0..5", "0,3,-1", "0,2..5" (default "0")
  -csv
        alias for -sep=','
  -sep string
        separator (default " ")
  -skip-header
        skip first line as header
  -tsv
        alias for -sep='\t'
```

## Example

```
$ echo "foo,bar\naaa,bbb" | printcol -csv -col -1
bar
bbb
```

## License

MIT
