package sampling

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"math"
	"sort"
	"strconv"
)

func Task1(decimalPlaces, n int, expectedValue, stdDeviation float64, pathResults, pathProfData string) bool {
	alpha := ReturnAlpha()
	seq := GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n)
	sort.Float64s(seq)
	cnst := GetRand(int(seq[0]), int(seq[len(seq)-1]))
	m := Average(seq)
	Sigma := Variance(seq)
	TStatisticVal := TStatistic(seq, float64(cnst))
	tcrit1 := GetStudent(1-alpha, n)
	tcrit2 := GetStudent(1-alpha/2, n) // Аккуратнее смотреть какую альфу передаём

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Задание")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	sheet1.SetColWidth(0, 0, 15)

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt1.Value = "Константа const"

	rowSt2 := sheet1.AddRow()
	cellSt2 := rowSt2.AddCell()
	cellSt2.Value = strconv.Itoa(cnst)

	rowSt3 := sheet1.AddRow()
	cellSt3 := rowSt3.AddCell()
	cellSt3.Value = "Выборка:"

	for j := range seq {
		rowSt4 := sheet1.AddRow()
		cellSt4 := rowSt4.AddCell()
		cellSt4.Value = strconv.FormatFloat(seq[j], 'f', -2, 64)
	}

	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Ответы")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}

	sheet2.SetColWidth(0, 0, 25)

	header := []string{"Выборочное среднее", "Выборочная дисперсия", "T-статистика",
		"t критическое порядка 1-a", "t критическое порядка 1-a/2", "Альфа"}
	result := []string{strconv.FormatFloat(m, 'f', -2, 64),
		strconv.FormatFloat(Sigma, 'f', -2, 64),
		strconv.FormatFloat(TStatisticVal, 'f', -2, 64),
		strconv.FormatFloat(tcrit1, 'f', -2, 64),
		strconv.FormatFloat(tcrit2, 'f', -2, 64),
		strconv.FormatFloat(alpha, 'f', -2, 64)}
	for h := range header {
		rowPr1 := sheet2.AddRow()
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
		cellPr2 := rowPr1.AddCell()
		cellPr2.Value = result[h]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task2(decimalPlaces, n int, expectedValue, stdDeviation float64, pathResults, pathProfData string) bool {
	alpha := ReturnAlpha()
	seq := GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n)
	m := Average(seq)
	Sigma := Variance(seq)
	tcrit := GetStudent(1-alpha/2, n)
	tcrit1 := tcrit
	leftBoard := LeftTboard(seq, tcrit)
	rightBoard := RightTboard(seq, tcrit)

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Задание")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt1.Value = "Выборка:"

	for j := range seq {
		rowSt2 := sheet1.AddRow()
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = strconv.FormatFloat(seq[j], 'f', -2, 64)
	}

	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Ответы")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}

	sheet2.SetColWidth(0, 0, 27)

	header := []string{"Выборочное среднее", "Дисперсия",
		"t критическое порядка 1-a/2", "Левая граница дов. интервала",
		"Правая граница дов. интервала", "Альфа"}
	result := []string{strconv.FormatFloat(m, 'f', -2, 64),
		strconv.FormatFloat(Sigma, 'f', -2, 64),
		strconv.FormatFloat(tcrit1, 'f', -2, 64),
		strconv.FormatFloat(leftBoard, 'f', -2, 64),
		strconv.FormatFloat(rightBoard, 'f', -2, 64),
		strconv.FormatFloat(alpha, 'f', -2, 64)}
	for h := range header {
		rowPr1 := sheet2.AddRow()
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
		cellPr2 := rowPr1.AddCell()
		cellPr2.Value = result[h]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task3(decimalPlaces, n int, expectedValue, stdDeviation float64, pathResults, pathProfData string) bool {
	alpha := ReturnAlpha()
	MAGIC_COEF := 1.05
	seq1 := GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n)
	seq2 := GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n)
	m1 := Average(seq1)
	m2 := Average(seq2)
	Sigma1 := Round(math.Sqrt(Variance(seq1))*MAGIC_COEF, 2)
	Sigma2 := Round(math.Sqrt(Variance(seq2))*MAGIC_COEF, 2)
	ZStatisticVal := ZStatistic(seq1, seq2)
	ucrit1 := GetNorm(1 - alpha)
	ucrit2 := GetNorm(1-alpha) / 2

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Задание")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	sheet1.SetColWidth(0, 1, 15)

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt2 := rowSt1.AddCell()
	cellSt1.Value = "Первая выборка:"
	cellSt2.Value = "Вторая выборка:"

	for j := range seq1 {
		rowSt2 := sheet1.AddRow()
		cellSt3 := rowSt2.AddCell()
		cellSt4 := rowSt2.AddCell()
		cellSt3.Value = strconv.FormatFloat(seq1[j], 'f', -2, 64)
		cellSt4.Value = strconv.FormatFloat(seq2[j], 'f', -2, 64)
	}

	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Ответы")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}

	sheet2.SetColWidth(0, 0, 35)

	header := []string{"Выборочное среднее первой выборки", "Выборочное среднее второй выборки",
		"Стандартное отклонение первой выборки", "Стандартное отклонение второй выборки", "Z статистика",
		"u критическое порядка 1-a", "u критическое порядка 1-a/2", "Альфа"}
	result := []string{strconv.FormatFloat(m1, 'f', -2, 64),
		strconv.FormatFloat(m2, 'f', -2, 64),
		strconv.FormatFloat(Sigma1, 'f', -2, 64),
		strconv.FormatFloat(Sigma2, 'f', -2, 64),
		strconv.FormatFloat(ZStatisticVal, 'f', -2, 64),
		strconv.FormatFloat(ucrit1, 'f', -2, 64),
		strconv.FormatFloat(ucrit2, 'f', -2, 64),
		strconv.FormatFloat(alpha, 'f', -2, 64)}
	for h := range header {
		rowPr1 := sheet2.AddRow()
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
		cellPr2 := rowPr1.AddCell()
		cellPr2.Value = result[h]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task4(pathResults, pathProfData string) bool {
	alpha := ReturnAlpha()
	n1 := GetRand(100, 1500)
	n2 := GetRand(100, 1500)
	m1 := GetRand(int(0.25*float64(n1)), int(0.85*float64(n1)))
	m2 := GetRand(int(0.25*float64(n2)), int(0.85*float64(n2)))
	p1 := Round(float64(m1)/float64(n1), 2)
	p2 := Round(float64(m2)/float64(n2), 2)
	ZStatisticVal := ZStatistic2(n1, n2, m1, m2, p1, p2)
	ucrit1 := GetNorm(1 - alpha)
	ucrit2 := GetNorm(1 - alpha/2)

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Задание")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	sheet1.SetColWidth(0, 0, 32)

	header1 := []string{"Количество элементов первой выборки", "Число из первой группы",
		"Количество элементов второй выборки", "Число из второй группы"}
	result1 := []string{strconv.Itoa(n1),
		strconv.Itoa(m1),
		strconv.Itoa(n2),
		strconv.Itoa(m2)}
	for h1 := range header1 {
		rowSt1 := sheet1.AddRow()
		cellSt1 := rowSt1.AddCell()
		cellSt1.Value = header1[h1]
		cellSt2 := rowSt1.AddCell()
		cellSt2.Value = result1[h1]
	}

	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Ответы")

	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}

	sheet2.SetColWidth(0, 0, 27)

	header2 := []string{"Выборочная доля первой выборки", "Выборочная доля второй выборки",
		"Z статистика", "u критическое порядка 1-a", "u критическое порядка 1-a/2", "Альфа"}
	result2 := []string{strconv.FormatFloat(p1, 'f', -2, 64),
		strconv.FormatFloat(p2, 'f', -2, 64),
		strconv.FormatFloat(ZStatisticVal, 'f', -2, 64),
		strconv.FormatFloat(ucrit1, 'f', -2, 64),
		strconv.FormatFloat(ucrit2, 'f', -2, 64),
		strconv.FormatFloat(alpha, 'f', -2, 64)}
	for h2 := range header2 {
		rowPr1 := sheet2.AddRow()
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header2[h2]
		cellPr2 := rowPr1.AddCell()
		cellPr2.Value = result2[h2]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task5(decimalPlaces, n int, expectedValue, stdDeviation float64, pathResults, pathProfData string) bool {
	alpha := ReturnAlpha()
	seq1 := GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n)
	seq2 := GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n)
	Sigma1 := Variance(seq1)
	Sigma2 := Variance(seq2)

	var f float64
	if Variance(seq1) > Variance(seq2) {
		f = Variance(seq1) / Variance(seq2)
	} else {
		f = Variance(seq2) / Variance(seq1)
	}

	F := Round(f, 2)
	fcrit1left := GetFisherLeft(alpha, n)
	fcrit1right := GetFisherRight(alpha/2, n)
	fcrit2left := GetFisherLeft(alpha/2, n)
	fcrit2right := GetFisherRight(alpha/2, n)
	fcrit3left := GetFisherLeft(1-alpha, n)
	fcrit3right := GetFisherRight(1-alpha, n)
	fcrit4left := GetFisherLeft(1-alpha/2, n)
	fcrit4right := GetFisherRight(1-alpha/2, n)

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Задание")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	sheet1.SetColWidth(0, 1, 15)

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt2 := rowSt1.AddCell()
	cellSt1.Value = "Первая выборка:"
	cellSt2.Value = "Вторая выборка:"

	for j := range seq1 {
		rowSt2 := sheet1.AddRow()
		cellSt3 := rowSt2.AddCell()
		cellSt4 := rowSt2.AddCell()
		cellSt3.Value = strconv.FormatFloat(seq1[j], 'f', -2, 64)
		cellSt4.Value = strconv.FormatFloat(seq2[j], 'f', -2, 64)
	}

	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Ответы")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}

	sheet2.SetColWidth(0, 0, 35)

	header := []string{"Выборочная дисперсия первой подвыборки", "Выборочная дисперсия второй подвыборки",
		"F статистика", "левое f критическое порядка a", "правое f критическое порядка a",
		"левое f критическое порядка a/2", "правое f критическое порядка a/2",
		"левое f критическое порядка 1-a", "правое f критическое порядка 1-a",
		"левое f критическое порядка 1-a/2", "правое f критическое порядка 1-a/2", "Альфа"}
	result := []string{strconv.FormatFloat(Sigma1, 'f', -2, 64),
		strconv.FormatFloat(Sigma2, 'f', -2, 64),
		strconv.FormatFloat(F, 'f', -2, 64),
		strconv.FormatFloat(fcrit1left, 'f', -2, 64),
		strconv.FormatFloat(fcrit1right, 'f', -2, 64),
		strconv.FormatFloat(fcrit2left, 'f', -2, 64),
		strconv.FormatFloat(fcrit2right, 'f', -2, 64),
		strconv.FormatFloat(fcrit3left, 'f', -2, 64),
		strconv.FormatFloat(fcrit3right, 'f', -2, 64),
		strconv.FormatFloat(fcrit4left, 'f', -2, 64),
		strconv.FormatFloat(fcrit4right, 'f', -2, 64),
		strconv.FormatFloat(alpha, 'f', -2, 64)}
	for h := range header {
		rowPr1 := sheet2.AddRow()
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
		cellPr2 := rowPr1.AddCell()
		cellPr2.Value = result[h]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task6(decimalPlaces, n1, n2, n3 int, expectedValue, stdDeviation float64, pathResults, pathProfData string) bool {
	alpha := ReturnAlpha()
	var seq1, seq2, seq3 []float64
	var stopCounter = 0
	var quit = false
	for ok := true; ok; ok = (quit || stopCounter >= 100000) {
		seq1 = GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n1)
		seq2 = GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n2)
		seq3 = GenerateSeq(expectedValue, stdDeviation, decimalPlaces, n3)
		quit = MaxOfThree(Average(seq1), Average(seq2), Average(seq3))/MinOfThree(Average(seq1), Average(seq2), Average(seq3)) <= 1.4
		stopCounter++
	}

	X1 := Average(seq1)
	X2 := Average(seq2)
	X3 := Average(seq3)
	Xc := Average(ConcatForThree(seq1, seq2, seq3))
	TSS := Tss(seq1, seq2, seq3)
	Q1 := Q1(seq1, seq2, seq3)
	Q2 := Q2(seq1, seq2, seq3)
	F := F(seq1, seq2, seq3)
	fcrit1left := GetFisherLeft(alpha, (n1+n2+n3-3)/3)
	fcrit1right := GetFisherRight(alpha, (n1+n2+n3-3)/3)

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Задание")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	sheet1.SetColWidth(0, 0, 15)

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt1.Value = "Первая выборка:"

	for s1 := range seq1 {
		rowSt2 := sheet1.AddRow()
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = strconv.FormatFloat(seq1[s1], 'f', -2, 64)
	}

	rowSt3 := sheet1.AddRow()
	rowSt3 = sheet1.AddRow()
	cellSt3 := rowSt3.AddCell()
	cellSt3.Value = "Вторая выборка:"

	for s2 := range seq2 {
		rowSt4 := sheet1.AddRow()
		cellSt4 := rowSt4.AddCell()
		cellSt4.Value = strconv.FormatFloat(seq2[s2], 'f', -2, 64)
	}

	rowSt5 := sheet1.AddRow()
	rowSt5 = sheet1.AddRow()
	cellSt5 := rowSt5.AddCell()
	cellSt5.Value = "Третья выборка:"

	for s3 := range seq3 {
		rowSt6 := sheet1.AddRow()
		cellSt6 := rowSt6.AddCell()
		cellSt6.Value = strconv.FormatFloat(seq3[s3], 'f', -2, 64)
	}

	err1 = fileSt.Save(pathResults)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	// Для преподавателя
	filePr := xlsx.NewFile()
	sheet2, err2 := filePr.AddSheet("Ответы")
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}

	sheet2.SetColWidth(0, 0, 35)

	header := []string{"Выборочное среднее первой подвыборки", "Выборочное среднее второй подвыборки",
		"Выборочное среднее третьей подвыборки", "Среднее всей выборки", "TSS", "Межгрупповой размах", "Внутригрупповой размах",
		"F статистика", "левое f критическое", "правое f критическое", "Альфа"}
	result := []string{strconv.FormatFloat(X1, 'f', -2, 64),
		strconv.FormatFloat(X2, 'f', -2, 64),
		strconv.FormatFloat(X3, 'f', -2, 64),
		strconv.FormatFloat(Xc, 'f', -2, 64),
		strconv.FormatFloat(TSS, 'f', -2, 64),
		strconv.FormatFloat(Q1, 'f', -2, 64),
		strconv.FormatFloat(Q2, 'f', -2, 64),
		strconv.FormatFloat(F, 'f', -2, 64),
		strconv.FormatFloat(fcrit1left, 'f', -2, 64),
		strconv.FormatFloat(fcrit1right, 'f', -2, 64),
		strconv.FormatFloat(alpha, 'f', -2, 64)}
	for h := range header {
		rowPr1 := sheet2.AddRow()
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
		cellPr2 := rowPr1.AddCell()
		cellPr2.Value = result[h]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}
