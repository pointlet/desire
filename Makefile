dev: 
	@make -j dev-tailwind templ

templ:
	@templ generate --watch --proxy="http://localhost:1323" --cmd="go run ./cmd/main.go" $(ARGS)

dev-tailwind:
	@make ARGS="--watch" tailwind 

tailwind:
	@npx tailwindcss -c ./view/config/tailwind.config.js -i ./view/config/tailwind.css -o ./static/css/styles.css $(ARGS)
