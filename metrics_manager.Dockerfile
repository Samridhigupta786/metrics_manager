# Start from the latest golang base image
FROM golang:latest
# Add Maintainer Info
LABEL maintainer="SAMRIDHI GUPTA <samridhigupta.100@gmail.com>"
# Set the Current Working Directory inside the container
WORKDIR /src
#copy everything to current directory
COPY  . ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
#build the go application
RUN cd /src/metrics_manager && go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./metrics_manager/main"]