package templates

var MigrationCreateTemplate = `package migrations

import (
	"github.com/lscgzwd/go-migrate/config"
	"github.com/lscgzwd/go-migrate/pkg/interfaces"
	"github.com/lscgzwd/go-migrate/pkg/lib/mysql"
)

func init() {
	config.Migrations = append(config.Migrations, Create%[2]sTable())
}

type %[2]sTable struct{}

func Create%[2]sTable() interfaces.Migration {
	return &%[2]sTable{}
}

func (t *%[2]sTable) Up() error {
	return %[1]s.Schema.Create("%[3]s", func(table interfaces.Blueprint) {
		table.Id("id", 22)
		table.Timestamps()
	})
}

func (t *%[2]sTable) Down() error {
	return %[1]s.Schema.DropIfExists("%[3]s")
}
`
