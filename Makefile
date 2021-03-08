UID=$(shell id -u)
GID=$(shell id -g)

proto:
	docker run -it --rm \
		-u "$(UID):$(GID)" \
		-v "$(PWD)"/proto:/watchdog/github.com/jchorl/watchdog/proto \
		-w /watchdog/github.com/jchorl/watchdog/proto \
		namely/protoc-all \
		-f watchdog.proto \
		-l go \
		-o /watchdog

deploy:
	docker run -it --rm \
		-v $(PWD):/watchdog \
		-w /watchdog \
		-v watchdogcreds:/root/.config/gcloud/ \
		gcr.io/google.com/cloudsdktool/cloud-sdk:330.0.0 \
		sh -c "echo \"gcloud auth login\ngcloud config set project watchdog-222905\ngcloud app deploy\ngcloud app deploy cron.yaml\" && \
		bash"

cmd-docker:
	docker build -t jchorl/wdping -f dockerfiles/wdping.Dockerfile .

.PHONY: proto
