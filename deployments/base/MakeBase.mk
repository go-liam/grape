
test-unit:
	cd ${BASEDIR}; \
	cd ../../ ; \
	go mod tidy; \
	export PROJECT_ENV="unit" ; \
	go test $(go list ./... | grep -v /proto/ | grep -v /temp-demo/ | grep -v /frontend/ ) -coverprofile=coverage.data ./...;

deploy-dev-config:
	cd ${BASEDIR}; \

deploy-test-config:
	cd ~/.kube ;  \
	cp -p config-minikube ./config \

deploy-pre-config:
	cd ~/.kube ;  \
	cp -p config-aliy-test-dev ./config \

deploy-prod-config:
	cd ~/.kube ;  \

k8s-dev-create:
	cd ${BASEDIR} ;  \
	kubectl create namespace dev-qps ;  \
	echo "deploy-creat-dev";  \

k8s-dev-delete:
	cd ${BASEDIR} ;  \
	kubectl delete namespace dev-qps ;  \
	echo "deploy-delete-dev";  \

get-config:
	echo ${DEPLOY_SERVER_NAME}; \
	echo ${DEPLOY_SERVER_VERSION}; \
	echo ${DEPLOY_DOCKER_NAME}; \
	echo ${DEPLOY_DOCKER_PSW}; \
	echo ${BASEDIR}

before-script:
	cd ${BASEDIR} ; \
	cd ../../ ; \
	go mod tidy; \

build-file:
	cd ${BASEDIR} ; \
	rm -rf build | echo "no build dir"; \
	mkdir build; \
    mkdir ./build/log ; \
    mkdir ./build/assets ; \
    mkdir ./build/assets/${DEPLOY_SERVER_NAME} ; \
    cp -R ../../assets/${DEPLOY_SERVER_NAME}/* ./build/assets/${DEPLOY_SERVER_NAME} | echo "no copy dir" ; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o build/main ../../app/${DEPLOY_SERVER_NAME}/main.go;\
	echo "build-file finish";

build-copy-cert:
	cd ${BASEDIR} ; \
    cp -R ../../assets/cert/* ./build/assets/cert ; \
	echo "build-copy-cert";

test-temp:
	echo ${DockerUsername}

build-image:
	echo "$$USER"
	DockerImageVersion=${DEPLOY_SERVER_VERSION} ;  \
	DockerImageFile="${DEPLOY_DOCKER_NAME}/${DEPLOY_SERVER_DOCKER_NAME}:v$$DockerImageVersion" ;  \
	echo $$DockerImageFile ; \
	docker rmi $$DockerImageFile ;  \
	cd ${BASEDIR} ;  \
	docker build -t $$DockerImageFile -f Dockerfile .; \
	docker login --username=${DEPLOY_DOCKER_NAME} -p ${DEPLOY_DOCKER_PSW} ;  \
	docker push $$DockerImageFile ;  \

deploy-local-apply:
	echo "deploy-local finish" ;  \

deploy-dev-apply:
	cd ${BASEDIR} ;  \
	cd ../docker/mn/ ;  \
    docker-compose up -d ;  \
    echo "deploy-dev finish" ;  \

deploy-test-apply:
	cd ${BASEDIR} ;  \
    kubectl apply -f ./minikube  -n test-mn ;  \
    echo "deploy-test finish" ;  \

deploy-pre-apply:
	cd ${BASEDIR} ;  \
	kubectl apply -f ./ali-k8s  -n dev-cbox ;  \
	echo "deploy-pre finish";  \

deploy-prod-apply:
	cd ${BASEDIR} ;  \
	echo "deploy-prod finish";  \

delete-dev-apply:
	cd ${BASEDIR} ;  \
	cd ../docker/mn/ ;  \
    docker-compose down ;  \
    echo "delete-dev-apply finish" ;  \
