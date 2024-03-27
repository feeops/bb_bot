package database

import (
	"bb_bot/ent"
	"context"
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"

	ents "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
)

var (
	once sync.Once
	// DBConn is a database connection singleton.
	DBConn *ent.Client
)

func InitDB() *ent.Client {
	once.Do(func() {
		db, err := sql.Open("mysql",
			viper.GetString("DBDSN"),
		)
		if err != nil {
			log.Fatal("failed open mysql: %v", err)
		}

		db.SetMaxIdleConns(6)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour * 2)
		drv := ents.OpenDB("mysql", db)

		// 可以添加ent.Debug()参数开启debug
		DBConn = ent.NewClient(ent.Driver(drv))

		ctx := context.Background()
		// Run migration.
		err = DBConn.Schema.Create(ctx)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

	})

	return DBConn
}
