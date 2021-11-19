# GoLang Clean Architecture

```
go-boilerplate-clean-arch
├── Makefile
├── README.md
├── app
│   └── bootstrap.go
├── config
│   └── config.go
├── domain
│   ├── models
│   │   ├── mail_model.go
│   │   └── response_model.go
│   └── stores
│       ├── user_activation.go
│       └── user_store.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── database
│   │   ├── connector.go
│   │   ├── migration.go
│   │   ├── mysql.go
│   │   └── postgresql.go
│   └── middleware
│       ├── authenticate.go
│       └── rate_limiter.go
├── main.go
├── modules
│   ├── auth
│   │   ├── domain
│   │   │   ├── interfaces
│   │   │   │   └── user_auth_service_interface.go
│   │   │   └── models
│   │   │       ├── user_auth_profile_model.go
│   │   │       ├── user_auth_request_model.go
│   │   │       ├── user_auth_response_model.go
│   │   │       └── user_auth_validation_model.go
│   │   ├── handlers
│   │   │   └── auth_handler.go
│   │   ├── module.go
│   │   ├── route.go
│   │   └── services
│   │       └── auth_service.go
│   └── user
│       ├── domain
│       │   ├── interfaces
│       │   │   ├── user_activation_factory_interface.go
│       │   │   ├── user_forgot_pass_factory_interface.go
│       │   │   ├── user_repository_interface.go
│       │   │   └── user_service_interface.go
│       │   └── models
│       │       ├── user_activation_request_model.go
│       │       ├── user_activation_request_validation_model.go
│       │       ├── user_create_reponse_model.go
│       │       ├── user_create_request_model.go
│       │       ├── user_create_request_validation_model.go
│       │       ├── user_forgot_pass_act_request_model.go
│       │       ├── user_forgot_pass_act_validation_model.go
│       │       ├── user_forgot_pass_request_model.go
│       │       ├── user_forgot_pass_validation_model.go
│       │       ├── user_reactivation_request_model.go
│       │       └── user_reactivation_validation_model.go
│       ├── handlers
│       │   └── user_handler.go
│       ├── module.go
│       ├── repositories
│       │   └── user_repository.go
│       ├── route.go
│       └── services
│           ├── factories
│           │   ├── activation_factory.go
│           │   ├── user_activation_factory.go
│           │   └── user_forgot_pass_factory.go
│           └── user_service.go
├── server
├── templates
│   ├── user_activation.html
│   └── user_forgot_password.html
└── utils
    ├── api_group.go
    ├── api_wrapper.go
    ├── hash.go
    ├── mail.go
    ├── rand_char.go
    └── validation_struct.go
```
