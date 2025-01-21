package validator

import (
    "fmt"
    "strings"
)

// ValidationError represents an error during cloud-init validation
type ValidationError struct {
    Field string
    Error string
}

// CloudInitValidator validates cloud-init configurations
type CloudInitValidator struct {
    MaxPackages    int
    MaxSSHKeySize int
}

// NewCloudInitValidator creates a new validator with default limits
func NewCloudInitValidator() *CloudInitValidator {
    return &CloudInitValidator{
        MaxPackages:    50,    // Maximum number of packages to install
        MaxSSHKeySize:  4096,  // Maximum SSH key size in bytes
    }
}

// ValidateSSHKey validates an SSH key format and size
func (v *CloudInitValidator) ValidateSSHKey(key string) *ValidationError {
    if len(key) == 0 {
        return &ValidationError{
            Field: "ssh_key",
            Error: "SSH key cannot be empty",
        }
    }

    if len(key) > v.MaxSSHKeySize {
        return &ValidationError{
            Field: "ssh_key",
            Error: fmt.Sprintf("SSH key exceeds maximum size of %d bytes", v.MaxSSHKeySize),
        }
    }

    parts := strings.Fields(key)
    if len(parts) < 2 {
        return &ValidationError{
            Field: "ssh_key",
            Error: "Invalid SSH key format",
        }
    }

    return nil
}

// ValidatePackages validates the package list
func (v *CloudInitValidator) ValidatePackages(packages []string) *ValidationError {
    if len(packages) > v.MaxPackages {
        return &ValidationError{
            Field: "packages",
            Error: fmt.Sprintf("Number of packages exceeds maximum of %d", v.MaxPackages),
        }
    }

    for _, pkg := range packages {
        if strings.TrimSpace(pkg) == "" {
            return &ValidationError{
                Field: "packages",
                Error: "Package name cannot be empty",
            }
        }

        // Basic package name validation
        if strings.ContainsAny(pkg, ";&|><$") {
            return &ValidationError{
                Field: "packages",
                Error: fmt.Sprintf("Invalid package name: %s", pkg),
            }
        }
    }

    return nil
}
