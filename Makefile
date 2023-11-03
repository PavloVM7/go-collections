revive:
	$(GOPATH)/bin/revive -config ./revive.toml -formatter friendly ./...
revive-no-tests:
	$(GOPATH)/bin/revive -config ./revive.toml \
    -exclude lists/linked_list_test.go \
    -exclude lists/list_item_test.go \
    -exclude lists/quick_sort_list_test.go \
    -exclude lists/quick_sort_list_benchmark_test.go \
    -formatter friendly ./...
