package instance

import (
    "fmt"
    "github.com/your-username/contabo-client/pkg/cloudinit"
)

// CreateInstanceOptions represents options for creating a new instance
type CreateInstanceOptions struct {
    ImageID     string
    ProductID   string
    Region      string
    DisplayName string
    CloudInit   *cloudinit.CloudInitConfig
}

// CreateInstance creates a new instance with the specified options
func CreateInstance(opts CreateInstanceOptions) error {
    if opts.CloudInit != nil {
        // Validate cloud-init configuration
        validator := cloudinit.NewCloudInitValidator()
        if err := validator.ValidateConfig(opts.CloudInit); err != nil {
            return fmt.Errorf("invalid cloud-init configuration: %w", err)
        }

        // Convert cloud-init to base64
        cloudInitBase64, err := opts.CloudInit.ToBase64()
        if err != nil {
            return fmt.Errorf("failed to encode cloud-init configuration: %w", err)
        }

        // TODO: Add the actual API call to create instance with cloud-init
        // This would involve calling the Contabo API with the cloudInitBase64
    }

    return nil
}

// CreateInstanceWithDefaultCloudInit creates an instance with default cloud-init settings
func CreateInstanceWithDefaultCloudInit(imageID, productID, region, displayName string) error {
    config := cloudinit.NewCloudInitConfig()
    config.PackageUpdate = true
    config.PackageUpgrade = true
    config.Packages = []string{"curl", "wget", "vim"}

    return CreateInstance(CreateInstanceOptions{
        ImageID:     imageID,
        ProductID:   productID,
        Region:      region,
        DisplayName: displayName,
        CloudInit:   config,
    })
}
