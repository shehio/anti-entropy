load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "client",
    srcs = ["client.go"],
    importpath = "github.com/shehio/anti-entropy/src/anti_entropy/client",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gorilla_mux//:mux",
    ],
)

go_test(
    name = "client_test",
    srcs = ["client_test.go"],
    embed = [":client"],
) 