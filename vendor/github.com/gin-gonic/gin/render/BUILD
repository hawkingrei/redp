package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "data.go",
        "html.go",
        "json.go",
        "msgpack.go",
        "redirect.go",
        "render.go",
        "text.go",
        "xml.go",
        "yaml.go",
    ],
    importmap = "vendor/github.com/gin-gonic/gin/render",
    importpath = "github.com/gin-gonic/gin/render",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/gin-gonic/gin/json:go_default_library",
        "//vendor/github.com/ugorji/go/codec:go_default_library",
        "//vendor/gopkg.in/yaml.v2:go_default_library",
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
