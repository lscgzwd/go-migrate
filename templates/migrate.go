package templates

var MigrateTemplate = `package main

import (
	"github.com/panda843/go-migrate/cmd"
	"github.com/panda843/go-migrate/config"
	_ "%[2]s/databases"
)

func init() {
	config.Config = config.DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "",
		Dbname:   "test",
	}

	config.Migrator = %[1]s.InitMigrator()
	config.Driver = "%[1]s"
}

func main() {
	cmd.Execute()
}
`
