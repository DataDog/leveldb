include build_flags/linux.mk
linux_sources := $(sort $(SOURCES))

include build_flags/darwin.mk
darwin_sources := $(sort $(SOURCES))

# for shell substitution
SHELL := /bin/bash

# Take the set union of the source files for all the supported platforms
all_sources := $(sort $(linux_sources) $(darwin_sources))

link_leveldb:
	@for sf in $(all_sources); do \
		f="vendor/leveldb/$$sf"; \
		rm -f $${f//\//_}; \
		ln -s $$f $${f//\//_}; \
	done

link_snappy:
	rm -f vendor_snappy*
	ln -s vendor/snappy/snappy.cc vendor_snappy.cc
	ln -s vendor/snappy/snappy-c.cc vendor_snappy-c.cc
	ln -s vendor/snappy/snappy-sinksource.cc vendor_snappy-sinksource.cc

link_lz4:
	rm -f vendor_lz4*
	ln -s vendor/lz4/lz4.c vendor_lz4.c

link: link_snappy link_leveldb link_lz4
