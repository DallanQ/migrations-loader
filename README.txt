REFRESH DEPENDENCIES - last refreshed 10 Sep 2015

git subtree pull --prefix src/github.com/golang/protobuf https://github.com/golang/protobuf master --squash
git subtree pull --prefix src/golang.org/x/net https://github.com/golang/net master --squash
git subtree pull --prefix src/golang.org/x/oauth2 https://github.com/golang/oauth2 master --squash
git subtree pull --prefix src/google.golang.org/grpc https://github.com/google/grpc master --squash
git subtree pull --prefix src/google.golang.org/api https://github.com/google/google-api-go-client master --squash
git subtree pull --prefix src/google.golang.org/cloud https://github.com/GoogleCloudPlatform/gcloud-golang master --squash

CHANGES

git rm -rf src/google.golang.org/cloud/examples
git rm -rf src/google.golang.org/api/examples
