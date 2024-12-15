**所有请求与响应数据格式均为JSON**

# 登录接口
接口功能：用户登录系统
请求方法：POST
请求URL ：http://localhost:8888/login
## 请求参数：
- username：字符串，必填，代表用户名
- password：字符串，必填，代表密码
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 400：请求错误
- 500：用户不存在

# 注册接口：
接口功能：用户注册系统
请求方法：POST
请求URL：http://localhost:8888/register
## 请求参数：
- username：字符串，必填，代表用户名
- password：字符串，必填，代表密码
- nickname：字符串，选填，代表昵称
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 400：请求错误

# 用户数据修改接口：
接口功能：用户昵称密码修改
请求方法：POST
请求URL：http://localhost:8888/profile
## 请求参数：
- password：字符串，选填，代表密码
- nickname：字符串，选填，代表昵称
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 400：请求错误
- 500：修改错误

# 留言查询接口：
接口功能：使用留言id查询所有留言，包括子留言
请求方法：GET
请求URL：http://localhost:8888/getmessage?id=
## 请求参数：
- id：整形，查询的留言id
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 500：查找错误

# 留言发布接口：
接口功能：发布留言
请求方法：POST
请求URL：http://localhost:8888/postmessage
## 请求参数：
- context：字符串，必填，留言内容
- parentid：整型，选填，父留言id    
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 400：请求错误
- 500：发布错误

# 留言删除接口：
接口功能：删除留言
请求方法：GET
请求URL：http://localhost:8888/deletemessage?id=
## 请求参数：
- id：整数，删除留言的id
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 400：请求错误
- 500：发布错误

# 点赞接口：
接口功能：给留言点赞，如果已经点赞则会取消点赞
请求方法：GET
请求URL：http://localhost:8888/likemessage?id=
## 请求参数：
- id：整数，点赞留言的id
## 响应数据：
- code：整数，状态响应码
- message：字符串，响应信息
## 状态码：
- 200：成功
- 400：请求错误
- 500：点赞错误

