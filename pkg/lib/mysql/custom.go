package mysql

import (
	"fmt"
	"strings"
)

type indexWithName struct {
	name    string
	columns []string
}

func (iwn *indexWithName) generateSql() string {
	return fmt.Sprintf("INDEX %s (%s)", iwn.name, strings.Join(iwn.columns, ","))
}
