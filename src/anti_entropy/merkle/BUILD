load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "merkle",
    srcs = ["merkle_node.go", "merkle_tree.go"],
    importpath = "github.com/shehio/anti-entropy/src/anti_entropy/merkle",
    visibility = ["//visibility:public"],
)

go_test(
    name = "merkle_test",
    srcs = ["merkle_test.go"],
    embed = [":merkle"],
) 