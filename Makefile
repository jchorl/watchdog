pkg-main:
	sed -i'' s/"package watchdog"/"package main"/ *.go	

pkg-watchdog:
	sed -i'' s/"package main"/"package watchdog"/ *.go	

serve: pkg-main
	docker run -it --rm \
		-v "$(PWD)":/watchdog \
		-w /watchdog \
		-p 8080:8080 \
		-p 8000:8000 \
		jchorl/appengine-go:latest \
		sh -c "dev_appserver.py --port=8080 --host=0.0.0.0 --admin_host=0.0.0.0 \$$(pwd)"

img:
	docker build -t jchorl/watchdog -f Dockerfile.proto .

proto:
	docker run -it --rm \
		-v $(PWD):/watchdog \
		-w /watchdog \
		jchorl/watchdog \
		sh -c "protoc --go_out=paths=source_relative:. watchdog.proto"

deploy: pkg-main proto
	docker run -it --rm \
		-v $(PWD):/watchdog \
		-w /watchdog \
		jchorl/appengine-go:latest \
		sh -c "go get ./... && \
		echo \"gcloud auth login\ngcloud config set project watchdog-222905\ngcloud app deploy\ngcloud app deploy cron.yaml\" && \
		bash"
