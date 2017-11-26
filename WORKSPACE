http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.7.0/rules_go-0.7.0.tar.gz",
    sha256 = "91fca9cf860a1476abdc185a5f675b641b60d3acf0596679a27b580af60bf19c",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

# new_local_repository(
#     name = "system_libs",
#     # pkg-config --variable=libdir x11
#     path = "/usr/lib/x86_64-linux-gnu",
#     build_file_content = """
# cc_library(
#     name = "x11",
#     srcs = ["libX11.so"],
#     visibility = ["//visibility:public"],
# )
# """,
# )

# new_git_repository(
#     name = "go_snappy",
#     build_file = "third_party/go/snappy.BUILD",
#     commit = "d9eb7a3d35ec988b8585d4a0068e462c27d28380",
#     remote = "https://github.com/golang/snappy.git",
# )

