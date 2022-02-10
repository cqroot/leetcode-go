package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Problem struct {
	Id   string
	Name string
	Dir  string
}

type Problems struct {
	Algorithms []Problem
	Database   []Problem
}

func getDirs(name string) []fs.FileInfo {
	dirs, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}
	return dirs
}

func getAlgorithmsProblems() []Problem {
	problems := []Problem{}

	dirs := getDirs("algorithms")
	for _, dir := range dirs {
		for _, problem := range getDirs(filepath.Join("algorithms", dir.Name())) {
			items := strings.Split(problem.Name(), ".")
			p := Problem{
				Id:   items[0][1:],
				Name: items[1],
				Dir:  fmt.Sprintf("%s/%s", dir.Name(), problem.Name()),
			}
			problems = append(problems, p)
		}
	}

	return problems
}

func getDatabaseProblems() []Problem {
	problems := []Problem{}

	dirs := getDirs("database")
	for _, file := range dirs {
		items := strings.Split(file.Name(), ".")
		p := Problem{
			Id:   items[0][1:],
			Name: items[1],
			Dir:  fmt.Sprintf("%s/%s", file.Name(), file.Name()),
		}
		problems = append(problems, p)
	}

	return problems
}

func main() {
	problems := Problems{
		Algorithms: getAlgorithmsProblems(),
		Database:   getDatabaseProblems(),
	}

	tmpl, err := template.ParseFiles("./README.tmpl")
	if err != nil {
		panic(err)
	}

	output, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	tmpl.Execute(output, problems)
}
