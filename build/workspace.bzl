# TensorFlow external dependencies that can be loaded in WORKSPACE files.

load("//vendor:repo.bzl", "tf_http_archive")


# Sanitize a dependency so that it works correctly from code that includes
# TensorFlow as a submodule.
def clean_dep(dep):
  return str(Label(dep))

# If TensorFlow is linked as a submodule.
# path_prefix is no longer used.
# tf_repo_name is thought to be under consideration.
def tf_workspace(path_prefix="", tf_repo_name=""):
  # Note that we check the minimum bazel version in WORKSPACE.
  tf_http_archive(
      name = "pcre",
      sha256 = "1d75ce90ea3f81ee080cdc04e68c9c25a9fb984861a0618be7bbf676b18eda3e",
      urls = [
          "http://ftp.exim.org/pub/pcre/pcre-8.40.tar.gz",
      ],
      strip_prefix = "pcre-8.40",
      build_file = clean_dep("//vendor:pcre.BUILD"),
  )
