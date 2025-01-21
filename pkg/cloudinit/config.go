package cloudinit

import (
    "encoding/base64"
    "fmt"
    "gopkg.in/yaml.v3"
)

// CloudInitConfig représente la configuration Cloud-Init
type CloudInitConfig struct {
    PackageUpdate  bool     `yaml:"package_update,omitempty"`
    PackageUpgrade bool     `yaml:"package_upgrade,omitempty"`
    Packages      []string `yaml:"packages,omitempty"`
    SSHKeys       []string `yaml:"ssh_authorized_keys,omitempty"`
    RunCmd        []string `yaml:"runcmd,omitempty"`
}

// NewCloudInitConfig crée une nouvelle configuration Cloud-Init
func NewCloudInitConfig() *CloudInitConfig {
    return &CloudInitConfig{}
}

// ToBase64 encode la configuration en base64 pour l'API Contabo
func (c *CloudInitConfig) ToBase64() (string, error) {
    data, err := yaml.Marshal(c)
    if err != nil {
        return "", fmt.Errorf("error marshaling cloud-init config: %w", err)
    }
    return base64.StdEncoding.EncodeToString(data), nil
}
