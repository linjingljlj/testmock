package record

import "testing"
import (
	"github.com/stretchr/testify/assert"

	"testmock/db"

)

func TestRecordSave(t *testing.T) {
	db := &db.MyDb{}
	r := &Record{Name:"lin"}
	err := r.Save(db)
	assert.Equal(t, nil, err)


	r = &Record{Name:"jing"}
	err = r.Save(db)
	assert.Equal(t, nil, err)
	// assert.NotEqual(t, nil, err)
}



/*
$ go test testmock/record -covermode=count -coverprofile /tmp/r.cov
ok  	testmock/record	0.001s	coverage: 60.0% of statements

$ go test testmock/record -covermode=count -coverprofile /tmp/r.cov
ok  	testmock/record	0.003s	coverage: 100.0% of statements
 */
