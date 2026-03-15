package domain

import "testing"

func TestCalcAlignmentGrade(t *testing.T) {
	tests := []struct {
		score int
		want  AlignmentGrade
	}{
		{200, GradeRoyalKnight},
		{100, GradeRoyalKnight},
		{99, GradeLawful},
		{50, GradeLawful},
		{49, GradeNeutral},
		{0, GradeNeutral},
		{-1, GradeCaution},
		{-30, GradeCaution},
		{-31, GradeChaotic},
		{-100, GradeChaotic},
	}
	for _, tt := range tests {
		got := CalcAlignmentGrade(tt.score)
		if got != tt.want {
			t.Errorf("CalcAlignmentGrade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}
