package sampling

import (
	"fmt"
	"github.com/tealeg/xlsx"
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
	for i := 0; i <= n; i++ {
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
		for i := 0; i <= len(seq1); i++ {
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
		for i := 0; i <= len(seq1); i++ {
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
		for i := 0; i <= len(seq); i++ {
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

func ReturnTask(count int, task map[string][]int, answer map[string]float64, compare map[string]bool, taskID int) bool {
	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
		os.Mkdir("./generated_data", 0755)
	}

	t := time.Now()
	timeTask := t.Format("Mon Jan _2 15:04:05 2006")
	pathHomework := "./generated_data/homework_" + timeTask
	os.Mkdir(pathHomework, 0755)

	for i := 0; i < count; i++ {
		path1 := "./generated_data/homework_" + timeTask + "/result"
		path2 := "./generated_data/homework_" + timeTask + "/professor_data"
		os.Mkdir(path1, 0755)
		os.Mkdir(path2, 0755)

		number := strconv.Itoa(i + 1)
		pathResults := path1 + "/resultsfile-" + number + ".xlsx"
		pathProfData := path2 + "/proffile-" + number + ".xlsx"

		if taskID == 1 {
			return task1XLSX(pathResults, pathProfData, task, answer, compare)
		}
		if taskID == 2 {
			return task2XLSX(pathResults, pathProfData, task, answer)
		}
		if taskID == 3 {
			return task3XLSX(pathResults, pathProfData, task, answer)
		}
		if taskID == 4 {
			return task4XLSX(pathResults, pathProfData, answer)
		}
		if taskID == 5 {
			return task5XLSX(pathResults, pathProfData, task, answer)
		}
		if taskID == 6 {
			return task6XLSX(pathResults, pathProfData, task, answer)
		}
	}
	return false
}

func task1XLSX(pathResults, pathProfData string, task map[string][]int, answer map[string]float64, compare map[string]bool) bool {
	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}
	seq := task["seq"]

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt1.Value = fmt.Sprintf("%f", answer["cnst"])

	rowSt2 := sheet1.AddRow()
	rowSt2.WriteSlice(seq, len(seq))
	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Sheet1")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	rowPr1 := sheet2.AddRow()
	rowPr1.WriteSlice([]string{"Sample mean", "Sample variance", "tStatistic", "tcrit1", "tcrit2", "Hypothesis true?:"}, 6)

	rowPr2 := sheet2.AddRow()
	rowPr2.WriteSlice([]string{fmt.Sprintf("%f", answer["m"]), fmt.Sprintf("%f", answer["Sigma"]), fmt.Sprintf("%f", answer["TStatistic"]), fmt.Sprintf("%f", answer["tcrit1"]), fmt.Sprintf("%f", answer["tcrit2"]), strconv.FormatBool(compare["answer"])}, 6)

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func task2XLSX(pathResults, pathProfData string, task map[string][]int, answer map[string]float64) bool {
	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}
	seq := task["seq"]

	rowSt1 := sheet1.AddRow()
	rowSt1.WriteSlice(seq, len(seq))
	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Sheet1")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	rowPr1 := sheet2.AddRow()
	rowPr1.WriteSlice([]string{"Sample mean", "Variance", "tcrit1", "leftBoard", "rightBoard"}, 5)

	rowPr2 := sheet2.AddRow()
	rowPr2.WriteSlice([]string{fmt.Sprintf("%f", answer["m"]), fmt.Sprintf("%f", answer["Sigma"]), fmt.Sprintf("%f", answer["tcrit1"]), fmt.Sprintf("%f", answer["leftBoard"]), fmt.Sprintf("%f", answer["rightBoard"])}, 5)

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func task3XLSX(pathResults, pathProfData string, task map[string][]int, answer map[string]float64) bool {
	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}
	seq1 := task["seq1"]
	seq2 := task["seq2"]

	rowSt1 := sheet1.AddRow()
	rowSt1.WriteSlice(seq1, len(seq1))

	rowSt2 := sheet1.AddRow()
	rowSt2.WriteSlice(seq2, len(seq2))
	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Sheet1")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	rowPr1 := sheet2.AddRow()
	rowPr1.WriteSlice([]string{"Sample mean1", "Sample mean2", "Sample variance1", "Sample variance2", "ZStatistic", "ucrit1", "ucrit2"}, 7)

	rowPr2 := sheet2.AddRow()
	rowPr2.WriteSlice([]string{fmt.Sprintf("%f", answer["m1"]), fmt.Sprintf("%f", answer["m2"]), fmt.Sprintf("%f", answer["Sigma1"]), fmt.Sprintf("%f", answer["Sigma2"]), fmt.Sprintf("%f", answer["ZStatistic"]), fmt.Sprintf("%f", answer["ucrit1"]), fmt.Sprintf("%f", answer["ucrit2"])}, 7)

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func task4XLSX(pathResults, pathProfData string, answer map[string]float64) bool {
	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	rowSt1.WriteSlice([]string{"n1", "M1", "n2", "M2"}, 4)

	rowSt2 := sheet1.AddRow()
	rowSt2.WriteSlice([]string{fmt.Sprintf("%f", answer["n1"]), fmt.Sprintf("%f", answer["m1"]), fmt.Sprintf("%f", answer["n2"]), fmt.Sprintf("%f", answer["m2"])}, 4)
	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Sheet1")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	rowPr1 := sheet2.AddRow()
	rowPr1.WriteSlice([]string{"p1", "p2", "ZStatistic", "ucrit1", "ucrit2"}, 5)

	rowPr2 := sheet2.AddRow()
	rowPr2.WriteSlice([]string{fmt.Sprintf("%f", answer["p1"]), fmt.Sprintf("%f", answer["p2"]), fmt.Sprintf("%f", answer["ZStatistic"]), fmt.Sprintf("%f", answer["ucrit1"]), fmt.Sprintf("%f", answer["ucrit2"])}, 5)

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func task5XLSX(pathResults, pathProfData string, task map[string][]int, answer map[string]float64) bool {
	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}
	seq1 := task["seq1"]
	seq2 := task["seq2"]

	rowSt1 := sheet1.AddRow()
	rowSt1.WriteSlice(seq1, len(seq1))

	rowSt2 := sheet1.AddRow()
	rowSt2.WriteSlice(seq2, len(seq2))
	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Sheet1")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	rowPr1 := sheet2.AddRow()
	rowPr1.WriteSlice([]string{"Sigma1", "Sigma2", "FStatistic", "fcrit1left", "fcrit1right", "fcrit2left", "fcrit2right", "fcrit3left", "fcrit3right", "fcrit4left", "fcrit4right"}, 11)

	rowPr2 := sheet2.AddRow()
	rowPr2.WriteSlice([]string{fmt.Sprintf("%f", answer["Sigma1"]), fmt.Sprintf("%f", answer["Sigma2"]), fmt.Sprintf("%f", answer["F"]), fmt.Sprintf("%f", answer["fcrit1left"]), fmt.Sprintf("%f", answer["fcrit1right"]), fmt.Sprintf("%f", answer["fcrit2left"]), fmt.Sprintf("%f", answer["fcrit2right"]), fmt.Sprintf("%f", answer["fcrit3left"]), fmt.Sprintf("%f", answer["fcrit3right"]), fmt.Sprintf("%f", answer["fcrit4left"]), fmt.Sprintf("%f", answer["fcrit4right"])}, 11)

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func task6XLSX(pathResults, pathProfData string, task map[string][]int, answer map[string]float64) bool {
	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}
	seq1 := task["seq1"]
	seq2 := task["seq2"]
	seq3 := task["seq3"]

	rowSt1 := sheet1.AddRow()
	rowSt1.WriteSlice(seq1, len(seq1))

	rowSt2 := sheet1.AddRow()
	rowSt2.WriteSlice(seq2, len(seq2))

	rowSt3 := sheet1.AddRow()
	rowSt3.WriteSlice(seq3, len(seq3))
	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Sheet1")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	rowPr1 := sheet2.AddRow()
	rowPr1.WriteSlice([]string{"X1", "X2", "X3", "Xc", "TSS", "Q1", "Q2", "F", "fcrit1left", "fcrit1right"}, 10)

	rowPr2 := sheet2.AddRow()
	rowPr2.WriteSlice([]string{fmt.Sprintf("%f", answer["X1"]), fmt.Sprintf("%f", answer["X2"]), fmt.Sprintf("%f", answer["X3"]), fmt.Sprintf("%f", answer["Xc"]), fmt.Sprintf("%f", answer["TSS"]), fmt.Sprintf("%f", answer["Q1"]), fmt.Sprintf("%f", answer["Q2"]), fmt.Sprintf("%f", answer["F"]), fmt.Sprintf("%f", answer["fcrit1left"]), fmt.Sprintf("%f", answer["fcrit1right"])}, 10)

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}
