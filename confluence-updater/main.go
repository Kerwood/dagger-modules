// A generated module for ConfluenceUpdater functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/confluence-updater/internal/dagger"
	"strings"
)

type ConfluenceUpdater struct {
	// +private
	Fqdn string

	// +private
	User string

	// +private
	Token *dagger.Secret

	// +private
	ConfigPath string

	// +private
	Labels []string

	// +private
	SrcDirectory *dagger.Directory

	// +private
	LogLevel string

	// +private
	CliVersion string
}

func New(
	// +optional
	fqdn string,

	// +optional
	user string,

	// +optional
	token *dagger.Secret,

	// +optional
	// +default="./confluence-updater.yaml"
	configPath string,

	// +optional
	labels string,

	// +optional
	// +defaultPath="/"
	srcDirectory *dagger.Directory,

	// +optional
	// +default="info"
	logLevel string,

	// +optional
	// +default="2.3.1"
	cliVersion string,

) *ConfluenceUpdater {
	var labelSlice []string
	if labels != "" {
		labelSlice = strings.Split(labels, ",")
	}

	c := &ConfluenceUpdater{
		Fqdn:         fqdn,
		User:         user,
		Token:        token,
		ConfigPath:   configPath,
		Labels:       labelSlice,
		SrcDirectory: srcDirectory,
		LogLevel:     logLevel,
		CliVersion:   cliVersion,
	}
	return c
}

// Run Confluence Updater.
func (m *ConfluenceUpdater) RunUpdate(ctx context.Context) (string, error) {
	var args []string

	for _, label := range m.Labels {
		args = append(args, "-l", label)
	}

	output, nil := dag.Container().
		From("ghcr.io/kerwood/confluence-updater:"+m.CliVersion).
		WithMountedDirectory("/workdir", m.SrcDirectory).
		WithWorkdir("/workdir").
		WithEnvVariable("CU_FQDN", m.Fqdn).
		WithEnvVariable("CU_USER", m.User).
		WithSecretVariable("CU_SECRET", m.Token).
		WithEnvVariable("CU_CONFIG_PATH", m.ConfigPath).
		WithEnvVariable("CU_LOG_LEVEL", m.LogLevel).
		WithExec(args, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Stdout(ctx)

	return output, nil
}

// The fully qualified domain name of your Atlassian Cloud. [REQUIRED]
func (m *ConfluenceUpdater) WithFqdn(fqdn string) *ConfluenceUpdater {
	m.Fqdn = fqdn
	return m
}

// Confluence user to login with. [REQUIRED]
func (m *ConfluenceUpdater) WithUser(user string) *ConfluenceUpdater {
	m.User = user
	return m
}

// The token/secret to use. https://id.atlassian.com/manage-profile/security/api-tokens [REQUIRED]
func (m *ConfluenceUpdater) WithToken(token *dagger.Secret) *ConfluenceUpdater {
	m.Token = token
	return m
}

// Config file path. [default: ./confluence-updater]
func (m *ConfluenceUpdater) WithConfigPath(configPath string) *ConfluenceUpdater {
	m.ConfigPath = configPath
	return m
}

// Set label, can be used multiple times.
func (m *ConfluenceUpdater) WithLabel(label string) *ConfluenceUpdater {
	m.Labels = append(m.Labels, label)
	return m
}

// Log Level. [default: info] [possible values: trace, debug, info, warn, error]
func (m *ConfluenceUpdater) WithLogLevel(logLevel string) *ConfluenceUpdater {
	m.LogLevel = logLevel
	return m
}

// Source directory [default: root directory of repository]
func (m *ConfluenceUpdater) WithSourceDirectory(srcDirectory *dagger.Directory) *ConfluenceUpdater {
	m.SrcDirectory = srcDirectory
	return m
}

// Set the version of the confluence-updater cli to use.
func (m *ConfluenceUpdater) WithCliVersion(version string) *ConfluenceUpdater {
	m.CliVersion = version
	return m
}
