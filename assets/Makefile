.PHONY: build
build:
	@sass --quiet \
		--no-source-map \
		stylesheets/application.scss:../internal/server/routes/assets/application.css
	@cp \
		bootstrap/dist/js/bootstrap.js \
		bootstrap/dist/js/bootstrap.js.map \
		../internal/server/routes/assets/
	@cp \
		htmx/dist/htmx.js \
		../internal/server/routes/assets
	@echo "done"
