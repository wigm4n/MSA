package sampling

import "MSA/data"

func GetStudent(alpha float64, freedomDegree int) (float64, float64) {
	val, semiVal, _ := data.GetStudentValue(alpha, freedomDegree)
	return val, semiVal
}

func GetNorm(alpha float64) (float64, float64) {
	val, semiVal, _ := data.GetNormValue(alpha)
	return val, semiVal
}

func GetFisher(alpha float64, freedomDegreeBig, freedomDegreeLittle int) float64 {
	if freedomDegreeBig >= 120 {
		freedomDegreeBig = 120
	} else if freedomDegreeBig > 10 && freedomDegreeBig < 12 {
		freedomDegreeBig = 10
	} else if freedomDegreeBig > 12 && freedomDegreeBig < 15 {
		freedomDegreeBig = 12
	} else if freedomDegreeBig > 15 && freedomDegreeBig < 20 {
		freedomDegreeBig = 15
	} else if freedomDegreeBig > 20 && freedomDegreeBig < 24 {
		freedomDegreeBig = 20
	} else if freedomDegreeBig > 24 && freedomDegreeBig < 30 {
		freedomDegreeBig = 24
	} else if freedomDegreeBig > 30 && freedomDegreeBig < 40 {
		freedomDegreeBig = 30
	} else if freedomDegreeBig > 40 && freedomDegreeBig < 60 {
		freedomDegreeBig = 40
	} else if freedomDegreeBig > 60 && freedomDegreeBig < 120 {
		freedomDegreeBig = 60
	}

	if freedomDegreeLittle >= 120 {
		freedomDegreeLittle = 120
	} else if freedomDegreeLittle > 30 && freedomDegreeLittle < 40 {
		freedomDegreeLittle = 30
	} else if freedomDegreeLittle > 40 && freedomDegreeLittle < 60 {
		freedomDegreeLittle = 40
	} else if freedomDegreeLittle > 60 && freedomDegreeLittle < 120 {
		freedomDegreeLittle = 60
	}

	val, _ := data.GetFisherValue(alpha, freedomDegreeBig, freedomDegreeLittle)
	return val
}
