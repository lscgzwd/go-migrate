package mysql

import (
	"fmt"

	sk "github.com/laijunbin/go-solve-kit"
)

type meta struct {
	Name          string
	Type          string
	Length        int
	Decimals      int
	Nullable      bool
	Unique        indexWithName
	Index         indexWithName
	Modify        bool
	IndexName     string
	Primary       bool
	AutoIncrement bool
	unsigned      bool
	Collate       string
	Default       interface{}
	Comment       string
	Custom        string
	Foreign       *foreignMeta
	TableComment  string
}

type operation interface {
	generateSql(table string, metadata []*meta) []string
}

type createOperation struct{}
type alterOperation struct{}

func (o *createOperation) generateSql(table string, metadata []*meta) []string {
	columns := sk.FromInterfaceArray(metadata)
	tableComment := ""
	proceedColumns := columns.Map(func(v sk.Type, i int) interface{} {
		m := v.ValueOf().(*meta)

		if m.Type == "DROP" {
			return nil
		}

		if m.Foreign != nil {
			return m.Foreign.generateSql(table, m.Name)
		}

		s := ""
		if m.Type != "" {
			s += fmt.Sprintf("`%s` %s", m.Name, m.Type)
		}
		if m.Custom != "" {
			return m.Custom
		}
		if m.Decimals != 0 {
			s += fmt.Sprintf("(%d,%d)", m.Length, m.Decimals)
		} else if m.Length != 0 {
			s += fmt.Sprintf("(%d)", m.Length)
		}

		if m.Collate != "" {
			s += " COLLATE " + m.Collate
		}

		if m.unsigned {
			s += " UNSIGNED"
		}

		if s != "" && !m.Nullable {
			s += " NOT NULL"
		}

		if m.AutoIncrement {
			s += " AUTO_INCREMENT"
		}

		if m.Comment != "" && m.TableComment == "" {
			s += fmt.Sprintf(` COMMENT "%s"`, m.Comment)
		}

		if m.Default != nil {
			s += fmt.Sprintf(" DEFAULT %v", m.Default)
		}

		if m.Primary {
			if s != "" {
				s += ", "
			}
			s += fmt.Sprintf("PRIMARY KEY (`%s`)", m.Name)
		}

		var indexNameSql string
		if m.Unique.b {
			if m.IndexName != "" {
				if len(m.Unique.columns) > 0 {
					return m.Unique.generateSql(true, m.IndexName, "")
				}
				indexNameSql = fmt.Sprintf(", %s", m.Unique.generateSql(true, m.IndexName, m.Name))
			} else {
				if s != "" {
					s += ", "
				}
				s += fmt.Sprintf("UNIQUE (`%s`)", m.Name)
			}
		}

		if m.Index.b {
			if m.IndexName != "" {
				if len(m.Index.columns) > 0 {
					return m.Index.generateSql(false, m.IndexName, "")
				}
				if indexNameSql == "" {
					indexNameSql = fmt.Sprintf(", %s", m.Index.generateSql(false, m.IndexName, m.Name))
				} else {
					if s != "" {
						s += ", "
					}
					s += fmt.Sprintf("INDEX (`%s`)", m.Name)
				}
			} else {
				if s != "" {
					s += ", "
				}
				s += fmt.Sprintf("INDEX (`%s`)", m.Name)
			}
		}

		if m.TableComment != "" {
			tableComment = m.TableComment
			return nil
		}
		return s + indexNameSql
	})

	columnsStr := proceedColumns.Filter(func(s sk.Type, i int) bool {
		return s.ValueOf() != nil
	}).ToStringArray().Join(", ").ValueOf()
	sql := fmt.Sprintf("CREATE TABLE `%s` (%s)", table, columnsStr)
	if tableComment != "" {
		sql += fmt.Sprintf(" comment='%s'", tableComment)
	}
	fmt.Println(sql + ";\n")
	return []string{sql + ";"}
}

