# Go Example HTTP Service

This project provides a simple example Go HTTP Service that can be used as a template

# Running with Docker Compose

1. Run `docker compose up`
   ```sh
   docker compose up -d
   ```
2. Start the Go webserver in `service` container
   ```sh
   docker compose exec service go run .
   ```
3. Access [http://localhost:8080/hello](http://localhost:8080/hello)
4. To clean-up containers, run `docker compose down`
   ```sh
   docker compose down
   ```

# Running with Dev Containers in VS Code

1. Make sure you have `Dev Containers` extension or equivalent installed along with `Docker`
2. Open this folder in VS Code
3. Run command `Dev Containers: Rebuid and Reopen in Container`
4. Start the Go webserver in the VS Code terminal
   ```sh
   go run .
   ```
5. Access [http://localhost:8080/hello](http://localhost:8080/hello)

# Resources

- [Ultimate Guide to Dev Containers](https://www.daytona.io/dotfiles/ultimate-guide-to-dev-containers)
- [How I write HTTP services in Go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
