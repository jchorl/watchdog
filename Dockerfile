FROM jchorl/appengine-go

RUN apt-get update && apt-get install -y unzip && \
    mkdir /proto && \
    wget -P /proto https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip /proto/protoc-3.6.1-linux-x86_64.zip -d /proto && \
    rm -rf /proto/protoc-3.6.1-linux-x86_64.zip && \
    go get -u github.com/golang/protobuf/protoc-gen-go
ENV PATH="/proto/bin:${PATH}"

CMD sh -c "protoc --go_out=plugins=grpc:$GOPATH/src watchdog.proto && \
go get ./... && \
cd server && \
dev_appserver.py app.yaml --host 0.0.0.0 --admin_host 0.0.0.0 --support_datastore_emulator=True"
