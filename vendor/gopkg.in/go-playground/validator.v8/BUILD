package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "baked_in.go",
        "cache.go",
        "doc.go",
        "regexes.go",
        "util.go",
        "validator.go",
    ],
    importmap = "vendor/gopkg.in/go-playground/validator.v8",
    importpath = "gopkg.in/go-playground/validator.v8",
    tags = ["automanaged"],
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
