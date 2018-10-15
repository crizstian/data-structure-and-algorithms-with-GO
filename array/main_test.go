package array

import (
	"fmt"
	"log"
	"testing"
)

func TestGrades(t *testing.T) {
	log.Println("### Excercise 1")
	log.Println("## Implementing reduce() func like javascript")
	gradeObj := []int{90, 89, 75, 90, 89, 75, 90, 89, 75, 90, 89, 75, 90, 89, 75, 90, 89, 75}
	grade := NewGrades(gradeObj)

	result := grade.reduce(func(prev interface{}, current interface{}) interface{} {
		p := prev.(int)
		p += current.(int)
		return p
	}, 0)

	fmt.Println("The result is: ", result)
}

func TestWords(t *testing.T) {
	log.Println("### Excercise 2")
	arrayWords := []string{"hello ", "my ", "friend "}
	words := NewWords(arrayWords)

	result := words.reduce(func(prev interface{}, current interface{}) interface{} {
		p := prev.(string)
		p += current.(string)
		return p
	}, "")

	result2 := words.reduceRight(func(prev interface{}, current interface{}) interface{} {
		p := prev.(string)
		p += current.(string)
		return p
	}, "")

	log.Println("displaying forward: ", result)
	log.Println("displaying backward: ", result2)
}

func TestWeekTemps(t *testing.T) {
	log.Println("### Excercise 3")

	randomMonth := [][]int{
		[]int{45, 23, 32, 12, 31, 21, 22},
		[]int{12, 12, 13, 11, 9, 34, 23},
		[]int{33, 34, 23, 25, 26, 12, 11},
		[]int{14, 15, 18, 19, 22, 24, 25},
	}

	thisMonth := NewWeekTemps(randomMonth)

	log.Println("Month Temp avg: ", thisMonth.displayMonthAvg())
	log.Println("Week 2 Temp avg: ", thisMonth.averageWeek(2))
	log.Println("Display All Week Avg")
	thisMonth.displayAllWeekAvg()

}
