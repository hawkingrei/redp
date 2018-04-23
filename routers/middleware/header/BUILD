package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
    "go_library",
)

go_test(
    name = "go_default_test",
    srcs = ["header_test.go"],
    embed = [":go_default_library"],
    tags = ["automanaged"],
    deps = [
        "//routers/middleware:go_default_library",
        "//vendor/github.com/gin-gonic/gin:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["header.go"],
    tags = ["automanaged"],
    deps = ["//vendor/github.com/gin-gonic/gin:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
