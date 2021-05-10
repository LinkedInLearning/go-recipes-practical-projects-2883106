package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	items := []Item{
		{"m183x", "Magic Wand"},
		{"m184y", "Invisibility Cape"},
		{"m185z", "Levitation Spell"},
	}

	if err := writeItems("items.csv", items); err != nil {
		log.Fatal(err)
	}
}

type Item struct {
	SKU  string
	Name string
}

func writeItems(fileName string, items []Item) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	row := []string{"sku", "name"}

	wtr := csv.NewWriter(file)
	defer wtr.Flush()

	if err := wtr.Write(row); err != nil {
		return err
	}

	for _, item := range items {
		row[0] = item.SKU
		row[1] = item.Name
		if err := wtr.Write(row); err != nil {
			return err
		}
	}

	return wtr.Error()
}
