package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
使用mysql 练习
*/
type people struct {
	id   int64  `db:"id"`
	name string `db:"name"`
	age  int    `db:"age"`
}

// 数据库对象只需要创建一次,因为这个数据库对象可以安全的供给多个goroutine使用,只需要全局初始化一次
//var db *sql.DB
var db *sqlx.DB

func main() {
	mysqlDsn := "root:123456@tcp(127.0.0.1:3306)/zhibi_rank?charset=utf8mb4&parseTime=True"
	//建立连接前先设置好 建立连接失败就关闭连接
	//defer db.Close()
	initConn(mysqlDsn)
	var user people
	//queryRow(3,&user)
	//queryRows("张三",user)
	user = people{
		name: "哈哈1回",
		age:  12,
	}
	add(user)
	//delete(4)
	//update(user)
}

/*
初始化数据库
*/
/*func initConn(mysqlDsn string) (err error) {
	//Open方法只验证其参数而不创建与数据库的连接。要验证数据源名称是否有效，请调用// Ping()
	db,err=sql.Open("mysql",mysqlDsn)
	if err !=nil {
		fmt.Println("连接数据库信息错误")
	}
	//与数据库建立连接
	err=db.Ping()
	if err !=nil {
		fmt.Println("与数据库建立连接失败")
	}
	//设置最大空闲连接 默认值为 2
	db.SetMaxIdleConns(8)
	//设置最多连接数 默认值为0(无限制)
	db.SetMaxOpenConns(0)
	return
}*/
/*
初始化数据库 使用第三方包 更方便一点
*/
func initConn(mysqlDsn string) (err error) {
	//调用 Connect() 方法内部会执行 open和ping方法
	db, err = sqlx.Connect("mysql", mysqlDsn)
	if err != nil {
		fmt.Println("与数据库建立连接失败", err)
	}
	//设置最大空闲连接 默认值为 2
	db.SetMaxIdleConns(8)
	//设置最多连接数 默认值为0(无限制)
	db.SetMaxOpenConns(0)
	return
}

/*
查询单行方法
*/
func queryRow(id int64, user people) {
	sqlstr := "SELECT * FROM myuser WHERE id = ?"
	//先预处理sql 避免sql注入问题 和提升性能
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("预编译sql失败!")
	}
	//记得关闭预处理对象
	defer stmt.Close()
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放 这里必须传递结构体的指针才可以赋值
	err = stmt.QueryRow(id).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Println("数据映射结构体失败!", err)
	}
	fmt.Printf("id:%d name:%s age:%d\n", user.id, user.name, user.age)
}

/*
查询多行方法
*/
func queryRows(name string, user people) {
	sqlstr := "SELECT * FROM myuser WHERE name = ?"
	//先预处理sql 避免sql注入问题 和提升性能
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("预编译sql失败!")
	}
	//记得关闭预处理对象
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("数据映射结构体失败!", err)
	}
	for rows.Next() {
		rows.Scan(&user.id, &user.name, &user.age)
		fmt.Printf("id:%d name:%s age:%d\n", user.id, user.name, user.age)
	}
}

/*
插入数据 并测试 事务
保证两条数据同时插入,有一条失败另一条也不插入
*/
func add(user people) {
	//开启事务
	tx, err := db.Begin()
	//检查事务开启是否报错
	if err != nil {
		//报错后看是否已经开启事务
		if tx != nil {
			//已经开启则立刻回滚
			tx.Rollback()
		}
		fmt.Println("开启事务失败!", err)
		return
	}
	sqlstr := "INSERT INTO myuser (name,age) VALUES (?,?)"
	//先预处理sql 避免sql注入问题 和提升性能
	//开启事务需要用事务对象 来进行预处理操作才能使事务生效
	stmt, err := tx.Prepare(sqlstr)
	if err != nil {
		fmt.Println("预编译sql失败!", err)
		return
	}
	//记得关闭预处理对象
	defer stmt.Close()
	res1, err1 := stmt.Exec(user.name, user.age)
	res2, err2 := stmt.Exec(user.name, "true")
	if err1 != nil || err2 != nil {
		fmt.Println("插入数据失败开始回滚!", err1, err2)
		err := tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败!")
		}
		fmt.Println("回滚成功")
		return
	}
	i1, err1 := res1.RowsAffected()
	i2, err2 := res2.RowsAffected()
	if err1 != nil || err2 != nil {
		fmt.Println("获取影响行数失败!", err)
		err := tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败!")
		}
		fmt.Println("回滚成功")
		return
	}
	if i1 == 1 && i2 == 1 {
		fmt.Printf("受影响的行数:%d,%d", i1, i2)
		//提交事务
		err := tx.Commit()
		if err != nil {
			fmt.Println("提交事务失败")
		}
		return
	} else {
		//回滚事务
		tx.Rollback()
	}
}

/*
删除数据
*/
func delete(id int64) {
	sqlstr := "DELETE FROM myuser WHERE id=?"
	//先预处理sql 避免sql注入问题 和提升性能
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("预编译sql失败!", err)
	}
	//记得关闭预处理对象
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("插入数据失败!", err)
	}
	i, err := res.RowsAffected()
	fmt.Printf("受影响的行数:%d", i)
	//删除数据不返回 删除数据的id
	idd, err := res.LastInsertId()
	fmt.Printf("删除数据的id为:%d", idd)
}

/*
修改数据
*/
func update(user people) {
	sqlstr := "UPDATE myuser SET name = ?, age = ? WHERE id = ?"
	//先预处理sql 避免sql注入问题 和提升性能
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("预编译sql失败!", err)
	}
	//记得关闭预处理对象
	defer stmt.Close()
	res, err := stmt.Exec(user.name, user.age, user.id)
	if err != nil {
		fmt.Println("插入数据失败!", err)
	}
	i, err := res.RowsAffected()
	fmt.Printf("受影响的行数:%d", i)
	//修改数据不返回 修改数据的id
	idd, err := res.LastInsertId()
	fmt.Printf("修改数据的id为:%d", idd)
}
