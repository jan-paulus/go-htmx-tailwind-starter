.PHONY: dev
dev:
	@echo "Starting air and npm watch"
	@{ \
		air 2>&1 | awk '{print "\033[1;34m[AIR]\033[0m " $$0}' & \
		npm run watch 2>&1 | awk '{print "\033[1;32m[NPM]\033[0m " $$0}' & \
		xdg-open "http://localhost:8081" & \
		open "http://localhost:8081" & \
		wait; \
	}

.PHONY: build
build:
	@echo "Building Go and Frontend"
	@rm -rf ./build
	@mkdir build
	@go build -o ./build/app 2>&1 | awk '{print "\033[1;34m[GO]\033[0m " $$0}'
	@npm run build 2>&1 | awk '{print "\033[1;32m[NPM]\033[0m " $$0}'
	@cp -r static/ build/static
	@cp -r templates/ build/templates
