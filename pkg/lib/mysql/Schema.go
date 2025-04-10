package mysql

import (
	"fmt"

	"github.com/lscgzwd/go-migrate/pkg/interfaces"
	mysql_interfaces "github.com/lscgzwd/go-migrate/pkg/lib/mysql/interfaces"
)

type schema struct{}
type Schema_test struct{}

func newSchema() interfaces.SchemaWithSeeder {
	return &schema{}
}

func (s *schema) Create(table string, schemaFunc func(interfaces.Blueprint)) interfaces.Seeder {
	driver, err := NewDriver()
	if err != nil {
		return NewSeeder("", err)
	}

	return NewSeeder(table, createWithDriver(driver, table, schemaFunc))
}

func (s *Schema_test) Create(driver mysql_interfaces.Driver, table string, schemaFunc func(interfaces.Blueprint)) interfaces.Seeder {
	return NewTestSeeder(driver, table, createWithDriver(driver, table, schemaFunc))
}

func createWithDriver(driver mysql_interfaces.Driver, table string, schemaFunc func(interfaces.Blueprint)) error {
	defer driver.Close()

	blueprint := NewBlueprint().(*Blueprint)
	schemaFunc(blueprint)
	sqls := blueprint.GetSqls(table, metaOperations.CREATE)
	for _, sql := range sqls {
		if _, err := driver.Execute(sql); err != nil {
			return err
		}
	}

	return nil
}

func (s *schema) Table(table string, schemaFunc func(interfaces.Blueprint)) error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}

	return tableWithDriver(driver, table, schemaFunc)
}

func (s *Schema_test) Table(driver mysql_interfaces.Driver, table string, schemaFunc func(interfaces.Blueprint)) error {
	return tableWithDriver(driver, table, schemaFunc)
}

func tableWithDriver(driver mysql_interfaces.Driver, table string, schemaFunc func(interfaces.Blueprint)) error {
	defer driver.Close()

	blueprint := NewBlueprint().(*Blueprint)
	schemaFunc(blueprint)

	sqls := blueprint.GetSqls(table, metaOperations.ALTER)
	for _, sql := range sqls {
		if _, err := driver.Execute(sql); err != nil {
			return err
		}
	}

	return nil
}

func (s *schema) DropIfExists(table string) error {
	driver, err := NewDriver()
	if err != nil {
		return err
	}

	return dropIfExistsWithDriver(driver, table)
}

func (s *Schema_test) DropIfExists(driver mysql_interfaces.Driver, table string) error {
	return dropIfExistsWithDriver(driver, table)
}

func dropIfExistsWithDriver(driver mysql_interfaces.Driver, table string) error {
	defer driver.Close()
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	_, err := driver.Execute(sql)
	return err
}

var Schema = newSchema()
