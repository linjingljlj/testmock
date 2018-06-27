package record

import (
	"testmock/db"
	"fmt"
)

type Record struct {
	Name string
}

func (r *Record) Save(db db.Db) error {
	err := db.Save(r)
	if err != nil {
		fmt.Printf("save err %s", err)
		return err
	}
	fmt.Print("save success")
	return nil
}
