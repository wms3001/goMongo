package model

type MOption struct {
	//ProjectionMap map[string]interface{} `json:"projectionMap" bson:"projectionMap"`
	Limit   int64                  `json:"limit" bson:"limit"`
	Skip    int64                  `json:"skip" bson:"skip"`
	SortMap map[string]interface{} `json:"sortMap" bson:"sortMap"`
}
