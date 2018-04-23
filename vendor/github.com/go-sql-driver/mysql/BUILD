package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "buffer.go",
        "collations.go",
        "connection.go",
        "connection_go18.go",
        "const.go",
        "driver.go",
        "dsn.go",
        "errors.go",
        "fields.go",
        "infile.go",
        "packets.go",
        "result.go",
        "rows.go",
        "statement.go",
        "transaction.go",
        "utils.go",
        "utils_go17.go",
        "utils_go18.go",
    ],
    importmap = "vendor/github.com/go-sql-driver/mysql",
    importpath = "github.com/go-sql-driver/mysql",
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
