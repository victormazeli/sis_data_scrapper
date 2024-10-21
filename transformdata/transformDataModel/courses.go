package transformdata

type InputCourse struct {
	ID              string           `json:"id"`
	Code            string           `json:"code"`
	Credits         int              `json:"credits"`
	Representations []Representation `json:"representations"`
	Attributes      []Attribute      `json:"attributes"`
}

type OutputCourse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CourseCode  string `json:"course_code"`
	Description string `json:"description"`
	CreditUnit  int    `json:"credit_unit"`
	ImageID     string `json:"image_id,omitempty"`
}

func TransformCourse(inputCourse InputCourse) OutputCourse {
	var name, description string
	if len(inputCourse.Representations) > 0 {
		name = inputCourse.Representations[0].LongName
		description = inputCourse.Representations[0].LongDescr
	}
	var imageID string
	for _, attr := range inputCourse.Attributes {
		if attr.Key == "image.id" {
			imageID = attr.Value
		}
	}

	return OutputCourse{
		ID:          inputCourse.ID,
		Name:        name,
		CourseCode:  inputCourse.Code,
		Description: description,
		CreditUnit:  inputCourse.Credits,
		ImageID:     imageID,
	}
}
