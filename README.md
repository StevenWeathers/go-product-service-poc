## Product Information Service POC 

This is an example RESTful API in [Go](https://golang.org/)
using [CouchBase](http://www.couchbase.com/)
and [Docker](https://www.docker.com/)
that returns Product Information by productId

## Installation

This API is best served with [Docker](https://www.docker.com/)
```bash
docker build -t product-services .
docker run -d -p 8080:8080 --name=product-services product-services
```
What about [Docker Compose](https://docs.docker.com/compose/)?
```bash
docker-compose up -d --build
```

What about example data
```bash
Load the data in data/products.txt using Query tool in Couchbase admin panel (for now)
```

## Open Source Projects
Project | License
--- | ---
[Docker](https://github.com/docker/docker) | [Apache-2.0](https://github.com/docker/docker/blob/master/LICENSE)
[CouchBase](http://www.couchbase.com/) | tbd
[Go](https://github.com/golang/go) | [BSD-3-Clause](https://github.com/golang/go/blob/master/LICENSE)


### ENDPOINT: /product/{productId}

#### Response
```json
{
	"product": {
		"specs": [
			"Compartment Type:Multi-use insert",
			"Length (Inches):21.25",
			"Width (Inches):11.5",
			"Maximum Drawer Height (Inches):2.5",
			"Number of Compartments:4.0",
			"Material:Plastic",
			"Color/Finish Family:White",
			"Installation Type:Drop in",
			"Package Contents:Drawer insert included",
			"Minimum Drawer Height (Inches):2.4"
		],
		"barcode": "764195301563",
		"brand": "Style Selections",
		"bullets": [
			"Features include soft rounded corners, interior overmold and branded icons",
			"Non-slip rubber feet keep tray from slipping"
		],
		"description": "21.25-in x 11.5-in Plastic Multi-use Insert Drawer Organizer",
		"imgId": "90713027845",
		"itemNum": "270866",
		"lgImg": "http://images.lowes.com/product/converted/090713/090713028002lg.jpg",
		"modelNum": "GUT-10W-52",
		"networkPrice": "25.12",
		"numRatings": "0",
		"productId": "3107239",
		"rating": "0",
		"smImg": "http://images.lowes.com/product/converted/090713/090713028002sm.jpg",
		"type": "other"
	}
}
```