# Keep this makefile!

sources := lz4.c lz4.h
replacements := LZ4_sizeofState LZ4_setStreamDecode LZ4_decompress_safe LZ4_loadDict LZ4_sizeofStreamState LZ4_uncompress_unknownOutputSize LZ4_saveDict LZ4_slideInputBuffer LZ4_create LZ4_decompress_fast LZ4_resetStream LZ4_compress LZ4_freeStreamDecode LZ4_freeStream LZ4_uncompress LZ4_createStream LZ4_createStreamDecode LZ4_versionNumber LZ4_versionString


default:
	@for r in $(replacements); do \
		for f in $(sources); do \
			sed -i "s,$${r},level_$${r},g" "$$f"; \
		done; \
	done

restore:
	@for r in $(replacements); do \
		for f in $(sources); do \
			sed -i "s,level_$${r},$${r},g" "$$f"; \
		done; \
	done
