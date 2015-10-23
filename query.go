package discoveryclient

type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type IntValue struct {
	Value int
}

type BoolValue struct {
	Value bool
}

type FloatValue struct {
	Value float64
}

type Criterion struct {
	Dimension                      string      `json:"dimension"`
	Id                             interface{} `json:"id,omitempty"`
	NotId                          interface{} `json:"notId,omitempty"`
	IsolatedId                     string      `json:"isolatedId,omitempty"`
	ExactRequires                  string      `json:"exactRequires,omitempty"`
	Operator                       string      `json:"operator,omitempty"`
	Criteria                       []Criterion `json:"criteria,omitempty"`
	MinStartsWithLength            *IntValue   `json:"minStartsWithLength,omitempty"`
	IgnoreFieldLength              *BoolValue  `json:"ignoreFieldLength,omitempty"`
	IgnoreInverseDocumentFrequency *BoolValue  `json:"ignoreInverseDocumentFrequency,omitempty"`
	SearchStyle                    string      `json:"searchStyle,omitempty"`
	Value                          interface{} `json:"value,omitempty"`
	DidYouMean                     interface{} `json:"didYouMean,omitempty"`
	Cull                           *BoolValue  `json:"cull,omitempty"`
	Min                            *FloatValue `json:"min,omitempty"`
	MinInclusive                   *BoolValue  `json:"minInclusive,omitempty"`
	Max                            *FloatValue `json:"max,omitempty"`
	MaxInclusive                   *BoolValue  `json:"maxInclusive,omitempty"`
	Latitude                       *FloatValue `json:"latitude,omitempty"`
	Longitude                      *FloatValue `json:"longitude,omitempty"`
	ZipCode                        string      `json:"zipcode,omitempty"`
	DistanceUnit                   string      `json:"distanceUnit,omitempty"`
	Fields                         []string    `json:"fields,omitempty"`
	IsoRelevanceDistance           interface{} `json:"isoRelevanceDistance,omitempty"`
	ExactMatch                     *BoolValue  `json:"exactMatch,omitempty"`
	NullExactMatch                 *BoolValue  `json:"nullExactMatch,omitempty"`
	ExactRelevance                 *FloatValue `json:"exactRelevance,omitempty"`
	FuzzyRelevance                 *FloatValue `json:"fuzzyRelevance,omitempty"`
	NullRelevance                  *FloatValue `json:"nullRelevance,omitempty"`
	ExactDistance                  *FloatValue `json:"exactDistance,omitempty"`
	CullDistance                   *FloatValue `json:"cullDistance,omitempty"`
	NormalDistance                 *FloatValue `json:"normalDistance,omitempty"`
	Weight                         *FloatValue `json:"weight,omitempty"`
	Placemarks                     *GeoJSON    `json:"placemarks,omitempty"`
}

type GroupBy struct {
	Dimension string `json:"dimension"`
	TopN      int    `json:"topN"`
	Legacy    bool   `json:"legacyGroupBy"`
}

type IndexValue struct {
	Dimension string `json:"dimension"`
}

type Query struct {
	Criteria         []Criterion  `json:"criteria,omitempty"`
	ExactMatchesOnly bool         `json:"exactMatchesOnly"`
	PageSize         int          `json:"pageSize"`
	Items            []string     `json:"items,omitempty"`
	NotItems         []string     `json:"notItems,omitempty"`
	ExactRelevance   *FloatValue  `json:"exactRelevance,omitempty"`
	StartIndex       int          `json:"startIndex"`
	Values           []string     `json:"values,omitempty"`
	Explain          string       `json:"explain,omitempty"`
	IndexValues      []IndexValue `json:"indexValues,omitempty"`
	GroupBy          *GroupBy     `json:"groupBy,omitempty"`
}

type Results struct {
	DatasetSize     int                               `json:"datasetSize"`
	PageSize        int                               `json:"pageSize"`
	CurrentPageSize int                               `json:"currentPageSize"`
	AvailableSize   int                               `json:"availableSize"`
	ExactSize       int                               `json:"exactSize"`
	TotalSize       int                               `json:"totalSize"`
	ExactMatches    []bool                            `json:"exactMatches"`
	RelevanceValues []float64                         `json:"relevanceValues"`
	ItemIds         []string                          `json:"itemIds"`
	IndexValues     map[string]map[string]interface{} `json:"indexValues"`
}
