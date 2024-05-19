.PHONY: image

image:
	docker run --rm -d -p 27017:27017 --name some-mongo \
        -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
        -e MONGO_INITDB_ROOT_PASSWORD=secret \
        mongo