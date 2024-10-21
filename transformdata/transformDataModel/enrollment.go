package transformdata

type Enrollment struct {
	Attributes             []Attribute             `json:"attributes"`
	BundleID               string                  `json:"bundleId"`
	EndDate                string                  `json:"endDate"`
	EnrolDate              string                  `json:"enrolDate"`
	ID                     string                  `json:"id"`
	Meta                   MetaData                `json:"meta"`
	OrgID                  string                  `json:"orgId"`
	ProgramLevelEnrolments []ProgramLevelEnrolment `json:"programLevelEnrolments"`
	Result                 interface{}             `json:"result"`
	SiteID                 string                  `json:"siteId"`
	StartDate              string                  `json:"startDate"`
	StateKey               string                  `json:"stateKey"`
	StudentID              string                  `json:"studentId"`
}

type ProgramLevelEnrolment struct {
	AtpID                   string        `json:"atpId"`
	Attributes              []Attribute   `json:"attributes"`
	BundleEnrolmentID       string        `json:"bundleEnrolmentId"`
	BundleID                string        `json:"bundleId"`
	CourseAcknowledgements  []interface{} `json:"courseAcknowledgements"`
	EndDate                 string        `json:"endDate"`
	EnrolDate               string        `json:"enrolDate"`
	ID                      string        `json:"id"`
	IntakeID                string        `json:"intakeId"`
	IntakeSectionID         interface{}   `json:"intakeSectionId"`
	Meta                    MetaData      `json:"meta"`
	MethodOfDeliveryTypeKey string        `json:"methodOfDeliveryTypeKey"`
	OrgID                   string        `json:"orgId"`
	PresentationTypeKey     string        `json:"presentationTypeKey"`
	ProgramLevelID          string        `json:"programLevelId"`
	Result                  interface{}   `json:"result"`
	SiteID                  string        `json:"siteId"`
	StartDate               string        `json:"startDate"`
	StateKey                string        `json:"stateKey"`
	StateReason             interface{}   `json:"stateReason"`
	StudentID               string        `json:"studentId"`
}
type OuputProgrammeLevelEnrollment struct {
	Id                string
	ProgrammeIntakeId string
	EndDate           string
	EnrollDate        string
	AtpId             string
	ProgrammeLevelId  string
	CreateAt          string
	UpdateAt          string
}
type OutputEnrollment struct {
	Id                       string
	BundleId                 string
	StudentID                string
	EnrollDate               string
	EndDate                  string
	ProgrammeLevelEnrollment []OuputProgrammeLevelEnrollment
	CreateAt                 string
	UpdateAt                 string
}

func TransformEnrollment(data Enrollment) OutputEnrollment {
	var programmeLevelEnrollment []OuputProgrammeLevelEnrollment
	for _, v := range data.ProgramLevelEnrolments {
		programmeLevelEnrollment = append(programmeLevelEnrollment, OuputProgrammeLevelEnrollment{
			ProgrammeIntakeId: v.IntakeID,
			EndDate:           v.EndDate,
			EnrollDate:        v.StartDate,
			ProgrammeLevelId:  v.ProgramLevelID,
			AtpId:             v.AtpID,
			Id:                v.ID,
			CreateAt:          v.Meta.CreateTime,
			UpdateAt:          v.Meta.UpdateTime,
		})

	}
	return OutputEnrollment{
		Id:                       data.ID,
		EnrollDate:               data.EndDate,
		EndDate:                  data.EndDate,
		BundleId:                 data.BundleID,
		StudentID:                data.StudentID,
		ProgrammeLevelEnrollment: programmeLevelEnrollment,
		CreateAt:                 data.Meta.CreateTime,
		UpdateAt:                 data.Meta.UpdateTime,
	}
}
