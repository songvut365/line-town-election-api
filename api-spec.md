<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [LINE Town Election API Spec](#line-town-election-api-spec)
  - [GET Candidate](#get-candidate)
    - [Response Detail](#response-detail)
  - [GET Candidate Detail](#get-candidate-detail)
    - [Parameter Detail](#parameter-detail)
    - [Response Detail](#response-detail-1)
  - [POST Create a new Candidate](#post-create-a-new-candidate)
    - [Parameter Detail](#parameter-detail-1)
    - [Response Detail](#response-detail-2)
  - [PUT Update a Candidate](#put-update-a-candidate)
    - [Parameter Detail](#parameter-detail-2)
    - [Response Detail](#response-detail-3)
  - [DELETE Delete a Candidate](#delete-delete-a-candidate)
    - [Parameter Detail](#parameter-detail-3)
    - [Response Detail](#response-detail-4)
  - [POST Check Vote status](#post-check-vote-status)
    - [Response Detail](#response-detail-5)
  - [POST Vote](#post-vote)
    - [Response Detail](#response-detail-6)
  - [POST Toggle Election](#post-toggle-election)
    - [Response Detail](#response-detail-7)
  - [GET Election Result](#get-election-result)
    - [Response Detail](#response-detail-8)
  - [GET Election Result](#get-election-result-1)
    - [Response Detail](#response-detail-9)
  - [GET Exported Result (download)](#get-exported-result-download)
  - [Real-time Vote Stream](#real-time-vote-stream)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# LINE Town Election API Spec

*NOTES*

This spec is designed to test candidate ability to design and implement backend system. Some might have security issues (and thus not production ready), please feel free to add security measures/improvements!

## GET Candidate

API to get Candidate list

Example curl:
```
curl {endpoint}/api/candidates \
--header 'Authorization: Bearer xxxx'
```

Expected Response:
```
[
  {
    "id": "1",
    "name": "Elon Musk",
    "dob": "June 28, 1971",
    "bioLink": "https://en.wikipedia.org/wiki/Elon_Musk",
    "imageLink": "https://upload.wikimedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 0
  },
  {
    "id": "2",
    "name": "Jeff Bezos",
    "dob": "January 12, 1964",
    "bioLink": "https://en.wikipedia.org/wiki/Jeff_Bezos",
    "imageLink": "https://pbs.twimg.com/profile_images/669103856106668033/UF3cgUk4_400x400.jpg",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 0
  },
  ...
]
```

### Response Detail
| Parameter | Description |
------------|--------------
| id | ID of the candidate |
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |
| votedCount | Candidate's vote count |


## GET Candidate Detail

API to get Candidate detail

Example curl:
```
# curl {endpoint}/api/candidates/:candidateId \
# --header 'Authorization: Bearer xxxx'

# Example
curl {endpoint}/api/candidates/1 \
--header 'Authorization: Bearer xxxx'
```

### Parameter Detail
| Parameter | Description |
------------|--------------
| candidateId | ID of the candidate |


Expected Response:
```
  {
    "id": "1",
    "name": "Elon Musk",
    "dob": "June 28, 1971",
    "bioLink": "https://en.wikipedia.org/wiki/Elon_Musk",
    "imageLink": "https://upload.wikimedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 0
  }
```

### Response Detail
| Parameter | Description |
------------|--------------
| id | ID of the candidate |
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |
| votedCount | Candidate's vote count |


## POST Create a new Candidate 

API to create a new Candidate

Example curl:
```
curl -X POST {endpoint}/api/candidates \
--header 'Authorization: Bearer xxxx' \
-d '{
  "name": "Brown",
  "dob": "August 8, 2011",
  "bioLink": "https://line.fandom.com/wiki/Brown",
  "imageLink": "https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630",
  "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown"
}'
```

### Parameter Detail
| Parameter | Description |
------------|--------------
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |


Expected Response:
```
  {
    "id": "3",
    "name": "Brown",
    "dob": "August 8, 2011",
    "bioLink": "https://line.fandom.com/wiki/Brown",
    "imageLink": "https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 0
  }
```

### Response Detail
| Parameter | Description |
------------|--------------
| id | ID of the candidate |
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |

## PUT Update a Candidate 

API to update a Candidate

Example curl:
```
curl -X PUT {endpoint}/api/candidates/3 \
--header 'Authorization: Bearer xxxx' \
-d '{
  "number": 3,
  "name": "LINE Brown",
  "dob": "August 8, 2011",
  "bioLink": "https://line.fandom.com/wiki/Brown",
  "imageLink": "https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630",
  "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown"
}'
```

### Parameter Detail
| Parameter | Description |
------------|--------------
| candidateId | ID of the candidate |
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |


Expected Response:
```
# Updated candidate's detail as response
  {
    "id": "3",
    "name": "LINE Brown",
    "dob": "August 8, 2011",
    "bioLink": "https://line.fandom.com/wiki/Brown",
    "imageLink": "https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 0
  }
```

### Response Detail
| Parameter | Description |
------------|--------------
| id | ID of the candidate |
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |

## DELETE Delete a Candidate 

API to delete a Candidate

Example curl:
```
curl -X DELETE {endpoint}/api/candidates/3 \
--header 'Authorization: Bearer xxxx'
```

### Parameter Detail
| Parameter | Description |
------------|--------------
| candidateId | ID of the candidate |


Expected Response:
```
# success
  {
    "status": "ok"
  }

# fail
  {
    "status": "error",
    "message": "Candidate not found"
  }
```

### Response Detail
| Parameter | Description |
------------|--------------
| status | status of request |
| message | (optional) error message |

## POST Check Vote status

API to check vote status for voter

Example curl:
```
curl -X POST {endpoint}/api/vote/status \
--header 'Authorization: Bearer xxxx' \
--header 'Content-Type: application/json' \
-d '{
  "nationalId": "1111111111114"
}'
```

Expected Response:
```
# Success
{
  "status": true
}

# Error
{
  "status": false
}
```

### Response Detail
| Parameter | Description |
------------|--------------
| status | Current vote status, if this national id can vote or not |


## POST Vote

API to vote

Example curl:
```
curl -X POST {endpoint}/api/vote \
--header 'Authorization: Bearer xxxx' \
--header 'Content-Type: application/json' \
-d '{
  "nationalId": "1111111111114",
  "candidateId": 1
}'
```

Expected Response:
```
# Success
{
  "status": "ok",
}

# Error
# In case, voter has already voted
{
  "status": "error",
  "message": "Already voted"
}

# In case, election is closed
{
  "status": "error",
  "message": "Election is closed"
}
```

### Response Detail
| Parameter | Description |
------------|--------------
| status | status of request |
| message | (optional) error message |



## POST Toggle Election
Example curl:
```
curl -X POST {endpoint}/api/election/toggle \
--header 'Authorization: Bearer xxxx' \
--header 'Content-Type: application/json' \
-d '{
  "enable": true
}'
```

Expected Response:
```
# Success
{
  "status": "ok",
  "enable": <true|false>
}
```


### Response Detail
| Parameter | Description |
------------|--------------
| status | status of request |
| enable | status of election (open or close as true/false) |
    

## GET Election Result

Example curl:
```
curl -X POST {endpoint}/api/election/toggle \
--header 'Authorization: Bearer xxxx' \
--header 'Content-Type: application/json' \
-d '{
  "enable": true
}'
```

Expected Response:
```
# Success
[
  {
    "id": "1",
    "votedCount": 10
  },
  ...
]
```

### Response Detail
| Parameter | Description |
------------|--------------
| id | id of candidate |
| votedCount | how many votes did the candidate get |
    
## GET Election Result

Example curl:
```
curl -X GET {endpoint}/api/election/result \
--header 'Authorization: Bearer xxxx' 

```

Expected Response:
```
[
  {
    "id": "1",
    "name": "Elon Musk",
    "dob": "June 28, 1971",
    "bioLink": "https://en.wikipedia.org/wiki/Elon_Musk",
    "imageLink": "https://upload.wikimedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 3,
    "percentage": "75%"
  },
  {
    "id": "2",
    "name": "Jeff Bezos",
    "dob": "January 12, 1964",
    "bioLink": "https://en.wikipedia.org/wiki/Jeff_Bezos",
    "imageLink": "https://pbs.twimg.com/profile_images/669103856106668033/UF3cgUk4_400x400.jpg",
    "policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
    "votedCount": 1,
    "percentage": "25%"
  }
]
```

### Response Detail
| Parameter | Description |
------------|--------------
| id | ID of the candidate |
| name | Candidate's name |
| dob | Candidate's Date of birth |
| bioLink | Candidate's biography link |
| imageLink | Candidate's image |
| policy | Candidate's policy |
| votedCount | Candidate's vote count |
| percentage | Candidate's vote percentage |

## GET Exported Result (download)

Example curl:
```
curl -X GET {endpoint}/api/election/export \
--header 'Authorization: Bearer xxxx' 

```

Expected Response:

CSV file download with columns:
- Candidate id
- National id


## Real-time Vote Stream

Websocket stream for real-time vote count

Update Speed real-time

Payload:
```
{
  "id": "1",     // Candidate ID
  "votedCount": 10      // Total number of votes for this particular candidate
}