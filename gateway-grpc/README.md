# API Gateway with gRPC-Gateway and Envoy

This project provides an API Gateway that converts REST API calls to gRPC calls for the PostService using two approaches:

1. **gRPC-Gateway**: Direct HTTP to gRPC conversion using Google API HTTP annotations
2. **Envoy Proxy**: HTTP to gRPC transcoding using Envoy's gRPC-JSON transcoder

## Architecture

```
Frontend (REST) → API Gateway → PostService (gRPC)
```

## Features

- REST API endpoints that automatically convert to gRPC calls
- OpenAPI/Swagger documentation generation
- CORS support
- Docker containerization
- Health check endpoints

## API Endpoints

The gateway exposes the following REST endpoints:

- `POST /api/v1/posts` - Create a new post
- `GET /api/v1/posts` - Get all posts
- `GET /api/v1/posts/{id}` - Get post by ID
- `PATCH /api/v1/posts/{id}` - Update post
- `DELETE /api/v1/posts/{id}` - Delete post
- `GET /health` - Health check

## Setup

### Prerequisites

- Go 1.24+
- Docker and Docker Compose
- Protocol Buffers compiler (protoc)

### Option 1: gRPC-Gateway (Recommended)

1. **Generate protobuf code**:

   ```bash
   cd gateway-grpc
   ./generate.sh
   ```

2. **Run with Docker Compose**:

   ```bash
   docker-compose up --build
   ```

3. **Or run locally**:
   ```bash
   go mod tidy
   go run main.go
   ```

The gRPC-Gateway will be available at `http://localhost:8080`

### Option 2: Envoy Proxy

1. **Generate protobuf descriptor**:

   ```bash
   protoc -I . -I third_party \
     --descriptor_set_out=proto/post/post.pb \
     --include_imports \
     proto/post/post.proto
   ```

2. **Run with Docker Compose**:
   ```bash
   docker-compose up envoy post-service postgres
   ```

The Envoy proxy will be available at `http://localhost:8081`

## Testing the API

### Create a Post

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{
    "post_by": "user123",
    "title": "My First Post",
    "detail": "This is a test post",
    "image_url": "https://example.com/image.jpg",
    "event_id": 1,
    "status": "active"
  }'
```

### Get All Posts

```bash
curl http://localhost:8080/api/v1/posts
```

### Get Post by ID

```bash
curl http://localhost:8080/api/v1/posts/1
```

### Update Post

```bash
curl -X PATCH http://localhost:8080/api/v1/posts/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Post Title",
    "detail": "Updated post content"
  }'
```

### Delete Post

```bash
curl -X DELETE http://localhost:8080/api/v1/posts/1
```

## Configuration

Environment variables can be set in `env.example`:

- `GATEWAY_PORT`: Port for the gateway server (default: 8080)
- `GRPC_SERVER_ENDPOINT`: gRPC server address (default: localhost:50066)
- `POST_SERVICE_ADDR`: Post service gRPC address

## Development

### Adding New Endpoints

1. Update the proto file with HTTP annotations:

   ```protobuf
   rpc NewMethod(NewMethodRequest) returns (NewMethodResponse) {
     option (google.api.http) = {
       post: "/api/v1/new-endpoint"
       body: "*"
     };
   }
   ```

2. Regenerate the code:

   ```bash
   ./generate.sh
   ```

3. Restart the gateway

### Swagger Documentation

The gateway automatically generates OpenAPI/Swagger documentation. Access it at:

- `http://localhost:8080/swagger-ui/` (if swagger-ui is configured)
- `http://localhost:8080/proto/post/post.swagger.json` (raw OpenAPI spec)

## Troubleshooting

### Common Issues

1. **Connection refused**: Ensure the PostService gRPC server is running on the correct port
2. **Proto compilation errors**: Check that all required proto files and dependencies are present
3. **CORS issues**: Verify CORS configuration in the gateway

### Logs

Check Docker logs:

```bash
docker-compose logs grpc-gateway
docker-compose logs envoy
docker-compose logs post-service
```

## Performance Considerations

- gRPC-Gateway provides better performance for direct HTTP-to-gRPC conversion
- Envoy proxy is better for complex routing, load balancing, and advanced features
- Both solutions support HTTP/2 and connection pooling
