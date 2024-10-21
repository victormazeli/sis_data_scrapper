package transformdata

type CourseResultTemplate struct {
	Attributes            []interface{}                        `json:"attributes"` // Keeping as empty interface since the structure is not defined
	DecimalPlaces         float64                              `json:"decimalPlaces"`
	ID                    string                               `json:"id"`
	Meta                  CourseResultTemplateMetaData         `json:"meta"`
	Representations       []CourseResultTemplateRepresentation `json:"representations"`
	ResultConfigurations  []ResultConfiguration                `json:"resultConfigurations"`
	TemplateResultTypeKey string                               `json:"templateResultTypeKey"`
}

type CourseResultTemplateMetaData struct {
	AuditFunction interface{} `json:"auditFunction"` // Keeping as empty interface
	CreateID      string      `json:"createId"`
	CreateTime    string      `json:"createTime"`
	UpdateID      string      `json:"updateId"`
	UpdateTime    string      `json:"updateTime"`
	VersionInd    int         `json:"versionInd"`
}

type CourseResultTemplateRepresentation struct {
	ID         string `json:"id"`
	Locale     string `json:"locale"`
	LongDescr  string `json:"longDescr"`
	LongName   string `json:"longName"`
	ShortDescr string `json:"shortDescr"`
	ShortName  string `json:"shortName"`
}

type ResultConfiguration struct {
	FromPoints    int    `json:"fromPoints"`
	ID            string `json:"id"`
	ResultTypeKey string `json:"resultTypeKey"`
	ToPoints      int    `json:"toPoints"`
}

type OuputResultTemplate struct {
	id                  string
	Name                string
	Description         string
	MarkFormart         float64
	ResultTemplateType  string
	ResultConfiguration []ResultConfiguration
	CreatedAt           string
	UpdatedAt           string
}

func TransformResultTemplate(data CourseResultTemplate) OuputResultTemplate {
	var name, description string
	if len(data.Representations) != 0 {
		name = data.Representations[0].ShortName
		description = data.Representations[0].LongDescr
	}
	return OuputResultTemplate{
		id:                  data.ID,
		Name:                name,
		Description:         description,
		MarkFormart:         data.DecimalPlaces,
		ResultTemplateType:  data.TemplateResultTypeKey,
		ResultConfiguration: data.ResultConfigurations,
		CreatedAt:           data.Meta.CreateTime,
		UpdatedAt:           data.Meta.UpdateID,
	}
}
