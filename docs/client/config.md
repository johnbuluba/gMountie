# Client Configuration

The gMountie client configuration file uses YAML format and supports various
options for connecting to servers, authentication, and mounting volumes.

## Configuration File Structure

The configuration file has three main sections:

- Server connection settings
- Authentication configuration
- Mount configuration

Basic example:

```yaml
server:
  address: 127.0.0.1
  port: 9449
  tls: false
auth:
  type: basic
  username: admin
  password: admin
mount:
  type: single
  volume: shared
  path: /mnt/shared
```

## Server Options

The `server` section configures the connection to the gMountie server:

| Option  | Type    | Default   | Description                   |
|---------|---------|-----------|-------------------------------|
| address | string  | 127.0.0.1 | Server IP address or hostname |
| port    | integer | 9449      | Server port number            |
| tls     | boolean | false     | Enable/disable TLS encryption |

Example:

```yaml
server:
  address: 192.168.1.100  # Remote server address
  port: 8080              # Custom port
  tls: true               # Enable TLS
```

## Authentication Options

The `auth` section configures client authentication:

| Option   | Type   | Required        | Description                             |
|----------|--------|-----------------|-----------------------------------------|
| type     | string | yes             | Authentication type ("none" or "basic") |
| username | string | yes (for basic) | Username for basic auth                 |
| password | string | yes (for basic) | Password for basic auth                 |

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
  username: admin
  password: admin
```

## Mount Configuration

The `mount` section defines how volumes are mounted. There are two mount types:

1. Single volume mount
2. VFS (Virtual File System) mount

### Single Volume Mount

Mounts a single volume at a specified path:

| Option | Type   | Required | Description            |
|--------|--------|----------|------------------------|
| type   | string | yes      | Must be "single"       |
| volume | string | yes      | Volume name to mount   |
| path   | string | yes      | Local mount point path |

Example:

```yaml
mount:
  type: single
  volume: documents
  path: /home/user/documents
```

### VFS Mount

Mounts multiple volumes under a single mount point:

| Option    | Type     | Required | Description                       |
|-----------|----------|----------|-----------------------------------|
| type      | string   | yes      | Must be "vfs"                     |
| path      | string   | yes      | Base mount point path             |
| mount_all | boolean  | no       | Mount all available volumes       |
| volumes   | []string | no       | List of specific volumes to mount |

Example:

```yaml
mount:
  type: vfs
  path: /mnt/gmountie
  mount_all: false
  volumes:
  - documents
  - media
  - backup
```

## Complete Configuration Examples

### Single Volume Mount Example

```yaml
server:
  address: 192.168.1.100
  port: 9449
  tls: false
auth:
  type: basic
  username: admin
  password: admin
mount:
  type: single
  volume: documents
  path: /home/user/documents
```

### VFS Mount Example

```yaml
server:
  address: 192.168.1.100
  port: 9449
  tls: false
auth:
  type: basic
  username: admin
  password: admin
mount:
  type: vfs
  path: /mnt/gmountie
  mount_all: false
  volumes:
    - documents
    - media
    - backup
```

