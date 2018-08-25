PROGRAM=hub-trend

all: usage

edit e:
	vi $(PROGRAM).go

edit-make em:
	vi Makefile

build b:
	go build -o $(PROGRAM) $(PROGRAM).go

run r:
	#./$(PROGRAM) -h
	./$(PROGRAM) -time=week -lang=go
	./$(PROGRAM) -time=week -lang=go -item=dev
	./$(PROGRAM) -time=week -lang=go -item=lang

install i:
	cp $(PROGRAM) $(GOPATH)/bin

git g:
	@echo ""
	@echo "make (git) [gn|gu]"
	@echo ""

new gn:
	echo "# hub-trend" >> README.md
	git init
	git add README.md
	git commit -m "first commit"
	git remote add origin https://sikang99@github.com/sikang99/$(PROGRAM).git
	git push -u origin master

update gu:
	git init
	git add README.md Makefile *.go vendor/
	git commit -m "upload with /vendor"
	git push -u https://sikang99@github.com/sikang99/$(PROGRAM) master

md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "make [edit|build|run]"
	@echo ""

