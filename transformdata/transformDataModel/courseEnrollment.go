package transformdata

type CourseEnrollment struct {
	AtpID                   string      `json:"atpId"`
	Attributes              []Attribute `json:"attributes"`
	CourseID                string      `json:"courseId"`
	CourseOfferingID        string      `json:"courseOfferingId"`
	CourseSectionID         string      `json:"courseSectionId"`
	CreditsAwarded          int         `json:"creditsAwarded"`
	EndDate                 string      `json:"endDate"`
	EnrolDate               string      `json:"enrolDate"`
	ID                      string      `json:"id"`
	LanguageTypeKey         string      `json:"languageTypeKey"`
	Meta                    MetaData    `json:"meta"`
	MethodOfDeliveryTypeKey string      `json:"methodOfDeliveryTypeKey"`
	OrgID                   string      `json:"orgId"`
	PresentationTypeKey     string      `json:"presentationTypeKey"`
	ReferenceID             string      `json:"referenceId"`
	ReferenceTypeKey        string      `json:"referenceTypeKey"`
	Result                  Result      `json:"result"`
	SiteID                  string      `json:"siteId"`
	StartDate               string      `json:"startDate"`
	StateKey                string      `json:"stateKey"`
	StudentID               string      `json:"studentId"`
}

type Result struct {
	Attributes        []Attribute `json:"attributes"`
	CourseCompleted   bool        `json:"courseCompleted"`
	ExamMark          float64     `json:"examMark"`
	FinalMark         float64     `json:"finalMark"`
	GradePoint        float64     `json:"gradePoint"`
	ID                string      `json:"id"`
	Meta              MetaData    `json:"meta"`
	ModuleMark        float64     `json:"moduleMark"`
	ParticipationMark float64     `json:"participationMark"`
	ResultDate        string      `json:"resultDate"`
	ResultStateKey    string      `json:"resultStateKey"`
	ResultTypeKey     string      `json:"resultTypeKey"`
	SecondExamMark    float64     `json:"secondExamMark"`
	Symbol            string      `json:"symbol"`
}

type OuputCourseEnrollment struct {
	Id                         string
	UpdatedAt                  string
	CreatedAt                  string
	CourseOfferingId           string
	AtpSemsterId               string
	Status                     string
	ProgrammeLevelEnrollmentId string
}

func TransformCourseEnrollment(data CourseEnrollment) OuputCourseEnrollment {
	return OuputCourseEnrollment{
		Id:               data.ID,
		AtpSemsterId:     data.AtpID,
		CreatedAt:        data.Meta.CreateTime,
		UpdatedAt:        data.Meta.UpdateTime,
		CourseOfferingId: data.CourseOfferingID,
	}
}
