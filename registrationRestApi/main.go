package main 
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"dbconnectionApi"
	"entities"
	"database/sql"
	"userAuthentication"
  _ "github.com/go-sql-driver/mysql"

)
var user []entities.User
var db *sql.DB
var err error

func userLogin(w http.ResponseWriter , req *http.Request){
	userAuthentication.GetUserDetailurl(w,req,db)
}

func CreateUserDetailurl(w http.ResponseWriter , req *http.Request){
	//params :=mux.Vars(req)
	var users entities.User
	_ = json.NewDecoder(req.Body).Decode(&users)
	//users.Email=params["email"]
	user = append(user,users)
    json.NewEncoder(w).Encode(user)
}
func DeleteUserDetailurl(w http.ResponseWriter , req *http.Request){
  params :=mux.Vars(req)
  for index,item :=range user{
	  if item.Email == params["email"]{
		  user = append(user[:index], user[index+1:]...)
		  break
	  }
	  json.NewEncoder(w).Encode(user)
  }	
}
func main() {
	 getDB() 
	 router:=mux.NewRouter()
     //user =append(user,User{Email:"sharad@gmail.com",Password:"Test",Phoneno:12344,Salary:3456,City:"Delhi",Country:"India",First_name:"Sharad",Last_name:"Maurya",})
	 router.HandleFunc("/getUser/{email}&&{password}",userLogin).Methods("GET")
	 router.HandleFunc("/createUser",CreateUserDetailurl).Methods("POST")
	 router.HandleFunc("/deleteUser/{email}",DeleteUserDetailurl).Methods("DELETE")
	 log.Fatal(http.ListenAndServe(":8002",router))
	 
}
func getDB(){
	db,err=dbconnectionApi.GetConnection()
	if err !=nil{
		panic(err.Error())
	}
	fmt.Print("Connected DB")
}

