package objects

type (
	Location struct {
		Name         string              `bson:"name" json:"name"`
		Geographic   *Geographic         `bson:"geographic" json:"geographic"`
		Provider     LocationProvider    `bson:"provider" json:"provider"`
		Features     LocationFeatureSets `bson:"features" json:"features"`
		Abbreviation string              `bson:"abbreviation" json:"abbreviation"`
		Annotations  map[string]string   `bson:"annotations" json:"annotations"`
	}

	Geographic struct {
		Lat     float64 `bson:"latitude" json:"latitude"`
		Long    float64 `bson:"longitude" json:"longitude"`
		City    string  `bson:"city" json:"city"`
		State   string  `bson:"state" json:"state"`
		Country string  `bson:"country" json:"country"`
		Region  string  `bson:"region" json:"region"`
	}

	LocationProvider struct {
		Location string `bson:"location" json:"location"` // the primary ID of how we reference this location
		Code     string `bson:"code" json:"code"`
	}

	LocationFeatureSets struct {
		Available LocationFeatures `bson:"available" json:"available"`
		Supported LocationFeatures `bson:"supported" json:"supported"`
	}

	LocationFeatures []LocationFeature

	LocationFeature string
)

const (
	LF_BARE_METAL LocationFeature = "bare-metal"
	LF_VM         LocationFeature = "virtual-machine"
)
