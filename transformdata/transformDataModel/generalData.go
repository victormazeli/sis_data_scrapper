package transformdata

import "strconv"

type MetaData struct {
	AuditFunction interface{} `json:"auditFunction"`
	CreateID      string      `json:"createId"`
	CreateTime    string      `json:"createTime"`
	UpdateID      string      `json:"updateId"`
	UpdateTime    string      `json:"updateTime"`
	VersionInd    int         `json:"versionInd"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Representation struct {
	ID         string `json:"id"`
	Locale     string `json:"locale"`
	LongDescr  string `json:"longDescr"`
	LongName   string `json:"longName"`
	ShortDescr string `json:"shortDescr"` 
	ShortName  string `json:"shortName"` 
}

func ChangeIntToString(value int) string {
	return strconv.Itoa(value)
}
