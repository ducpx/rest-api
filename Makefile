.PHONY: local

local:
	echo "Starting docker environment"
	docker-compose -f docker-compose.local.yml up --build