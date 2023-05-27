package database

import "fmt"

type Database struct{}

func (d *Database) DBAddress() string {
	return fmt.Sprintf("%p", d)
}
