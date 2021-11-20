Run playyoshimi in background to detect midi device plugged-in or out, then we launch yoshimi or stop its process.

Quick run:
```
go run playyoshimi.go yoshimi --state=/path/to/1.state
```

Build :
```
sudo docker build . -t playyoshimi
# Start playyoshimi
sudo docker run --name playyoshimi_container playyoshimi

# Get playyoshimi from docker image
sudo docker cp playyoshimi_container:/go/src/app/playyoshimi .
```

Run playyoshimi on linux device
```

```