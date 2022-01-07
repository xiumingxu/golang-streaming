package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//func openConn() *sql.DB{
//	dbConn, err := sql.Open("mysql", "root:test@tcp(localhost:3306)")
//	if err != nil{
//		panic(err.Error())
//	}
//	return dbConn //最好不要总openConn, 否则很多redudant
//} ==> move到 conn.go里了

func AddUserCredential(loginName string, pwd string) error {
	//不要用 +
	// Prepare也是为了安全
	//"INSERT INTO users(login_name, pwd) VALUES (?, ?)"
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil

}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	// defer 要少用 , 只在站推出的时候才调用
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

// open connection 一般是放到config cmdb 后期, config 要节后
