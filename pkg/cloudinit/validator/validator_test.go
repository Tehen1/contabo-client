package validator

import (
    "testing"
)

func TestValidateSSHKey(t *testing.T) {
    validator := NewCloudInitValidator()

    tests := []struct {
        name    string
        key     string
        wantErr bool
    }{
        {
            name:    "valid ssh key",
            key:     "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0g+ZTxC user@host",
            wantErr: false,
        },
        {
            name:    "empty key",
            key:     "",
            wantErr: true,
        },
        {
            name:    "invalid format",
            key:     "not-a-valid-key",
            wantErr: true,
        },
        {
            name:    "oversized key",
            key:     string(make([]byte, 5000)),
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validator.ValidateSSHKey(tt.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateSSHKey() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestValidatePackages(t *testing.T) {
    validator := NewCloudInitValidator()

    tests := []struct {
        name     string
        packages []string
        wantErr  bool
    }{
        {
            name:     "valid packages",
            packages: []string{"nginx", "docker", "vim"},
            wantErr:  false,
        },
        {
            name:     "empty packages",
            packages: []string{},
            wantErr:  false,
        },
        {
            name:     "too many packages",
            packages: make([]string, 51),
            wantErr:  true,
        },
        {
            name:     "invalid package name",
            packages: []string{"nginx", "invalid package name"},
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validator.ValidatePackages(tt.packages)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidatePackages() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
