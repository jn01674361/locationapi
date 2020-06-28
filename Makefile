local: 
	docker build . -t locationapi
	cd client && docker build . -t frontend && cd ..
	docker-compose up -d
