package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestEmptyCategoriesBehavior(t *testing.T) {
	gcA := NewGradeCalculator()
	gcA.AddGrade("a", 100, Assignment)
	if got := gcA.GetFinalGrade(); got != "F" {
		t.Fatalf("only assignments=100: want F, got %s", got)
	}

	gcE := NewGradeCalculator()
	gcE.AddGrade("e", 100, Exam)
	if got := gcE.GetFinalGrade(); got != "F" {
		t.Fatalf("only exams=100: want F, got %s", got)
	}

	gcS := NewGradeCalculator()
	gcS.AddGrade("s", 100, Essay)
	if got := gcS.GetFinalGrade(); got != "F" {
		t.Fatalf("only essays=100: want F, got %s", got)
	}

	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("no grades: want F, got %s", got)
	}
}

func TestGradeTypeNames(t *testing.T) {
	tests := map[GradeType]string{
		Assignment: "assignment",
		Exam:       "exam",
		Essay:      "essay",
	}

	for gradeType, expected := range tests {
		if got := gradeType.String(); got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	}
}

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 70, Assignment)
	gc.AddGrade("e1", 70, Exam)
	gc.AddGrade("s1", 70, Essay)

	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("want C, got %s", got)
	}
}

func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 65, Assignment)
	gc.AddGrade("e", 65, Exam)
	gc.AddGrade("s", 65, Essay)

	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("want D, got %s", got)
	}
}