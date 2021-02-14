# DB Model Document

## sites

|name|type|comment|
|:---|:---|:------|
|id |INT, PRIMARY KEY, AUTO INC |监控站点ID|
|tag |CHAR(36),PRIMARY KEY|uuid，用于客户端注册|
|name |VARCHAR(63)|站点备注 |
|public |INT |是否公开展示|
|last_check   |datetime|上次检查时间|
|first_check  |datetime|本次连续在线，初次检查时间(如果为null则为不在线)|

## group

|name|type|comment|
|:---|:---|:------|
|id |INT, PRIMARY KEY, AUTO INC |GroupID|
|link|CHAR(36)| 群组唯一链接|

## site_group

|site_id|INT, FOREIGN KEY -> sites.id|site id|
|group_id |INT, FOREIGN KEY -> group.id|group id|

## conditions

|name|type|comment|
|:---|:---|:------|
|id |INT, PRIMARY KEY, AUTO INC |ConditonID|
|site_id|INT, FOREIGN KEY -> sites.id|site id|
|is_enabled|INT|该检查条件是否被启用|
|link |TEXT|要检查的链接 |
|res_type|INT| 检查的资源类型|
|res_threshold|FLOAT|要检查的资源阈值|

## notification

|name|type|comment|
|:---|:---|:------|
|id |INT, PRIMARY KEY, AUTO INC |NotificationID|
|info|VCHAR(1023), in json|通知类型及参数|


## site_notification

|name|type|comment|
|:---|:---|:------|
|id |INT, PRIMARY KEY, AUTO INC |NotificationID|
|site_id|INT, FOREIGN KEY -> sites.id|site id|

## log
|name|type|comment|
|:---|:---|:------|
|id |INT, PRIMARY KEY, AUTO INC |LogID|
|content| TEXT|监控日志信息|
|create_at| datetime| 创建时间|