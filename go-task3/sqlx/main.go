package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	var db *sqlx.DB
	var err error
	db, err = connectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	//1.编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	fmt.Println(getEmployeesByDepartment(db, "技术部"))
	//2.编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	fmt.Println(getHighestPaidEmployee(db))

	//3.编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	fmt.Println(getBooksAbovePrice(db, 30.0))
	if err != nil {
		fmt.Println("Error querying employees:", err)
	}
	defer db.Close()

}

// sqlx连接mysql数据库
func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 emplsoyees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func getHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// 设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func getBooksAbovePrice(db *sqlx.DB, price float64) ([]Book, error) {
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err := db.Select(&books, query, price)
	if err != nil {
		return nil, err
	}
	return books, nil
}
