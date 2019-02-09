package data

import (
	"encoding/csv"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func SomeCode(count int, length int, from int, to int) (res int, err error) {
	// Создания вариантов в зависимости от вводимого числа N и параметров: n, max, min
	os.RemoveAll("./result")
	for i := 1; i < count; i++ {
		createTask(i, length, to, from)
	}
	res = 1
	return
}

// Создание варианта
func createTask(i, n, max, min int) {
	os.Mkdir("./result", 0755)

	path := "./result/resultsfile" + strconv.Itoa(i) + ".csv"
	outfile, err := os.Create(path)
	if err != nil {
		log.Fatal("Unable to open output")
	}
	defer outfile.Close()
	writer := csv.NewWriter(outfile)

	seq := generateSeq(n, max, min)
	for i := range seq {
		writer.Write([]string{strconv.Itoa(seq[i])})
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// Генерация чисел
func generateSeq(n, max, min int) (seq []int) {
	//rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		seq = append(seq, min+int(rand.Float64()*float64((max-min)+1)))
	}
	return
}

// Сумма сгенерированных чисел
func sum(seq []int) (sum int) {
	for i := 0; i < len(seq); i++ {
		sum = sum + seq[i]
	}
	return
}

// Среднее число сгенерированных чисел
func average(seq []int) (aver float64) {
	aver = float64(sum(seq)) / float64(len(seq))
	return
}

// Поиск дисперсии
func variance(seq []int) float64 {
	sumDiffsSquared := 0.0
	avg := average(seq)
	for idx := range seq {
		diff := float64(seq[idx]) - avg
		diff *= diff
		sumDiffsSquared += diff
	}
	return sumDiffsSquared / float64(len(seq)-1)
}

// Табличное значение из Таблицы Стъюдента
func student() int {
	return 1 // табличное значение из БД
}

// Значение левой границы
func leftTboard(seq []int) float64 {
	return variance(seq) - float64(student())*variance(seq)/math.Sqrt(float64(len(seq)))
}

// Значение правой границы
func rightTboard(seq []int) float64 {
	return variance(seq) + float64(student())*variance(seq)/math.Sqrt(float64(len(seq)))
}
