## List cluster

### url

GET /v1/clusters

### header

Basic Auth(default admin/admin)

### response

| Code | Description |
| :--: | :---------: |
| 200  |     OK      |

```json
[
    {
        "Id": "1",
        "Name": "k8s",
        "Url": "https://",
        "Token": "eyJhbGciOiJSUzI1NiIsImtpZCI6ImRubm1RaEdmWmNPTGhsdzUwWDdOTUhadWJvRUdzZ05vcXNrajNsS3VUaHMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjYXR0bGUtc3lzdGVtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImNhdHRsZS10b2tlbi02a3NmYyIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJjYXR0bGUiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJmNjAzZjg0MS1lOGRmLTRlNTgtOTU4YS1iN2M2Y2Q2NTkwYTYiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6Y2F0dGxlLXN5c3RlbTpjYXR0bGUifQ.gQ0cTf9QHPEKRLckzHbsotOWKhUedcZfcYucEd8tY7lVvbaEQ3lXRriAYPgadqevzOojiX3vouRxNmt8FUyL89SbbQY5tUy7timVpKKzav6xVqIA86mUzMHqCd72ssOhSzjtiHQfzDW0ZhTIdZLYxffcQy7J6UnryLLRQpKTAusDKSWZFbuKLVkSUFKG28c1HwfRUZIV_fqyG419bvdKV7GtcpM-gak77-RxjIe0RtdGVZYd-M2EGtNmkkS-TOWp2lQXgiic57oclxnTdKMgBtngThwl7ZjQkLiPZT77LjLADsbZlhdW9MPMc0lGSIjCXqmULPnVGq6pzn3NINIwbg"
    },
    {
        "Id": "2",
        "Name": "k8s1",
        "Url": "https://",
        "Token": "eyJhbGciOiJSUzI1NiIsImtpZCI6ImRubm1RaEdmWmNPTGhsdzUwWDdOTUhadWJvRUdzZ05vcXNrajNsS3VUaHMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJtdWx0aS1jbHVzdGVyLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJtdWx0aS1jbHVzdGVyLXRva2VuLWo2cnFzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Im11bHRpLWNsdXN0ZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJjYzYwOGQ3Yy05MzhhLTRhZjUtYTU1ZC0xYjlhYmEwMGY1MDciLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6bXVsdGktY2x1c3Rlci1zeXN0ZW06bXVsdGktY2x1c3RlciJ9.u_NXHR2zr_MWDVc_1kZuH59cNO0D5-I8NzirWqA1VE--9iMBXNQ0SZr4u7jEO62sjZMhpuynWzpwtOcKbtjjgPE64Z7W_OgMd2faJvkdwU-5uhLlGxgrVJW8E_TgXRo_LxsCdYJLPZbPOQ6lDyIYw2aonBkvpZxsJnmEQmDeTKoMuyF9D-zeIVEJBNGZSY9GH4PZtYUor8N1TZS9RcVGwK4IyLlh9oD9evuXqzHM45TDp7_7Wh31kLTxGtS_n18lZgi3Sm5rudoGS3Enn5gY8RbLaNyNo3fAylDitXJ7GhXuFbn0gEGjL2qFalMO4JlHqo4EMsn_lp3UyqyKrF9l0A"
    }
]
```

## Get cluster

### url

GET /v1/clusters/{cluster}

### path

|  Field  | Description  |
| :-----: | :----------: |
| Cluster | cluster name |

### header

Basic Auth(default admin/admin)

### response

| Code | Description |
| :--: | :---------: |
| 200  |     OK      |

```json
{
    "Id": "3",
    "Name": "k8s",
    "Url": "https://",
    "Token": "eyJhbGciOiJSUzI1NiIsImtpZCI6ImRubm1RaEdmWmNPTGhsdzUwWDdOTUhadWJvRUdzZ05vcXNrajNsS3VUaHMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJtdWx0aS1jbHVzdGVyLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJtdWx0aS1jbHVzdGVyLXRva2VuLWo2cnFzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Im11bHRpLWNsdXN0ZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJjYzYwOGQ3Yy05MzhhLTRhZjUtYTU1ZC0xYjlhYmEwMGY1MDciLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6bXVsdGktY2x1c3Rlci1zeXN0ZW06bXVsdGktY2x1c3RlciJ9.u_NXHR2zr_MWDVc_1kZuH59cNO0D5-I8NzirWqA1VE--9iMBXNQ0SZr4u7jEO62sjZMhpuynWzpwtOcKbtjjgPE64Z7W_OgMd2faJvkdwU-5uhLlGxgrVJW8E_TgXRo_LxsCdYJLPZbPOQ6lDyIYw2aonBkvpZxsJnmEQmDeTKoMuyF9D-zeIVEJBNGZSY9GH4PZtYUor8N1TZS9RcVGwK4IyLlh9oD9evuXqzHM45TDp7_7Wh31kLTxGtS_n18lZgi3Sm5rudoGS3Enn5gY8RbLaNyNo3fAylDitXJ7GhXuFbn0gEGjL2qFalMO4JlHqo4EMsn_lp3UyqyKrF9l0A"
}
```

