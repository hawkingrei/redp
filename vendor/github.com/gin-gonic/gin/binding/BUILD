package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "binding.go",
        "default_validator.go",
        "form.go",
        "form_mapping.go",
        "json.go",
        "msgpack.go",
        "protobuf.go",
        "query.go",
        "xml.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/gin-gonic/gin/json:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/ugorji/go/codec:go_default_library",
        "//vendor/gopkg.in/go-playground/validator.v8:go_default_library",
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
