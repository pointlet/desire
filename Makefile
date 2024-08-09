dev: 
	@make -j dev-tailwind templ sync-assets

sync-assets:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "static" \
		--build.include_ext "js,css"

templ:
	@templ generate --watch --proxy="http://localhost:1323" --cmd="go run ./cmd/main.go" $(ARGS)

dev-tailwind:
	@make ARGS="--watch" tailwind 

tailwind:
	@npx tailwindcss -c ./view/config/tailwind.config.js -i ./view/config/tailwind.css -o ./static/css/styles.css $(ARGS)