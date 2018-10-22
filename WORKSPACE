load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
	name = "io_bazel_rules_go",
	urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.0/rules_go-0.16.0.tar.gz"],
	sha256 = "ee5fe78fe417c685ecb77a0a725dc9f6040ae5beb44a0ba4ddb55453aad23a8a",
)

http_archive(
	name = "bazel_gazelle",
	urls = ["https://github.com/bazelbuild/bazel-gazelle/archive/master.zip"],
	strip_prefix = "bazel-gazelle-master",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "go_repository", "gazelle_dependencies")
gazelle_dependencies()

go_repository(
	name = "com_18f_hmacauth",
	importpath = "github.com/18F/hmacauth",
	commit = "9232a6386b737d7d1e5c1c6e817aa48d5d8ee7cd",
)

go_repository(
	name = "com_benbjohnson_clock",
	importpath = "github.com/benbjohnson/clock",
	commit = "7dc76406b6d3c05b5f71a86293cbcf3c4ea03b19",
)

go_repository(
	name = "com_datadog_go",
	importpath = "github.com/datadog/datadog-go",
	commit = "281ae9f2d8950e694e810c78daf74201d869b0d0",
)

go_repository(
	name = "com_go_yaml",
	importpath = "gopkg.in/yaml.v2",
	commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
)

go_repository(
	name = "com_google_cloud_go",
	importpath = "cloud.google.com/go",
	commit = "deb9345c9d7f5f3d9d07726b582d39105f57feb1",
)

go_repository(
	name = "com_google_api_go",
	importpath = "google.golang.org/api",
	commit = "a2651947f503a1793446d4058bb073a6fdf99e53",
)

go_repository(
	name = "com_imdario_mergo",
	importpath = "github.com/imdario/mergo",
	commit = "9f23e2d6bd2a77f959b2bf6acdbefd708a83a4a4",
)

go_repository(
	name = "com_kelseyhightower_envconfig",
	importpath = "github.com/kelseyhightower/envconfig",
	commit = "dd1402a4d99de9ac2f396cd6fcb957bc2c695ec1",
)

go_repository(
	name = "com_miscreant_go",
	importpath = "github.com/miscreant/miscreant-go",
	commit = "325cbd69228b6af571a635f7502586a920a2749a",
)

go_repository(
	name = "com_rakyll_statik",
	importpath = "github.com/rakyll/statik",
	commit = "1355192d24db2566a83c3914e187e2a7e7679832",
)

go_repository(
	name = "com_sirupsen_logrus",
	importpath = "github.com/sirupsen/logrus",
	commit = "4fabf2fffcecfd47f802869b7b92d75e43c5a095",
)

go_repository(
	name = "org_golang_x_crypto",
	importpath = "golang.org/x/crypto",
	commit = "0c41d7ab0a0ee717d4590a44bcb987dfd9e183eb",
)

go_repository(
	name = "org_golang_x_oauth2",
	importpath = "golang.org/x/oauth2",
	commit = "9dcd33a902f40452422c2367fefcb95b54f9f8f8",
)

# Docker support.
http_archive(
	name = "io_bazel_rules_docker",
	urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.5.1.tar.gz"],
	sha256 = "29d109605e0d6f9c892584f07275b8c9260803bf0c6fcb7de2623b2bedc910bd",
	strip_prefix = "rules_docker-0.5.1"
)
load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
_go_image_repos()
