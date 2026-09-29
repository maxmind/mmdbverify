# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-09-29

### Changed

- Go 1.26 or later is now required to build from source.
- mmdbverify now rejects extended type bytes 0 and 250 through 255. The MaxMind
  DB spec does not define them. Before, they decoded as other types.
- mmdbverify now rejects a data pointer that points into the middle of a field.
- mmdbverify now also checks the metadata section. Its pointers must point to
  the start of a field, and all data after the metadata map must be valid.

### Fixed

- mmdbverify no longer rejects a search-tree record that points to a value
  nested in another data record. The MaxMind DB spec permits these records, and
  mmdbwriter can write them.

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

[unreleased]: https://github.com/maxmind/mmdbverify/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/maxmind/mmdbverify/releases/tag/v1.1.0
[1.0.0]: https://github.com/maxmind/mmdbverify/releases/tag/v1.0.0
