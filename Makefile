
.PHONY: release-dry
release-dry:
		GO111MODULE=on goreleaser --snapshot --skip-publish --rm-dist