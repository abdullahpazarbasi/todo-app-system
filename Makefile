up:
	sh .cd/up.sh
down:
	sh .cd/down.sh
build:
	sh .cd/build.sh
deploy:
	sh .cd/deploy.sh
undeploy:
	sh .cd/undeploy.sh
dash:
	minikube dashboard
delete:
	sh .cd/down.sh && minikube delete
port-forward-todo-mysql:
	kubectl port-forward pods/todo-mysql-ss-0 3306:3306
