host = "localhost:3000"

items 			= "$(host)/items"
categories 	= "$(host)/categories"
cashes			= "$(host)/cashes"

all: test

test:
	DBURL="dbname=balance_test sslmode=disable" go test ./models ./rest/...

items.index:
	curl $(items) | jq .

items.create:
	curl $(items) | jq .

categories.index:
	curl $(categories) | jq .

categories.create:
	curl -d 'category[name]=$(name)' $(categories) | jq .

categories.delete:
	curl -X DELETE $(categories)/$(id) | jq .

categories.update:
	curl -X PATCH -d 'category[name]=$(name)&category[income]=$(income)' $(categories)/$(id) | jq .

cashes.index:
	curl $(cashes) | jq .

cashes.create:
	curl -d "$(params)" $(cashes) | jq .

cashes.update:
	curl -X PATCH -d "$(params)" $(cashes)/$(id) | jq .

cashes.destroy:
	curl -X DELETE $(cashes)/$(id) | jq .
