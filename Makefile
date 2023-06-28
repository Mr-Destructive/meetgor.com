clean:
	markata clean

build:
	markata build

serve:
	python -m http.server -d markout/

cleanbuild: clean build

crun: cleanbuild serve

run: build serve
