package transformdata

type ProgramIntake struct {
	ApplicationPreRequisites []interface{} `json:"applicationPreRequisites"`
	AtpID                    string        `json:"atpId"`
	Attributes               []Attribute   `json:"attributes"`
	BundleID                 string        `json:"bundleId"`
	Contributors             []string      `json:"contributors"`
	CourseID                 *string       `json:"courseId"`
	Duration                 int           `json:"duration"`
	DurationUnitTypeKey      string        `json:"durationUnitTypeKey"`
	EnrolmentPreRequisites   []interface{} `json:"enrolmentPreRequisites"`
	ID                       string        `json:"id"`
	IntakeSections           []string      `json:"intakeSections"`
	LanguageTypeKey          *string       `json:"languageTypeKey"`
	MaxNoOfStudents          int           `json:"maxNoOfStudents"`
	Meta                     MetaData      `json:"meta"`
	MethodOfDeliveryTypeKey  string        `json:"methodOfDeliveryTypeKey"`
	MinNoOfStudents          int           `json:"minNoOfStudents"`
	OrgID                    string        `json:"orgId"`
	PresentationTypeKey      string        `json:"presentationTypeKey"`
	ProgramLevelID           string        `json:"programLevelId"`
	SiteID                   string        `json:"siteId"`
}

type OuputProgrammeIntake struct {
	Id                       string
	ProgrammeLevelId         string
	UpdatedAt                string
	CreatedAt                string
	AtpId                    string
	Duration                 string
	DurationUnit             string
	MinNoOfStudent           int
	MaxNoOfStudent           int
	BundleId                 string
	ProgrammeId              string
	RequiredApplication      bool
	RequiredManualApproval   bool
	Restrict_course_offering bool
	Grade_scale_id           string
	ResultTemplateId         string
}

func TransformProgrammeIntake(data ProgramIntake) OuputProgrammeIntake {
	output := OuputProgrammeIntake{
		Id:               data.ID,
		ProgrammeLevelId: data.ProgramLevelID,
		UpdatedAt:        data.Meta.UpdateTime,
		CreatedAt:        data.Meta.CreateTime,
		AtpId:            data.AtpID,
		Duration:         ChangeIntToString(data.Duration),
		DurationUnit:     data.DurationUnitTypeKey,
		MinNoOfStudent:   data.MinNoOfStudents,
		MaxNoOfStudent:   data.MaxNoOfStudents,
		BundleId:         data.BundleID,
	}
	for _, attr := range data.Attributes {
		switch attr.Key {
		case "application.type":
			if attr.Value == "application.process.required" {
				output.RequiredApplication = true
			}
		case "approval.type":
			if attr.Value == "application.approval.type.manual" {
				output.RequiredManualApproval = true
			}
		case "course.offering.restriction.type":
			if attr.Value == "course.offering.restriction.list" {
				output.Restrict_course_offering = true
			}
		case "program.offering.result.calc.type":
			output.ResultTemplateId = attr.Value
		case "intake.result.calc.type":
			output.Grade_scale_id = attr.Value
		}
	}

	return output
}
