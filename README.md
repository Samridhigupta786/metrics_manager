# metrics_manager

Go Application for exposing custom metrics that checks the external urls are up and gets the response time in miliseconds with output in prometheus format 
e.g. sample_external_url_up{url="https://httpstat.us/503"} 0

To achieve the expected output, I used prometheus, promauto, and promhttp Go libraries, to expose my custom metrics along with the go metrics.


# Getting Started

To run the code locally 

1. clone the git repository 
2. In the same directory run "go mod download" 
3. change directory "cd metrics_manager"
4. run go main.go collector.go
5. Application will start running in the local
6. curl -l localhost:8080 to see output with expected result in the end 


To run as docker container
1. In the root directory, you will find Dockerfile named as metrics_manager.Dockerfile
2. run the cmd "docker build -t metrics_manager -f metrics_manager.Dockerfile ." to build the image
3. After that run the image using "docker run -d --network host --name metrics_manager -p  8080:8080 -it metrics_manager" 
4. curl -l localhost:8080 to see the output 

# OUTPUT
#HELP sample_external_url_response_ms response time for url  
#TYPE sample_external_url_response_ms counter  

sample_external_url_response_ms{url="https://httpstat.us/200"} 3.55105526e+08  
sample_external_url_response_ms{url="https://httpstat.us/503"} 7.09015772e+08  

#HELP sample_external_url_up shows if url is up  
#TYPE sample_external_url_up counter  

sample_external_url_up{url="https://httpstat.us/200"} 1  
sample_external_url_up{url="https://httpstat.us/503"} 0


