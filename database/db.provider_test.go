package database

import "testing"

func Test_InitDatabase(t *testing.T) {
	NewProvider("./db")
}
