# Citrix Sharefile ReverseProxy Mesh API (as Win64 Service)

Currently the Sharefile API doesn't support proper REST Full API interface, so I decided to create a ReverseProxy which 
can be registered as Win64 Service and perform reverse proxy back to the server.


```shell script

# To compile on Windows 
./build.cmd   # Creates a Windows Executable Service

# To compile on your Lovely Mac
./build.sh

# output exe file will be under arget/win64/sharefile.exe

# copy the exe file to your windows machine

# Install Service
sharefile.exe install

# Start Service
sharefile.exe start

# Stop Service
sharefile.exe Stop

# Remove Service
sharefile remove

# fole logs look inside windows audit event

# run it in debug mode (make sure you stop service before running in debug mode)
sharefile debug 


```

By default for now, the service listen on localhost port 8080.


# What about Linux?

We have systemctl and the process can be registered as a service or run as no hup.

# Testing

```shell script
curl localhost:8080/version
cul localhost:8080/employees
cul localhost:8080/roles
```

# Still work in progress 
