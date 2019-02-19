
.PHONY: release-test
release-dry:
		GO111MODULE=on goreleaser --snapshot --skip-publish --rm-dist

.PHONY: release-dry
release-dry:
		GO111MODULE=on goreleaser --skip-publish --rm-dist
