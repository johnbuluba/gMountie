<div align="center">
  <img src="assets/logo-full-cropped-min.png" alt="gMountie Logo" width="300"/>
  <p><i>Simplifying remote filesystem mounting with FUSE and gRPC</i></p>
</div>

![coverage](https://raw.githubusercontent.com/johnbuluba/gMountie/badges/.badges/main/coverage.svg)

## Overview

gMountie is a robust filesystem mounting utility designed to seamlessly connect remote directories to your local system. Built on modern technology with gRPC and FUSE, it offers a reliable and efficient solution for managing remote filesystem mounts, complete with a user-friendly desktop application.

### Key Features

- **Client/Server Architecture**: Enables secure and efficient remote filesystem access
- **FUSE Implementation**: Provides a flexible userspace filesystem interface
- **gRPC Communication**: Ensures fast and reliable communication between client and server using unary operations
- **Desktop Application**: Offers an intuitive graphical interface for easy management
- **Cross-platform Compatibility**: Supports Linux (macOS and Windows coming soon)

## Quick Start Guide

1. **Install gMountie**
   ```bash
   # Installation instructions coming soon
   ```
2. **Start the Server**
   ```bash
   ./gmountie-server
   ```

3. **Mount a Remote Directory**
   ```bash
   ./gmountie-client mount /path/to/mount/point
   ```

## Configuration

gMountie offers flexible configuration options:

1. Edit the \`config.yaml\` file for persistent settings
2. Use environment variables for runtime configuration
3. Apply command-line flags for temporary adjustments

Refer to our [configuration guide](docs/configuration.md) for detailed options.

## Troubleshooting

If you encounter any issues:

1. Check the server and client logs for error messages
2. Verify your network connection and firewall settings
3. Ensure you have the necessary permissions for mounting
4. Consult our [troubleshooting guide](docs/troubleshooting.md) for common solutions

If problems persist, please [open an issue](https://github.com/yourusername/gmountie/issues) with a detailed description.

## Contributing

We welcome contributions to gMountie! If you'd like to help improve the project:

1. Fork the repository
2. Create a feature branch
3. Make your changes and add tests
4. Ensure all tests pass
5. Submit a pull request

Please review our [contribution guidelines](CONTRIBUTING.md) for more information.

## License

gMountie is released under the Apache License. See the [LICENSE](LICENSE) file for details.

---

gMountie: Making remote filesystem mounting simple and efficient.
