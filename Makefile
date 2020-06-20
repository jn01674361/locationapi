local: 
	docker build . -t locationapi
	docker run -p 8080:8080 locationapi

