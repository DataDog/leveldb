include build_flags/linux.mk
linux_sources := $(sort $(SOURCES))

include build_flags/darwin.mk
darwin_sources := $(sort $(SOURCES))

# Take the set union of the source files for all the supported platforms
all_sources := $(sort $(linux_sources) $(darwin_sources))

source_files:
	@for sf in $(all_sources); do \
		echo $$sf; \
	done
