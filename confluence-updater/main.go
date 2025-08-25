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
	"fmt"
)

type ConfluenceUpdater struct {
	// +private
	SrcDirectory *dagger.Directory

	// +private
	ApiUser string

	// +private
	ApiToken *dagger.Secret

	// +private
	ApiFqdn string

	// +private
	ApiLogLevel string
}

func New(
	// +optional
	// +defaultPath="/"
	srcDirectory *dagger.Directory,

	// +optional
	user string,

	// +optional
	token *dagger.Secret,

	// +optional
	fqdn string,

	// +optional
	// +default="info"
	logLevel string,
) *ConfluenceUpdater {
	c := &ConfluenceUpdater{
		SrcDirectory: srcDirectory,
		ApiUser:      user,
		ApiToken:     token,
		ApiFqdn:      fqdn,
		ApiLogLevel:  logLevel,
	}
	return c
}

// Run Confluence Updater.
func (m *ConfluenceUpdater) RunUpdate(ctx context.Context) (string, error) {

	if m.ApiFqdn == "" {
		return "", fmt.Errorf("Confluence FQDN not set.")
	}

	if m.ApiUser == "" {
		return "", fmt.Errorf("Confluence user not set.")
	}

	//args := []string{"-u", m.ApiUser, "--fqdn", m.ApiFqdn, "-c", "/workdir/confluence-updater.yaml", "--log-level", m.ApiLogLevel}
	args := []string{"-u", m.ApiUser, "--fqdn", m.ApiFqdn, "--log-level", m.ApiLogLevel}

	output, nil := dag.Container().
		From("ghcr.io/kerwood/confluence-updater:2.3.1").
		WithMountedDirectory("/workdir", m.SrcDirectory).
		WithWorkdir("/workdir").
		WithSecretVariable("CU_SECRET", m.ApiToken).
		WithExec(args, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Stdout(ctx)

	return output, nil
}

// The fully qualified domain name of your Atlassian Cloud. [REQUIRED]
func (m *ConfluenceUpdater) Fqdn(fqdn string) *ConfluenceUpdater {
	m.ApiFqdn = fqdn
	return m
}

// Confluence user to login with. [REQUIRED]
func (m *ConfluenceUpdater) User(user string) *ConfluenceUpdater {
	m.ApiUser = user
	return m
}

// The token/secret to use. https://id.atlassian.com/manage-profile/security/api-tokens [REQUIRED]
func (m *ConfluenceUpdater) Token(token *dagger.Secret) *ConfluenceUpdater {
	m.ApiToken = token
	return m
}

// Log Level. [default: info] [possible values: trace, debug, info, warn, error]
func (m *ConfluenceUpdater) LogLevel(LogLevel string) *ConfluenceUpdater {
	m.ApiLogLevel = LogLevel
	return m
}

// Source directory [default: root directory of repository]
func (m *ConfluenceUpdater) SourceDirectory(srcDirectory *dagger.Directory) *ConfluenceUpdater {
	m.SrcDirectory = srcDirectory
	return m
}
