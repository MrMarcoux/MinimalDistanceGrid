package main

import (
    "os"
    "fmt"
	"io/ioutil"
    "gopkg.in/yaml.v3"
	"matchanalysis/distancegrid"
)
 
func main() {
    
    filePath := os.Args[1]

	yfile, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		panic(err2)
	}

	// Type classification is successful here, just need to setup parsing and call constructors
	switch v := data["substitution"].(type) {
		case int:
			grid := distancegrid.NewGrid(
				data["insertion"].(int),
				data["deletion"].(int),
				data["substitution"].(int),
				data["correspondance"].(int),
				data["s1"].(string),
				data["s2"].(string))

			fmt.Printf("It's the second!", v)
			fmt.Printf("It's the second!", grid.Insertion + 1)
		default:
			substitutions := make(map[rune]distancegrid.SubstitutionSet)
			
			for key, val := range data["substitution"].(map[string]interface{}) {
				runes := []rune(key)
				leftSideRune := runes[0]
				rightSideRune := runes[2]
				
				if _, ok := substitutions[leftSideRune]; !ok {
					// Rune wasnt registered
					substitution := distancegrid.SubstitutionSet{
						Costs: make(map[rune]int),
					}

					substitutions[leftSideRune] = substitution
				}

				substitutions[leftSideRune].Costs[rightSideRune] = val.(int)
			}

			grid := distancegrid.Grid{
				Insertion: data["insertion"].(int),
				Deletion: data["deletion"].(int),
				Substitutions: substitutions,
				Comparand: data["s1"].(string),
				Comparator: data["s2"].(string)}
			
			fmt.Printf("It's the first!", grid.Insertion)
	}

	for k, v := range data {
		fmt.Printf("%s -> %d\n", k, v)
	}
}