package templates

var MigrationTemplate = `package migrations

import (
	"github.com/panda843/go-migrate/config"
	"github.com/panda843/go-migrate/pkg/interfaces"
)

func init() {
	config.Migrations = append(config.Migrations, Create%[2]s())
}

type %[2]s struct{}

func Create%[2]s() interfaces.Migration {
	return &%[2]s{}
}

func (t *%[2]s) Up() error {
	
}

func (t *%[2]s) Down() error {

}
`
