# EV102

### The EV101 program is designed to implement the following key modules:

- [x] **1) INSERT DATA**
- [x] **2) SIM QUERY API**
- [x] **3) Testing** (Optional)

The program incorporates a runtime configuration through a dedicated config file with the following settings:

```yaml
app:
  version: "0.0.1"
rdbs:
  addr: "127.0.0.1:6370,127.0.0.1:6371,127.0.0.1:6372,127.0.0.1:6373,127.0.0.1:6374"
crdb_addr: "127.0.0.1:6379"
api_server: "127.0.0.1"
enable_load: true
data_file: "/a/b/c/data.csv"
```

- The program should expect config file as an argument
```go
go run main.go -config config.yaml
```


___

### 1) INSERT DATA
For the data insertion module, a CSV file containing information for 2364 SIMs is provided. Each line in the CSV file consists of 8 columns, as illustrated in the example below:

> Example:

| MSISDN  | IMSI |ICCID |Secret |TAC |EID |CID |IMEI | BundleID
| ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
|811502214250| 217500013105250| 26524849111319248| b59c45efc7dc9022| 64| 76| 6089| 349554520566217| x |
|811502210200| 217500013101200| 26524849110335100| 6110dfd138c5136f| 28| 63| 5428| 906356868756491| x |
|811502213300| 217500013104300| 26524849111088400| 5d672e3fa975e368| 46| 49| 8457| 655918658555720| x |


- Parse the data.csv file.
- Save each line into the respective Redis instance based on manual sharding (e.g., MSISDN 811502214250 saved in Redis node #0).
- Ensure the key does not exist before insertion. `(Optional)`
- Utilize proper Redis data types for storing the data (Hashes). [Redis Docs](https://redis.io/docs/data-types/hashes/)
- Randomly assign a BundleID from the provided bundle data.

> Insert bundles manually:
- You need also to insert the bundle data into the database, you can insert it into a specefic redis called `crdb`
- `docker run --rm --name crdb -v /home/ev101/redis/crdb/data:/data -p 6379:6379 redis redis-server`
- `redis-cli`
- `>hset bundle:1021 ul 4194304 dl 2097152 quota 3221225472 duration 2629800 label Capela type Data`


### Bundles
| ID  | UL | DL |Quota |Duration |Label |Type |
| ---- | ------- | ------- | ---------- | ---------- | ---------- | ---------- |
| 1021 | 4194304 | 2097152 | 3221225472 | 2629800 | Capela | Data
| 1022 | 2621440 | 10485760 | 1073741824 | 2629800 | Acrux | Data
| 1023 | 5242880 | 5242880 | 59055800320 | 2629800 | Castor | Data
| 1024 | 1048576 | 1048576 | 838860800 | 2629800 | Vega | Data
| 1025 | 5242880 | 5242880 | 1072668082176 | 2629800 | Antares | Data
| 1026 | 2621440 | 2097152 | 32212254720 | 2629800 | Rigel | Data
| 1027 | 5242880 | 5242880 | 8589934592 | 2629800 | Altair | Data
| 1028 | 1048576 | 1048576 | 10737418240 | 2629800 | Spica | Data


- After the insertion process is complete, generate a brief report containing:
  * Duration
  * Number of inserted records
___

### 2) SIM QUERY API

This module should provide a simple rest api with the following endpoints, for the http you can use 

- **GET** `/ev101/api/sims/811502214250`

`Success Response`

  ```json
  {
    "msisdn": 811502214250,
    "imsi": 217500013105250,
    "iccid": 26524849111319250,
    "secret": "b59c45efc7dc9022",
    "tac": 64,
    "eid": 76,
    "cid": 6089,
    "imei": 349554520566217,
    "bundle": {
        "id": 1025,
        "ul": 5242880,
        "dl": 5242880,
        "quota": 1072668082176,
        "duration": 2629800,
        "label": "Antares",
        "type": "data"
    }
}
  ```

`Failed Response`
```json
{
    "status": "failed",
    "error": "sim does not exist"
}
```

- **DELETE** `/ev101/api/sims/811502214250`

`Success Response`
```json
{
    "status": "success"
}
```
___
### 3) Testing (Optional)

To ensure the functionality of the program, a simple test unit has been incorporated. This test involves inserting a few records and verifying the successful insertion. Execute the following command to run the test:
```go
go test ./...
```
Ensure that the test passes with the result **OK**

___

### Prerequisites

Ensure that 5 Redis nodes are operational and running on the local machine. Docker can be utilized for this purpose.

> Example for running a docker container on local machine:

```bash
docker run --rm --name redis0 -v /home/ev101/redis/redis0/data:/data -p 6370:6379 redis redis-server
docker run --rm --name redis1 -v /home/ev101/redis/redis1/data:/data -p 6371:6379 redis redis-server
.
.
.
docker run --rm --name redis4 -v /home/ev101/redis/redis4/data:/data -p 6374:6379 redis redis-server
```
