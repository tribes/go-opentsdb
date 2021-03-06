package opentsdb

type Query struct {
	Aggregator string            `json:"aggregator"`
	Metric     string            `json:"metric"`
	Downsample string            `json:"downsample,omitempty"`
	Rate       bool              `json:"rate,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
}

type QueryResult struct {
	Metric        string             `json:"metric"`
	AggregateTags []string           `json:"aggregateTags,omitempty"`
	Tags          map[string]string  `json:"tags"`
	Dps           map[string]float64 `json:"dps"`
}

type QueryParams struct {
	Start             interface{} `json:"start"`
	End               interface{} `json:"end,omitempty"`
	Queries           []Query     `json:"queries,omitempty"`
	NoAnnotations     bool        `json:"no_annotations,omitempty"`
	GlobalAnnotations bool        `json:"global_annotations,omitempty"`
	MsResolution      bool        `json:"ms,omitempty"`
	ShowTSUIDs        bool        `json:"show_tsuids,omitempty"`
	ShowSummary       bool        `json:"show_summary,omitempty"`
	ShowQuery         bool        `json:"show_query,omitempty"`
	Delete            bool        `json:"delete,omitempty"`
}

func NewQueryParams() (*QueryParams, error) {
	return &QueryParams{}, nil
}

type SuggestParams struct {
	Type  string `json:"type"`
	Match string `json:"q,omnitempty"`
	Max   int    `json:"max,omitempty"`
}