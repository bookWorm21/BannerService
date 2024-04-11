package banner

type Banner struct{}

type UserBannerRequest struct {
	TagId          int   `json:"tag_id"`
	FeatureId      int   `json:"feature_id"`
	UseLastVersion *bool `json:"use_last_version"`
}

type UserBannerResponse struct {
	BannerJson string
}
