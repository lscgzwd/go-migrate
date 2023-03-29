package mysql

import (
	"fmt"

	sk "github.com/laijunbin/go-solve-kit"
	"github.com/panda843/go-migrate/pkg/interfaces"
)

type Blueprint struct {
	metadata []*meta
}

func NewBlueprint() interfaces.Blueprint {
	return &Blueprint{}
}

func (bp *Blueprint) Id(name string, length int) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:          name,
		Type:          "BIGINT",
		Length:        length,
		AutoIncrement: true,
		Primary:       true,
		unsigned:      true,
		Comment:       "索引ID",
	})
	return bp
}

func (bp *Blueprint) String(name string, length int) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "VARCHAR",
		Length: length,
	})
	return bp
}

func (bp *Blueprint) Text(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "TEXT",
	})
	return bp
}

func (bp *Blueprint) CustomSql(sql string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Custom: sql,
	})
	return bp
}

func (bp *Blueprint) MediumText(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "MEDIUMTEXT",
	})
	return bp
}

func (bp *Blueprint) LongText(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "LONGTEXT",
	})
	return bp
}

func (bp *Blueprint) BigInt(name string, length int) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "BIGINT",
		Length: length,
	})
	return bp
}

func (bp *Blueprint) Collate(collate string) interfaces.Blueprint {
	if len(bp.metadata) != 0 {
		bp.metadata[len(bp.metadata)-1].Collate = collate
	}
	return bp
}

func (bp *Blueprint) Decimal(name string, length, decimals int) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:     name,
		Type:     "DECIMAL",
		Length:   length,
		Decimals: decimals,
	})
	return bp
}

func (bp *Blueprint) Integer(name string, length int) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "INT",
		Length: length,
	})
	return bp
}

func (bp *Blueprint) Date(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "DATE",
	})
	return bp
}

func (bp *Blueprint) Boolean(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "TINYINT",
	})
	return bp
}

func (bp *Blueprint) Comment(value string) interfaces.Blueprint {
	if len(bp.metadata) != 0 {
		bp.metadata[len(bp.metadata)-1].Comment = value
	}
	return bp
}

func (bp *Blueprint) TableComment(value string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		TableComment: value,
	})
	return bp
}

func (bp *Blueprint) DateTime(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "DATETIME",
	})
	return bp
}

func (bp *Blueprint) Timestamps() {
	bp.metadata = append(bp.metadata, &meta{
		Name:    "created_at",
		Type:    "DATETIME",
		Default: "CURRENT_TIMESTAMP",
		Comment: "创建时间",
	})

	bp.metadata = append(bp.metadata, &meta{
		Name:     "updated_at",
		Type:     "DATETIME",
		Nullable: true,
		Default:  "NULL",
		Comment:  "更新时间",
	})
}

func (bp *Blueprint) Nullable() interfaces.Blueprint {
	if len(bp.metadata) != 0 {
		bp.metadata[len(bp.metadata)-1].Nullable = true
	}
	return bp
}

func (bp *Blueprint) Unsigned() interfaces.Blueprint {
	if len(bp.metadata) != 0 {
		bp.metadata[len(bp.metadata)-1].unsigned = true
	}
	return bp
}

func (bp *Blueprint) Unique(column ...string) interfaces.Blueprint {
	if len(column) == 0 {
		if len(bp.metadata) != 0 {
			bp.metadata[len(bp.metadata)-1].Unique = true
		}
	} else {
		for _, c := range column {
			bp.metadata = append(bp.metadata, &meta{
				Name:   c,
				Unique: true,
			})
		}
	}
	return bp
}

func (bp *Blueprint) Index(column ...string) interfaces.Blueprint {
	if len(column) == 0 {
		bp.metadata[len(bp.metadata)-1].Index = true
	} else {
		for _, c := range column {
			bp.metadata = append(bp.metadata, &meta{
				Name:  c,
				Index: true,
			})
		}
	}
	return bp
}

func (bp *Blueprint) Default(value interface{}) interfaces.Blueprint {
	if len(bp.metadata) != 0 {
		bp.metadata[len(bp.metadata)-1].Default = fmt.Sprintf("'%v'", value)
	}
	return bp
}

func (bp *Blueprint) Foreign(name string) interfaces.ForeignBlueprint {
	fb := newForeignBlueprint().(*foreignBlueprint)
	bp.metadata = append(bp.metadata, &meta{
		Name:    name,
		Foreign: fb.meta,
	})
	return fb
}

func (bp *Blueprint) Primary(name ...string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:    sk.FromStringArray(name).Join("`, `").ValueOf(),
		Primary: true,
	})
	return bp
}

func (bp *Blueprint) DropColumn(column string) {
	bp.metadata = append(bp.metadata, &meta{
		Name: column,
		Type: "DROP",
	})
}

func (bp *Blueprint) DropUnique(name string) {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "DROP",
		Unique: true,
	})
}
func (bp *Blueprint) DropIndex(name string) {
	bp.metadata = append(bp.metadata, &meta{
		Name:  name,
		Type:  "DROP",
		Index: true,
	})
}
func (bp *Blueprint) DropForeign(name string) {
	bp.metadata = append(bp.metadata, &meta{
		Name:    name,
		Type:    "DROP",
		Foreign: newForeignBlueprint().(*foreignBlueprint).meta,
	})
}
func (bp *Blueprint) DropPrimary() {
	bp.metadata = append(bp.metadata, &meta{
		Type:    "DROP",
		Primary: true,
	})
}

func (bp *Blueprint) GetSqls(table string, operation operation) []string {
	return operation.generateSql(table, bp.metadata)
}
