package templates

var MigrationCreateTemplate = `package migrations

import (
	"github.com/panda843/go-migrate/config"
	"github.com/panda843/go-migrate/pkg/interfaces"
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
		table.Id("id", 10)
		table.Timestamps()
	})
}

func (t *%[2]sTable) Down() error {
	return %[1]s.Schema.DropIfExists("%[3]s")
}
`
