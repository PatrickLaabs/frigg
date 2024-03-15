---
name: 🚋 Release cycle tracking
about: Create a new release cycle tracking issue for a Frigg release
about: "[Only for release team lead] Create an issue to track tasks for a Frigg release."
title: Tasks for v<release-tag> release cycle
labels: ''
assignees: ''

---

## Tasks

* [ ] Check Milestone for the specific Release and if everything is completed, or if anything needs to be added or removed.
* [ ] Create a new feature-branch for the release-preperation
* [ ] Bump Versions in consts Package
* [ ] Update go version
* [ ] Update go mod
* [ ] Update go version for goreleaser
* [ ] Update Versions referenced in the README.md
* [ ] Test everything :) 
* [ ] If everything is fine, push local changes to the prevs. created feature-branch
* [ ] git tag x.x.x && git push --tags