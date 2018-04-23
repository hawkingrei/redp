package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["routers.go"],
    tags = ["automanaged"],
    deps = [
        "//routers/middleware/auth:go_default_library",
        "//server:go_default_library",
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
        "//routers/middleware:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
