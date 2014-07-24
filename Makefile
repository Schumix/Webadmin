GOC ?= go
GOCFLAGS =
BUILDPATH = build
TARGETNAME = webadmin

all: $(TARGETNAME)

$(TARGETNAME):
	$(GOC) build -o $(BUILDPATH)/$(TARGETNAME) src/main.go
	cp -r www build

clean:
	rm -rf build

.PHONY = all webadmin clean
