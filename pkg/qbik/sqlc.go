package qbik

import (
	"regexp"
	"strings"
)

func SQLC_RAW(val string) Statement {
	reg := regexp.MustCompile("^--(.*)\n")
	clear := strings.TrimSpace(reg.ReplaceAllString(val, ""))
	return Statement{
		SQL:  clear,
		Args: []interface{}{},
	}
}
