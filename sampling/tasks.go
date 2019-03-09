package sampling

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"math"
	"strconv"
)

func Task1(min, max, n int, alpha float64, pathResults, pathProfData string) bool {
	seq := GenerateSeq(max, min, n)
	cnst := GetRand(min, max)
	m := Average(seq)
	Sigma := Variance(seq)
	TStatisticVal := TStatistic(seq, float64(cnst))
	tcrit1 := GetStudent(1-alpha, n)
	tcrit2 := GetStudent(1-alpha/2, n) // Аккуратнее смотреть какую альфу передаём
	compare := (GetStudent(1-alpha, n) < TStatistic(seq, float64(cnst))) && (GetStudent(1-alpha/2, n) > TStatistic(seq, float64(cnst)))

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	cellSt1 := rowSt1.AddCell()
	cellSt1.Value = strconv.Itoa(cnst)

	rowSt2 := sheet1.AddRow()
	for j := range seq {
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = strconv.Itoa(seq[j])
	}

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
	header := []string{"Sample mean", "Sample variance", "tStatistic", "tcrit1", "tcrit2", "Hypothesis true?:"}
	for h := range header {
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
	}

	rowPr2 := sheet2.AddRow()
	result := []string{fmt.Sprintf("%f", m),
		fmt.Sprintf("%f", Sigma),
		fmt.Sprintf("%f", TStatisticVal),
		fmt.Sprintf("%f", tcrit1),
		fmt.Sprintf("%f", tcrit2),
		strconv.FormatBool(compare)}
	for r := range result {
		cellPr2 := rowPr2.AddCell()
		cellPr2.Value = result[r]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task2(min, max, n int, alpha float64, pathResults, pathProfData string) bool {
	seq := GenerateSeq(max, min, n)
	cnst := GetRand(min, max)
	m := Average(seq)
	Sigma := Variance(seq)
	tcrit := GetStudent(1-alpha/2, n)
	tcrit1 := tcrit
	leftBoard := LeftTboard(seq, tcrit)
	rightBoard := RightTboard(seq, tcrit)
	compare := (LeftTboard(seq, tcrit) < TStatistic(seq, float64(cnst))) && (RightTboard(seq, tcrit) > TStatistic(seq, float64(cnst)))

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	for j := range seq {
		cellSt1 := rowSt1.AddCell()
		cellSt1.Value = strconv.Itoa(seq[j])
	}

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
	header := []string{"Sample mean", "Variance", "tcrit1", "leftBoard", "rightBoard", "Hypothesis true?:"}
	for h := range header {
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
	}

	rowPr2 := sheet2.AddRow()
	result := []string{fmt.Sprintf("%f", m),
		fmt.Sprintf("%f", Sigma),
		fmt.Sprintf("%f", tcrit1),
		fmt.Sprintf("%f", leftBoard),
		fmt.Sprintf("%f", rightBoard),
		strconv.FormatBool(compare)}
	for r := range result {
		cellPr2 := rowPr2.AddCell()
		cellPr2.Value = result[r]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task3(min, max, n int, alpha float64, pathResults, pathProfData string) bool {
	MAGIC_COEF := 1.05
	seq1 := GenerateSeq(max, min, n)
	seq2 := GenerateSeq(max, min, n)
	m1 := Average(seq1)
	m2 := Average(seq2)
	Sigma1 := math.Sqrt(Variance(seq1)) * MAGIC_COEF
	Sigma2 := math.Sqrt(Variance(seq2)) * MAGIC_COEF
	ZStatisticVal := ZStatistic(seq1, seq2)
	ucrit1 := GetNorm(1 - alpha)
	ucrit2 := GetNorm(1-alpha) / 2

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	for j := range seq1 {
		cellSt1 := rowSt1.AddCell()
		cellSt1.Value = strconv.Itoa(seq1[j])
	}

	rowSt2 := sheet1.AddRow()
	for k := range seq2 {
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = strconv.Itoa(seq2[k])
	}

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
	header := []string{"Sample mean1", "Sample mean2", "Sample variance1", "Sample variance2", "ZStatistic", "ucrit1", "ucrit2"}
	for h := range header {
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
	}

	rowPr2 := sheet2.AddRow()
	result := []string{fmt.Sprintf("%f", m1),
		fmt.Sprintf("%f", m2),
		fmt.Sprintf("%f", Sigma1),
		fmt.Sprintf("%f", Sigma2),
		fmt.Sprintf("%f", ZStatisticVal),
		fmt.Sprintf("%f", ucrit1),
		fmt.Sprintf("%f", ucrit2)}
	for r := range result {
		cellPr2 := rowPr2.AddCell()
		cellPr2.Value = result[r]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task4(alpha float64, pathResults, pathProfData string) bool {
	n1 := GetRand(100, 1500)
	n2 := GetRand(100, 1500)
	m1 := GetRand(int(0.25*float64(n1)), int(0.85*float64(n1)))
	m2 := GetRand(int(0.25*float64(n2)), int(0.85*float64(n2)))
	p1 := float64(m1) / float64(n1)
	p2 := float64(m2) / float64(n2)
	ZStatisticVal := ZStatistic2(n1, n2, m1, m2)
	ucrit1 := GetNorm(1 - alpha)
	ucrit2 := GetNorm(1 - alpha/2)

	// Для студента
	fileSt := xlsx.NewFile()
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	header1 := []string{"n1", "M1", "n2", "M2"}
	for h1 := range header1 {
		cellSt1 := rowSt1.AddCell()
		cellSt1.Value = header1[h1]
	}

	rowSt2 := sheet1.AddRow()
	result1 := []string{strconv.Itoa(n1),
		strconv.Itoa(m1),
		strconv.Itoa(n2),
		strconv.Itoa(m2)}
	for r1 := range result1 {
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = result1[r1]
	}

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
	header2 := []string{"p1", "p2", "ZStatistic", "ucrit1", "ucrit2"}
	for h2 := range header2 {
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header2[h2]
	}

	rowPr2 := sheet2.AddRow()
	result2 := []string{fmt.Sprintf("%f", p1),
		fmt.Sprintf("%f", p2),
		fmt.Sprintf("%f", ZStatisticVal),
		fmt.Sprintf("%f", ucrit1),
		fmt.Sprintf("%f", ucrit2)}
	for r2 := range result2 {
		cellPr2 := rowPr2.AddCell()
		cellPr2.Value = result2[r2]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task5(min, max, n int, alpha float64, pathResults, pathProfData string) bool {
	seq1 := GenerateSeq(max, min, n)
	seq2 := GenerateSeq(max, min, n)
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
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	for s1 := range seq1 {
		cellSt1 := rowSt1.AddCell()
		cellSt1.Value = strconv.Itoa(seq1[s1])
	}

	rowSt2 := sheet1.AddRow()
	for s2 := range seq2 {
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = strconv.Itoa(seq2[s2])
	}

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
	header := []string{"Sigma1", "Sigma2", "FStatistic", "fcrit1left", "fcrit1right",
		"fcrit2left", "fcrit2right", "fcrit3left", "fcrit3right", "fcrit4left", "fcrit4right"}
	for h := range header {
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
	}

	rowPr2 := sheet2.AddRow()
	result := []string{fmt.Sprintf("%f", Sigma1),
		fmt.Sprintf("%f", Sigma2),
		fmt.Sprintf("%f", F),
		fmt.Sprintf("%f", fcrit1left),
		fmt.Sprintf("%f", fcrit1right),
		fmt.Sprintf("%f", fcrit2left),
		fmt.Sprintf("%f", fcrit2right),
		fmt.Sprintf("%f", fcrit3left),
		fmt.Sprintf("%f", fcrit3right),
		fmt.Sprintf("%f", fcrit4left),
		fmt.Sprintf("%f", fcrit4right)}
	for r := range result {
		cellPr2 := rowPr2.AddCell()
		cellPr2.Value = result[r]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}

func Task6(min, max, n1, n2, n3 int, alpha float64, pathResults, pathProfData string) bool {
	var seq1, seq2, seq3 []int
	var stopCounter = 0
	var quit = false
	for ok := true; ok; ok = (quit || stopCounter >= 100000) {
		seq1 = GenerateSeq(max, min, n1)
		seq2 = GenerateSeq(max, min, n2)
		seq3 = GenerateSeq(max, min, n3)
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
	sheet1, err1 := fileSt.AddSheet("Sheet1")
	if err1 != nil {
		fmt.Printf(err1.Error())
		return false
	}

	rowSt1 := sheet1.AddRow()
	for s1 := range seq1 {
		cellSt1 := rowSt1.AddCell()
		cellSt1.Value = strconv.Itoa(seq1[s1])
	}

	rowSt2 := sheet1.AddRow()
	for s2 := range seq2 {
		cellSt2 := rowSt2.AddCell()
		cellSt2.Value = strconv.Itoa(seq2[s2])
	}

	rowSt3 := sheet1.AddRow()
	for s3 := range seq3 {
		cellSt3 := rowSt3.AddCell()
		cellSt3.Value = strconv.Itoa(seq3[s3])
	}

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
	header := []string{"X1", "X2", "X3", "Xc", "TSS", "Q1", "Q2", "F", "fcrit1left", "fcrit1right"}
	for h := range header {
		cellPr1 := rowPr1.AddCell()
		cellPr1.Value = header[h]
	}

	rowPr2 := sheet2.AddRow()
	result := []string{fmt.Sprintf("%f", X1),
		fmt.Sprintf("%f", X2),
		fmt.Sprintf("%f", X3),
		fmt.Sprintf("%f", Xc),
		fmt.Sprintf("%f", TSS),
		fmt.Sprintf("%f", Q1),
		fmt.Sprintf("%f", Q2),
		fmt.Sprintf("%f", F),
		fmt.Sprintf("%f", fcrit1left),
		fmt.Sprintf("%f", fcrit1right)}
	for r := range result {
		cellPr2 := rowPr2.AddCell()
		cellPr2.Value = result[r]
	}

	err2 = filePr.Save(pathProfData)
	if err2 != nil {
		fmt.Printf(err2.Error())
		return false
	}
	return true
}
