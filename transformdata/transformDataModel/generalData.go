package transformdata

import "strconv"

type MetaData struct {
	AuditFunction interface{} `json:"auditFunction"` // Keeping as empty interface
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
	ShortDescr string `json:"shortDescr"` // Nullable
	ShortName  string `json:"shortName"`  // Nullable
}

func ChangeIntToString(value int) string {
	return strconv.Itoa(value)
}
