# pingpong
A simple 2-container script that when running the docker compose will make both endpoints run.

Step 1:
Clone the repository: `git clone https://github.com/Robert076/pingpong`

Step 2:
Run docker compose: `docker-compose up -d`

Step 3:
Go to *localhost:1234/ping* or *localhost:1235/pong*

The purpose of this repo:
The purpose of this repo was to learn **Docker Compose**. The whole idea is that the *ping/* and *pong/* directories act as the "services" of what would be a real application in the real world.

Each of the two "services" have a *Dockerfile* in them that instruct Docker on how to build the image for each service.

Then comes the *docker-compose.yml* in the root, which looks something like this:

```
services:
  ping:
    build:                     # tells Compose how to (re)build the application
      context: ./ping          # this is where our Dockerfile is at
      dockerfile: Dockerfile   # this is the name of it, although it is optional and Docker will find it without us specifying it
    ports:
      - "1234:1234"            # host:container port mapping. Going in our browser on port 1234 will be like going in the container on port 1234
    networks:                  
      - pingpong               # network name. In this case it is optional since our driver is a bridge (bridge is the default driver for docker networks)
  pong:
    build:
      context: ./pong
      dockerfile: Dockerfile
    ports:
      - "1235:1235"
    networks:
      - pingpong

networks:
  pingpong:
    driver: bridge
```
