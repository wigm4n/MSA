package sampling

import (
	"MSA/data"
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

// Генерация чисел
func GenerateSeq(expectedValue, stdDeviation float64, decimalPlaces, n int) (seq []float64) {
	for i := 0; i < n; i++ {
		seq = append(seq, getNormVal(expectedValue, stdDeviation, decimalPlaces))
	}
	return
}

func getNormVal(expectedValue, stdDeviation float64, decimalPlaces int) float64 {
	val := Round(rand.NormFloat64()*stdDeviation+expectedValue, decimalPlaces)
	if expectedValue >= 0 && val < 0 {
		for ok := true; ok; ok = (expectedValue >= 0 && val < 0) {
			val = Round(rand.NormFloat64()*stdDeviation+expectedValue, decimalPlaces)
		}
	}
	return val
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

// Сумма сгенерированных чисел
func Sum(seq []float64) (sum float64) {
	for i := 0; i < len(seq); i++ {
		sum = sum + seq[i]
	}
	return
}

// Среднее число сгенерированных чисел
func Average(seq []float64) (aver float64) {
	aver = Round(float64(Sum(seq))/float64(len(seq)), 2)
	return
}

// Поиск дисперсии
func Variance(seq []float64) float64 {
	sumDiffsSquared := 0.0
	avg := Average(seq)
	for _, val := range seq {
		diff := float64(val) - avg
		diff *= diff
		sumDiffsSquared += diff
	}
	return Round(sumDiffsSquared/float64(len(seq)-1), 2)
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
func ConcatForThree(seq1, seq2, seq3 []float64) (finalSeq []float64) {
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

func Q1(seq1, seq2, seq3 []float64) (q float64) {
	seq := ConcatForThree(seq1, seq2, seq3)
	avg := Average(seq)
	q = float64(len(seq1))*(Average(seq1)-avg)*(Average(seq1)-avg) + float64(len(seq2))*(Average(seq2)-avg)*(Average(seq2)-avg) + float64(len(seq3))*(Average(seq3)-avg)*(Average(seq3)-avg)
	return Round(q, 2)
}

func Q2(seq1, seq2, seq3 []float64) float64 {
	arr := make([][]float64, 0, 0)
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
	return Round(q, 2)
}

func F(seq1, seq2, seq3 []float64) float64 {
	q2 := Q2(seq1, seq2, seq3)
	q1 := Q1(seq1, seq2, seq3)
	return Round((q2/2)/(q1/float64(len(seq1)+len(seq2)+len(seq3)-2)), 2)
}

// Значение левой границы
func LeftTboard(seq []float64, tcrit float64) float64 {
	return Average(seq) - tcrit*Variance(seq)/math.Sqrt(float64(len(seq)))
}

// Значение правой границы
func RightTboard(seq []float64, tcrit float64) float64 {
	return Average(seq) + tcrit*Variance(seq)/math.Sqrt(float64(len(seq)))
}

// Рассчёт t-статистики
func TStatistic(seq []float64, a0 float64) float64 {
	return Round(((Average(seq)-a0)*math.Sqrt(float64(len(seq))))/math.Sqrt(Variance(seq)), 2)
}

// Рассчёт z-статистики
func ZStatistic(seq1, seq2 []float64) float64 {
	return (Average(seq1) - Average(seq2)) / math.Sqrt(Variance(seq1)/float64(len(seq1))+Variance(seq2)/float64(len(seq2)))
}

func ZStatistic2(n1, n2, m1, m2 int) float64 {
	p1 := float64(m1) / float64(n1)
	p2 := float64(m2) / float64(n2)
	p := float64(m1 + m2/n1 + n2)
	return p1 - p2/math.Sqrt(p*(1-p)*(float64(1)/float64(n1)+float64(1)/float64(n2)))
}

func Tss(seq1, seq2, seq3 []float64) (q float64) {
	seq := ConcatForThree(seq1, seq2, seq3)
	avg := Average(seq)
	for _, val := range seq {
		diff := float64(val) - avg
		diff *= diff
		q += diff
	}
	return Round(q, 2)
}

func ReturnTask1(taskExtended data.TaskExtended) bool {
	if _, err := os.Stat("./Homeworks"); os.IsNotExist(err) {
		os.Mkdir("./Homeworks", 0755)
	}

	t := time.Now()
	timeTask := t.Format("(02-Jan-2006,15:04)")
	pathHomework := "./Homeworks/Homework-1_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./Homeworks/Homework-1_" + timeTask + "/Tasks"
	path2 := "./Homeworks/Homework-1_" + timeTask + "/Answers"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < taskExtended.Count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/Task-" + number + ".xlsx"
		pathProfData := path2 + "/Answer-" + number + ".xlsx"
		if !Task1(taskExtended.DecimalPlaces, taskExtended.Size, taskExtended.ExpectedValue, taskExtended.StdDeviation, taskExtended.Alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask2(taskExtended data.TaskExtended) bool {
	if _, err := os.Stat("./Homeworks"); os.IsNotExist(err) {
		os.Mkdir("./Homeworks", 0755)
	}

	t := time.Now()
	timeTask := t.Format("(02-Jan-2006,15:04)")
	pathHomework := "./Homeworks/Homework-2_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./Homeworks/Homework-2_" + timeTask + "/Tasks"
	path2 := "./Homeworks/Homework-2_" + timeTask + "/Answers"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < taskExtended.Count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/Task-" + number + ".xlsx"
		pathProfData := path2 + "/Answer-" + number + ".xlsx"
		if !Task2(taskExtended.DecimalPlaces, taskExtended.Size, taskExtended.ExpectedValue, taskExtended.StdDeviation, taskExtended.Alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask3(taskExtended data.TaskExtended) bool {
	if _, err := os.Stat("./Homeworks"); os.IsNotExist(err) {
		os.Mkdir("./Homeworks", 0755)
	}

	t := time.Now()
	timeTask := t.Format("(02-Jan-2006,15:04)")
	pathHomework := "./Homeworks/Homework-3_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./Homeworks/Homework-3_" + timeTask + "/Tasks"
	path2 := "./Homeworks/Homework-3_" + timeTask + "/Answers"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < taskExtended.Count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/Task-" + number + ".xlsx"
		pathProfData := path2 + "/Answer-" + number + ".xlsx"
		if !Task3(taskExtended.DecimalPlaces, taskExtended.Size, taskExtended.ExpectedValue, taskExtended.StdDeviation, taskExtended.Alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask4(taskExtended data.TaskExtended) bool {
	if _, err := os.Stat("./Homeworks"); os.IsNotExist(err) {
		os.Mkdir("./Homeworks", 0755)
	}

	t := time.Now()
	timeTask := t.Format("(02-Jan-2006,15:04)")
	pathHomework := "./Homeworks/Homework-4_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./Homeworks/Homework-4_" + timeTask + "/Tasks"
	path2 := "./Homeworks/Homework-4_" + timeTask + "/Answers"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < taskExtended.Count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/Task-" + number + ".xlsx"
		pathProfData := path2 + "/Answer-" + number + ".xlsx"
		if !Task4(taskExtended.Alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask5(taskExtended data.TaskExtended) bool {
	if _, err := os.Stat("./Homeworks"); os.IsNotExist(err) {
		os.Mkdir("./Homeworks", 0755)
	}

	t := time.Now()
	timeTask := t.Format("(02-Jan-2006,15:04)")
	pathHomework := "./Homeworks/Homework-5_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./Homeworks/Homework-5_" + timeTask + "/Tasks"
	path2 := "./Homeworks/Homework-5_" + timeTask + "/Answers"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < taskExtended.Count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/Task-" + number + ".xlsx"
		pathProfData := path2 + "/Answer-" + number + ".xlsx"
		if !Task5(taskExtended.DecimalPlaces, taskExtended.Size, taskExtended.ExpectedValue, taskExtended.StdDeviation, taskExtended.Alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}

func ReturnTask6(taskExtended data.TaskExtended) bool {
	if _, err := os.Stat("./Homeworks"); os.IsNotExist(err) {
		os.Mkdir("./Homeworks", 0755)
	}

	t := time.Now()
	timeTask := t.Format("(02-Jan-2006,15:04)")
	pathHomework := "./Homeworks/Homework-6_" + timeTask
	os.Mkdir(pathHomework, 0755)

	path1 := "./Homeworks/Homework-6_" + timeTask + "/Tasks"
	path2 := "./Homeworks/Homework-6_" + timeTask + "/Answers"
	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	for i := 0; i < taskExtended.Count; i++ {
		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/Task-" + number + ".xlsx"
		pathProfData := path2 + "/Answer-" + number + ".xlsx"
		if !Task6(taskExtended.DecimalPlaces, taskExtended.Size, taskExtended.Size2,
			taskExtended.Size3, taskExtended.ExpectedValue, taskExtended.StdDeviation,
			taskExtended.Alpha, pathResults, pathProfData) {
			return false
		}
	}
	return true
}
