package array

import "log"

// ArrIntfce ...
type ArrIntfce interface {
	add(interface{})
	reduce(fn func(values interface{}, value interface{}) interface{}, i interface{}) interface{}
	reduceRight(fn func(values interface{}, value interface{}) interface{}, i interface{}) interface{}
}

// Grades ...
type Grades struct {
	Values []int
}

// Add ...
func (grades Grades) add(grade interface{}) {
	grades.Values = append(grades.Values, grade.(int))
}

// reduce ...
func (grades Grades) reduce(fn func(values interface{}, value interface{}) interface{}, i interface{}) interface{} {
	t := grades.Values[0]
	for i := 1; i < len(grades.Values); i++ {
		t = fn(t, grades.Values[i]).(int)
	}

	return t
}

// reduceRight ...
func (grades Grades) reduceRight(fn func(values interface{}, value interface{}) interface{}, i interface{}) interface{} {
	return nil
}

// NewGrades ...
func NewGrades(arg []int) ArrIntfce {
	return &Grades{
		Values: arg,
	}
}

// Words ...
type Words struct {
	values []string
}

// NewWords ...
func NewWords(arg []string) ArrIntfce {
	return &Words{
		values: arg,
	}
}

// Add ...
func (words Words) add(word interface{}) {
	words.values = append(words.values, word.(string))
}

// reduce ...
func (words Words) reduce(fn func(values interface{}, value interface{}) interface{}, i interface{}) interface{} {
	t := words.values[0]
	for i := 1; i < len(words.values); i++ {
		t = fn(t, words.values[i]).(string)
	}

	return t
}

// reduceRight ...
func (words Words) reduceRight(fn func(values interface{}, value interface{}) interface{}, i interface{}) interface{} {
	t := words.values[len(words.values)-1]
	for i := len(words.values) - 2; i >= 0; i-- {
		t = fn(t, words.values[i]).(string)
	}

	return t
}

// WeekTemps ...
type WeekTemps struct {
	dataStore [][]int
}

// averageWeek ...
func (w WeekTemps) averageWeek(week int) float32 {
	total := 0
	totalDays := len(w.dataStore[week])

	for _, v := range w.dataStore[week] {
		total += v
	}

	return float32(total) / float32(totalDays)
}

// displayMonthAvg ...
func (w WeekTemps) displayMonthAvg() float32 {
	sum := 0

	for _, v := range w.dataStore {
		for _, w := range v {
			sum += w
		}
	}

	return float32(sum) / float32(len(w.dataStore)*4)
}

// displayAllWeekAvg ...
func (w WeekTemps) displayAllWeekAvg() {
	sum := 0

	for i, week := range w.dataStore {
		for _, day := range week {
			sum += day
		}
		log.Println("Week ", i+1, " Temp Avg: ", float32(sum)/float32(len(week)))
		sum = 0
	}
}

// NewWeekTemps ...
func NewWeekTemps(arg [][]int) *WeekTemps {
	return &WeekTemps{
		dataStore: arg,
	}
}
