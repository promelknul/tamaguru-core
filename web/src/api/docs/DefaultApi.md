# DefaultApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**authLoginPost**](#authloginpost) | **POST** /auth/login | |
|[**authRegisterPost**](#authregisterpost) | **POST** /auth/register | |

# **authLoginPost**
> AuthResp authLoginPost(loginReq)


### Example

```typescript
import {
    DefaultApi,
    Configuration,
    LoginReq
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let loginReq: LoginReq; //

const { status, data } = await apiInstance.authLoginPost(
    loginReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loginReq** | **LoginReq**|  | |


### Return type

**AuthResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **authRegisterPost**
> AuthResp authRegisterPost(registerReq)


### Example

```typescript
import {
    DefaultApi,
    Configuration,
    RegisterReq
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let registerReq: RegisterReq; //

const { status, data } = await apiInstance.authRegisterPost(
    registerReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **registerReq** | **RegisterReq**|  | |


### Return type

**AuthResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

