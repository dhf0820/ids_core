USER ?= dhf0820
NS ?= dhf0820
TAG ?= 250714.0
TEST ?= dhf0820
PROD ?= vertisoft
ARC ?= amd64
ARCH ?= $(ARC)
IMG_NAME = ids_core 
VERSION ?= $(TAG)
IMAGE_NAME = ids_core
BINARY = $(IMAGE_NAME)
IMAGE_TEST_NAME ?= $(IMAGE_NAME)_test
MAC_IMAGE_NAME= $(IMAGE_NAME)
LINUX_IMAGE_NAME= $(IMAGE_NAME)
LINUX_VERSION_NAME=$(IMAGE_NAME)
BINARY_NAME=$(IMAGE_NAME)
BINARY_UNIX=$(BINARY_NAME)
DOCKER_NAME=$(IMAGE_NAME)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_MAC=$(IMAGE_NAME)
MAC_IMAGE_NAME= $(BINARY_MAC)


.PHONY: all api server client cert pb1

all: server client


#api/api.pb.go:
#protoc -I ./ --proto_path=./ --go_out=./ pkg/proto/*.proto
api:
#	@protoc -I ./protobufs/ \
		--proto_path=./ \
		--go_out=plugins=grpc:./ \
		./protobufs/*.proto


#	@protoc -I ./protobufs \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=./ \
		--go_out=plugins=grpc:./ \
		./protobufs/*.proto



#api: api/api.pb.go api/api.pb.gw.go api/api.swagger.json ## Auto-generate grpc go sources


dep: ## Get the dependencies
	@go get -v -d ./...

tidy: # add all new includes
	@go mod tidy -compat=1.17

build:
	$(GOBUILD) -o $(BINARY_NAME)

mac:
	CGO_ENABLED=0 $(GOBUILD) -o $(MAC_IMAGE_NAME) -v


	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(IMAGE_NAME) 
	docker build -t $(TEST)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	#docker push $(TEST)/$(DOCKER_NAME):$(VERSION)

linux_test:
	echo "Building for linux_test"
	ARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(IMAGE_NAME) 
	docker build -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(IMAGE_NAME):$(VERSION)


linux_local:
	echo "Building for linux_local"
	CGO_ENABLED=0  $(GOBUILD) -o $(IMAGE_NAME) 
	docker build --no-cache -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
	#docker push $(NS)/$(IMG_NAME):$(VERSION)

mac_test:
	echo "Building for mac_test"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(IMAGE_NAME)
	docker build --no-cache -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	docker push $(TEST)/$(IMAGE_NAME):$(VERSION)
	
mac_local:
	echo "Building for mac_local"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(IMAGE_NAME)_$(ARCH)
	docker build -t $(TEST)/$(IMAGE_NAME):mac$(VERSION) -f Dockerfile_$(ARCH) .
	#docker push $(TEST)/$(IMAGE_NAME):mac$(VERSION)


test_linux:
	echo "Building for test_linux"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOHOSTARCH=amd64 GOHOSTOS=linux $(GOBUILD) -o $(IMAGE_NAME) 
	docker build -t $(TEST)/$(IMAGE_NAME):$(VERSION) --no-cache -f Dockerfile .
	docker push $(TEST)/$(IMAGE_NAME):$(VERSION)

test_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(IMG_NAME) 
	docker build -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile_amd64 .
	docker push $(TEST)/$(IMAGE_NAME):$(VERSION)

prod:
	@echo "line: 109 Building for prod"
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v
	#@echo "line: 111 Docker Build"
	docker build -t $(PROD)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	#@echo "Line:113  creating buildImage " 
	buildImage =  docker.io/$(PROD)/$(DOCKER_NAME):$(VERSION)
	#@echo "Line: 115 " buildImage 
	#imageName = "docker.io/$(PROD)/$(DOCKER_NAME):$(VERSION)"
	#@echo "line: 117  Pushing image: " $(buildImage)
	docker push $(buildImage)

	#@docker push $(PROD)/$(DOCKER_NAME):$(VERSION)

# test:
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v
# 	docker build -t $(TEST)/$(DOCKER_NAME):$(VERSION) -f Dockerfile .
# 	docker push $(TEST)/$(DOCKER_NAME):$(VERSION)

# prod:
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v
# 	docker build -t $(PROD)/$(DOCKER_NAME):$(VERSION) -f Dockerfile .
# 	docker push $(PROD)/$(DOCKER_NAME):$(VERSION)


release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v
	docker build -t $(NS)/$(DOCKER_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(DOCKER_NAME):$(VERSION)

it:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v
	docker build -t $(NS)/$(DOCKER_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(DOCKER_NAME):$(VERSION)

delivery:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v
	docker build -t $(NS)/$(LINUX_IMAGE_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(LINUX_IMAGE_NAME):$(VERSION)

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_IMAGE_NAME) -v

docker-build:
	docker build -t $(NS)/$(LINUX_IMAGE_NAME):$(VERSION) -f Dockerfile .

docker-push: # push to docker
	docker push $(NS)/$(LINUX_IMAGE_NAME):$(VERSION)

client: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(PB_OUT) $(API_REST_OUT) $(API_SWAG_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run-server:
	go run main.go -port 8080

run-server-tls:
	go run cmd/server/main.go -port 9901 -tls server

run-server-mutual-tls:
	go run cmd/server/main.go -port 7777 -tls mutual

run-client-do3:
	go run src/client/main.go  -address docker-test.vertisoft.com -port 8080

run-client:
	go run src/client/main.go  -address localhost -port 9001

run-client-do-test:
	go run src/client/main.go  -address 161.35.229.137 -port 30001

run-client-tls:
	go run cmd/client/main.go  -address 0.0.0.0:7777 -tls server

run-client-mutual-tls:
	go run cmd/client/main.go  -address 0.0.0.0:7777 -tls mutual

cert:
	cd cert; ./gen.sh; cd ..


# d38f3380a1de
