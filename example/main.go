package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"runtime"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Sleeper interface {
	Sleep()
}

type Eater interface {
	Eat(foodName string)
}

type LazyAnimal interface {
	Sleeper
	Eater
}

type Animal struct {
	Name string `json:"name,omitempty"  xml:"name"`
	Age  int    `json:"age,omitempty"  xml:"age"`
}

type Dog struct {
	Animal
	DogLittleName string `json:"dog_little_name,omitempty"  xml:"dog_little_name"`
}

type Cat struct {
	Animal
	CatLittleName string `json:"cat_little_name,omitempty"  xml:"cat_little_name"`
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping...%s\n", d.Name, d.DogLittleName)
}

func (d Dog) Eat(foodName string) {
	fmt.Printf("dog %s is eating %s\n", d.Name, foodName)
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping...%s\n", c.Name, c.CatLittleName)
}

func (c Cat) Eat(foodName string) {
	fmt.Printf("cat %s is eating %s\n", c.Name, foodName)
}

func fibonacci(ch, quit chan int) {
	x, y := 1, 1
	for true {
		select {
		case ch <- x:
			fmt.Println(strconv.Itoa(runtime.NumCPU()) + "***")
			x, y = y, x+y
		case <-quit:
			fmt.Println("done")
			break
		}
	}
}

func test9() {
	fmt.Println(net.ParseIP("2002:c0e8:82e7:0:0:0:c0e8:82e7"))
}

func test8() {

	db, err := sql.Open("mysql", "vgb_process:vgb_process@tcp(10.36.188.50:5002)/vgb_process")
	if err != nil {
		panic("建立db连接错误")
	}
	rows, err := db.Query("select process_task_id,process_type,origin_file_id,new_file_id,process_content from vgb_process.process_record")
	if err != nil {
		panic("sql语句错误")
	}

	for rows.Next() {
		type model struct {
			processTaskId  int64  `json:"process_task_id,omitempty"  db:"process_task_id"`
			processType    int    `json:"process_type,omitempty"  db:"process_type"`
			originFileId   string `json:"origin_file_id,omitempty"  db:"origin_file_id"`
			newFileId      string `json:"new_file_id,omitempty"  db:"new_file_id"`
			processContent string `json:"process_content,omitempty"  db:"process_content"`
		}
		m := new(model)
		rows.Scan(&m.processTaskId, &m.processType, &m.originFileId, &m.newFileId, &m.processContent)
		fmt.Println(m)
	}
}

func main() {
	db, err := sql.Open("mysql", "vgb_core_develop:vgb_core_develop@tcp(10.36.188.50:5002)/vgb_core_develop")
	if err != nil {
		panic("建立db连接错误")
	}
	stmt, err := db.Prepare("insert into vgb_core_develop.test (a,b) values (?,?)")
	for i := 28891; i < 50000; i++ {
		stmt.Exec(i, i*2)
		fmt.Println("ok")
	}
	fmt.Println("done")
	//if err != nil {
	//	fmt.Println(err)
	//	panic("创建sql语句错误")
	//}
	//res, err := stmt.Exec("123123", "456456")
	//if err != nil {
	//	panic("执行sql错误")
	//}
	//id, err := res.LastInsertId()
	//if err != nil {
	//	panic("插入数据失败")
	//}
	//fmt.Println(id)
}

func test6() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println(r.Method)
		fmt.Println(r.URL.Scheme)
		template.ParseFiles()
		for k, v := range r.Form {
			fmt.Println(k, "=", v)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func test5(m map[string]string) {
	slice1 := make([]int, 10)
	slice1[1] = 2
	fmt.Println(slice1)
	test4(slice1)
	fmt.Println(slice1)

	m = make(map[string]string)
	m["a"] = "zzz"
	fmt.Println(m)
	test5(m)
	fmt.Println(m)
}

func test4(s []int) {
	s[1] = 99
}

func test3() {
	ch := make(chan int)
	quit := make(chan int)
	go fibonacci(ch, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	quit <- 1
	fmt.Println("done main")
}

func test2() {
	d := Dog{Animal: Animal{"doggg", 10}, DogLittleName: "sd"}
	c := Cat{Animal: Animal{"cattt", 15}, CatLittleName: "sc"}
	lazyAnimals := []LazyAnimal{d, c}
	for _, animal := range lazyAnimals {
		animal.Eat("anc")
		animal.Sleep()
	}

}

func test() {
	d := Dog{Animal: Animal{"doggg", 10}, DogLittleName: "sd"}
	c := Cat{Animal: Animal{"cattt", 15}, CatLittleName: "sc"}

	d.Sleep()
	d.Eat("noodles")
	fmt.Println()
	c.Sleep()
	c.Eat("beef")
}
