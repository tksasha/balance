DEST=../internal/server/routes/assets/

.PHONY: build
build: clean install
	@echo "done"

@PHONY: clean
clean:
	@echo "clean"
	@rm -f $(DEST)/*.js $(DEST)/*.css

.PHONY: install
install: bootstrap datepicker htmx hryvnia application
	@echo "install"

define hashname
$(notdir $(basename $(1))-$(shell sha1sum $(1)|cut -d' ' -f1)$(suffix $(1)))
endef

.PHONY: bootstrap
bootstrap:
	@cp \
		bootstrap/dist/js/bootstrap.js \
		$(DEST)/$(call hashname,bootstrap/dist/js/bootstrap.js)
	@cp \
		bootstrap/dist/js/bootstrap.js.map \
		$(DEST)

.PHONY: datepicker
datepicker:
	@cp \
		datepicker/dist/css/datepicker.min.css \
		$(DEST)/$(call hashname,datepicker/dist/css/datepicker.min.css)
	@cp \
		datepicker/dist/css/datepicker-bs5.min.css \
		$(DEST)/$(call hashname,datepicker/dist/css/datepicker-bs5.min.css)
	@cp \
		datepicker/dist/js/datepicker.min.js \
		$(DEST)/$(call hashname,datepicker/dist/js/datepicker.min.js)
	@cp \
		datepicker/dist/js/locales/uk.js \
		$(DEST)/$(call hashname,datepicker/dist/js/locales/uk.js)

.PHONY: htmx
htmx:
	@cp \
		htmx/dist/htmx.min.js \
		$(DEST)/$(call hashname,htmx/dist/htmx.min.js)

.PHONY: hryvnia
hryvnia:
	@cp images/hryvnia.png $(DEST)/$(call hashname,images/hryvnia.png)

.PHONY: compile
compile:
	@sass --quiet \
		--no-source-map \
		stylesheets/application.scss:stylesheets/application.css
	@cat \
		javascripts/_common.js \
		javascripts/backoffice.js \
		javascripts/cashes.js \
		javascripts/categories.js \
		javascripts/items.js \
		> javascripts/application.js

.PHONY: application
application: compile
	@mv \
		stylesheets/application.css \
		$(DEST)/$(call hashname,stylesheets/application.css)
	@mv \
		javascripts/application.js \
		$(DEST)/$(call hashname,javascripts/application.js)
