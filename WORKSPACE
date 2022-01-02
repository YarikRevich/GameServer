workspace(name = "hide-seek-server")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

http_archive(
    name = "rules_proto",
    sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
    strip_prefix = "rules_proto-4.0.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "59536e6ae64359b716ba9c46c39183403b01eabfbd57578e84398b4829ca499a",
    strip_prefix = "rules_docker-0.22.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.22.0/rules_docker-v0.22.0.tar.gz"],
)

http_archive(
    name = "rules_python",
    sha256 = "cd6730ed53a002c56ce4e2f396ba3b3be262fd7cb68339f0377a45e8227fe332",
    url = "https://github.com/bazelbuild/rules_python/releases/download/0.5.0/rules_python-0.5.0.tar.gz",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "io_bazel_rules_prometheus",
    sha256 = "c9980c638cba01015f31bc610788d5396b05b67c99c1f065fef17427cb8459fe",
    strip_prefix = "rules_prometheus-0.0.4",
    urls = ["https://github.com/5h4d0w4rt/rules_prometheus/archive/0.0.4.zip"],
)

http_archive(
    name = "rules_proto_grpc",
    sha256 = "dfc624227d7da02abeaa649d6629bf17b68db7a371a83e3c60fd4a5d60b180df",
    strip_prefix = "rules_proto_grpc-4.1.1",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/refs/tags/4.1.1.zip"],
)


#Proto workspace

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()


#GRPC workspace

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")
rules_proto_grpc_toolchains()
rules_proto_grpc_repos()



load("@rules_proto_grpc//:repositories.bzl", "bazel_gazelle", "io_bazel_rules_go")  # buildifier: disable=same-origin-load
io_bazel_rules_go()
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(
    version = "1.17.1",
)


bazel_gazelle()
load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")
rules_proto_grpc_go_repos()


#Gazelle workspace 

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_repositories")
# gazelle:repository_macro deps.bzl%go_repositories
go_repositories()
gazelle_dependencies()


#Docker workspace

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()
load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")
container_deps()
load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")
container_pull(
    name = "prometheus_base",
    registry = "index.docker.io",
    repository = "prom/prometheus",
)
load("@io_bazel_rules_docker//repositories:go_repositories.bzl", "go_deps")
go_deps()


#Prometheus workspace

load("@io_bazel_rules_prometheus//:deps.bzl", "prometheus_repositories")
prometheus_repositories()


#Grafana workspace

git_repository(
    name = "io_bazel_rules_grafana",
    commit = "5f6ada63da3efa52536f59e7833788d4f5ababaa",
    remote = "https://github.com/etsy/rules_grafana.git",
)
load("@io_bazel_rules_grafana//grafana:workspace.bzl", grafana_repositories = "repositories")
grafana_repositories()

