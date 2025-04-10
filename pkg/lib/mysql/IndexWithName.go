package mysql

import (
	"fmt"
	"strings"
)

type indexWithName struct {
	b       bool
	columns []string
}

func (iwn *indexWithName) generateSql(isUnique bool, name string, column string) string {
	index := ""
	if isUnique {
		index = "UNIQUE "
	}
	if len(iwn.columns) > 0 {
		for i := range iwn.columns {
			if !strings.ContainsAny(iwn.columns[i], "`(") {
				iwn.columns[i] = fmt.Sprintf("`%s`", iwn.columns[i])
			}
		}
		return fmt.Sprintf("%sINDEX %s (%s)", index, name, strings.Join(iwn.columns, ","))
	}
	return fmt.Sprintf("%sINDEX %s (`%s`)", index, name, column)
}
