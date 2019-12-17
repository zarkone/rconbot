export RCONBOT_QUAKE3_PASSWORD = freeforall

build:
	go build -v

run: build
	./rconbot
