# Skate Ipsum

😱

```bash
curl -I http://skateipsum.com/
HTTP/1.1 404 Not Found
```

An attempt to get the "gnarlier lorem ipsum generator" back. And ... to experiment with any release and deployment method for a Go application on the planet.

## Using the CLI

```markdown
A gnarlier ipsum generator

Usage:
skate [flags]
skate [command]

Available Commands:
completion Generate the autocompletion script for the specified shell
help Help about any command
server Start a skate ipsum api server

Flags:
-h, --help help for skate-ipsum
-p, --paragraphs int number of paragraphs (default 10)

Use "skate [command] --help" for more information about a command.
```

## Install the CLI

### Binaries and Packages

- Head on over to https://github.com/madflow/skate-ipsum/releases/ and download the latest release.

### Go install

- You can [install Go](https://golang.org/dl/) and build from source (requires Go 1.17+):

```bash
go install github.com/madflow/skate-ipsum@latest
```

### Run the CLI using Docker

```bash
docker run -it --rm ghcr.io/madflow/skate-ipsum:latest
```

### Homebrew

```bash
brew install madflow/skate-ipsum/skate-ipsum
```

## Deploy the ipsum server

### Deploy via Docker

```bash
docker run -it --rm -p 6969:6969 ghcr.io/madflow/skate-ipsum:latest server
```

### Serverless Function on Netlify

- A serverless function is located in `./netlify/functions/skate-ipsum/`
- An example deployment can be found here: https://graceful-sorbet-30de07.netlify.app/ with the serverless function https://graceful-sorbet-30de07.netlify.app/api/skate-ipsum

### Deploy to Render

- Just use the Docker image and the command/entrypoint `/usr/local/bin/skate-ipsum server`
- Example: https://skate-ipsum.onrender.com/
