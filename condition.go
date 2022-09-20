package goMongo

type Filter struct {
	FMap map[string]interface{} `json:"fMap" bson:"fMap"`
	Type string                 `json:"type" bson:"type"`
}

type UData struct {
	UMap map[string]interface{} `json:"uMap" bson:"uMap"`
}

type Option struct {
	Limit int64                  `json:"limit" bson:"limit"`
	Skip  int64                  `json:"skip" bson:"skip"`
	SMap  map[string]interface{} `json:"sMap" bson:"sMap"`
}

type FU struct {
	Filter Filter `json:"filter" bson:"filter"`
	UData  UData  `json:"uData" bson:"uData"`
}

type FO struct {
	Filter Filter `json:"filter" bson:"filter"`
	Option Option `json:"option" bson:"option"`
}