## Create cluster

### url

POST /v1/clusters

### header

Basic Auth(default admin/admin)

### body

```json
{
    "id":"3",
    "name":"k8s",
    "url":"https://",
    "token":"eyJhbGciOiJSUzI1NiIsImtpZCI6ImRubm1RaEdmWmNPTGhsdzUwWDdOTUhadWJvRUdzZ05vcXNrajNsS3VUaHMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJtdWx0aS1jbHVzdGVyLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJtdWx0aS1jbHVzdGVyLXRva2VuLWo2cnFzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Im11bHRpLWNsdXN0ZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJjYzYwOGQ3Yy05MzhhLTRhZjUtYTU1ZC0xYjlhYmEwMGY1MDciLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6bXVsdGktY2x1c3Rlci1zeXN0ZW06bXVsdGktY2x1c3RlciJ9.u_NXHR2zr_MWDVc_1kZuH59cNO0D5-I8NzirWqA1VE--9iMBXNQ0SZr4u7jEO62sjZMhpuynWzpwtOcKbtjjgPE64Z7W_OgMd2faJvkdwU-5uhLlGxgrVJW8E_TgXRo_LxsCdYJLPZbPOQ6lDyIYw2aonBkvpZxsJnmEQmDeTKoMuyF9D-zeIVEJBNGZSY9GH4PZtYUor8N1TZS9RcVGwK4IyLlh9oD9evuXqzHM45TDp7_7Wh31kLTxGtS_n18lZgi3Sm5rudoGS3Enn5gY8RbLaNyNo3fAylDitXJ7GhXuFbn0gEGjL2qFalMO4JlHqo4EMsn_lp3UyqyKrF9l0A"
}
```

### response

| Code | Description |
| :--: | :---------: |
| 200  |     OK      |

```json
{
    "message": "success"
}
```

## Update cluster

### url

PUT /v1/clusters/{cluster}

### path

|    Field     | Description |
| :----------: | :---------: |
| cluster name |             |

### header

Basic Auth(default admin/admin)

### body

```json
{
    "id":"3",
    "name":"k8s",
    "url":"https://",
    "token":"eyJhbGciOiJSUzI1NiIsImtpZCI6ImRubm1RaEdmWmNPTGhsdzUwWDdOTUhadWJvRUdzZ05vcXNrajNsS3VUaHMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJtdWx0aS1jbHVzdGVyLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJtdWx0aS1jbHVzdGVyLXRva2VuLWo2cnFzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Im11bHRpLWNsdXN0ZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJjYzYwOGQ3Yy05MzhhLTRhZjUtYTU1ZC0xYjlhYmEwMGY1MDciLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6bXVsdGktY2x1c3Rlci1zeXN0ZW06bXVsdGktY2x1c3RlciJ9.u_NXHR2zr_MWDVc_1kZuH59cNO0D5-I8NzirWqA1VE--9iMBXNQ0SZr4u7jEO62sjZMhpuynWzpwtOcKbtjjgPE64Z7W_OgMd2faJvkdwU-5uhLlGxgrVJW8E_TgXRo_LxsCdYJLPZbPOQ6lDyIYw2aonBkvpZxsJnmEQmDeTKoMuyF9D-zeIVEJBNGZSY9GH4PZtYUor8N1TZS9RcVGwK4IyLlh9oD9evuXqzHM45TDp7_7Wh31kLTxGtS_n18lZgi3Sm5rudoGS3Enn5gY8RbLaNyNo3fAylDitXJ7GhXuFbn0gEGjL2qFalMO4JlHqo4EMsn_lp3UyqyKrF9l0A"
}
```

### response

| Code | Description |
| :--: | :---------: |
| 200  |     OK      |

```json
{
    "message": "success"
}
```

## Delete cluster

### url

DELETE /v1/clusters/{cluster}

### path

|    Field     | Description |
| :----------: | :---------: |
| cluster name |             |

### header

Basic Auth(default admin/admin)

### response

| Code | Description |
| :--: | :---------: |
| 200  |     OK      |

```json
{
    "message": "success"
}
```

