package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
type SingleGame struct{
	name string
	reit float64
	platform string
	date string
}
var (
	d=0

id=&d)
var g SingleGame
func write(n string,r float64,p string,d string) {
	*id++
	db, _ := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer db.Close()
	result,err:=db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)",*id,n,r,p,d)
	if err != nil{
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
func read (ch int){
	var pustishka int
	db, _ := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	res, _ := db.Query("select * from TopGames where id = $1",ch)
	for res.Next(){
		res.Scan(&pustishka,&g.name, &g.reit, &g.platform, &g.date)
	}
	res.Close()
	db.Close()
}
func main() {
	var ch int
	i:=1
	for i!=3{
		fmt.Println("Выберите действие:\n1. Запись\n2. Чтение\n3. Выход")
		fmt.Scan(&i)
		switch i{
		case 1:
			fmt.Println("Введите название, рейтинг, платформу и дату релиза игры:")
			fmt.Scan(&g.name,&g.reit,&g.platform,&g.date)
			write(g.name,g.reit,g.platform,g.date)
		case 2:
			fmt.Println("Какую строку вывести?")
			fmt.Scan(&ch)
			read(ch)
			fmt.Println("Название: ",g.name,"\nРейтинг: ",g.reit,"\nПлатформу: ",g.platform,"\nДата релиза игры: ",g.date)
		case 3:
			break
		default:
			fmt.Println("Неверный выбор!!!")
		}
	}

}
