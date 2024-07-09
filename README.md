# voting application

- voters can register A_ID (verified)
- candidates (uploaded)
- voters have to come to center
- voter will get a unique code (verification)
- summary 
  - candidates with voting accuracy
  - demographic of voters (gender, age)
  - total active voting population

domain modelling

voter
    id
    name
    a_id index
    voting_status HAS_REGISTERED, HAS_APPLIED, HAS_VOTED
    gender
    age
    code 
    center_id

candidate
    id
    name
    party_id

center
    id
    location
    name 
    code

voter_candidate_mapping
    id
    candidate_id
    voter_id

system_internal #1
API -> /register
payload

```json
{
    "name": "lorem",
    "gender": "Male",
    "age": 34,
    "AID": "ipsum"
}
```

```json
{
    "AID": "ipsum",
    "status": "REGISTERED"
}
```

API -> /applyForVoting
payload 

```json
{
    "AID": "lorem",
    "center_id": 12
}
```

response

```json
{
    "UNIQUE_CODE": "123"
}
```

API -> /vote
payload

```json
{
    "AID": "lorem",
    "candidate_id": 123,
    "UNIQUE_CODE": "123"
}
```

API -> /summary

```json
{
  
}
```

