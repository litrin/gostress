INSTALL_PATH=/usr/bin
APPLICATION=gostress
BUILDTAGS=./gostress.go

all:
	go build -tags "$(BUILDTAGS)" -o $(APPLICATION)

install:
	cp $(APPLICATION) $(INSTALL_PATH)

clean:
	rm $(APPLICATION)

