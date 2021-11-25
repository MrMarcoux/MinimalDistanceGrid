package distancegrid

import (
	"fmt"
	"strings"
	"math"
)


type SubstitutionSet struct {
	Costs map[rune]int
}

type Grid struct {
	Insertion int
	Deletion int
	Substitutions map[rune]SubstitutionSet
	
	Comparand string
	Comparator string
}

func CreateRuneSet(s1 string, s2 string) []rune {
	a1 := []rune(s1)
	a2 := []rune(s2)
	
	unfilteredSet := make([]rune, len(a1) + len(s2))
	copy(unfilteredSet[:], a1[:])
	copy(unfilteredSet[len(a1):], a2[:])

	allKeys := make(map[rune]bool)
    filteredSet := []rune{}

    for _, item := range unfilteredSet {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            filteredSet = append(filteredSet, item)
        }
    }

    return filteredSet
}

func NewGrid(insertion int, deletion int, substitution int, correspondance int, comparand string, comparator string) *Grid {
	grid := new(Grid)
	grid.Insertion = insertion
	grid.Deletion = deletion
	grid.Comparand = comparand
	grid.Comparator = comparator
	grid.Substitutions = make(map[rune]SubstitutionSet)

	runeSet := CreateRuneSet(comparand, comparator)

	for _, key := range runeSet {
		if _, ok := grid.Substitutions[key]; !ok {
			substitutionSet := new(SubstitutionSet)
			substitutionSet.Costs = make(map[rune]int)
			
			for _, value := range runeSet {
				if key == value {
					substitutionSet.Costs[value] = correspondance
				} else {
					substitutionSet.Costs[value] = substitution
				}
			}

			grid.Substitutions[key] = *substitutionSet
		}
	}

	return grid
}

func(grid Grid) Values() [][]int {
	values := make([][] int, len(s1), len(s2))
	a1 := []rune(s1)
	a2 := []rune(s2)

	for i:=0; i<len(s1); i++ {
		for j:=0; j<len(s2); j++ {
			if i == 0 {
				values[i][j] = grid.Insertion * j
			} else if j == 0 {
				values[i][j] = grid.Insertion * i
			} else {
				insertionCandidate := values[i-1][j] + grid.Insertion
				deletionCandidate := values[i][j-1] + grid.Deletion
				substitutionCandidate := values[i-1][j-1] + grid.Substitutions[a1[i]].Costs[a2[i]]
				
				values[i][j] = math.Max(insertionCandidate, deletionCandidate, substitutionCandidate)
			}
		}
    }

	return values
}

func(grid Grid) Print() {

	numericalValues = grid.Values()

	fmt.Printf("|ϕ|")
	
	for _, value := range []rune(s1) {
		fmt.Printf("%d|", value)
	}

	for index, value := range []rune(s2) {
		fmt.Printf("|%d|", value)

		// Print the numerical values here
		for _, distance := range numericalValues[index] {
			fmt.Printf("%d|", distance)
		}
	}
}

