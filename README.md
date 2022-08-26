Simple DAG problem in Go


### API Endpoint(s)

| Action |                 Endpoint                  | Method |
|:------:|:-----------------------------------------:|:-------|
|||
| find the total flight paths <br/>starting and ending airports |            /api/process-flight-list             | POST   |

### Example

```shell
curl --location --request POST 'localhost:50000/api/process-flight-list' \
--header 'Content-Type: application/json' \
--data-raw '[
    [
        "IND",
        "EWR"
    ],
    [
        "SFO",
        "ATL"
    ],
    [
        "GSO",
        "IND"
    ],
    [
        "ATL",
        "GSO"
    ]
]'
```
output :
```json
{
    "starting": "SFO",
    "ending": "EWR"
}
```