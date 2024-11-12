# gMountie Documentation

## Quick Start Guide

### Server Setup

1. Create a basic server configuration file (config.yaml):

    ```yaml
    server:
      address: 0.0.0.0
      port: 9449
      metrics: true
    auth:
      type: basic
      users:
      - username: admin
        password: admin
    volumes:
    - name: shared
      path: /path/to/shared/directory
    ```

1. Start the server:
    ```bash
    ./gMountie serve -c config.yaml
    ```

### Client Setup

1. Create a client configuration file (client.yaml):

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
      path: /home/john/gMountie
    ```

1. Mount the remote filesystem:
    ```bash
    ./gMountie mount -c client.yaml
    ```

### Using the Desktop Application

1. Launch the desktop application:
    ```bash
    ./gMountie desktop
    ```
2. Connect to your server using the GUI interface
3. Select volumes to mount
4. Choose mount points
5. Click "Mount" to connect

## Configuration Reference

For detailed configuration options, see:

- [Server Configuration](server/config.md)
- [Client Configuration](client/config.md)

