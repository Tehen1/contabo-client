package cloudinit

import (
    "encoding/base64"
    "testing"
    "gopkg.in/yaml.v3"
)

func TestNewCloudInitConfig(t *testing.T) {
    config := NewCloudInitConfig()
    if config == nil {
        t.Error("NewCloudInitConfig should return a non-nil config")
    }
}

func TestCloudInitConfigToBase64(t *testing.T) {
    tests := []struct {
        name    string
        config  CloudInitConfig
        wantErr bool
    }{
        {
            name: "basic config",
            config: CloudInitConfig{
                PackageUpdate:  true,
                PackageUpgrade: true,
                Packages:      []string{"nginx", "docker"},
                SSHKeys:       []string{"ssh-rsa AAAA..."},
                RunCmd:        []string{"systemctl start nginx"},
            },
            wantErr: false,
        },
        {
            name: "empty config",
            config: CloudInitConfig{},
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := tt.config.ToBase64()
            if (err != nil) != tt.wantErr {
                t.Errorf("CloudInitConfig.ToBase64() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            // Decode and verify the base64 string
            decoded, err := base64.StdEncoding.DecodeString(got)
            if err != nil {
                t.Errorf("Failed to decode base64 string: %v", err)
                return
            }

            // Unmarshal and verify the YAML
            var result CloudInitConfig
            if err := yaml.Unmarshal(decoded, &result); err != nil {
                t.Errorf("Failed to unmarshal YAML: %v", err)
                return
            }
        })
    }
}
