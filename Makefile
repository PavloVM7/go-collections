COLLECTIONS = pkg/collections
LISTS = $(COLLECTIONS)/lists
revive:
	$(GOPATH)/bin/revive -config ./revive.toml -formatter friendly ./...
revive-no-tests:
	$(GOPATH)/bin/revive -config ./revive.toml \
    -exclude $(LISTS)/linked_list_test.go \
    -exclude $(LISTS)/list_item_test.go \
    -exclude $(LISTS)/quick_sort_list_test.go \
    -exclude $(LISTS)/quick_sort_list_benchmark_test.go \
    -exclude $(COLLECTIONS)/collection_utils_test.go \
    -exclude $(COLLECTIONS)/set_test.go \
    -formatter friendly ./...
