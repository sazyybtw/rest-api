package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		DBName:               Envs.DBName,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStoage := NewMySQLStorage(cfg)

	db, err := sqlStoage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	api := NewAPIServer(":3000", store)
	api.Serve()
}
