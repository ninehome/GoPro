package schema

import (
	"GeeOrm/dialect"
	"testing"
)

var TestDial, _ = dialect.GetDialect("sqlite3")

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
