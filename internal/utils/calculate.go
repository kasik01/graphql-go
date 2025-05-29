package utils

import "graphql-hasura-demo/graph/model"

func CalculateAcademicPerformance(overallAverage float64) model.AcademicPerformance {
	switch {
	case overallAverage >= 9.0:
		return model.AcademicPerformanceExcellent
	case overallAverage >= 8.0:
		return model.AcademicPerformanceGood
	case overallAverage >= 6.5:
		return model.AcademicPerformanceAverage
	case overallAverage >= 5.0:
		return model.AcademicPerformanceWeak
	default:
		return model.AcademicPerformancePoor
	}
}
