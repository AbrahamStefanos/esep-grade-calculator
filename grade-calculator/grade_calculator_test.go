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

func TestGetGradeC(t *testing.T) {
	expected := "C"
	gc := NewGradeCalculator()
	// 0.5*70 + 0.35*70 + 0.15*70 = 70 -> C
	gc.AddGrade("a1", 70, Assignment)
	gc.AddGrade("e1", 70, Exam)
	gc.AddGrade("s1", 70, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeD(t *testing.T) {
	expected := "D"
	gc := NewGradeCalculator()
	// 0.5*60 + 0.35*60 + 0.15*60 = 60 -> D
	gc.AddGrade("a1", 60, Assignment)
	gc.AddGrade("e1", 60, Exam)
	gc.AddGrade("s1", 60, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Assignment.String() = %q, want %q", Assignment.String(), "assignment")
	}
	if Exam.String() != "exam" {
		t.Errorf("Exam.String() = %q, want %q", Exam.String(), "exam")
	}
	if Essay.String() != "essay" {
		t.Errorf("Essay.String() = %q, want %q", Essay.String(), "essay")
	}
}
