clean:
	rm -f sync



sync: clean
	go generate ./...
	CGO_ENABLED=0 go build ./cmd/sync

sync-image: sync
	go get github.com/openshift/imagebuilder/cmd/imagebuilder
	imagebuilder -f Dockerfile.sync -t docker.io/jimminter/sync:latest .

sync-push: sync-image
	docker push docker.io/jimminter/sync:latest

.PHONY: clean sync-image sync-push

# docker pull docker.io/jimminter/sync:latest
# docker run --dns=8.8.8.8 -i -v /root/.kube:/.kube:z -e KUBECONFIG=/.kube/config docker.io/jimminter/sync:latest
