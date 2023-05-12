package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shaurya947/gophercises-phone/store"
)

func main() {
	phoneStore, err := store.NewStore(os.Getenv("ELEPHANT_SQL_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	// insertSampleData(phoneStore)
	// queryAndDisplay(phoneStore)
	// normalizeNumbers(phoneStore)
	// queryAndDisplay(phoneStore)
	// removeDupes(phoneStore)
	queryAndDisplay(phoneStore)
}

func insertSampleData(s *store.PhoneStore) {
	s.Insert([]*store.PhoneNumber{
		{Number: "1234567890"},
		{Number: "123 456 7891"},
		{Number: "(123) 456 7892"},
		{Number: "(123) 456-7893"},
		{Number: "123-456-7894"},
		{Number: "123-456-7890"},
		{Number: "1234567892"},
		{Number: "(123)456-7892"},
	})
}

func queryAndDisplay(s *store.PhoneStore) {
	phoneNums, err := s.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	for _, pn := range phoneNums {
		fmt.Println(pn.Number)
	}
}

func normalizeNumbers(s *store.PhoneStore) {
	err := s.Normalize(normalizePhone)
	if err != nil {
		log.Fatalln(err)
	}
}

func removeDupes(s *store.PhoneStore) {
	err := s.RemoveDupes()
	if err != nil {
		log.Fatalln(err)
	}
}
