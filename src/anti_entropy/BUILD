load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "node_lib",
    srcs = ["node/node.go"],
    importpath = "github.com/shehio/anti-entropy/src/anti_entropy/node",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "anti_entropy",
    srcs = ["main.go"],
    deps = [":node_lib"],
) 