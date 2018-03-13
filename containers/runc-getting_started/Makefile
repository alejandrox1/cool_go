run: all
	runc start

all:
	docker create --name cont1 alpine sh
	docker export cont1 > alpine.tar
	mkdir rootfs
	tar -xf alpine.tar -C rootfs/
	runc spec

clean:
	docker rm -f cont1
	rm -rf rootfs/ alpine.tar config.json
