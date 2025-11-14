# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-11-14

### Added

- Initial official release of mmdbverify
- Verify MaxMind DB file validity
- Search tree validation
- Data section validation
- Metadata validation and sanity checks
- Command-line interface with `-file` flag
- Cross-platform builds (Linux, macOS, Windows for amd64 and arm64)
- Debian and RPM package support
- Clear error messages to stderr
- Zero exit code on success, non-zero on failure

[unreleased]: https://github.com/maxmind/mmdbverify/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/maxmind/mmdbverify/releases/tag/v1.0.0
