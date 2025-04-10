package templates

var MigrationAlterTemplate = `package migrations

import (
	"github.com/lscgzwd/go-migrate/config"
	"github.com/lscgzwd/go-migrate/pkg/interfaces"
	"github.com/lscgzwd/go-migrate/pkg/lib/mysql"
)

func init() {
	config.Migrations = append(config.Migrations, Create%[2]s())
}

type %[2]s struct{}

func Create%[2]s() interfaces.Migration {
	return &%[2]s{}
}

func (t *%[2]s) Up() error {
	return %[1]s.Schema.Table("%[3]s", func(table interfaces.Blueprint) {

	})
}

func (t *%[2]s) Down() error {
	return %[1]s.Schema.Table("%[3]s", func(table interfaces.Blueprint) {
		
	})
}
`
