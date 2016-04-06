package discoveryclient

type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type Criterion struct {
	Dimension                      string      `json:"dimension"`
	Id                             interface{} `json:"id,omitempty"`
	NotId                          interface{} `json:"notId,omitempty"`
	IsolatedId                     string      `json:"isolatedId,omitempty"`
	ExactRequires                  string      `json:"exactRequires,omitempty"`
	Operator                       string      `json:"operator,omitempty"`
	Criteria                       []Criterion `json:"criteria,omitempty"`
	MinStartsWithLength            *int        `json:"minStartsWithLength,omitempty"`
	IgnoreFieldLength              *bool       `json:"ignoreFieldLength,omitempty"`
	IgnoreInverseDocumentFrequency *bool       `json:"ignoreInverseDocumentFrequency,omitempty"`
	SearchStyle                    string      `json:"searchStyle,omitempty"`
	Value                          interface{} `json:"value,omitempty"`
	DidYouMean                     interface{} `json:"didYouMean,omitempty"`
	Cull                           *bool       `json:"cull,omitempty"`
	Min                            *float64    `json:"min,omitempty"`
	MinInclusive                   *bool       `json:"minInclusive,omitempty"`
	Max                            *float64    `json:"max,omitempty"`
	MaxInclusive                   *bool       `json:"maxInclusive,omitempty"`
	Latitude                       *float64    `json:"latitude,omitempty"`
	Longitude                      *float64    `json:"longitude,omitempty"`
	ZipCode                        string      `json:"zipcode,omitempty"`
	DistanceUnit                   string      `json:"distanceUnit,omitempty"`
	Fields                         []string    `json:"fields,omitempty"`
	IsoRelevanceDistance           interface{} `json:"isoRelevanceDistance,omitempty"`
	ExactMatch                     *bool       `json:"exactMatch,omitempty"`
	NullExactMatch                 *bool       `json:"nullExactMatch,omitempty"`
	ExactRelevance                 *float64    `json:"exactRelevance,omitempty"`
	FuzzyRelevance                 *float64    `json:"fuzzyRelevance,omitempty"`
	NullRelevance                  *float64    `json:"nullRelevance,omitempty"`
	ExactDistance                  *float64    `json:"exactDistance,omitempty"`
	CullDistance                   *float64    `json:"cullDistance,omitempty"`
	NormalDistance                 *float64    `json:"normalDistance,omitempty"`
	Weight                         *float64    `json:"weight,omitempty"`
	Placemarks                     *GeoJSON    `json:"placemarks,omitempty"`
}

// http://docs.discoverysearchengine.com/release/4.0/query/query_api.html#sort-by-criterion
type SortBy struct {
	// Dimension is mandatory unless Builtin is provided
	Dimension string `json:"dimension,omitempty"`
	// Builtin should be exactMatch, relevance, id or random
	Builtin string `json:"builtin,omitempty"`
	// Order defaults to asc for everything other than
	//    builtin exactMatch which defaults to desc (True > False)
	//    builtin relevance which defaults to desc
	Reverse *bool `json:"reverse,omitempty"`

	// Only valid for geoloc dimensions
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	// DistanceUnit should be miles or km
	DistanceUnit string `json:"distanceUnit,omitempty"`

	// Only valid for builtin random
	Seed         *float64 `json:"seed,omitempty"`
	MinRelevance *float64 `json:"minRelevance,omitempty"`
	MaxRelevance *float64 `json:"maxRelevance,omitempty"`
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
	ExactRelevance   *float64     `json:"exactRelevance,omitempty"`
	StartIndex       int          `json:"startIndex"`
	Values           []string     `json:"values,omitempty"`
	Explain          string       `json:"explain,omitempty"`
	IndexValues      []IndexValue `json:"indexValues,omitempty"`
	GroupBy          *GroupBy     `json:"groupBy,omitempty"`
	Properties       []string     `json:"properties"`
	SortBy           []SortBy     `json:"sortBy,omitempty"`
}

type PropertyMap map[string]interface{}

type Groups struct {
	ItemIDs         [][]string  `json:"itemIds"`
	ExactMatches    [][]bool    `json:"exactMatches"`
	RelevanceValues [][]float64 `json:"relevanceValues"`
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
	Properties      []PropertyMap                     `json:"properties"`
	IndexValues     map[string]map[string]interface{} `json:"indexValues"`
	IsGrouped       bool                              `json:"isGrouped"`
	Groups          *Groups                           `json:"groups,omitempty"`
}
