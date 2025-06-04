package notification

type (
	Grade struct {
		Id        uint    `json:"id"`
		StudentId uint    `json:"student_id"`
		SubjectId uint    `json:"subject_id"`
		GradeType string  `json:"grade_type"`
		Score     float64 `json:"score"`
	}
)
