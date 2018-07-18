load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["stamp.go"],
    importpath = "github.com/antonovvk/bazel_stamp/go",
    visibility = ["//visibility:public"],
    x_defs = {
        "github.com/antonovvk/bazel_stamp/go.StableGitTag": "{STABLE_GIT_TAG}",
        "github.com/antonovvk/bazel_stamp/go.StableBuildDate": "{STABLE_BUILD_DATE}",
        "github.com/antonovvk/bazel_stamp/go.BuildUser": "{BUILD_USER}",
        "github.com/antonovvk/bazel_stamp/go.BuildHost": "{BUILD_HOST}",
    },
)
