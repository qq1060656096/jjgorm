package jjmgorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func sqlLite3Config(name string) Config {
	dataSource := sqlLite3ConfigDataSource(name)
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	return conf
}

func sqlLite3ConfigDataSource(name string) string {
	dir, _ := os.Getwd()
	dataSource := dir + "/testdata/test-sqlite3-" + name + ".db"
	return dataSource
}

func TestManager_Add(t *testing.T) {
	m := NewManager()
	m.Add("managerAdd", sqlLite3Config("managerAdd"))
	m.Add("managerAdd2", sqlLite3Config("managerAdd2"))
	assert.Equal(t, 2, m.Len())
}

func TestManager_Remove(t *testing.T) {
	m := NewManager()
	m.Add("managerRemove", sqlLite3Config("managerRemove"))
	m.Add("managerRemove2", sqlLite3Config("managerRemove2"))
	assert.Equal(t, 2, m.Len())
	m.Remove("managerRemove")
	assert.Equal(t, 1, m.Len())
}

func TestManager_Change(t *testing.T) {
	m := NewManager()
	m.Add("managerChange", sqlLite3Config("managerChange"))
	m.Add("managerChange2", sqlLite3Config("managerChange2"))
	assert.Equal(t, 2, m.Len())
	m.Change("managerChange2", sqlLite3Config("managerChange2"))
	assert.Equal(t, 2, m.Len())
	m.Change("managerChange3", sqlLite3Config("managerChange3"))
	assert.Equal(t, 3, m.Len())
}

func TestManager_Get(t *testing.T) {
	m := NewManager()
	m.Add("managerGet", sqlLite3Config("managerGet"))
	m.Add("managerGet2", sqlLite3Config("managerGet2"))
	assert.Equal(t, 2, m.Len())
	conn := m.Get("managerGet2")
	assert.Equal(t, sqlLite3ConfigDataSource("managerGet2"), conn.Config().DataSource)
	conn = m.Get("managerGet")
	assert.Equal(t, sqlLite3ConfigDataSource("managerGet"), conn.Config().DataSource)
}

func TestManager_Exist(t *testing.T) {
	m := NewManager()
	m.Add("managerExist", sqlLite3Config("managerExist"))
	m.Add("managerExist2", sqlLite3Config("managerExist2"))
	assert.Equal(t, 2, m.Len())
	assert.True(t, m.Exist("managerExist2"))
	assert.False(t, m.Exist("managerExist3"))
}

func TestManager_StringNilConnections(t *testing.T) {
	m := &Manager{}
	fmt.Println(m.String())
	assert.Equal(t, m.String(), "")
}

func TestManager_String0Len(t *testing.T) {
	m := NewManager()
	fmt.Println(m.String())
	assert.Equal(t, m.String(), "")
}

func TestManager_String(t *testing.T) {
	m := NewManager()
	m.Add("managerString", sqlLite3Config("managerString"))
	m.Add("managerString2", sqlLite3Config("managerString2"))
	fmt.Println(m.String())
	assert.True(t, len(m.String()) > 0)
}
