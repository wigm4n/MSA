package data

import "log"

func GetStudentValue(alpha float64, degree int) (val, semiVal float64, err error) {
	err = db.QueryRow("SELECT val_alpha, semi_val_alpha FROM student WHERE alpha = $1 and freedom_degree = $2",
		alpha, degree).Scan(&val, &semiVal)
	if err != nil {
		log.Println("in GetStudentValue exception:", err)
		return
	}
	return
}

func GetNormValue(alpha float64) (val, semiVal float64, err error) {
	err = db.QueryRow("SELECT val_alpha, semi_val_alpha FROM norm WHERE alpha = $1",
		alpha).Scan(&val, &semiVal)
	if err != nil {
		log.Println("in GetNormValue exception:", err)
		return
	}
	return
}

func GetFisherValue(alpha float64, degreeBig, degreeLittle int) (val float64, err error) {
	err = db.QueryRow("SELECT val_alpha FROM fisher WHERE alpha = $1 and freedom_degreeBig = $2 and freedom_degreeLittle = $3",
		alpha, degreeBig, degreeLittle).Scan(&val)
	if err != nil {
		log.Println("in GetFisherValue exception:", err)
		return
	}
	return
}
