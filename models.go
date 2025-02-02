package manssh

import (
	"fmt"

	"github.com/StephenBrown2/ssh_config"
)

// HostConfig struct include alias, connect string and other config
type HostConfig struct {
	// Alias alias
	Alias string
	// Path found in which file
	Path string
	// PathMap key is file path, value is the alias's hosts
	PathMap map[string][]*ssh_config.Host
	// OwnConfig own config
	OwnConfig map[string]string
	// ImplicitConfig implicit config
	ImplicitConfig map[string]string
}

// NewHostConfig new HostConfig
func NewHostConfig(alias, path string, host *ssh_config.Host) *HostConfig {
	return &HostConfig{
		Alias:          alias,
		Path:           path,
		PathMap:        map[string][]*ssh_config.Host{path: {host}},
		OwnConfig:      map[string]string{},
		ImplicitConfig: map[string]string{},
	}
}

// ConnectionStr return the connection string
func (hc *HostConfig) ConnectionStr() string {
	if !hc.Display() {
		return ""
	}

	var (
		user, hostname, port string
		ok                   bool
	)

	if user, ok = hc.OwnConfig["user"]; !ok {
		user = hc.ImplicitConfig["user"]
		delete(hc.ImplicitConfig, "user")
	} else {
		delete(hc.OwnConfig, "user")
	}

	if hostname, ok = hc.OwnConfig["hostname"]; !ok {
		delete(hc.ImplicitConfig, "hostname")
		hostname = hc.ImplicitConfig["hostname"]
	} else {
		delete(hc.OwnConfig, "hostname")
	}

	if port, ok = hc.OwnConfig["port"]; !ok {
		port = hc.ImplicitConfig["port"]
		delete(hc.ImplicitConfig, "port")
	} else {
		delete(hc.OwnConfig, "port")
	}

	return fmt.Sprintf("%s@%s:%s", user, hostname, port)
}

// Display Whether to display connection string
func (hc *HostConfig) Display() bool {
	hostname := hc.OwnConfig["hostname"]
	if hostname == "" {
		hostname = hc.ImplicitConfig["hostname"]
	}

	return hostname != ""
}
