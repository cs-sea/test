linux:
	@sh shell/linux.sh

mac:
	@sh shell/mac.sh

test:
	go test -v ./...

auto-tag:
	@sh shell/auto_tag.sh "$(msg)"