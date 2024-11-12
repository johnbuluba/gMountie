# Server CLI Options

The gMountie server can be started using the `serve` command with various
options to customize its behavior.

## Basic Usage

The basic syntax for starting the server is:

```bash
gmountie serve [flags]
```

## Command Flags

| Flag      | Short | Default                        | Description                |
|-----------|-------|--------------------------------|----------------------------|
| --config  | -c    | ~/.config/gmountie/server.yaml | Path to configuration file |
| --verbose | -v    | false                          | Enable verbose logging     |

## Configuration File

If no configuration file is specified, gMountie will:

1. Look for a config file at `~/.config/gmountie/server.yaml`
2. If not found, create a default configuration file at that location
3. Use the following default configuration:

```yaml
server:
  address: 127.0.0.1
  port: 9449
  metrics: true
auth:
  type: basic
  users:
    - username: admin
      password: admin
  ```

## Examples

1. Start server with default configuration:
   ```bash
   gmountie serve
   ```

2. Start server with custom config file:
   ```bash
   gmountie serve -c /path/to/config.yaml
   ```

3. Start server with verbose logging:
   ```bash
   gmountie serve -v
   ```

## Security Considerations

1. Default configuration uses basic authentication
2. Production deployments should:
    - Use a custom configuration file
    - Change default credentials
    - Consider using TLS
    - Bind to specific network interfaces

## See Also
- [Server Configuration](server/config.md) - Detailed configuration file options
