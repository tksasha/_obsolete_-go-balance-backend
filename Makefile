host = "localhost:3000"

items 			= "$(host)/items"
categories 	= "$(host)/categories"

all: test

test:
	DBURL="dbname=balance_test sslmode=disable" go test ./models ./rest/...

#
# Items
#
items.index:
	curl $(items) | jq .

items.create:
	curl $(items) | jq .

#
# Categories
#
categories.index:
	curl $(categories) | jq .

categories.create:
	curl -d 'category[name]=$(name)' $(categories) | jq .

categories.delete:
	curl -X DELETE $(categories)/$(id) | jq .

categories.update:
	curl -X PATCH -d 'category[name]=$(name)&category[income]=$(income)' $(categories)/$(id) | jq .
