package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "store.go",
        "version.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//conf:go_default_library",
        "//internal/version:go_default_library",
        "//store:go_default_library",
        "//vendor/github.com/gin-gonic/gin:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//routers/middleware/auth:all-srcs",
        "//routers/middleware/header:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
