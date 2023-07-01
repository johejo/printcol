package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var (
	sep        string
	col        string
	csv        bool
	tsv        bool
	skipHeader bool
)

func init() {
	flag.StringVar(&sep, "sep", " ", "separator")
	flag.StringVar(&col, "col", "0", `0-based column indexes (negative index is supported), for example "0", "0..5", "0,3,-1", "0,2..5"`)
	flag.BoolVar(&csv, "csv", false, "alias for -sep=','")
	flag.BoolVar(&tsv, "tsv", false, "alias for -sep='\\t'")
	flag.BoolVar(&skipHeader, "skip-header", false, "skip first line as header")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if csv && tsv {
		log.Fatal("cannot specify -csv and -tsv")
	}

	if csv {
		sep = ","
	}
	if tsv {
		sep = "\t"
	}

	cols, err := parseCol(col)
	if err != nil {
		log.Fatal(err)
	}

	signal.Ignore(syscall.SIGPIPE)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s := bufio.NewScanner(os.Stdin)
	first := true
	for s.Scan() {
		if skipHeader && first {
			first = false
			continue
		}
		line := s.Text()
		ss := strings.Split(line, sep)
		for i, s := range ss {
			found := false
			for _, c := range cols {
				if c < 0 {
					c += len(ss)
				}
				if len(ss)-1 < c {
					log.Fatalf("invalid col %d, line has %d columns", c, len(ss)-1)
				}
				if i == c {
					w.WriteString(strings.TrimSpace(s))
					found = true
					break
				}
			}
			if !found {
				continue
			}
			w.WriteString(sep)
		}
		w.WriteString("\n")
	}
}

func parseCol(s string) ([]int, error) {
	list := strings.Split(s, ",")
	cols := make([]int, 0, len(list))
	for _, v := range list {
		if v == "" {
			continue
		}
		_cols, err := _parseCol(v)
		if err != nil {
			return nil, err
		}
		cols = append(cols, _cols...)
	}

	return cols, nil
}

func _parseCol(s string) ([]int, error) {
	i, err := strconv.Atoi(s)
	if err == nil {
		return []int{i}, nil
	}

	if _begin, _end, found := strings.Cut(s, ".."); found {
		if _begin == "" || _end == "" {
			return nil, fmt.Errorf("invalid format %s, begin or end are empty", s)
		}
		begin, err := strconv.Atoi(_begin)
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(_end)
		if err != nil {
			return nil, err
		}
		if end < begin {
			begin, end = end, begin
		}
		cols := make([]int, 0, end-begin+1)
		for i := begin; i <= end; i++ {
			cols = append(cols, i)
		}
		return cols, nil
	}
	return []int{}, nil
}
