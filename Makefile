run-backend-manager:
	cd backend && make run-manager

run-backend-worker:
	cd backend && make run-worker

run-frontend:
	cd frontend && make run-dev

update-manager:
	git pull 
	docker compose -f=docker-compose-manager.yaml down 
	docker compose -f=docker-compose-manager.yaml up -d --build 
	docker compose -f=docker-compose-manager.yaml logs -f 

update-worker:
	git pull 
	docker compose -f=docker-compose-worker.yaml down 
	docker compose -f=docker-compose-worker.yaml up -d --build 
	docker compose -f=docker-compose-worker.yaml logs -f 
