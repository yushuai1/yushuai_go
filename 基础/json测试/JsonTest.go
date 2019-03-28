package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors  []string
}


func main() {

	movie:=new(Movie)
	movie.Title = "南京"
	movie.Year = 100
	movie.Color = false
	movie.Actors = []string{"ni","hao"}


	data,err := json.Marshal(movie)

	if  err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	movie1 :=new(Movie)

	if err := json.Unmarshal(data,&movie1);err !=nil{
		fmt.Println(err)
	}

	fmt.Println(movie1.Actors)


}
