include platform.mk

source_files:
	@for sf in $(SOURCES); do \
		echo $$sf; \
	done
