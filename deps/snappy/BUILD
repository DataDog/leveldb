cc_library(
    name = "snappy",
    srcs = [
        "snappy-internal.h",
        "snappy-sinksource.cc",
        "snappy-sinksource.h",
        "snappy-stubs-internal.cc",
        "snappy-stubs-internal.h",
        "snappy-stubs-public.h",
        "snappy.cc",
    ],
    hdrs = ["snappy.h"],
    includes = ["."],  # allow `#include "snappy.h"` rather than `#include "deps/snappy/snappy.h`
    visibility = ["//:__pkg__", "//deps/leveldb:__pkg__"],
)
