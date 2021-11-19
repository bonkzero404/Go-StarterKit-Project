# GoLang Clean Architecture

If you have read through the article Uncle bob, he said that:

> Each has at least one layer for business rules, and another for interfaces.

So, each of these architectures produce systems that are:

-   Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
-   Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
-   Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
-   Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
-   Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

From Uncle Bob’s Architecture we can divide our code in 4 layers :

-   Entities: encapsulate enterprise wide business rules. An entity in Go is a set of data structures and functions.
-   Use Cases: the software in this layer contains application specific business rules. It encapsulates and implements all of the use cases of the system.
-   Controller: the software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web
-   Framework & Driver: this layer is generally composed of frameworks and tools such as the Database, the Web Framework, etc.

> In this project trying to implement a clean architecture using uncle bob's method, it's not 100% correct for implementation process, there are still scattered files for the <b>UseCase</b>, but their have the same approach, afterwards in this project the unit test has not been implemented.

> This project uses a modular approach to isolate features, making them easier to manage.

So this is the existing file structure in this project.

```tree
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

# How to run this project?

To run this project copy the .env.example file into .env, then do the configuration as you need, here is the env file

```conf
# Application port
APP_PORT=3000

# Endpoint
API_WRAP=api
API_VERSION=v1

# Database Connection
# mysql | pgsql
DB_DRIVER=mysql
DB_HOST=localhost
DB_NAME=your_database_name
DB_USER=root
DB_PASSWORD=yout_db_password
DB_PORT=3306

# Database Pool
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100

#JWT
JWT_SECRET=rahasiabanget

# Mail
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=2525
MAIL_USERNAME=your_mailtrap_username
MAIL_PASSWORD=your_mailtrap_password
MAIL_FROM=Administrator <admin@example.com>

```

After you create the configuration file, create a database in MySQL or PostgreSQL with the appropriate name in the configuration file above.

Run the command in the root directory

```
go run main.go
```

or if you use makefile run the following command

```
make watch
```

This command has a "hot reload" feature, but you will need the <b>reflect</b> library to run the command
