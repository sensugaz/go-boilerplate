# GO Boilerplate

#### Hexagonal Architecture
This golang project adopt the hexagonal architecture to grouping the code by domain and context of business problems.

**Core**

Everything surrounded the `core` this layer contains the concrete core business logics.

The core could be viewed as a “box” (represented as a hexagon) capable of resolve all the business logic independently.

**Domain**

The domain entity it contains the go struct definition of each entity that is part of the domain problem and can be used across the application.

>**Seperate of concern**
For example, this layer should not be known how the data keeps its, or how the caching store in the redis.

**Port** is the interfaces are the ports that external actors will plug their adapters into to drive the application.

**Adapter** is the implementation of the ports

**Service**
Service is through which the client will interact with actual business logic/domain objects. The client can be rest handlers, test agents, etc.

**Repository**
The repository interface allows us to connect to multiple data sources.

**Dependency Injection (Port and Adaptor Design Patterns)**
Outside concerns, such as repositories and user interfaces, use the adapters to plug into the business domain services.

#### Go Philosophy
- Reliability 
- Scalability
- Durability 
- Simplicity 
- Performance

#### Grouping code by domain
Go does not force gopher about framework, project structure.
[Please read my blog carefully](https://goangle.medium.com/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-gophercon-2018-kat-zien-how-do-you-structure-your-go-apps-a96faca9e8f1)

Don't stick to the best practices of other companies that are very successful., just make good practice engineering for Bitkub.

# Features
1. docker live-reloader
2. pre-commit
   1. golang-ci linter
   2. commit linter
   3. custom rule

# Getting Started
## Setup local
At root directory.
1. run `$ make init` to install pre-requisite tools.
2. run `$ precommit.rehooks` to hook the pre-commit.
3. Change some file, then `$ git add .`
4. trigger the pre-commit with following command `$ git commit`
5. the pre-commit features will be trigger.
6. check the linter, and push your code.

# Contributing
## Commitlint
You can reach all configs about commit lint under `./commitlint.config.js`
## TODO
- Self shell command to repository for pre-commit.
