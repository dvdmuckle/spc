# Releasing

While there is a decent amount of automation in terms of creating a new release, there's still a couple manual actions that need to happen to both trigger the automation and make sure said automation functions correctly.

## Versioning

This repo adheres to semver, or at least tries to. Patches, or 1.0.X releases, do not break anything nor add new major features. This may be tweaking an existing command in terms of its output, or tweaking a command in a way that does not change any current functionality. In other words, the changes are either purely additive, or are done in such a way that a user can use the tool in the same manner as before. A minor version bump, or 1.X.0 release, may add a new feature, such as a new command. Still, nothing changes existing commands in such a way to break functionality of older features. A major release, X.0.0, can break functionality in adding new features or changing existing features.

That all being said, some releases come down to gut feeling in terms of what kind of version bump is necessary. So semver is more of a guideline than an actual rule.

## Bumping Package Specs

First and foremost, Linux package build specs need to have versions bumped to the new version. The following make targets need to be run. This can be done either on the `dev` branch or in a branch that is ultimately merged into the `dev` branch.

```bash
make deb-bump-version
make rpm-bump-version
```

These will add some new sections to the changelog for both `spc.spec` and `debian/changelog`. These still have to be manually changed after the make targets are run. `spc.spec` needs to have the version manually changed at the `%tag` area, with the `Release` field changed back from `2%{?dist}` to `1%{?dist}`. The changelog at the bottom of `spc.spec` should also be edited accordingly, both with the right version and any changes from the last version. These changes should also be made in a similar manner to `debian/changelog` by changing the version in the header for the changes and describing the changes.

If the version is not bumped for Linux package changes, the builds for these packages will fail automatically. Fixing this would mean redoing the release process.

## Merging

The only method for commiting to the `main` branch is via PR from the `dev` branch. The changes in the dev branch should include all the actual changes in the proceeding version, as well as the version bumps for package specs as described above. Once a PR has been merged into the `main` branch, the changes should be pulled in locally, and tagged. This tag takes the form of `X.X.X` where each is an integer. Note the lack of a preceeding `v` in the tag. This is due to Linux packaging constraints, specifically with Fedora packages. A tag with the `v` as a prefix will be created by automation for Golang module purposes. Once the tag has been created with `git tag`, it can be pushed with `git push --tags`.

## Publishing A Release

Once the tag has been pushed, it will kick off drafting a new release. If everything looks good with the release, it can then be published. This will kick off creating archives with the `spc` binary for all supported operating systems and architectures, as well as send build data to Copr and Launchpad for Linux package builds. Of note, the Copr job watches build progress in Copr, and will fail if the Copr build fails. The Launchpad job is a fire and forget, if the Launchpad build fails there will be a notification via email.

## Recovering From A Failed Release

In the case of a failed release, for example failing to bump the version for Linux packages, the release process will need to be restarted. The failed release will need to be deleted, as well as its associated tag.

```bash
git tag -d <failed release tag>
```

From there, the release process can be restarted with creating a new PR into `main`, merging the PR, tagging the merge commit in `main` with the tag, and pushing the tag. This will also recreate the release draft, which can then be published.
