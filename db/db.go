package db

type Db interface {
	Save(interface{}) error
}

type MyDb struct {
}

func (db *MyDb) Save(interface{}) error {
	// 就是不出错， 你打我啊
	return nil
}
