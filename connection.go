package connection

import (
	"github.com/jmoiron/sqlx"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

var (
	pgsql *sqlx.DB
	err error
)

func initPgsql(){
	pgInfo := configuration.postgresql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	pgInfo.Host, pgInfo.Port, pgInfo.User, pgInfo.Password, pgInfo.Dbname)
	if pgsql, err = sqlx.Connect("postgres", psqlInfo); err != nil{
		log.Println("Connection Err : ", err.Error())
	} else{
		log.Println("connection Created Success")
	}
}

func SetPgsql(pgsql Pgsql){
	configuration.postgresql = pgsql
}

func GetConnectionPsql()(*sqlx.DB) {
	if pgsql ==nil{
		initPgsql()
	}
	return pgsql
}
