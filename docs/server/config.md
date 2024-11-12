# Server Configuration

The gMountie server configuration file uses YAML format and supports various
options for customizing server behavior, authentication, and volume management.

## Configuration File Structure

The configuration file has three main sections:

- Server configuration
- Authentication configuration
- Volumes configuration

Basic example:

```yaml
server:
  address: 0.0.0.0
  port: 9449
authentication:
  type: none
volumes:
  - name: shared
    path: /shared
```

## Server Options

The `server` section configures the core server settings:

| Option  | Type    | Default      | Description                       |
|---------|---------|--------------|-----------------------------------|
| address | string  | "0\.0\.0\.0" | IP address the server listens on  |
| port    | integer | 9449         | Port number for the gRPC server   |
| metrics | boolean | true         | Enable/disable Prometheus metrics |

Example:

```yaml
server:
  address: 192.168.1.100 # Listen on specific interface 
  port: 8080 # Custom port 
  metrics: false # Disable metrics
```

## Authentication Options

The `auth` section configures user authentication:

| Option | Type   | Required        | Description                             |
|--------|--------|-----------------|-----------------------------------------|
| type   | string | yes             | Authentication type ("none" or "basic") |
| users  | array  | yes (for basic) | List of user credentials                |

### None Authentication

Disables authentication (not recommended for production):

```yaml
auth:
  type: none
```

### Basic Authentication

Enables username/password authentication:

```yaml
auth:
  type: basic
  users:
    - username: admin
      password: admin
    - username: user1
      password: pass123
```

## Volume Configuration

The `volumes` section defines shared directories:

| Option | Type   | Required | Description                       |
|--------|--------|----------|-----------------------------------|
| name   | string | yes      | Unique volume identifier          |
| path   | string | yes      | Absolute path to shared directory |

Example with multiple volumes:

```yaml
volumes:
  - name: documents
    path: /srv/documents
  - name: media
    path: /srv/media
  - name: backup
    path: /srv/backup
```

## Example Configuration

Here's an example configuration file that enables basic authentication and
exposes two volumes:

```yaml
server:
  address: 0.0.0.0
  port: 9449
  metrics: true
authentication:
  type: basic
  users:
    - username: admin
      password: admin
volumes:
  - name: shared
    path: /shared
  - name: private
    path: /private
```
