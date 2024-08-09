dev: 
	@make -j dev-tailwind templ

air-test:
	@go run github.com/cosmtrek/air@v1.51.0

air:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "go build ." --build.delay "500" \
		--build.exclude_dir "node_modules" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

sync-assets:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "static" \
		--build.include_ext "js,css"

templ:
	@templ generate --watch --proxy="http://localhost:1323" --open-browser=false -v

dev-tailwind:
	@make ARGS="--watch" tailwind 

tailwind:
	@npx tailwindcss -c ./view/config/tailwind.config.js -i ./view/config/tailwind.css -o ./static/css/styles.css $(ARGS)

live:
	@make -j dev-tailwind templ air-test sync-assets