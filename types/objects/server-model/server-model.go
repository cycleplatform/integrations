package serverModelObj

type (
	ServerModel struct {
		Name        string   `bson:"name" json:"name"`
		Drivers     []string `bson:"drivers" json:"drivers"`
		Description string   `bson:"description" json:"description"`
		Specs       Specs    `bson:"specs" json:"specs"`
		Provider    Provider `bson:"provider" json:"provider"`
		Pricing     Pricing  `bson:"pricing" json:"pricing"`
		Compatible  *bool    `bson:"compatible" json:"compatible"`
	}

	Pricing struct {
		HourlyMils  int `json:"hourly"` // $1 = 1000 Mills, $0.10 = 100 Mills
		MonthlyMils int `json:"monthly"`
	}

	Provider struct {
		Category          string            `bson:"category" json:"category"`
		Class             string            `bson:"class" json:"class,omitempty"`
		ModelId           string            `bson:"model_id" json:"model_id"` // the primary ID of how we reference this server / plan / model
		Locations         []string          `bson:"locations" json:"locations"`
		AvailabilityZones AvailabilityZones `bson:"availability_zones" json:"availability_zones"`
	}

	AvailabilityZones map[string][]string

	Specs struct {
		Cpu      CPU       `bson:"cpu" json:"cpu"`
		Gpu      *GPU      `bson:"gpu" json:"gpu"`
		Memory   Memory    `bson:"memory" json:"memory"`
		Storage  []Storage `bson:"storage" json:"storage"`
		Network  []NIC     `bson:"network" json:"network"`
		Features Features  `bson:"features" json:"features"`
	}

	Features struct {
		RAID       *string                `bson:"raid" json:"raid"`
		Hypervisor *bool                  `bson:"hypervisor" json:"hypervisor"`
		Extra      map[string]interface{} `bson:"extra" json:"extra,omitempty"` // provider specific features
	}

	CPU struct {
		Count   uint              `bson:"count" json:"count"`
		Cores   *uint             `bson:"cores" json:"cores"`
		Threads *uint             `bson:"threads" json:"threads"`
		Type    string            `bson:"type" json:"type"`
		Shared  *bool             `bson:"shared" json:"shared"`
		Extra   map[string]string `bson:"extra" json:"extra,omitempty"`
	}

	GPU struct {
		Count  uint              `bson:"count" json:"count"`
		Cores  *uint             `bson:"cores" json:"cores"`
		VRAMGB *uint             `bson:"vram_gb" json:"vram_gb"`
		Type   string            `bson:"type" json:"type"`
		Shared *bool             `bson:"shared" json:"shared"`
		Extra  map[string]string `bson:"extra" json:"extra,omitempty"`
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
		Scope      string `bson:"scope" json:"scope"`
		Type       string `bson:"type" json:"type"`
		Throughput string `bson:"throughput" json:"throughput"`
	}
)
