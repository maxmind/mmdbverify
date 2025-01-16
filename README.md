MaxMind DB Verifier
-------------------

This is a utility to verify that a MaxMind DB is valid. The search tree,
data section, and metadata are all checked for validity and sanity. If the
database cannot be parsed or is otherwise invalid, a short description will
be sent to `stderr` and the utility will exit with a non-zero exit code. If
successful, no output will be printed and an exit code of zero will be
returned.

Note that this tool may flag a database as invalid even though it is
completely parseable. This may happen if there is unexpected data in the
data section or metadata that is not well-formed.

Usage
=====

Required:

* -file=[FILENAME] - The path to the database to test.

Copyright and License
=====================

This software is Copyright (c) 2015-2025 by MaxMind, Inc.

This is free software, licensed under the Apache License, Version 2.0.
