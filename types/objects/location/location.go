package locationObj

type (
	Location struct {
		Name         string            `bson:"name" json:"name"`
		Geographic   *Geographic       `bson:"geographic" json:"geographic"`
		Provider     Provider          `bson:"provider" json:"provider"`
		Features     FeatureSets       `bson:"features" json:"features"`
		Abbreviation string            `bson:"abbreviation" json:"abbreviation"`
		Annotations  map[string]string `bson:"annotations" json:"annotations"`
	}

	Geographic struct {
		Lat     float64 `bson:"latitude" json:"latitude"`
		Long    float64 `bson:"longitude" json:"longitude"`
		City    string  `bson:"city" json:"city"`
		State   string  `bson:"state" json:"state"`
		Country string  `bson:"country" json:"country"`
		Region  string  `bson:"region" json:"region"`
	}

	Provider struct {
		LocationId string `bson:"location_id" json:"location_id"` // the primary ID of how we reference this location
		Code       string `bson:"code" json:"code"`
	}

	FeatureSets struct {
		Available Features `bson:"available" json:"available"`
		Supported Features `bson:"supported" json:"supported"`
	}

	Features []Feature

	Feature string
)

const (
	LF_BARE_METAL Feature = "bare-metal"
	LF_VM         Feature = "virtual-machine"
)
