# docker-registry-api
light docker registry api such as harbor

## Support Version
* v2.0.1

## Target
use docker registry as harbor simple, implement such as function
* paginate
* search
* oci

## Data Defination

#### Project
```golang
type Project struct {
    Base
}
```
#### Repository
```golang
type Repository struct {
    Base
}
```
#### Artifact
```golang
type Artifact struct {
    Base
}
```
## API
#### Project
##### List Projects
* Request
    | Attribute | Params Type | Required | Description |
    |-----------|-------------|----------| ------------|
    | `page` | query | N | list project current page |
    | `pageSize` | query | N | list project page size |
* Response

#### Repository
##### List Repository
* Request
    | Attribute | Params Type | Required | Description |
    |-----------|-------------|----------| ------------|
    | `page` | query | N | list repository current page |
    | `pageSize` | query | N | list repository page size |
* Response
#### Artifact
##### List Artifact
* Request
    | Attribute | Params Type | Required | Description |
    |-----------|-------------|----------| ------------|
    | `page` | query | N | list project current page |
    | `pageSize` | query | N | list project page size |
* Response
##### Get Artifact
* Request
    | Attribute | Params Type | Required | Description |
    |-----------|-------------|----------| ------------|
    | `name` | path | Y | get artifact with artifact name |
* Response
##### Delete Artifact
* Request
    | Attribute | Params Type | Required | Description |
    |-----------|-------------|----------| ------------|
    | `page` | query | N | list project current page |
    | `pageSize` | query | N | list project page size |
* Response

## How to use
* deployment