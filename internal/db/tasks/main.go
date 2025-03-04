package main

import "log"

func main() {
	if err := setCategoriesSlug(); err != nil {
		log.Fatal(err)
	}
}
