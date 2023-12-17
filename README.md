# go-poc

Proof of Concept (PoC) Objective:
The primary goal of this PoC is to showcase the implementation of a Go service capable of logging reverse IPs from incoming requests. Additionally, the project includes a Helm chart for efficient deployment on a Google Kubernetes Engine (GKE) cluster and Infrastructure as Code (IaC) in Terraform to automate the deployment of GKE and Cloud SQL with PostgreSQL.

This PoC have the following: 

* go service doing saving reverse IP in all requests
** project builded on go with gin framework and gorm for storing on DBMS
** healthcheck added / not counting 
** ip saved over postgres database 
** added basic CRUD for IP operations
* dockerfile
* helm chart to deploy to GKE cluster
* Iac code in terraform to deploy GKE and Cloudsql/Postgres
* 2 github actions pipelines:
** 1 builds and deploys the code
** 1 create Iac if detects changes over `iac` folder

---
out of the scope:
* GCP: docker artifactory, apis GCP/CLOUDSQL enabled manually and bucket to save TFSTATE
* I create a github action secret to save DB_URL for application and pass it to helm chart

findings : 

* I was force to create an exception for `/healthz`  path to avoid filling the database  of the cluster requests for the liveness and readinessprobes.

* after testing application deployed on GKE I find out the IP getting saved was the one of GKE nodes hitting LB, I update chart to use `externalTrafficPolicy: Local` to fix it.

### example requests 

right now Ingress for the application is the following : `34.122.29.18`

* any `IP/path`
```
curl --location 'http://34.122.29.18' \
--data ''
Route not found, but saving IP
{"post":{"ID":1606,"CreatedAt":"2023-12-17T10:36:53.580318431Z","UpdatedAt":"2023-12-17T10:36:53.580318431Z","DeletedAt":null,"Ip":"24.0.10.10"}}%

curl --location 'http://34.122.29.18/deel' \
--data ''
Route not found, but saving IP
{"post":{"ID":1607,"CreatedAt":"2023-12-17T10:37:29.410626879Z","UpdatedAt":"2023-12-17T10:37:29.410626879Z","DeletedAt":null,"Ip":"22.0.10.10"}}%
```

 * GET `IP/posts`
```
curl --location 'http://34.122.29.18/posts' | jq 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   649  100   649    0     0   3892      0 --:--:-- --:--:-- --:--:--  4133
{
  "posts": [
    {
      "ID": 1,
      "CreatedAt": "2023-12-17T10:42:18.008231Z",
      "UpdatedAt": "2023-12-17T10:42:18.008231Z",
      "DeletedAt": null,
      "Ip": "22.0.10.10"
    },
    {
      "ID": 2,
      "CreatedAt": "2023-12-17T10:42:21.575775Z",
      "UpdatedAt": "2023-12-17T10:42:21.575775Z",
      "DeletedAt": null,
      "Ip": "22.0.10.10"
    },
    {
      "ID": 3,
      "CreatedAt": "2023-12-17T10:42:23.268852Z",
      "UpdatedAt": "2023-12-17T10:42:23.268852Z",
      "DeletedAt": null,
      "Ip": "22.0.10.10"
    },
    {
      "ID": 4,
      "CreatedAt": "2023-12-17T10:42:28.644528Z",
      "UpdatedAt": "2023-12-17T10:42:28.644528Z",
      "DeletedAt": null,
      "Ip": "22.0.10.10"
    },
    {
      "ID": 5,
      "CreatedAt": "2023-12-17T10:42:32.94638Z",
      "UpdatedAt": "2023-12-17T10:42:32.94638Z",
      "DeletedAt": null,
      "Ip": "22.0.10.10"
    }
  ]
}
    ...
```

* POST `IP/posts`
```
curl --location --request POST 'http://34.122.29.18/posts'
{"post":{"ID":9,"CreatedAt":"2023-12-17T10:47:37.044518327Z","UpdatedAt":"2023-12-17T10:47:37.044518327Z","DeletedAt":null,"Ip":"22.0.10.10"}}%
```

 * PUT `IP/posts`
 ```
 curl --location --request PUT 'http://34.122.29.18/posts/2' \
--data ''
{"post":{"ID":2,"CreatedAt":"2023-12-17T10:42:21.575775Z","UpdatedAt":"2023-12-17T10:44:20.051604222Z","DeletedAt":null,"Ip":"22.0.10.10"}}%
 ```

 * DEL `IP/posts`
 ```
 curl --location --request DELETE 'http://34.122.29.18/posts/2'
 ```

