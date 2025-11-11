# Confluence Updater

https://github.com/Kerwood/confluence-updater

## Module Configuration
Install the `confluence-updater` module, making sure to replace the version number with the latest available release.

```sh
dagger install github.com/kerwood/dagger-modules/confluence-updater@v0.1.0
```

Implement the module using the example below.
```go
func (m *MyModule) SyncToConfluence(
	ctx context.Context,
	token *dagger.Secret,
	// +optional
	// +defaultPath="/"
	srcDirectory *dagger.Directory,
) string {
	output, err := dag.ConfluenceUpdater().
		WithFqdn("https://your-domain.atlassian.net"). // Required
		WithUser("your-user@example.org").             // Required
		WithToken(token).                              // Required
		WithSourceDirectory(srcDirectory).             // Required
		WithConfigPath("./confluence-updater.yaml").   // Default value
		WithLabel("first-label").                      // Can be used multiple times
		WithLabel("second-label").
		WithLogLevel("info").                          // Default value
		WithCliVersion("2.3.2").                       // Default value
		RunUpdate(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return output
}
```

Export your Atlassian API token as an environment variable and run the function.
```sh
export TOKEN=<your-atlassian-token>
dagger -c 'sync-to-confluence env://TOKEN'
```

## Install as Blueprint

Blueprints are a relatively new feature in Dagger that let you run modules without any boilerplate code and cannot be used alongside regular Dagger modules.
Running the command below generates a `dagger.json` file that defines the `confluence-updater` module as a blueprint,
skipping the usual boilerplate and making its functions directly accessible from the Dagger CLI.

Be sure to replace the version number with the latest available release.
```sh
dagger init --blueprint=github.com/kerwood/dagger-modules/confluence-updater@v0.1.0
```


Export your Atlassian API token as an environment variable and run the Dagger CLI.

*Note: `with-source-drectory` is not required when using blueprints, it defaults to the repository root directory.*
```sh
export TOKEN=<your-atlassian-token>
dagger -c 'with-fqdn https://your-domain.atlassian.net | with-user your-user@example.org | with-token env://TOKEN | run-update'
```
