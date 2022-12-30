### Building the docker image

`docker build . --tag=brunty/go-hello:<version_here> --platform=linux/arm64`

_Note that we're building for `linux/arm64` because I'm running this on my k3s cluster that's running on Raspberry Pis

### Pushing the docker image

`docker push brunty/go-hello:<version_here>`

### Running stuff

- `kubectl apply -f namespace.yaml`
- `kubectl apply -f deployment.yaml`
- `kubectl apply -f service.yaml`
