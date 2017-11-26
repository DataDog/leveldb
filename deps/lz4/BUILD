cc_library(
    name = "lz4",
    srcs = [
        "lz4.c",
    ],
    hdrs = ["lz4.h"],
    includes = ["."],  # allow `#include "whatever.h"` rather than `#include "deps/lz4/whatever.h`
    visibility = ["//:__pkg__", "//deps/leveldb:__pkg__"],  # TODO not sure if this is the right way to allow golz4 to use this but not anything else
)
