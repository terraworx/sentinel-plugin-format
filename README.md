# Sentinel Environment Variable Plugin

## Overview

This repository contains a simple Go package named `env` that serves as a plugin for interacting with environment variables. It is designed to be used with the Sentinel framework.

## Configuration

To use this plugin, you need to have Sentinel installed. If you haven't set up Sentinl environment, please visit [Install Sentinel CLI](https://docs.hashicorp.com/sentinel/intro/getting-started/install).

Once Sentinel is installed, you can configure the plugin by adding the following to the Sentinel configuration:

```terraform
import "plugin" "env" {
  source = "https://github.com/terraworx/sentinel-plugin-env/releases/download/v0.0.2/sentinel-env-v0.0.2-linux-amd64"
}
```

## Usage

### Importing the Plugin

```sentinel
import "env"
```

### Available Functions

#### 1. `get(key string) `

This function retrieves the value of the specified environment variable.

Example:

```sentinel
import "env"

colorTerm = env.get("COLORTERM") // truecolor

main = rule {
    colorTerm is "truecolor"
}
```

#### 2. `list()`

This function returns a map of all environment variables.

Example:

```sentinel
import "env"

main = rule {
    env.list().COLORTERM is "truecolor"
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- [HashiCorp Sentinel SDK](https://github.com/hashicorp/sentinel-sdk)

---
