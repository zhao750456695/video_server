package dbops

import (
	"testing"
)
import _ "github.com/go-sql-driver/mysql"



func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}

func TestMain(m *testing.M)  {
    clearTables()
    m.Run()
    clearTables()
}

func TestUserWorkFlow(t *testing.T){
    t.Run("Add", TestAddUser)
    t.Run("Get", TestGetUser)
    t.Run("Del", TestDeleteUser)
    t.Run("Reget", TestRegetUser)
}

func TestAddUser(t *testing.T)  {
    err := AddUserCredential("root", "root")
    if err != nil{
    	t.Errorf("Error of Adduser:%v", err)
	}
}

func TestGetUser(t *testing.T) {
    pwd, err := GetUserCredential("root")
    if pwd != "root" || err != nil{
		t.Errorf("Error of Getuser")
	}
}

func TestDeleteUser(t *testing.T) {
    err := DeleteUser("root", "root")
	if err != nil{
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func TestRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("root")
	if err != nil{
		t.Errorf("Error of RegetUser")
	}
	if pwd != ""{
		t.Errorf("Delete failed")
	}
}