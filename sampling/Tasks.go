package sampling

import "math"

func Task1(count, min, max, n int, alpha float64) bool {
	answer := make(map[string]float64)
	task := make(map[string][]int)
	compare := make(map[string]bool)
	seq := GenerateSeq(max, min, n)
	cnst := GetRand(min, max)

	task["seq"] = seq

	answer["cnst"] = float64(cnst)
	answer["m"] = Average(seq)
	answer["Sigma"] = Variance(seq)
	answer["TStatistic"] = TStatistic(seq, float64(cnst))
	answer["tcrit1"] = GetStudent(1-alpha, n)
	answer["tcrit2"] = GetStudent(1-alpha/2, n) // Аккуратнее смотреть какую альфу передаём

	compare["answer"] = (GetStudent(1-alpha, n) < TStatistic(seq, float64(cnst))) && (GetStudent(1-alpha/2, n) > TStatistic(seq, float64(cnst)))

	return ReturnTask(count, task, answer, compare, 1)
}

func Task2(count, min, max, n int, cnst, alpha float64) bool {
	answer := make(map[string]float64)
	task := make(map[string][]int)
	compare := make(map[string]bool)
	seq := GenerateSeq(max, min, n)

	task["seq"] = seq

	answer["m"] = Average(seq)
	answer["Sigma"] = Variance(seq)
	tcrit := GetStudent(1-alpha/2, n)
	answer["tcrit1"] = tcrit
	answer["leftBoard"] = LeftTboard(seq, tcrit)
	answer["rightBoard"] = RightTboard(seq, tcrit)

	compare["answer"] = (LeftTboard(seq, tcrit) < TStatistic(seq, cnst)) && (RightTboard(seq, tcrit) > TStatistic(seq, cnst))

	return ReturnTask(count, task, answer, compare, 2)
}

func Task3(count, min, max, n int, alpha float64) bool {
	MAGIC_COEF := 1.05
	compare := make(map[string]bool)
	answer := make(map[string]float64)
	task := make(map[string][]int)
	seq1 := GenerateSeq(max, min, n)
	seq2 := GenerateSeq(max, min, n)

	task["seq1"] = seq1
	task["seq2"] = seq2

	answer["m1"] = Average(seq1)
	answer["m2"] = Average(seq2)
	answer["Sigma1"] = math.Sqrt(Variance(seq1)) * MAGIC_COEF
	answer["Sigma2"] = math.Sqrt(Variance(seq2)) * MAGIC_COEF
	answer["ZStatistic"] = ZStatistic(seq1, seq2)
	answer["ucrit1"] = GetNorm(1 - alpha)
	answer["ucrit2"] = GetNorm(1-alpha) / 2

	return ReturnTask(count, task, answer, compare, 3)
}

func Task4(count int, alpha float64) bool {
	answer := make(map[string]float64)
	task := make(map[string][]int)
	compare := make(map[string]bool)
	n1 := GetRand(100, 1500)
	n2 := GetRand(100, 1500)
	m1 := GetRand(int(0.25*float64(n1)), int(0.85*float64(n1)))
	m2 := GetRand(int(0.25*float64(n2)), int(0.85*float64(n2)))
	answer["n1"] = float64(n1)
	answer["n2"] = float64(n2)
	answer["m1"] = float64(m1)
	answer["m2"] = float64(m2)
	answer["p1"] = float64(m1 / n1)
	answer["p2"] = float64(m2 / n2)
	answer["ZStatistic"] = ZStatistic2(n1, n2, m1, m2)
	answer["ucrit1"] = GetNorm(1 - alpha)
	answer["ucrit2"] = GetNorm(1 - alpha/2)

	return ReturnTask(count, task, answer, compare, 4)
}

func Task5(count, min, max, n int, alpha float64) bool {
	answer := make(map[string]float64)
	task := make(map[string][]int)
	compare := make(map[string]bool)
	seq1 := GenerateSeq(max, min, n)
	seq2 := GenerateSeq(max, min, n)
	answer["Sigma1"] = Variance(seq1)
	answer["Sigma2"] = Variance(seq2)
	var f float64
	if Variance(seq1) > Variance(seq2) {
		f = Variance(seq1) / Variance(seq2)
	} else {
		f = Variance(seq2) / Variance(seq1)
	}

	task["seq1"] = seq1
	task["seq2"] = seq2

	answer["F"] = Round(f, 2)
	answer["fcrit1left"] = GetFisherLeft(alpha, n)
	answer["fcrit1right"] = GetFisherRight(alpha/2, n)
	answer["fcrit2left"] = GetFisherLeft(alpha/2, n)
	answer["fcrit2right"] = GetFisherRight(alpha/2, n)
	answer["fcrit3left"] = GetFisherLeft(1-alpha, n)
	answer["fcrit3right"] = GetFisherRight(1-alpha, n)
	answer["fcrit4left"] = GetFisherLeft(1-alpha/2, n)
	answer["fcrit4right"] = GetFisherRight(1-alpha/2, n)

	return ReturnTask(count, task, answer, compare, 5)
}

func Task6(count, min, max, n1, n2, n3 int, alpha float64) bool {
	answer := make(map[string]float64)
	task := make(map[string][]int)
	compare := make(map[string]bool)
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

	task["seq1"] = seq1
	task["seq2"] = seq2
	task["seq3"] = seq3

	answer["X1"] = Average(seq1)
	answer["X2"] = Average(seq2)
	answer["X3"] = Average(seq3)
	answer["Xc"] = Average(ConcatForThree(seq1, seq2, seq3))
	answer["TSS"] = Tss(seq1, seq2, seq3)
	answer["Q1"] = Q1(seq1, seq2, seq3)
	answer["Q2"] = Q2(seq1, seq2, seq3)
	answer["F"] = F(seq1, seq2, seq3)
	answer["fcrit1left"] = GetFisherLeft(alpha, (n1+n2+n3-3)/3)
	answer["fcrit1right"] = GetFisherRight(alpha, (n1+n2+n3-3)/3)

	return ReturnTask(count, task, answer, compare, 6)
}
