package bootstrap

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
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

	host := viper.GetString("database.postgres." + domain + ".hostname")
	dbname := viper.GetString("database.postgres." + domain + ".dbname")
	username := viper.GetString("database.postgres." + domain + ".username")
	password := viper.GetString("database.postgres." + domain + ".password")

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=disable", host, username, password, dbname)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Panic(err)
	}
	err = conn.Ping()
	if err != nil {
		logrus.Panic("Error connecting to SQL Server: ", err.Error())
	}
	fmt.Println("Connected to SQL Server!")
	return conn
}
