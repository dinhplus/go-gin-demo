# GraphQL Integration

This project now includes GraphQL API alongside the existing REST API endpoints.

## GraphQL Endpoints

- **GraphQL API**: `POST /graphql`
- **GraphQL Playground**: `GET /playground` (for development)

## Features

### ✅ Implemented
- Complete GraphQL schema for all models (User, Department, Position, Stack, Role, Permission, Resource, FeatureFlag)
- Query resolvers for fetching data
- Mutation resolvers for creating users
- JWT authentication middleware for GraphQL
- GraphQL Playground for development
- Integration with existing service layer

### 🔄 In Progress
- Complete CRUD mutations for all models
- Field-level permissions
- Subscription support
- Advanced pagination and filtering

## Usage

### Starting the Server
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` with:
- REST API endpoints (existing)
- GraphQL endpoint at `/graphql`
- GraphQL Playground at `/playground`

### GraphQL Playground
Visit `http://localhost:8080/playground` to explore the GraphQL schema and test queries.

### Example Queries

#### Get All Users
```graphql
query {
  users {
    id
    name
    email
    department {
      name
    }
    position {
      name
    }
    stack {
      name
    }
    role {
      name
      permissions {
        name
        action
      }
    }
  }
}
```

#### Get User by Email
```graphql
query {
  userByEmail(email: "user@example.com") {
    id
    name
    email
    department {
      name
    }
  }
}
```

#### Create New User
```graphql
mutation {
  createUser(input: {
    name: "John Doe"
    email: "john@example.com"
    password: "password123"
    departmentId: 1
    positionId: 1
    stackId: 1
    roleId: 1
  }) {
    id
    name
    email
  }
}
```

#### Simple Health Check
```graphql
query {
  ping
}
```

### Authentication

For protected operations, include JWT token in headers:
```
Authorization: Bearer <your-jwt-token>
```

## File Structure

```
internal/
├── graphql/
│   ├── generated/          # Auto-generated GraphQL code
│   │   ├── generated.go
│   │   └── models_gen.go
│   ├── resolvers/          # Custom resolver implementations
│   │   ├── resolver.go
│   │   └── schema.resolvers.go
│   └── schema/             # GraphQL schema definitions
│       └── schema.graphql
├── handler/
│   ├── graphql_handler.go  # GraphQL HTTP handlers
│   └── graphql_auth.go     # GraphQL authentication middleware
└── ...
```

## Development Workflow

1. **Modify Schema**: Edit `internal/graphql/schema/schema.graphql`
2. **Generate Code**: Run `go run github.com/99designs/gqlgen generate`
3. **Implement Resolvers**: Update `internal/graphql/resolvers/schema.resolvers.go`
4. **Test**: Use GraphQL Playground at `/playground`

## Configuration

GraphQL configuration is in `gqlgen.yml`:
- Schema location
- Generated code location
- Resolver configuration
- Model bindings

## Best Practices

- Keep business logic in the service layer
- Use existing repository functions
- Handle authentication in resolvers
- Return appropriate GraphQL errors
- Use context for request-scoped data
- Follow established naming conventions
- Don't modify generated files

## Apollo Client Integration

This GraphQL server is fully compatible with Apollo Client. Configure Apollo Client to use:
- Endpoint: `http://localhost:8080/graphql`
- Include JWT token in authorization headers for protected operations

Example Apollo Client setup:
```javascript
import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';

const httpLink = createHttpLink({
  uri: 'http://localhost:8080/graphql',
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem('token');
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
});

const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache()
});
```
