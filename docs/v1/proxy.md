## proxy

### url

anyMethod   /v1/proxy/clusters/{cluster}/*any

### path

|  Field  | Description  |
| :-----: | :----------: |
| cluster | cluster name |

## example

### List namespace

#### url

GET /v1/proxy/clusters/{cluster}/api/v1/namespaces

#### path

|  Field  | Description  |
| :-----: | :----------: |
| cluster | cluster name |

#### response

| Code | Description |
| :--: | :---------: |
| 200  |     OK      |

```json
{
    "kind": "Namespace",
    "apiVersion": "v1",
    "metadata": {
        "selfLink": "/api/v1/namespaces",
        "resourceVersion": "194313498"
    },
    "items": [
        {
            "metadata": {
                "name": "testxxx",
                "selfLink": "/api/v1/namespaces/yorktest999",
                "uid": "1c1f6779-8dda-46e9-a60c-e7ddb9ab2b0f",
                "resourceVersion": "192919535",
                "creationTimestamp": "2021-07-23T07:48:59Z",
                "labels": {
                    "aa": "aa"
                },
                "annotations": {
                    "dce.daocloud.io/limitRangeInited": "true"
                }
            },
            "spec": {
                "finalizers": [
                    "kubernetes"
                ]
            },
            "status": {
                "phase": "Active"
            }
        }
    ]
}
```

