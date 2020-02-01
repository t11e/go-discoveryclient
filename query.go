package discoveryclient

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type Criterion struct {
	Dimension                      string      `json:"dimension,omitempty"`
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
	NotValue                       interface{} `json:"notValue,omitempty"`
	DidYouMean                     interface{} `json:"didYouMean,omitempty"`
	Cull                           *bool       `json:"cull,omitempty"`
	Min                            interface{} `json:"min,omitempty"`
	MinInclusive                   *bool       `json:"minInclusive,omitempty"`
	Max                            interface{} `json:"max,omitempty"`
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

// GroupBy criterion.
type GroupBy struct {
	Dimension  string   `json:"dimension"`
	TopN       int      `json:"topN"`
	Properties []string `json:"properties"`
	Groups     []string `json:"groups"`
	NotGroups  []string `json:"notGroups"`

	// LegacyGroupBy must currently always be false, since we only support the
	// non-legacy format. However, we need to include this field because the
	// engine defaults to true here.
	LegacyGroupBy bool `json:"legacyGroupBy"`

	// TODO: Is this correct?
	Highlighting *Highlighting `json:"highlighting"`

	// TODO: Support "indexValues"
}

type IndexValue struct {
	Dimension string `json:"dimension"`
}

type Query struct {
	Criteria         []Criterion               `json:"criteria,omitempty"`
	ExactMatchesOnly bool                      `json:"exactMatchesOnly"`
	PageSize         int                       `json:"pageSize"`
	Items            []string                  `json:"items,omitempty"`
	NotItems         []string                  `json:"notItems,omitempty"`
	ExactRelevance   *float64                  `json:"exactRelevance,omitempty"`
	StartIndex       int                       `json:"startIndex"`
	Values           []string                  `json:"values,omitempty"`
	Explain          string                    `json:"explain,omitempty"`
	IndexValues      []IndexValue              `json:"indexValues,omitempty"`
	GroupBy          *GroupBy                  `json:"groupBy,omitempty"`
	Properties       []string                  `json:"properties"`
	SortBy           []SortBy                  `json:"sortBy,omitempty"`
	Facets           map[string]FacetCriterion `json:"facets,omitempty"`
	Highlighting     *Highlighting             `json:"highlighting"`
}

type Highlighting struct {
	Merge             *string  `json:"merge,omitempty"`
	Template          []string `json:"template,omitempty"`
	IncludeDimensions []string `json:"includeDimensions,omitempty"`
	ExcludeDimensions []string `json:"excludeDimensions,omitempty"`
}

type FacetCriterion struct {
	MinCount      *int             `json:"minCount,omitempty"`
	TopN          *int             `json:"topN,omitempty"`
	Dynamic       *json.RawMessage `json:"dynamic,omitempty"` // TODO
	NavChildIDs   *string          `json:"navChildIds,omitempty"`
	Dimension     *string          `json:"dimension,omitempty"`
	DataIds       []string         `json:"dataIds,omitempty"`
	IncludeBounds *bool            `json:"includeBounds,omitempty"`
	IncludeLabel  *bool            `json:"includeLabel,omitempty"`
	Navigable     *bool            `json:"navigable,omitempty"`
	SortBy        *string          `json:"sortBy,omitempty"`
	IsolatedID    *string          `json:"isolatedId,omitempty"`
	RootID        *string          `json:"rootId,omitempty"`
	Depth         *int             `json:"depth,omitempty"`
	CountType     *string          `json:"countType,omitempty"`
}

type Groups struct {
	ItemIDs         [][]string          `json:"itemIds"`
	ExactMatches    [][]bool            `json:"exactMatches"`
	RelevanceValues [][]float64         `json:"relevanceValues"`
	Properties      [][]json.RawMessage `json:"properties"`
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
	Properties      []json.RawMessage                 `json:"properties"`
	IndexValues     map[string]map[string]interface{} `json:"indexValues"`
	IsGrouped       bool                              `json:"isGrouped"`
	Groups          *Groups                           `json:"groups,omitempty"`
	Highlighting    []map[string]*HighlightingData    `json:"highlighting"`
	Facets          *json.RawMessage                  `json:"facets"`
}

type HighlightingData struct {
	Fragments []string
}

func (h *HighlightingData) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		if err := json.Unmarshal(b, &h.Fragments); err != nil {
			return errors.Wrap(err, "array of highlighting fragments expected")
		}
	} else {
		var fragment string
		if err := json.Unmarshal(b, &fragment); err != nil {
			return errors.Wrap(err, "single highlighting fragment expected")
		}
		h.Fragments = []string{fragment}
	}
	return nil
}

func (h *HighlightingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.Fragments)
}
