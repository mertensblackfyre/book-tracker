run:
	cd cmd && go run main.go

build:
	cd cmd && go install main.go
	
migrateup:
	migrate -database sqlite://./app.db -source file://./pkg\db\schema up


migratedrop:
	migrate -database sqlite://./app.db -source file://./migrations down
	
generate:
	.\sqlc.exe generate
