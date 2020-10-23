package model

import (
	"encoding/json"
	"errors"
	"log"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//表明SelectUserByName 是User的一个方法
func (user *User)SelectUserByName(name string) error  {
	stmt,err :=DB.Prepare("SELECT user_name,password FROM user WHERE user_name=?")
	if nil != err {
		return err
	}
	defer stmt.Close()
	rows,err :=stmt.Query(name)
	defer rows.Close()
	if nil != err {
		return err
	}
	// 数据处理
	for rows.Next() {
		rows.Scan(&user.UserName,&user.Password)
	}
	if err :=rows.Err();err != nil {
		return err
	}
	return nil
}

func (u *User) Validate() error {
	if u.UserName == "" || u.Password == "" {
		return errors.New("用户名和密码不能为空")
	}
	return nil
}

func (user *User) Create() (int64,error) {
	id,err := Insert("Insert into user(user_name,password) values (?,?)",user.UserName,user.Password)
	if err != nil {
		return 1,err
	}
	return id,err
}

func  (user *User) UserToJson() string {
	jsonStr,err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	return string(jsonStr)
}

func (user *User)JsonToUser(jsonBlob string)error  {
	err := json.Unmarshal([]byte(jsonBlob), &user)
	if err != nil {
		return err
	}
	return nil
}