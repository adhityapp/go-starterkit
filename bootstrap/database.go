package bootstrap

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c *Container) Dbw() *sqlx.DB {
	if c.db == nil {
		c.db = c.DBMustConnect(true)
	}
	return c.db
}

func (c *Container) Dbr() *sqlx.DB {
	if c.db == nil {
		c.db = c.DBMustConnect(false)
	}
	return c.db
}

func (c *Container) DBMustConnect(isWrite bool) *sqlx.DB {
	domain := "write"
	if !isWrite {
		domain = "read"
	}

	var dsn string
	host := viper.GetString("database.postgres." + domain + ".hostname")
	dbname := viper.GetString("database.postgres." + domain + ".dbname")
	username := viper.GetString("database.postgres." + domain + ".username")
	password := viper.GetString("database.postgres." + domain + ".password")

	dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable", host, username, password, dbname)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Panic(err)
	}
	if err := conn.Ping(); err != nil {
		logrus.Panic("Error connecting to SQL Server: ", err.Error())
	}
	fmt.Println("Connected to SQL Server!")
	return conn
}
