default:
	docker run -it --rm \
		-v $(PWD):/go/src/github.com/jchorl/watchdog \
		-w /go/src/github.com/jchorl/watchdog \
		-p 8080:8080 \
		-p 8000:8000 \
		jchorl/watchdog

img:
	docker build -t jchorl/watchdog .
