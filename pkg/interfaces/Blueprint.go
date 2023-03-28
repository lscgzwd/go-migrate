package interfaces

type Blueprint interface {
	Id(name string, length int) Blueprint
	String(name string, length int) Blueprint
	Text(name string) Blueprint
	Integer(name string, length int) Blueprint
	Date(name string) Blueprint
	Comment(value string) Blueprint
	TableComment(value string) Blueprint
	Boolean(name string) Blueprint
	DateTime(name string) Blueprint
	Nullable() Blueprint
	Unique(column ...string) Blueprint
	Index(column ...string) Blueprint
	Default(value interface{}) Blueprint
	Foreign(name string) ForeignBlueprint
	Primary(name ...string) Blueprint
	DropColumn(column string)
	DropUnique(name string)
	DropIndex(name string)
	DropForeign(name string)
	DropPrimary()
	Timestamps()
}
