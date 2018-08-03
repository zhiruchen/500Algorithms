CURRENT_GIT_REPO := github.com/zhiruchen/500Algorithms

test:
	go test -v $(CURRENT_GIT_REPO)/all/array
	go test -v $(CURRENT_GIT_REPO)/all/stack
	go test -v $(CURRENT_GIT_REPO)/all/binarytree 
