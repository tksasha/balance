DEST=../internal/server/routes/assets/

.PHONY: build
build: clean install
	@echo "done"

@PHONY: clean
clean:
	@echo "clean"
	@rm -f $(DEST)/*.js $(DEST)/*.css

.PHONY: install
install: bootstrap datepicker htmx hryvnia jquery application
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
		bootstrap-datepicker/dist/css/bootstrap-datepicker.min.css \
		$(DEST)/$(call hashname,bootstrap-datepicker/dist/css/bootstrap-datepicker.min.css)
	@cp \
		bootstrap-datepicker/dist/js/bootstrap-datepicker.min.js \
		$(DEST)/$(call hashname,bootstrap-datepicker/dist/js/bootstrap-datepicker.min.js)
	@cp \
		bootstrap-datepicker/dist/locales/bootstrap-datepicker.uk.min.js \
		$(DEST)/$(call hashname,bootstrap-datepicker/dist/locales/bootstrap-datepicker.uk.min.js)

.PHONY: htmx
htmx:
	@cp \
		htmx/dist/htmx.min.js \
		$(DEST)/$(call hashname,htmx/dist/htmx.min.js)

.PHONY: hryvnia
hryvnia:
	@cp images/hryvnia.png $(DEST)/$(call hashname,images/hryvnia.png)

.PHONY: jquery
jquery:
	@cp \
		jquery/dist/jquery.min.js \
		$(DEST)/$(call hashname,jquery/dist/jquery.min.js)

.PHONY: compile
compile:
	@sass --quiet \
		--no-source-map \
		stylesheets/application.scss:stylesheets/application.css
	@cat \
		javascripts/common.js \
		javascripts/cashes.js \
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
