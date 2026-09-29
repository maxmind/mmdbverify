# Prereqs

- You must have the [GitHub CLI tool (gh)](https://cli.github.com/) installed,
  in your path, and logged into an account that can push release tags to the
  repo.
- Your environment also must have `bash`, `git`, `go`, and `sed` available.

# Releasing

- Review open issues and PRs to see if anything needs to be addressed before
  release.
- Create a release branch and switch to it. `main` is protected.
- Set the release version and today's date in `CHANGELOG.md`, for example
  `## [1.1.0] - 2026-09-30`. The version must follow
  [Semantic Versioning](https://semver.org/).
  - Update the link references at the end of `CHANGELOG.md`.
- Commit these changes.
- Run `dev-bin/release.sh`. It pushes the branch, then pushes an annotated tag
  with the notes from `CHANGELOG.md`.
- The tag push starts the Release workflow. The authorized releasers get an
  email to review the pending deployment. An authorized releaser must approve
  the run. Then GoReleaser builds the binaries and packages and creates the
  GitHub release with the notes from the tag.
- Verify the release on the GitHub Releases page.
- Make a PR and get it merged.
