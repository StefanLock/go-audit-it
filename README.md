# Go-Audit-IT

## API Usage

__Note:__ Replace localhost with your servers URL

### Create system
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"host":"my_test_server", "status":"pass"}' \
  localhost:5000/systems
```
### Get system
``` 
curl localhost:5000/systems 
```
### Update system
```
curl -X PATCH \
  -H 'Content-Type: application/json' \
  -d '{"Host":"my_updated_test_server"}' \
  localhost:5000/systems/<id>

```
### Delete system
```
curl -X DELETE localhost:5000/systems/<id>
```