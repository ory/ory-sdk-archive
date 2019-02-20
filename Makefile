
.PHONY: release-test
release-test:
		GO111MODULE=on goreleaser --snapshot --skip-publish --rm-dist

.PHONY: release-dry
release-dry:
		GO111MODULE=on goreleaser --skip-publish --rm-dist

.PHONY: release
release:
		GO111MODULE=on goreleaser --rm-dist
