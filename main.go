package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	ID        int
	FirstName string
	LastName  string
	Subject   string
	Score     int64
}

func main() {
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Argument needs to be an Integer", err)
	}

	file, err := os.Create("data.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write([]string{"ID", "FirstName", "LastName", "Subject", "Score"}); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	for i := 0; i < size; i++ {
		record := createRow(i)
		row := []string{strconv.Itoa(record.ID),
			record.FirstName,
			record.LastName,
			record.Subject,
			strconv.Itoa(int(record.Score))}

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	log.Printf("Done")
}

func createRow(i int) Row {
	subjects := [6]string{"Python", "C", "C++", "Java", "Go", "Rust"}
	//seed := time.Now().UTC().UnixNano()

	firstName := getFirstName()
	firstName = strings.Trim(firstName, "\"")
	lastName := getLastName()
	lastName = strings.Trim(lastName, "\"")

	subject := subjects[rand.Intn(len(subjects))]
	score := rand.Int63n(100)
	row := Row{
		ID:        i,
		FirstName: firstName,
		LastName:  lastName,
		Subject:   subject,
		Score:     score,
	}
	return row
}

func getFirstName() string { return getName("first") }
func getLastName() string  { return getName("last") }

func getName(endpoint string) string {
	response, err := http.Get(fmt.Sprintf("http://127.0.0.1:8000/name/%s", endpoint))
	if err != nil {
		return ""
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}
