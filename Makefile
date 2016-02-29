all: test

test:
	DBURL="dbname=balance_test sslmode=disable" go test ./models ./rest/...
