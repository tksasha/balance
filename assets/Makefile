DEST=../internal/server/routes/assets/

.PHONY: build
build: clean compile copy
	@echo "done"

@PHONY: clean
clean:
	@echo "clean"
	@rm -f $(DEST)/*.js $(DEST)/*.css

.PHONY: compile
compile:
	@echo "compile"
	@sass --quiet \
		--no-source-map \
		stylesheets/application.scss:application.css

.PHONY: copy
copy:
	@echo "copy"
	@mv \
		application.css \
		$(DEST)/application-$(shell sha1sum application.css|cut -d' ' -f1).css
	@cp \
		bootstrap/dist/js/bootstrap.js \
		$(DEST)/bootstrap-$(shell sha1sum bootstrap/dist/js/bootstrap.js|cut -d' ' -f1).js
	@cp \
		bootstrap/dist/js/bootstrap.js.map \
		$(DEST)
	@cp \
		htmx/dist/htmx.js \
		$(DEST)/htmx-$(shell sha1sum htmx/dist/htmx.js|cut -d' ' -f1).js
