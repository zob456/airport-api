# -*- mode: Python -*-
docker_build(
    'db-image',
    '.',
    dockerfile='db/deployments/local/Dockerfile'
)

docker_build('rest-api-image', '.', dockerfile='rest-api/deployments/local/Dockerfile')
docker_build('airport-service-image', '.', dockerfile='airport-service/deployments/local/Dockerfile')

k8s_yaml([
 './db/deployments/local/deployment.yaml',
 './rest-api/deployments/local/kubernetes.yaml',
 './airport-service/deployments/local/kubernetes.yaml',
])

k8s_resource(workload='db', port_forwards=5432)
k8s_resource(workload='rest-api', resource_deps=['db'], port_forwards=8000)
k8s_resource(workload='airport-service', resource_deps=['db'], port_forwards=8080)
