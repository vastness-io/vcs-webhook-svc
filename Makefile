generate:
	@echo "=> generating stubs"
	protoc -I ${PWD}/webhook/github --proto_path=${PWD}/webhook/github/ ${PWD}/webhook/github/*.proto --go_out=plugins=grpc:${PWD}/webhook/github
	protoc -I ${PWD}/webhook/bitbucketserver --proto_path=${PWD}/webhook/bitbucketserver/ ${PWD}/webhook/bitbucketserver/*.proto --go_out=plugins=grpc:${PWD}/webhook/bitbucketserver
	@echo "=> generating mocks"
	mockgen github.com/vastness-io/vcs-webhook-svc/webhook/github GithubWebhookClient,GithubWebhookServer > mock/webhook/github/webhook_mock.go
	mockgen github.com/vastness-io/vcs-webhook-svc/webhook/bitbucketserver BitbucketServerPostWebhookClient,BitbucketServerPostWebhookServer > mock/webhook/bitbucketserver/webhook_mock.go
