# metrics_manager

Go Application for exposing custom metrics that checks the external urls are up and gets the response time in miliseconds with output in prometheus format 
e.g. sample_external_url_up{url="https://httpstat.us/503"} 0

To achieve the expected output, I used prometheus, promauto, and promhttp Go libraries, to expose my custom metrics along with the go metrics.

# Before you begin
Download and install golang  
  * Download the installer from the page https://golang.org/dl/ and install the package.  
  * Create a Go workspace and set GO PATH  

# Getting Started

To run the code locally 
1. clone this git repository in the Go workspace
2. In the root directory of the git repo, run "go mod download" or you can run "go mod tidy"
3. change directory "cd metrics_manager"
4. run "go run main.go collector.go"
5. Application will start running in the local (make sure no other application is running on 8080)
7. curl -l localhost:8080 to see output with expected result in the end 


To run as docker container
1. In the root directory, you will find Dockerfile named as metrics_manager.Dockerfile
2. run the cmd "docker build -t metrics_manager -f metrics_manager.Dockerfile ." to build the image
3. After that run the image using "docker run -d --network host --name metrics_manager -p  8080:8080 -it metrics_manager" 
4. curl -l localhost:8080 to see the output 

# OUTPUT
#HELP sample_external_url_response_ms response time for url  
#TYPE sample_external_url_response_ms counter  

sample_external_url_response_ms{url="https://httpstat.us/200"} 387  
sample_external_url_response_ms{url="https://httpstat.us/503"} 803  

#HELP sample_external_url_up shows if url is up  
#TYPE sample_external_url_up counter  

sample_external_url_up{url="https://httpstat.us/200"} 1  
sample_external_url_up{url="https://httpstat.us/503"} 0

# Metrics In Prometheus and Grafana

Steps to See the custom metrics on Prometheus Dashboard
 1. From root diectory of the repo, do "cd Prometheus"
 2. Build the prometheus image using cmd "docker build -t my-prometheus -f prom.Dockerfile ."
 3. Addition Info : We have a "prometheus.yml" file in the /prometheus folder, which is custom config file contains the scrape config for the go application, the "target" field is set as localhost:8080 if running locally else set it as ClusterIP/name:port of the kubernetes Service created for it.In the dockerfile, we copy our custom "prometheus.yml" to get the metrics.
 4. Run the image using cmd "docker run  -d  --network host --name my-prometheus-p 9090:9090  my-prometheus"
 5. Go to http://localhost:9090 , you will be able to see the custom metrics added on the go application.
 
 Steps to See the custom metrics on Grafana Dashboard
 
 1. Run the cmd "docker run --network host -d -p 3000:3000 --name grafana grafana/grafana:6.5.0", to create the container
 2. Go to http://localhost:3000, login to the the account using admin as username and password.
 3. create a datasource as prometheus.
 4. add custom queries to the panel, and you'll be able to see the graphs for them.
 
 
 # Kubernetes Deployment
 
 Created the kubernetes deployment files for the Go Application and prometheus pod with service as ClusterIP and For grafana created a deployment and exposed it as loadBalancer.
 All the files are provided in the kuberentes folder.
 

 
# Screen Shots of output from prometheus, grafana and local application run


![image](https://user-images.githubusercontent.com/25012913/110841120-879b3580-82cb-11eb-86b0-e3bda2f488aa.png)
![image](https://user-images.githubusercontent.com/25012913/110841566-0d1ee580-82cc-11eb-9a1b-80df6f447790.png)
![image](https://user-images.githubusercontent.com/25012913/110841620-23c53c80-82cc-11eb-808e-51bf82a6f5a5.png)






