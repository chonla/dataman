GIT_COMMIT=`git rev-parse HEAD`
build:
	go build -ldflags="-X 'main.GitCommitID=${GIT_COMMIT}'" -o bin/dataman main.go