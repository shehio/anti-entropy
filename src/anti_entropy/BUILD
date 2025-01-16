load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "anti_entropy_lib",
    srcs = ["main.go"],
    importpath = "github.com/shehio/anti-entropy/src/anti_entropy",
    visibility = ["//visibility:private"],
    deps = [
        "//src/anti_entropy/node",
        "//src/anti_entropy/merkle",
    ],
)

go_binary(
    name = "anti_entropy",
    embed = [":anti_entropy_lib"],
    visibility = ["//visibility:public"],
) 