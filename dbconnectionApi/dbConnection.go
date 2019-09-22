package dbconnectionApi
import (
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
)
func GetConnection()(db *sql.DB,err error){
	dbDriver :="mysql"
	dbUName :="root"
	dbPassword :="root"
	dbName :="userinfo"
	db,err = sql.Open(dbDriver,dbUName+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName)
	return db,err
}
func GetDBUser(db *sql.DB)(rows1 *sql.Rows){
	rows1,_ =db.Query("select email,password from userinfo")
	return rows1
}