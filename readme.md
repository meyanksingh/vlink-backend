video-calling-backend/
├── cmd/
│   └── server/
│       └── main.go               # Entry point of the application, initializes Gin router
├── config/
│   └── config.go                 # Configuration settings (e.g., environment variables)
├── internal/
│   ├── app/
│   │   ├── controllers/
│   │   │   ├── auth_controller.go    # Authentication-related handlers
│   │   │   └── friend_controller.go  # Friend management handlers
│   │   ├── models/
│   │   │   ├── user.go               # User model definition
│   │   │   └── friend.go             # Friend model definition
│   │   ├── services/
│   │   │   ├── auth_service.go       # Business logic for authentication
│   │   │   └── friend_service.go     # Business logic for friend management
│   │   └── repository/
│   │       ├── user_repository.go    # Database functions for user data
│   │       └── friend_repository.go  # Database functions for friend data
│   ├── db/
│   │   └── db.go                     # Database connection setup (e.g., PostgreSQL)
│   └── middleware/
│       └── auth_middleware.go        # JWT authentication middleware
├── pkg/
│   └── utils/
│       ├── jwt.go                    # JWT utilities for token generation and validation
│       └── hash.go                   # Utilities for password hashing and validation
├── routes/
│   └── routes.go                     # Registers routes and handlers with the Gin router
├── scripts/
│   └── migrate.sh                    # Script for setting up database tables
├── .env                              # Environment variables (database URL, JWT secret, etc.)
├── go.mod                            # Go module dependencies
└── README.md                         # Project documentation
