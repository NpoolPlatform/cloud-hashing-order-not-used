# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-order.proto](#npool/cloud-hashing-order.proto)
    - [CreateGasPayingRequest](#cloud.hashing.order.v1.CreateGasPayingRequest)
    - [CreateGasPayingResponse](#cloud.hashing.order.v1.CreateGasPayingResponse)
    - [CreateGoodPayingRequest](#cloud.hashing.order.v1.CreateGoodPayingRequest)
    - [CreateGoodPayingResponse](#cloud.hashing.order.v1.CreateGoodPayingResponse)
    - [CreateOrderRequest](#cloud.hashing.order.v1.CreateOrderRequest)
    - [CreateOrderResponse](#cloud.hashing.order.v1.CreateOrderResponse)
    - [GasPaying](#cloud.hashing.order.v1.GasPaying)
    - [GetGasPayingRequest](#cloud.hashing.order.v1.GetGasPayingRequest)
    - [GetGasPayingResponse](#cloud.hashing.order.v1.GetGasPayingResponse)
    - [GetGoodPayingRequest](#cloud.hashing.order.v1.GetGoodPayingRequest)
    - [GetGoodPayingResponse](#cloud.hashing.order.v1.GetGoodPayingResponse)
    - [GetOrderDetailRequest](#cloud.hashing.order.v1.GetOrderDetailRequest)
    - [GetOrderDetailResponse](#cloud.hashing.order.v1.GetOrderDetailResponse)
    - [GetOrderRequest](#cloud.hashing.order.v1.GetOrderRequest)
    - [GetOrderResponse](#cloud.hashing.order.v1.GetOrderResponse)
    - [GetOrdersByAppRequest](#cloud.hashing.order.v1.GetOrdersByAppRequest)
    - [GetOrdersByAppResponse](#cloud.hashing.order.v1.GetOrdersByAppResponse)
    - [GetOrdersByAppUserRequest](#cloud.hashing.order.v1.GetOrdersByAppUserRequest)
    - [GetOrdersByAppUserResponse](#cloud.hashing.order.v1.GetOrdersByAppUserResponse)
    - [GetOrdersByGoodRequest](#cloud.hashing.order.v1.GetOrdersByGoodRequest)
    - [GetOrdersByGoodResponse](#cloud.hashing.order.v1.GetOrdersByGoodResponse)
    - [GetOrdersDetailByAppRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppRequest)
    - [GetOrdersDetailByAppResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppResponse)
    - [GetOrdersDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest)
    - [GetOrdersDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse)
    - [GetOrdersDetailByGoodRequest](#cloud.hashing.order.v1.GetOrdersDetailByGoodRequest)
    - [GetOrdersDetailByGoodResponse](#cloud.hashing.order.v1.GetOrdersDetailByGoodResponse)
    - [GoodPaying](#cloud.hashing.order.v1.GoodPaying)
    - [Order](#cloud.hashing.order.v1.Order)
    - [OrderDetail](#cloud.hashing.order.v1.OrderDetail)
    - [UpdateGasPayingRequest](#cloud.hashing.order.v1.UpdateGasPayingRequest)
    - [UpdateGasPayingResponse](#cloud.hashing.order.v1.UpdateGasPayingResponse)
    - [UpdateGoodPayingRequest](#cloud.hashing.order.v1.UpdateGoodPayingRequest)
    - [UpdateGoodPayingResponse](#cloud.hashing.order.v1.UpdateGoodPayingResponse)
    - [UpdateOrderRequest](#cloud.hashing.order.v1.UpdateOrderRequest)
    - [UpdateOrderResponse](#cloud.hashing.order.v1.UpdateOrderResponse)
    - [VersionResponse](#cloud.hashing.order.v1.VersionResponse)
  
    - [CloudHashingOrder](#cloud.hashing.order.v1.CloudHashingOrder)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-order.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-order.proto



<a name="cloud.hashing.order.v1.CreateGasPayingRequest"></a>

### CreateGasPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGasPayingResponse"></a>

### CreateGasPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGoodPayingRequest"></a>

### CreateGoodPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGoodPayingResponse"></a>

### CreateGoodPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateOrderRequest"></a>

### CreateOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.CreateOrderResponse"></a>

### CreateOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.GasPaying"></a>

### GasPaying



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| PlatformTransactionID | [string](#string) |  |  |
| DurationMinutes | [uint32](#uint32) |  |  |
| Used | [bool](#bool) |  |  |






<a name="cloud.hashing.order.v1.GetGasPayingRequest"></a>

### GetGasPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetGasPayingResponse"></a>

### GetGasPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.GetGoodPayingRequest"></a>

### GetGoodPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetGoodPayingResponse"></a>

### GetGoodPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.GetOrderDetailRequest"></a>

### GetOrderDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrderDetailResponse"></a>

### GetOrderDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) |  |  |






<a name="cloud.hashing.order.v1.GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppRequest"></a>

### GetOrdersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppResponse"></a>

### GetOrdersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppUserRequest"></a>

### GetOrdersByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppUserResponse"></a>

### GetOrdersByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersByGoodRequest"></a>

### GetOrdersByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByGoodResponse"></a>

### GetOrdersByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppRequest"></a>

### GetOrdersDetailByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppResponse"></a>

### GetOrdersDetailByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest"></a>

### GetOrdersDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse"></a>

### GetOrdersDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByGoodRequest"></a>

### GetOrdersDetailByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByGoodResponse"></a>

### GetOrdersDetailByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GoodPaying"></a>

### GoodPaying



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| PlatformTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| Discount | [uint32](#uint32) |  |  |
| SpecialReductionAmount | [double](#double) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| GoodPayID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| CompensateMinutes | [uint32](#uint32) |  |  |
| CompensateElapsedMinutes | [uint32](#uint32) |  |  |
| GasStart | [uint32](#uint32) |  |  |
| GasEnd | [uint32](#uint32) |  |  |
| GasPayIDs | [string](#string) | repeated |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.OrderDetail"></a>

### OrderDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| Discount | [uint32](#uint32) |  |  |
| SpecialReductionAmount | [double](#double) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| GoodPaying | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| CompensateMinutes | [uint32](#uint32) |  |  |
| CompensateElapsedMinutes | [uint32](#uint32) |  |  |
| GasStart | [uint32](#uint32) |  |  |
| GasEnd | [uint32](#uint32) |  |  |
| GasPayings | [GasPaying](#cloud.hashing.order.v1.GasPaying) | repeated |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.UpdateGasPayingRequest"></a>

### UpdateGasPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.UpdateGasPayingResponse"></a>

### UpdateGasPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.UpdateGoodPayingRequest"></a>

### UpdateGoodPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.UpdateGoodPayingResponse"></a>

### UpdateGoodPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.UpdateOrderRequest"></a>

### UpdateOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.UpdateOrderResponse"></a>

### UpdateOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.order.v1.CloudHashingOrder"></a>

### CloudHashingOrder
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#cloud.hashing.order.v1.VersionResponse) | Method Version |
| CreateGoodPaying | [CreateGoodPayingRequest](#cloud.hashing.order.v1.CreateGoodPayingRequest) | [CreateGoodPayingResponse](#cloud.hashing.order.v1.CreateGoodPayingResponse) |  |
| GetGoodPaying | [GetGoodPayingRequest](#cloud.hashing.order.v1.GetGoodPayingRequest) | [GetGoodPayingResponse](#cloud.hashing.order.v1.GetGoodPayingResponse) |  |
| UpdateGoodPaying | [UpdateGoodPayingRequest](#cloud.hashing.order.v1.UpdateGoodPayingRequest) | [UpdateGoodPayingResponse](#cloud.hashing.order.v1.UpdateGoodPayingResponse) |  |
| CreateGasPaying | [CreateGasPayingRequest](#cloud.hashing.order.v1.CreateGasPayingRequest) | [CreateGasPayingResponse](#cloud.hashing.order.v1.CreateGasPayingResponse) |  |
| GetGasPaying | [GetGasPayingRequest](#cloud.hashing.order.v1.GetGasPayingRequest) | [GetGasPayingResponse](#cloud.hashing.order.v1.GetGasPayingResponse) |  |
| UpdateGasPaying | [UpdateGasPayingRequest](#cloud.hashing.order.v1.UpdateGasPayingRequest) | [UpdateGasPayingResponse](#cloud.hashing.order.v1.UpdateGasPayingResponse) |  |
| CreateOrder | [CreateOrderRequest](#cloud.hashing.order.v1.CreateOrderRequest) | [CreateOrderResponse](#cloud.hashing.order.v1.CreateOrderResponse) |  |
| GetOrder | [GetOrderRequest](#cloud.hashing.order.v1.GetOrderRequest) | [GetOrderResponse](#cloud.hashing.order.v1.GetOrderResponse) |  |
| GetOrderDetail | [GetOrderDetailRequest](#cloud.hashing.order.v1.GetOrderDetailRequest) | [GetOrderDetailResponse](#cloud.hashing.order.v1.GetOrderDetailResponse) |  |
| UpdateOrder | [UpdateOrderRequest](#cloud.hashing.order.v1.UpdateOrderRequest) | [UpdateOrderResponse](#cloud.hashing.order.v1.UpdateOrderResponse) |  |
| GetOrdersByAppUser | [GetOrdersByAppUserRequest](#cloud.hashing.order.v1.GetOrdersByAppUserRequest) | [GetOrdersByAppUserResponse](#cloud.hashing.order.v1.GetOrdersByAppUserResponse) |  |
| GetOrdersByApp | [GetOrdersByAppRequest](#cloud.hashing.order.v1.GetOrdersByAppRequest) | [GetOrdersByAppResponse](#cloud.hashing.order.v1.GetOrdersByAppResponse) |  |
| GetOrdersByGood | [GetOrdersByGoodRequest](#cloud.hashing.order.v1.GetOrdersByGoodRequest) | [GetOrdersByGoodResponse](#cloud.hashing.order.v1.GetOrdersByGoodResponse) |  |
| GetOrdersDetailByAppUser | [GetOrdersDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest) | [GetOrdersDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse) |  |
| GetOrdersDetailByApp | [GetOrdersDetailByAppRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppRequest) | [GetOrdersDetailByAppResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppResponse) |  |
| GetOrdersDetailByGood | [GetOrdersDetailByGoodRequest](#cloud.hashing.order.v1.GetOrdersDetailByGoodRequest) | [GetOrdersDetailByGoodResponse](#cloud.hashing.order.v1.GetOrdersDetailByGoodResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

