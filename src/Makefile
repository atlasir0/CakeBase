GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
BINARY0=readDB
BINARY1=compareDB
BINARY2=compareFS
SRC0=cmd/readDB/main.go
SRC1=cmd/compareD/main.go
SRC2=cmd/compareF/main.go
OLDSNAPSHOT=cmd/compareF/snapshot_v1.txt
NEWSNAPSHOT=cmd/compareF/snapshot_v2.txt
OLDDB=cmd/compareD/original_database.xml
NEWDB=cmd/compareD/stolen_database.json
FORMAT=4


all: task0 task1 task2


task0: $(SRC0)
	$(GOBUILD) -o $(BINARY0) $(SRC0)
	./$(BINARY0) -f jdatabase.json


task1: $(SRC1)
	$(GOBUILD) -o $(BINARY1) $(SRC1)
	./$(BINARY1) --old $(OLDDB) --new $(NEWDB)


task2: $(SRC2)
	$(GOBUILD) -o $(BINARY2) $(SRC2)
	./$(BINARY2) --old $(OLDSNAPSHOT) --new $(NEWSNAPSHOT)


clean:
	$(GOCLEAN)
	rm -f $(BINARY0) $(BINARY1) $(BINARY2)

