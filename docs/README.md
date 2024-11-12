<div align="center">
  <img class="logo" src="assets/logo-full-cropped.png" alt="gMountie Logo" width="300"/>
  <h1>gMountie: Your Filesystem's Best Friend 🤠</h1>
  <p><i>Because remote filesystems shouldn't feel so... remote</i></p>
</div>

![coverage](https://raw.githubusercontent.com/johnbuluba/gMountie/badges/.badges/master/coverage.svg)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## What's This All About? 🤔

Ever wished you could access your remote files as easily as if they were right there on your computer? That's exactly what gMountie does! It's like having a really efficient courier service for your files, but instead of waiting days for delivery, everything happens instantly.

Built with modern tech (FUSE and gRPC) and a lot of ❤️, gMountie makes remote filesystems feel local. No more complicated mounting procedures or slow network transfers - just smooth, efficient file access.

## Features That'll Make You Smile 😊

- **Lightning Fast**: Built with gRPC for speedy communication
- **Rock Solid**: Extensive test coverage ensures reliability
- **User Friendly**: Both CLI and GUI options available
- **Secure**: Built-in authentication and encryption options
- **Modern**: Uses FUSE for flexible filesystem operations
- **Cross-Platform**: Linux support (macOS coming soon!)

## Installation 📦

Detailed installation instructions are available in our [documentation](https://gmountie.docs.com).

## Documentation 📚

Our comprehensive documentation covers everything you need to know:

- [Server Configuration](docs/server/config.md)
- [Client Configuration](docs/client/config.md)
- [Server CLI Reference](docs/server/cli.md)
- [Client CLI Reference](docs/client/cli.md)

## Architecture 🏗️

gMountie uses a client-server architecture:

```
Client (Your Computer) <-> gRPC <-> Server (Remote System)
↓                                   ↓
FUSE Mount                          Real Filesystem
```

## Contributing 🤝

We love contributions! Whether it's:

- 🐛 Bug reports
- 💡 Feature suggestions
- 📝 Documentation improvements
- 🔧 Code contributions

Check out our [Contributing Guide](CONTRIBUTING.md) to get started.

## Support & Community 💬

- 📫 [GitHub Issues](https://github.com/johnbuluba/gMountie/issues) for bug reports and feature requests
- ⭐ [Star us on GitHub](https://github.com/johnbuluba/gMountie) to show your support
- 💖 [Become a sponsor](https://github.com/sponsors/johnbuluba) to support development

## License 📜

gMountie is proudly open source, licensed under the [Apache License 2.0](LICENSE).

---

<div align="center">
  <i>Happy Mounting! 🎉</i>
</div>


