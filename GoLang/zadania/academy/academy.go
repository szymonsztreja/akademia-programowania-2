package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	gradesSum := 0

	if len(grades) == 0 {
		return 0
	}

	for counter := 0; counter < len(grades); counter++ {
		gradesSum += grades[counter]
	}

	averageGradeFloat := float64(gradesSum) / float64(len(grades))
	averageGrade := int(math.Round(averageGradeFloat))

	return averageGrade
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	numOfClassAttended := 0

	if len(attendance) == 0 {
		return 0.0
	}

	for counter := 0; counter < len(attendance); counter++ {
		if attendance[counter] {
			numOfClassAttended++
		}
	}
	attendancePercantage := float64(numOfClassAttended) / float64(len(attendance))

	return attendancePercantage
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {

	finalGradeSum := s.Project + AverageGrade(s.Grades)
	finalGradeFloat := math.Round(float64(finalGradeSum) / 2.0)
	finalGrade := int(finalGradeFloat)

	if AttendancePercentage(s.Attendance) < 0.8 {
		finalGrade--
	}

	if AttendancePercentage(s.Attendance) < 0.6 ||
		AverageGrade(s.Grades) == 1 || s.Project == 1 {
		finalGrade = 1
	}

	return finalGrade
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	studentsMap := make(map[string]uint8)

	for counter := 0; counter < len(students); counter++ {
		studentsMap[students[counter].Name] = uint8(FinalGrade(students[counter]))
	}

	return studentsMap
}
