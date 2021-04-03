LD_FLAGS = -X .com/cosmos/cosmos-sdk/version.Name=MetaId \
	-X github.com/cosmos/cosmos-sdk/version.AppName=mbcorecrd \
	-X github.com/cosmos/cosmos-sdk/version.Version=0.1.0
BUILD_FLAGS = -mod=readonly -ldflags='$(LD_FLAGS)'

.PHONY: all
all: build

.PHONY: build
build:
	go build $(BUILD_FLAGS) ./cmd/mbcorecrd