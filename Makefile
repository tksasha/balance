.PHONY: all
all: stop assets_clobber log_clear tmp_clear assets_precompile start

.PHONY: stop
stop:
	launchctl stop balance

.PHONY: start
start:
	launchctl start balance

.PHONY: assets_clobber
assets_clobber:
	rake assets:clobber

.PHONY: log_clear
log_clear:
	rake log:clear

.PHONY: tmp_clear
tmp_clear:
	rake tmp:clear

.PHONY: assets_precompile
assets_precompile:
	RAILS_ENV=production rake assets:precompile

