# Deployment Procedure

## Release Checklist

Before releasing a new version:

1. **Test the build locally**
   ```bash
   go build -o dist/ccfix.exe -ldflags="-s -w" main.go
   ```

2. **Run tests**
   ```bash
   go test -v
   ```

3. **Test the executable manually** - copy prefixed text, run ccfix.exe, verify clipboard is fixed

## Release Steps

### 1. Update CHANGELOG.md

Add a new section at the top (below the header):

```markdown
## [X.Y.Z] - YYYY-MM-DD

### Added
- New features

### Fixed
- Bug fixes

### Changed
- Changes to existing functionality
```

### 2. Update README.md

Update the version number in the badge line:

```markdown
> **Latest Version**: X.Y.Z | [See What's New](CHANGELOG.md)
```

### 3. Commit and Tag

```bash
# Stage and commit your changes (if not already committed)
git add .
git commit -m "Description of changes"

# Commit the version bump
git add CHANGELOG.md README.md
git commit -m "Release X.Y.Z"

# Create the tag (no 'v' prefix - required for GitHub Actions)
git tag -a X.Y.Z -m "Release X.Y.Z"

# Push commits and tag
git push
git push origin X.Y.Z
```

## Tag Format

- Correct: `1.0.0`, `1.1.0`, `2.0.0`
- Wrong: `v1.0.0`, `1.0`, `release-1.0.0`

## What Happens Automatically

When you push a correctly formatted tag, GitHub Actions will:

1. Check out the code
2. Build the Windows executable
3. Run tests
4. Create a GitHub Release with attached binary: `ccfix.exe`

## Verifying the Release

After pushing the tag:

1. Check the Actions tab: https://github.com/fanis/claude-clipboard-fix/actions
2. Verify the release was created: https://github.com/fanis/claude-clipboard-fix/releases
