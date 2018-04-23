package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_test",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":go_default_library"],
    tags = ["automanaged"],
)

go_test(
    name = "go_default_test",
    srcs = ["test_test.go"],
    embed = [":go_default_library"],
    tags = ["automanaged"],
    deps = [
        "//conf:go_default_library",
        "//vendor/github.com/WindomZ/testify/assert:go_default_library",
        "//vendor/github.com/appleboy/gofight:go_default_library",
        "//vendor/github.com/buger/jsonparser:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["redp-cli.go"],
    tags = ["automanaged"],
    deps = [
        "//conf:go_default_library",
        "//internal/version:go_default_library",
        "//model:go_default_library",
        "//routers:go_default_library",
        "//routers/middleware:go_default_library",
        "//routers/middleware/header:go_default_library",
        "//store:go_default_library",
        "//store/datastore:go_default_library",
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/gin-gonic/contrib/ginrus:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/golang.org/x/sync/errgroup:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
