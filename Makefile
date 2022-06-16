
# 二进制文件名称
Binary=pupil

# 版本信息
vDir="github.com/soonio/pupil/pkg/utils"

VERSION=v1.0
gitTag=$(shell git log --pretty=format:'%h' -n 1)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date "+%F %T")
gitCommit=$(shell git rev-parse --short HEAD)

ldflags="-s -w -X '${vDir}.version=${VERSION}' -X '${vDir}.gitBranch=${gitBranch}' -X '${vDir}.gitTag=${gitTag}' -X '${vDir}.gitCommit=${gitCommit}' -X '${vDir}.buildDate=${buildDate}'"

# 编译http服务包
GoBuildServe=go build -ldflags ${ldflags} -trimpath -o $(Binary)

# 编译命令包
GoBuildCli=go build -ldflags ${ldflags} -trimpath -o cli ./cli

GoClean=go clean

# 删除编译的包
DeleteBinary=rm -rf v3.new && rm -rf cli

# 发送二进制文件到服务器，并重启服务
define deploy
	scp v3.new $(1):$(2)/v3.new \
        && scp cli $(1):$(2)/cli \
        && rsync -rzvt $(shell pwd)/database $(1):$(2) --delete \
        && ssh -p22 $(1) "supervisorctl stop $(3)" \
        && ssh -p22 $(1) "cd $(2) && mv v3.new v3" \
        && ssh -p22 $(1) "cd $(2) && ./cli migrate:up" \
        && ssh -p22 $(1) "supervisorctl start $(3)"
endef

test:
	$(GOCLEAN)

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GoBuildServe)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GoBuildCli)
	$(call deploy,root@192.168.1.110,/data/pupil/,pupil)

	$(DeleteBinary)

prod:
	$(GOCLEAN)

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GoBuildServe)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GoBuildCli)
	$(call deploy,root@root@192.168.1.110,/data/pupil/,pupil)

	$(DeleteBinary)

build:
	$(GOCLEAN)

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GoBuildServe)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GoBuildCli)