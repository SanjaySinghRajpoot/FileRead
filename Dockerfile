FROM golang:alpine AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and build the binary
COPY . .

# Add static project files from the tmp folder
COPY tmp /tmp

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

### 
## Step 2: Runtime stage
FROM scratch

# Copy only the binary from the build stage to the final image
COPY --from=builder /app/myapp /
# Copy the txt files
COPY --from=builder /tmp /tmp


EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/myapp"]