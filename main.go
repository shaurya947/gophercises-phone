package main

import (
	"log"
	"os"

	"github.com/shaurya947/gophercises-phone/store"
)

func main() {
	phoneStore, err := store.NewStore(os.Getenv("ELEPHANT_SQL_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	phoneStore.Insert([]*store.PhoneNumber{
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
