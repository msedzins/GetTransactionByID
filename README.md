# GetTransactionByID

Example of how to parse qscc.GetTransactionByID output.

# Instruction

### Step 1 - Call GetTransacionById 

```
peer chaincode query -o 127.0.0.1:6050 -C mychannel -n qscc -c '{"function":"GetTransactionByID","Args":["mychannel", "fd30a31a7acd893c6bd0bac9eb68005df6de7311c9b8a196c8c500cd428abfa7"]}' --tls --cafile "${PWD}"/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt --hex > hex.txt
```

**Note:**
Examplary hex.txt file with response attached

### Step 2 - Run application

```
cat hex.txt | go run main.go
```

### Step 3 - Output

Application will print out quite a few lines with the analysis of the response. 

For example, this will show input parameters to the chaincode call (args are []byte, should be cast to string):
```
Input parameters to the chaincode:
{
    "chaincode_spec": {
        "type": 1,
        "chaincode_id": {
            "name": "basic_noinvoke2"
        },
        "input": {
            "args": [
                "Q3JlYXRlQXNzZXQ=",
                "NA==",
                "Ymx1ZQ==",
                "MzU=",
                "amVycnk=",
                "MTAwMA=="
            ]
        }
    }
}
```

R/W set:
```
{
    "reads": [
        {
            "key": "namespaces/fields/basic_noinvoke2/Sequence",
            "version": {
                "block_num": 52
            }
        }
    ]
}

{
    "reads": [
        {
            "key": "4"
        }
    ],
    "writes": [
        {
            "key": "4",
            "value": "eyJBcHByYWlzZWRWYWx1ZSI6MTAwMCwiQ29sb3IiOiJibHVlIiwiSUQiOiI0IiwiT3duZXIiOiJqZXJyeSIsIlNpemUiOjM1fQ=="
        }
    ]
}
```


