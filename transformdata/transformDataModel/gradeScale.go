package transformdata

type Grade struct {
	Attributes        []interface{} `json:"attributes"` // Assuming attributes can vary, keeping as empty interface
	AveragePercentage float64       `json:"averagePercentage"`
	ID                string        `json:"id"`
	Meta              MetaData      `json:"meta"`
	PercentageFrom    float64       `json:"percentageFrom"`
	PercentageTo      float64       `json:"percentageTo"`
	Point             float64       `json:"point"`
	Symbol            string        `json:"symbol"`
}

type GradeScale struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Description    string           `json:"description,omitempty"`
	CreatedAt      string           `json:"created_at,omitempty"`
	UpdatedAt      string           `json:"updated_at,omitempty"`
	GradeScaleItem []GradeScaleItem `json:"grade_scale_item"`
}

type GradeScaleItem struct {
	ID                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	PercentageFrom    float64 `json:"percentageFrom"`
	PercentageTo      float64 `json:"percentageTo"`
	AveragePercentage float64 `json:"averagePercentage,omitempty"`
	Point             float64 `json:"point"`
	Remark            string  `json:"remark,omitempty"`
	CreatedAt         string  `json:"created_at,omitempty"`
	UpdatedAt         string  `json:"updated_at,omitempty"`
}

func TransformGradeScaleItem(grade Grade) GradeScaleItem {
	return GradeScaleItem{
		ID:                grade.ID,
		Symbol:            grade.Symbol,
		PercentageFrom:    grade.PercentageFrom,
		PercentageTo:      grade.PercentageTo,
		Point:             grade.Point,
		AveragePercentage: grade.AveragePercentage,
		CreatedAt:         grade.Meta.CreateTime,
		UpdatedAt:         grade.Meta.UpdateTime,
	}
}
