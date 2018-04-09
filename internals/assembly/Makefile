GOOS=linux
GOARCH=amd64

SRC=$(wildcard *.go)
OBJ=$(SRC:.go=.o)
EXEC=$(OBJ:.o=.bin)

.SECONDARY: ${OBJ}

all: ${EXEC}
	GOOS=${GOOS} GOARCH=${GOARCH} go tool compile $^
	@objdump -j .text -t direct_topfunc_call.bin

%.o: %.go
	GOOS=${GOOS} GOARCH=${GOARCH} go tool compile $<

%.bin: %.o
	GOOS=${GOOS} GOARCH=${GOARCH} go tool link -o $@ $<

clean:
	rm -f ${OBJ} 
	rm -f ${EXEC}
