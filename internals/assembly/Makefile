GOOS=linux
GOARCH=amd64

SRC=$(wildcard *.go)
OBJ=$(SRC:.go=.o)
EXEC=$(OBJ:.o=.bin)

.SECONDARY: ${OBJ}

all: ${EXEC}

%.o: %.go
	GOOS=${GOOS} GOARCH=${GOARCH} go tool compile $<

%.bin: %.o
	GOOS=${GOOS} GOARCH=${GOARCH} go tool link -o $@ $<

clean:
	rm -f ${OBJ} 
	rm -f ${EXEC}
