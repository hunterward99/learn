package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Developer struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Active     bool   `json:"active"`
	TechSkills struct {
		Languages    []string
		Achievements []string
	} `json:"tech_skills"`
}

func main() {
	var dev Developer
	LoadData("cmd/developers.json", &dev)

	fmt.Printf("Имя: %s\nВозраст: %d\nАктивный: %t\n", dev.Name, dev.Age, dev.Active)
	fmt.Println("Навыки:")
	fmt.Printf("  Языки: %v\n", dev.TechSkills.Languages)
	fmt.Printf("  Достижения: %v\n", dev.TechSkills.Achievements)
}

func LoadData(path string, v interface{}) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(v); err != nil {
		log.Fatal(err)
	}

}