func (o *alterOperation) generateSql(table string, metadata []*meta) []string {
	columns := sk.FromInterfaceArray(metadata)
	tableComment := ""
	sql := fmt.Sprintf("ALTER TABLE `%s` %s", table, columns.Map(func(v sk.Type, i int) interface{} {
		m := v.ValueOf().(*meta)

		if m.Type == "DROP" {
			if m.Primary {
				return "DROP PRIMARY KEY"
			}

			if m.Index.b || m.Unique.b {
				return fmt.Sprintf("DROP INDEX `%s`", m.Name)
			}

			if m.Foreign != nil {
				return fmt.Sprintf("DROP FOREIGN KEY `%[1]s`, DROP INDEX `%[1]s`", fmt.Sprintf("fk_%s_%s", table, m.Name))
			}

			return fmt.Sprintf("DROP `%s`", m.Name)
		}

		if m.Foreign != nil {
			return fmt.Sprintf("ADD %s", m.Foreign.generateSql(table, m.Name))
		}

		s := "ADD "
		if m.Modify {
			s = "MODIFY "
		}

		if m.Type != "" {
			s += fmt.Sprintf("`%s` %s", m.Name, m.Type)
		}
		if m.Custom != "" {
			return m.Custom
		}
		if m.Decimals != 0 {
			s += fmt.Sprintf("(%d,%d)", m.Length, m.Decimals)
		} else if m.Length != 0 {
			s += fmt.Sprintf("(%d)", m.Length)
		}

		if m.Collate != "" {
			s += " COLLATE " + m.Collate
		}

		if m.unsigned {
			s += " UNSIGNED"
		}

		if s != "" && !m.Nullable {
			s += " NOT NULL"
		}

		if m.AutoIncrement {
			s += " AUTO_INCREMENT"
		}

		if m.Default != nil {
			s += fmt.Sprintf(" DEFAULT %v", m.Default)
		}

		if m.Comment != "" && m.TableComment == "" {
			s += fmt.Sprintf(` COMMENT "%s"`, m.Comment)
		}

		if m.Primary {
			if s != "" {
				s += ", "
			}
			s += fmt.Sprintf("ADD PRIMARY KEY (`%s`)", m.Name)
		}

		var indexNameSql string
		if m.Unique.b {
			if m.IndexName != "" {
				if len(m.Unique.columns) > 0 {
					return fmt.Sprintf("ADD %s", m.Unique.generateSql(true, m.IndexName, ""))
				}
				indexNameSql = fmt.Sprintf(", ADD %s", m.Unique.generateSql(true, m.IndexName, m.Name))
			} else {
				if s != "" {
					s += ", "
				}
				s += fmt.Sprintf("ADD UNIQUE (`%s`)", m.Name)
			}
		}

		if m.Index.b {
			if m.IndexName != "" {
				if len(m.Index.columns) > 0 {
					return fmt.Sprintf("ADD %s", m.Index.generateSql(false, m.IndexName, ""))
				}
				if indexNameSql == "" {
					indexNameSql = fmt.Sprintf(", ADD %s", m.Index.generateSql(false, m.IndexName, m.Name))
				} else {
					if s != "" {
						s += ", "
					}
					s += fmt.Sprintf("ADD INDEX (`%s`)", m.Name)
				}

			} else {
				if s != "" {
					s += ", "
				}
				s += fmt.Sprintf("ADD INDEX (`%s`)", m.Name)
			}
		}
		if m.TableComment != "" {
			tableComment = m.TableComment
			return nil
		}
		return s + indexNameSql
	}).ToStringArray().Join(", ").ValueOf())
	if tableComment != "" {
		sql += fmt.Sprintf(" comment='%s'", tableComment)
	}
	fmt.Println(sql + ";\n")
	return []string{
		sql + ";",
	}
}

var metaOperations = struct {
	CREATE operation
	ALTER  operation
}{
	CREATE: &createOperation{},
	ALTER:  &alterOperation{},
}
