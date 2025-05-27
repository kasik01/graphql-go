package database

type Grade struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	StudentID uint    `gorm:"index;not null"`
	SubjectID uint    `gorm:"index;not null"`
	GradeType string  `gorm:"type:varchar(20);not null"`
	Score     float64 `gorm:"not null"`

	Student Student //`gorm:"foreignKey:StudentID"`
	Subject Subject // `gorm:"foreignKey:SubjectID"`
}
