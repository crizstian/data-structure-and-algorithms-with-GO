# Arrays

### Manipulating arrays in go

```Go
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
```

### Results

```bash
$ go test -v
=== RUN   TestGrades
2018/10/14 17:50:38 ### Excercise 1
2018/10/14 17:50:38 ## Implementing reduce() func like javascript
The result is:  1524
--- PASS: TestGrades (0.00s)
=== RUN   TestWords
2018/10/14 17:50:38 ### Excercise 2
2018/10/14 17:50:38 displaying forward:  hello my friend
2018/10/14 17:50:38 displaying backward:  friend my hello
--- PASS: TestWords (0.00s)
=== RUN   TestWeekTemps
2018/10/14 17:50:38 ### Excercise 3
2018/10/14 17:50:38 Month Temp avg:  37.5625
2018/10/14 17:50:38 Week 2 Temp avg:  23.428572
2018/10/14 17:50:38 Display All Week Avg
2018/10/14 17:50:38 Week  1  Temp Avg:  26.571428
2018/10/14 17:50:38 Week  2  Temp Avg:  16.285715
2018/10/14 17:50:38 Week  3  Temp Avg:  23.428572
2018/10/14 17:50:38 Week  4  Temp Avg:  19.571428
--- PASS: TestWeekTemps (0.00s)
PASS
ok      github.com/data-structures/array        0.008s
```
