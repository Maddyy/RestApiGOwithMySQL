package userAuthentication
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"database/sql"
	"net/http"
	"dbconnectionApi"
	"entities"
  _ "github.com/go-sql-driver/mysql"

)
var err error
//var db *sql.DB
func GetUserDetailurl(w http.ResponseWriter , req *http.Request,db *sql.DB){
	item1 :=	dbconnectionApi.GetDBUser(db)
	 params :=mux.Vars(req)
	 for item1.Next(){
		 var usr entities.User
		 err=item1.Scan(&usr.Email,&usr.Password)
			  if err!=nil{
				 panic(err.Error())
				}   
			 if usr.Email==params["email"] && usr.Password==params["password"]{
				 fmt.Print("user found\n",usr.Email,usr.Password)
				 json.NewEncoder(w).Encode(usr)
				 return
			 } 
			 
	 }
	 fmt.Print("user not found\n") 	
	 json.NewEncoder(w).Encode(&entities.User{})
 }