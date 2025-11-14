# mmdbverify

A command-line utility to verify the validity of MaxMind DB (MMDB) files.

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.24%2B-00ADD8?logo=go)](https://golang.org)

## Overview

This utility verifies that a MaxMind DB file is valid by checking:

- **Search tree** - Validates the tree structure used for IP lookups
- **Data section** - Ensures all data is properly formatted and accessible
- **Metadata** - Verifies metadata is well-formed and contains expected fields

If the database is valid, the tool exits silently with status code `0`. If
invalid, an error description is printed to `stderr` and the tool exits with a
non-zero status code.

**Note:** This tool may flag a database as invalid even if it can be parsed.
This occurs when there is unexpected or malformed data in the data section or
metadata.

## Installation

### Binary Releases (Recommended)

Download pre-built binaries from the
[GitHub Releases page](https://github.com/maxmind/mmdbverify/releases).

> **Architecture Guide:**
>
> - `amd64` = x86-64 / x64 (most common for Intel/AMD processors)
> - `arm64` = ARM 64-bit (Apple Silicon, AWS Graviton, Raspberry Pi 4+)
> - `darwin` = macOS

#### Linux

**Using .deb package (Debian/Ubuntu):**

```bash
# Download the .deb file for your architecture from the releases page
sudo dpkg -i mmdbverify_<VERSION>_linux_<ARCH>.deb
```

**Using .rpm package (RedHat/CentOS/Fedora):**

```bash
# Download the .rpm file for your architecture from the releases page
sudo rpm -i mmdbverify_<VERSION>_linux_<ARCH>.rpm
```

**Using tar.gz archive:**

```bash
# Download and extract
tar -xzf mmdbverify_<VERSION>_linux_<ARCH>.tar.gz
sudo mv mmdbverify/mmdbverify /usr/local/bin/
```

#### macOS

```bash
# Download the appropriate file for your Mac:
# - darwin_arm64 for Apple Silicon (M1/M2/M3/M4)
# - darwin_amd64 for Intel Macs
tar -xzf mmdbverify_<VERSION>_darwin_<ARCH>.tar.gz
sudo mv mmdbverify/mmdbverify /usr/local/bin/
```

#### Windows

1. Download the Windows zip file for your architecture from the releases page
2. Extract the zip file
3. Add `mmdbverify.exe` to your PATH or run it directly

**Using PowerShell:**

```powershell
# Extract
Expand-Archive -Path mmdbverify_<VERSION>_windows_<ARCH>.zip -DestinationPath .

# Run
.\mmdbverify\mmdbverify.exe -file C:\path\to\database.mmdb
```

### From Source

```bash
go install github.com/maxmind/mmdbverify@latest
```

### Build Locally

```bash
git clone https://github.com/maxmind/mmdbverify.git
cd mmdbverify
go build
```

## Usage

```bash
# Verify a database file (silent on success)
mmdbverify -file /path/to/GeoIP2-City.mmdb

# Verify with verbose output
mmdbverify -file /path/to/GeoIP2-City.mmdb -verbose
# Verifying /path/to/GeoIP2-City.mmdb...
# /path/to/GeoIP2-City.mmdb is valid

# Check exit code
echo $?
# 0

# Failure (error printed to stderr with filename, non-zero exit code)
mmdbverify -file /path/to/invalid.mmdb
# Error verifying /path/to/invalid.mmdb: invalid database metadata
echo $?
# 1
```

### Command-Line Flags

- `-file` - **Required.** Path to the MaxMind DB file to verify
- `-verbose` - Print verification status messages

## Examples

### Verify GeoIP2 Database

```bash
# Silent mode (default)
mmdbverify -file GeoIP2-City.mmdb

# Verbose mode
mmdbverify -file GeoIP2-City.mmdb -verbose
```

### Verify Multiple Databases

```bash
# Verify all MMDB files with verbose output
for db in *.mmdb; do
    mmdbverify -file "$db" -verbose
done
```

### Use in CI/CD Pipeline

```bash
# Exit immediately if validation fails
mmdbverify -file GeoIP2-Country.mmdb || exit 1
```

## Use Cases

- **Pre-deployment validation** - Verify databases before deploying to
  production
- **Download verification** - Ensure downloaded MMDB files are not corrupted
- **CI/CD pipelines** - Validate databases in automated testing workflows
- **Quality assurance** - Check custom MMDB files created with
  [mmdbwriter](https://github.com/maxmind/mmdbwriter)

## Requirements

- Go 1.24 or later (for building from source)
- MaxMind DB files to verify (GeoIP2, GeoLite2, or custom MMDB files)

## Related Tools

- [mmdbinspect](https://github.com/maxmind/mmdbinspect) - Inspect and query MMDB
  files
- [mmdbwriter](https://github.com/maxmind/mmdbwriter) - Create custom MMDB files
- [mmdbconvert](https://github.com/maxmind/mmdbconvert) - Convert MMDB to
  CSV/Parquet

## License

Copyright (c) 2015-2025 by MaxMind, Inc.

This software is licensed under the Apache License, Version 2.0. See
[LICENSE](LICENSE) for details.

## Support

- **Issues:** [GitHub Issues](https://github.com/maxmind/mmdbverify/issues)
- **MaxMind Support:**
  [https://support.maxmind.com](https://support.maxmind.com)
