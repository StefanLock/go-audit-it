# Go-Audit-IT

## API Usage

__Note:__ Replace localhost with your servers URL

### Create system
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"hostname":"my_test_server"}' \
  -d '{"status": "true"}' \
  localhost:5000/systems
```
### Get system
``` 
curl localhost:5000/systems 
```
### Update system
```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"hostname":"my_updated_test_server"}' \
  localhost:5000/systems/<id>

```
### Delete system
```
curl -X DELETE localhost:5000/systems/<id>
```