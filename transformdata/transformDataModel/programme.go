package transformdata

type Program struct {
	Accreditations       []interface{}    `json:"accreditations"`
	Attributes           []Attribute      `json:"attributes"`
	EndDate              *string          `json:"endDate"`
	FieldOfStudyCode     string           `json:"fieldOfStudyCode"`
	ID                   string           `json:"id"`
	Levels               []Level          `json:"levels"`
	MaxPeriod            int              `json:"maxPeriod"`
	MaxPeriodUnitTypeKey string           `json:"maxPeriodUnitTypeKey"`
	Meta                 MetaData         `json:"meta"`
	MinPeriod            int              `json:"minPeriod"`
	MinPeriodUnitTypeKey string           `json:"minPeriodUnitTypeKey"`
	PreRequisites        []interface{}    `json:"preRequisites"`
	PresentingOrgs       []PresentingOrg  `json:"presentingOrgs"`
	ProgramCode          string           `json:"programCode"`
	ProgramTypeKey       string           `json:"programTypeKey"`
	Representations      []Representation `json:"representations"`
	Specifications       []interface{}    `json:"specifications"`
	StartDate            string           `json:"startDate"`
}

type Level struct {
	ID             string `json:"id"`
	LevelCode      string `json:"levelCode"`
	LevelStatusKey string `json:"levelStatusKey"`
}

type PresentingOrg struct {
	Attributes []interface{} `json:"attributes"`
	ID         string        `json:"id"`
	Meta       MetaData      `json:"meta"`
	OrgID      string        `json:"orgId"`
	SiteID     string        `json:"siteId"`
}

type OuputProgramme struct {
	Id              string
	Name            string
	ShortName       string
	Duration        string
	DurationUnit    string
	FacultyID       string
	MinimumDuration string
	Maximumduration string
	ProgrammeType   string
	Image           string
	Description     string
	CreatedAt       string
	UpdateAt        string
	Level           []ProgrammeLevel
	ProgrammeCode   string
}

type ProgrammeLevel struct {
	Id          string
	LevelNumber string
}

func TransformProgrammeData(data Program) OuputProgramme {
	var name, shortName, description, imageID string
	var level []ProgrammeLevel
	for _, v := range data.Levels {
		level = append(level, ProgrammeLevel{
			LevelNumber: v.LevelCode,
			Id:          v.ID,
		})
	}
	if len(data.Representations) != 0 {
		name = data.Representations[0].LongName
		shortName = data.Representations[0].ShortName
		description = data.Representations[0].LongDescr
	}
	for _, attr := range data.Attributes {
		if attr.Key == "image.id" {
			imageID = attr.Value
		}
	}
	return OuputProgramme{
		Level:           level,
		ProgrammeType:   data.ProgramTypeKey,
		ProgrammeCode:   data.ProgramCode,
		MinimumDuration: ChangeIntToString(data.MinPeriod),
		Maximumduration: ChangeIntToString(data.MaxPeriod),
		Name:            name,
		ShortName:       shortName,
		Description:     description,
		Image:           imageID,
		CreatedAt:       data.Meta.CreateTime,
		UpdateAt:        data.Meta.UpdateTime,
	}
}
