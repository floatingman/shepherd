package client

const (
	GKENodeConfigType                = "gkeNodeConfig"
	GKENodeConfigFieldBootDiskKmsKey = "bootDiskKmsKey"
	GKENodeConfigFieldDiskSizeGb     = "diskSizeGb"
	GKENodeConfigFieldDiskType       = "diskType"
	GKENodeConfigFieldImageType      = "imageType"
	GKENodeConfigFieldLabels         = "labels"
	GKENodeConfigFieldLocalSsdCount  = "localSsdCount"
	GKENodeConfigFieldMachineType    = "machineType"
	GKENodeConfigFieldOauthScopes    = "oauthScopes"
	GKENodeConfigFieldPreemptible    = "preemptible"
	GKENodeConfigFieldServiceAccount = "serviceAccount"
	GKENodeConfigFieldTags           = "tags"
	GKENodeConfigFieldTaints         = "taints"
)

type GKENodeConfig struct {
	BootDiskKmsKey string               `json:"bootDiskKmsKey,omitempty" yaml:"bootDiskKmsKey,omitempty"`
	DiskSizeGb     int64                `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`
	DiskType       string               `json:"diskType,omitempty" yaml:"diskType,omitempty"`
	ImageType      string               `json:"imageType,omitempty" yaml:"imageType,omitempty"`
	Labels         map[string]string    `json:"labels,omitempty" yaml:"labels,omitempty"`
	LocalSsdCount  int64                `json:"localSsdCount,omitempty" yaml:"localSsdCount,omitempty"`
	MachineType    string               `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	OauthScopes    []string             `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`
	Preemptible    bool                 `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`
	ServiceAccount string               `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	Tags           []string             `json:"tags,omitempty" yaml:"tags,omitempty"`
	Taints         []GKENodeTaintConfig `json:"taints,omitempty" yaml:"taints,omitempty"`
}
