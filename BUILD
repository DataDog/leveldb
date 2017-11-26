load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "batch.go",
        "cache.go",
        "cgo_flags_common.go",
        "cgo_snappy_common.go",
        "comparator.go",
        "conv.go",
        "db.go",
        "doc.go",
        "env.go",
        "filterpolicy.go",
        "iterator.go",
        "options.go",
        "seek_to.c",
        "seek_to.go",
        "seek_to.h",
        "snappy_supported.cc",
        "snappy_supported.go",
        "snappy_supported.h",
        "version.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "cgo_flags_darwin.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "cgo_flags_linux.go",
        ],
        "//conditions:default": [],
    }),
    cgo = True,
    cdeps = ["//deps/snappy", "//deps/lz4", "//deps/leveldb"],
    importpath = "github.com/DataDog/leveldb",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "leveldb_test.go",
        "snappy_test.go",
    ],
    importpath = "github.com/DataDog/leveldb",
    library = ":go_default_library",
)
