# POC-HEXAGONAL-GO-V2
Another approach for Golang hexagonal project

## SHORT RECAP

> The domain is the core of the hexagon, containing primary business logic, free of any infrastructure and framework boilerplate.

> Adapters are either external APIs of your application or clients to other systems. They translate the interfaces of external systems (e.g. a search index or a file server API) to the interfaces required or exposed by the domain. Those interfaces are called ports.

> Ports allow plugging in the adapters into the core domain. An example could be a repository interface with a method returning article content as a simple String. By declaring a port, e.g. as an plain Java interface, the domain declares the contract saying: ”I give an id and I expect text in return, where and how you get it from is your business”. The domain here deals with articles, which have a text title and a text content. Not with JSON or binary files. It does not want to hear about S3, ElasticSearch or an SFTP server.

## HOW TO GENERATE MOCKS
> First install: go get github.com/golang/mock/gomock

Example of how to generate:
> mockgen  -source=src/domain/product/find_product_data_output_port.go -destination=src/test/mocks/product/find_product_data_output_port.go -package=mocks_product 