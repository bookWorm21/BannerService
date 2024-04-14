package banner

import "encoding/json"

type Banner struct {
	Content json.RawMessage
}

type Info struct {
	TagId           int
	FeatureId       int
	UseLastRevision bool
}
