# Help
help:
	@echo " Docker-compose Group: make dc-u|dc-d|dc-stp|dc-stt\n"
#	@echo "  clean:\t remove all generated files"
#	@echo "  init:\t\t run terraform init"
#	@echo "  set-name:\t set prefix and network name for instances"
#	@echo "  plan:\t\t run terraform plan"
#	@echo "  apply:\t create or update instances and inventory file via terraform"
#	@echo "  list-ips:\t list instance ips from GCP"
#	@echo "  destroy:\t destroy all resources via terraform"
#	@echo "\n"

# Docker-compose commands
dc-u:
	docker-compose up -d
dc-d:
	docker-compose down
dc-stp:
	docker-compose stop
dc-stt:
	docker-compose start