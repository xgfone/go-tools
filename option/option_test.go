package option

import (
	"database/sql"
	"testing"
)

func TestOption(t *testing.T) {
	if NONE.IsSome() || NONE.SomeOr(123).(int) != 123 {
		t.Fail()
	}

	if Some(123).IsNone() || Some(123).SomeOr(456).(int) != 123 {
		t.Fail()
	}
}

func TestNamedOption(t *testing.T) {
	if NamedNone("").IsSome() || NamedNone("").SomeOr(123).(int) != 123 {
		t.Fail()
	}

	if NamedSome("", 123).IsNone() || NamedSome("", 123).SomeOr(456).(int) != 123 {
		t.Fail()
	}

	db, err := sql.Open("mysql", `root:singularity0618@tcp(10.241.230.112:3306)/snda_cloud_accounts`)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		t.Error(err)
	}
}
