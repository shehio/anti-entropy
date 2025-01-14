load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "node",
    srcs = ["node.go"],
    importpath = "github.com/shehio/anti-entropy/src/anti_entropy/node",
    visibility = ["//visibility:public"],
)

go_test(
    name = "node_test",
    srcs = ["node_test.go"],
    embed = [":node"],
) 