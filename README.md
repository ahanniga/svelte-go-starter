# Svelte Go Starter

This is a starter template that provides a Svelte front-end coupled to a Go backend. All the code is packaged into a single, self-contained binary making it ideal for container deployments. 

It works by building the frontend using Svelte's `adapter-static`. The compiled output is embedded in the Go app using the `//go:embed` directive. Single or multiple page apps are possible using simple Svelte [routing rules](https://learn.svelte.dev/tutorial/layouts).

The pages are served using the [Fiber](https://gofiber.io/) http library. 

## Building

```bash
git clone https://github.com/ahanniga/svelte-go-starter.git
cd svelte-go-starter
make init
make build
./svelte-go-starter
```

## Debugging

Uses Vite for the frontend and an IDE for the backend (or both).

### Start Vite
```bash
make debug 
```

Press <kbd>o</kbd> to open a browser window or just navigate to http://localhost:5173. To connect from VSCode too, here's an example VSCode launch config.

```json
{
    "name": "Launch Chrome Debugger",
    "type": "chrome",
    "request": "launch",
    "url": "http://localhost:5173",
    "webRoot": "${workspaceFolder}/frontend/src",
    "runtimeExecutable": "/path/to/your/browser"
}
```

> Note! As Vite is serving your pages when debugging, `vite.config.js` is configured to proxy all calls prefixed with `/api/` to the backend at port (default 4000). 

### Start Go Debugger

```json
{
    "name": "Launch Go Debug",
    "type": "go",
    "request": "launch",
    "mode": "auto",
    "program": "${workspaceFolder}",
    "cwd": "${workspaceFolder}",
    "env": {
        "APP_PORT": "4000"
    }
}
```

