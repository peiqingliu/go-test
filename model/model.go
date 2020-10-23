package model

import (
	"fmt"
	"log"
)

//插入操作  ...可变类型参数 interface{}表示任意类型
func Insert(sql string,args ...interface{}) (int64,error) {
	stmt,err :=DB.Prepare(sql)
	if err != nil {
		return 1,err
	}

	result,err :=stmt.Exec(args...)
	if err != nil {
		return 1,nil
	}

	id,err :=result.LastInsertId()
	if err != nil {
		return 1,err
	}

	fmt.Printf("插入成功，ID为%v\n",id)
	return id,nil
}

//删除操作
func Delete(sql string,args ...interface{})  {
	stmt,err :=DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err,"SQL语句设置失败")
	result,err :=stmt.Exec(args...)
	CheckErr(err, "参数添加失败")
	num,err :=result.RowsAffected()
	CheckErr(err,"删除失败")
	fmt.Printf("删除成功，删除行数为%d\n",num)
}


//	Update 修改操作
func Update(sql string,args... interface{})  {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err, "SQL语句设置失败")
	result, err := stmt.Exec(args...)
	CheckErr(err, "参数添加失败")
	num, err := result.RowsAffected()
	CheckErr(err,"修改失败")
	fmt.Printf("修改成功，修改行数为%d\n",num)
}

//用来校验 err是否为空
func CheckErr(err error,msg string)  {
	if err != nil {
		log.Panicln(msg,err)
	}
}