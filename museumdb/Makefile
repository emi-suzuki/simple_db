GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run

build: 
	$(GOBUILD) import_data.go object.go fields.go schema.go

run:
	$(GORUN) import_data.go