package gradecalculators

import (
	"strconv"
)

// Assessment is a struct contains a grade and weight of a
// course.
type Assessment struct {
	// if scale is percentage or points, grade will be converted
	// into numbers. if scale is letters, grade will be string
	// like 'A' or 'A+'
	Grade string
	// if scale is percentage or letters, weigth is weight
	// if scale is points, weight will be max points
	Weight int
}

// CalculateGrade calculate grades by the following formula:
// for percentage and letters:
// grade = (g1xw1 + g2xw2 ... + gnxwn) / (w1 + w2 + ...+ wn)
// for points:
// grade = (g1 + g2 + ... + gn) / (w1 + w2 + ... + wn)
func CalculateGrade(scale string, assessments []Assessment) float64 {

	var gTotal float64
	var wTotal int
	for i := 0; i < len(assessments); i++ {
		var g float64
		var err error

		if scale == "letters" {
			gTotal += letter2Point(assessments[i].Grade)
		} else {
			g, err = strconv.ParseFloat(assessments[i].Grade, 32)
			if err != nil {
				return 0
			}
			gTotal += g * float64(assessments[i].Weight)
		}
		wTotal += assessments[i].Weight
	}

	return gTotal / float64(wTotal)
}

func letter2Point(letter string) float64 {
	letters := []string{"A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-", "D+", "D", "D-", "F"}
	points := []float64{4.33, 4, 3.67, 3.33, 3, 2.67, 2.33, 2, 1.67, 1.33, 1, 0.67, 0}

	for i := 0; i < len(letters); i++ {
		if letter == letters[i] {
			return points[i]
		}
	}

	return 0
}
