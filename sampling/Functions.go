package sampling

import (
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	return rounder / pow
}

func GetRand(min, max int) int {
	return min + int(rand.Float64()*float64((max-min)+1))
}

func MinOfThree(x1, x2, x3 float64) float64 {
	arr := make([]float64, 0, 0)
	arr = append(arr, x1, x2, x3)
	var min = arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

func MaxOfThree(x1, x2, x3 float64) float64 {
	arr := make([]float64, 0, 0)
	arr = append(arr, x1, x2, x3)
	var max = arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}

// Генерация чисел
func GenerateSeq(max, min, n int) (seq []int) {
	for i := 0; i < n; i++ {
		seq = append(seq, GetRand(min, max))
	}
	return
}

// Сумма сгенерированных чисел
func Sum(seq []int) (sum int) {
	for i := 0; i < len(seq); i++ {
		sum = sum + seq[i]
	}
	return
}

// Среднее число сгенерированных чисел
func Average(seq []int) (aver float64) {
	aver = float64(Sum(seq)) / float64(len(seq))
	return
}

// Поиск дисперсии
func Variance(seq []int) float64 {
	sumDiffsSquared := 0.0
	avg := Average(seq)
	for _, val := range seq {
		diff := float64(val) - avg
		diff *= diff
		sumDiffsSquared += diff
	}
	return sumDiffsSquared / float64(len(seq)-1)
}

// Объединение двух последовательностей
func ConcatForTwo(seq1, seq2 []int) (finalSeq []int) {
	for j := 0; j < 2; j++ {
		for i := 0; i < len(seq1); i++ {
			if j == 0 {
				finalSeq = append(finalSeq, seq1[i])
			} else {
				finalSeq = append(finalSeq, seq2[i])
			}
		}
	}
	return
}

// Объединение трёх последовательностей
func ConcatForThree(seq1, seq2, seq3 []int) (finalSeq []int) {
	for j := 0; j < 3; j++ {
		for i := 0; i < len(seq1); i++ {
			switch {
			case j == 0:
				finalSeq = append(finalSeq, seq1[i])
			case j == 1:
				finalSeq = append(finalSeq, seq2[i])
			default:
				finalSeq = append(finalSeq, seq3[i])
			}
		}
	}
	return
}

func Q1(seq1, seq2, seq3 []int) (q float64) {
	seq := ConcatForThree(seq1, seq2, seq3)
	avg := Average(seq)
	q = float64(len(seq1))*(Average(seq1)-avg)*(Average(seq1)-avg) + float64(len(seq2))*(Average(seq2)-avg)*(Average(seq2)-avg) + float64(len(seq3))*(Average(seq3)-avg)*(Average(seq3)-avg)
	return
}

func Q2(seq1, seq2, seq3 []int) float64 {
	arr := make([][]int, 0, 0)
	arr = append(arr, seq1, seq2, seq3)
	var q = 0.0
	for j := 0; j < 3; j++ {
		seq := arr[j]
		for i := 0; i < len(seq); i++ {
			diff := float64(seq[i]) - Average(seq)
			diff *= diff
			q += diff
		}
	}
	return q
}

func F(seq1, seq2, seq3 []int) float64 {
	q2 := Q2(seq1, seq2, seq3)
	q1 := Q1(seq1, seq2, seq3)
	return (q2 / 2) / (q1 / float64(len(seq1)+len(seq2)+len(seq3)-2))
}

// Значение левой границы
func LeftTboard(seq []int, tcrit float64) float64 {
	return Average(seq) - tcrit*Variance(seq)/math.Sqrt(float64(len(seq)))
}

// Значение правой границы
func RightTboard(seq []int, tcrit float64) float64 {
	return Average(seq) + tcrit*Variance(seq)/math.Sqrt(float64(len(seq)))
}

// Рассчёт t-статистики
func TStatistic(seq []int, a0 float64) float64 {
	return ((Average(seq) - a0) * math.Sqrt(float64(len(seq)))) / math.Sqrt(Variance(seq))
}

// Рассчёт z-статистики
func ZStatistic(seq1, seq2 []int) float64 {
	return (Average(seq1) - Average(seq2)) / math.Sqrt(Variance(seq1)/float64(len(seq1))+Variance(seq2)/float64(len(seq2)))
}

func ZStatistic2(n1, n2, m1, m2 int) float64 {
	p1 := float64(m1) / float64(n1)
	p2 := float64(m2) / float64(n2)
	p := float64(m1 + m2/n1 + n2)
	return p1 - p2/math.Sqrt(p*(1-p)*(float64(1)/float64(n1)+float64(1)/float64(n2)))
}

func Tss(seq1, seq2, seq3 []int) (q float64) {
	seq := ConcatForThree(seq1, seq2, seq3)
	avg := Average(seq)
	for _, val := range seq {
		diff := float64(val) - avg
		diff *= diff
		q += diff
	}
	return
}

func ReturnTask1(count, min, max, n int, alpha float64) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./generated_data/homework_" + timeTask + "/result"
	path2 := "./generated_data/homework_" + timeTask + "/professor_data"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"
		if !Task1(min, max, n, alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask2(count, min, max, n int, alpha float64) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./generated_data/homework_" + timeTask + "/result"
	path2 := "./generated_data/homework_" + timeTask + "/professor_data"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"
		Task2(min, max, n, alpha, pathResults, pathProfData)
		if !Task2(min, max, n, alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask3(count, min, max, n int, alpha float64) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./generated_data/homework_" + timeTask + "/result"
	path2 := "./generated_data/homework_" + timeTask + "/professor_data"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"
		if !Task3(min, max, n, alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask4(count int, alpha float64) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./generated_data/homework_" + timeTask + "/result"
	path2 := "./generated_data/homework_" + timeTask + "/professor_data"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"
		if !Task4(alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask5(count, min, max, n int, alpha float64) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./generated_data/homework_" + timeTask + "/result"
	path2 := "./generated_data/homework_" + timeTask + "/professor_data"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"
		if !Task5(min, max, n, alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask6(count, min, max, n1, n2, n3 int, alpha float64) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./generated_data/homework_" + timeTask + "/result"
	path2 := "./generated_data/homework_" + timeTask + "/professor_data"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"
		if !Task6(min, max, n1, n2, n3, alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}
