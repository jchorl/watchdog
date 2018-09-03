default:
	docker run -it --rm \
		-v $(PWD):/go/src/github.com/jchorl/watchdog \
		-w /go/src/github.com/jchorl/watchdog \
		-p 8080:8080 \
		-p 8000:8000 \
		jchorl/watchdog

img:
	docker build -t jchorl/watchdog .

deploy:
	docker run -it --rm \
		-v $(PWD):/go/src/github.com/jchorl/watchdog \
		-w /go/src/github.com/jchorl/watchdog \
		jchorl/watchdog \
		sh -c "protoc --go_out=plugins=grpc:\$$GOPATH/src watchdog.proto && \
		go get ./... && \
		cd server && \
		echo \"gcloud auth login\ngcloud config set project watchdog-215220\ngcloud app deploy\ngcloud app deploy cron.yaml\" && \
		bash"
