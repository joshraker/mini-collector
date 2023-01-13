FROM golang:1.9

WORKDIR /go/src/github.com/aptible/mini-collector

ENV PYENV_VERSION=3.6
ENV PROTOC_VERSION="3.5.1"
ENV PATH="/tmp/protoc/bin:$PATH"

COPY .codegen/requirements.txt .codegen/requirements.txt
RUN apt-get update && \
  apt-get install -y python3 python3-pip unzip && \
  pip3 install --user -r .codegen/requirements.txt && \
  mkdir -p /tmp/protoc && \
  cd /tmp/protoc && \
  PROTOC_ARCHIVE="protoc-${PROTOC_VERSION}-linux-$(uname -m | sed -E 's/_?64/_64/').zip" && \
  curl -fsSLO "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ARCHIVE}" && \
  unzip "$PROTOC_ARCHIVE" && \
  rm "$PROTOC_ARCHIVE" && \
  go get -u github.com/twitchtv/retool && \
  go get -u github.com/golang/dep/cmd/dep

# I think tools.json is overwritten when retool is installed so we have to wait for it to be installed first.
COPY tools.json ./
RUN retool sync

COPY ./ ./
RUN make deps
