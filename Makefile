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
purge:
	sh .cd/down.sh && minikube delete
port-forward-todo-mysql:
	kubectl port-forward svc/todo-mysql-s 3306:3306
port-forward-todo-service:
	kubectl port-forward svc/todo-service-s 30083:80
