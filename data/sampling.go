package data

//
//import (
//	"encoding/csv"
//	"fmt"
//	"log"
//	"math"
//	"math/rand"
//	"os"
//	"strconv"
//	"time"
//)
//
//func SomeCode(count int, length int, from int, to int) (res int, err error) {
//	// Создания вариантов в зависимости от вводимого числа N и параметров: n, max, min
//	if _, err := os.Stat("./generated_data"); os.IsNotExist(err) {
//		os.Mkdir("./generated_data", 0755)
//	}
//	t := time.Now()
//	pathHomework := "./generated_data/homework_" + t.Format("Mon Jan _2 15:04:05 2006")
//	os.Mkdir(pathHomework, 0755)
//	for i := 0; i < count; i++ {
//		createTask(i, length, to, from, t.Format("Mon Jan _2 15:04:05 2006"))
//	}
//	res = 1
//	return
//}
//
//// Создание варианта
//func createTask(i, n, max, min int, pathHom string) {
//	number := strconv.Itoa(i + 1)
//	path1 := "./generated_data/homework_" + pathHom + "/result"
//	path2 := "./generated_data/homework_" + pathHom + "/professor_data"
//	os.Mkdir(path1, 0755)
//	os.Mkdir(path2, 0755)
//	pathResults := path1 + "/resultsfile-" + number + ".csv"
//	pathProfData := path2 + "/proffile-" + number + ".csv"
//	outfile1, err := os.Create(pathResults)
//	outfile2, err := os.Create(pathProfData)
//	if err != nil {
//		log.Fatal("Unable to open output")
//	}
//	defer outfile1.Close()
//	defer outfile2.Close()
//
//	writer1 := csv.NewWriter(outfile1)
//	seq := generateSeq(n, max, min)
//	for i := range seq {
//		writer1.Write([]string{strconv.Itoa(seq[i])})
//	}
//
//	writer2 := csv.NewWriter(outfile2)
//	writer2.WriteAll([][]string{
//		{"Average", fmt.Sprintf("%f", average(seq))}, {"Variance", fmt.Sprintf("%f", variance(seq))}, {"Student value", strconv.Itoa(student())}, {"leftTboard", fmt.Sprintf("%f", leftTboard(seq))}, {"rightTboard", fmt.Sprintf("%f", rightTboard(seq))}})
//
//	writer1.Flush()
//	writer2.Flush()
//	if err := writer1.Error(); err != nil {
//		log.Fatalln("error writing csv:", err)
//	}
//}
//
//// Генерация чисел
//func generateSeq(n, max, min int) (seq []int) {
//	//rand.Seed(time.Now().Unix())
//	for i := 0; i < n; i++ {
//		seq = append(seq, min+int(rand.Float64()*float64((max-min)+1)))
//	}
//	return
//}
//
//// Сумма сгенерированных чисел
//func sum(seq []int) (sum int) {
//	for i := 0; i < len(seq); i++ {
//		sum = sum + seq[i]
//	}
//	return
//}
//
//// Среднее число сгенерированных чисел
//func average(seq []int) (aver float64) {
//	aver = float64(sum(seq)) / float64(len(seq))
//	return
//}
//
//// Поиск дисперсии
//func variance(seq []int) float64 {
//	sumDiffsSquared := 0.0
//	avg := average(seq)
//	for idx := range seq {
//		diff := float64(seq[idx]) - avg
//		diff *= diff
//		sumDiffsSquared += diff
//	}
//	return sumDiffsSquared / float64(len(seq)-1)
//}
//
//// Табличное значение из Таблицы Стъюдента
//func student() int {
//	return 1 // табличное значение из БД
//}
//
//// Значение левой границы
//func leftTboard(seq []int) float64 {
//	return variance(seq) - float64(student())*variance(seq)/math.Sqrt(float64(len(seq)))
//}
//
//// Значение правой границы
//func rightTboard(seq []int) float64 {
//	return variance(seq) + float64(student())*variance(seq)/math.Sqrt(float64(len(seq)))
//}
