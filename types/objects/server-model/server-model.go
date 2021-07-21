package serverModelObj

type (
	ServerModel struct {
		Name        string   `bson:"name" json:"name"`
		Description string   `bson:"description" json:"description"`
		Specs       Specs    `bson:"specs" json:"specs"`
		Provider    Provider `bson:"provider" json:"provider"`
		Pricing     Pricing  `bson:"pricing" json:"pricing"`
	}

	Pricing struct {
		HourlyMils  int `json:"hourly"` // $1 = 1000 Mills, $0.10 = 100 Mills
		MonthlyMils int `json:"monthly"`
	}

	Provider struct {
		Category  string   `bson:"category" json:"category"`
		Class     string   `bson:"class" json:"class,omitempty"`
		ModelId   string   `bson:"model_id" json:"model_id"` // the primary ID of how we reference this server / plan / model
		Locations []string `bson:"locations" json:"locations"`
	}

	Specs struct {
		Cpu      CPU       `bson:"cpu" json:"cpu"`
		Gpu      *GPU      `bson:"gpu" json:"gpu"`
		Memory   Memory    `bson:"memory" json:"memory"`
		Storage  []Storage `bson:"storage" json:"storage"`
		Network  []NIC     `bson:"network" json:"network"`
		Features Features  `bson:"features" json:"features"`
	}

	Features struct {
		RAID  *string                `bson:"raid" json:"raid"`
		Extra map[string]interface{} `bson:"extra" json:"extra,omitempty"` // provider specific features
	}

	CPU struct {
		Count   uint              `bson:"count" json:"count"`
		Cores   *uint             `bson:"cores" json:"cores"`
		Threads *uint             `bson:"threads" json:"threads"`
		Type    string            `bson:"type" json:"type"`
		Extra   map[string]string `bson:"extra" json:"extra,omitempty"`
	}

	GPU struct {
		Count uint              `bson:"count" json:"count"`
		Cores *uint             `bson:"cores" json:"cores"`
		Type  string            `bson:"type" json:"type"`
		Extra map[string]string `bson:"extra" json:"extra,omitempty"`
	}

	Memory struct {
		SizeGB uint              `bson:"size_gb" json:"size_gb"`
		Type   string            `bson:"type" json:"type"`
		Extra  map[string]string `bson:"extra" json:"extra,omitempty"`
	}

	Storage struct {
		Count  uint              `bson:"count" json:"count"`
		SizeGB uint              `bson:"size_gb" json:"size_gb"`
		Type   string            `bson:"type" json:"type"`
		Extra  map[string]string `bson:"extra" json:"extra,omitempty"`
	}

	NIC struct {
		Count          uint     `bson:"count" json:"count"`
		Scope          NICScope `bson:"scope" json:"scope"`
		Type           string   `bson:"type" json:"type"`
		ThroughputMbps uint     `bson:"throughput_mbps" json:"throughput_mbps"`
	}

	NICScope string
)

const (
	NICS_PUBLIC  NICScope = "public"
	NICS_PRIVATE NICScope = "private"
	NICS_SHARED  NICScope = "shared"
)
