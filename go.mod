module github.com/Nedson202/book-api-go

go 1.12

require (
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.1
	github.com/lib/pq v1.1.0
	github.com/nedson202/book-api-go v0.0.0-20190426113148-428ec46f50b9
	github.com/subosito/gotenv v1.1.1
)

replace github.com/nedson202/book-api-go v0.0.0-20190426113148-428ec46f50b9 => ./../book-api-go
