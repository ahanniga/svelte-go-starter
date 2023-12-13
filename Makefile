EXEC=svelte-go-starter
APP_PORT=4000

init:
	cd frontend && pnpm i
	go mod tidy
clean:
	go clean
	cd frontend && if [ -d build ]; then rm -rf build; fi

debug:
	cd frontend && pnpm dev

build: clean
	cd frontend && pnpm build
	go build -o $(EXEC)

run: 
	APP_PORT=$(APP_PORT) ./$(EXEC)

.PHONY: clean build run debug

