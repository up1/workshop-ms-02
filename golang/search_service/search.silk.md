# Search

## POST /search

Search a name.

* Content-Type: "application/json"
* Accept: "appliaction/json"

Include the `name` text in the body:

```json
{
  "name": "Santi"
}
```

===

### Example response

* Status: 200
* Content-Type: "application/json"

```json
{
  "name": "Santi"
}
```
