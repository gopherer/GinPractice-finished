# GinProjectInstance-unfinished

![image](https://github.com/gopherer/GinProjectInstance-unfinished/raw/main/readmePhoto/structure-chart.png)

须知：1、将用户注册数据插入数据库表user_Register同时会在user_info表中将新建一条数据把id置为用户注册时的id号,以此关联二者 UserName 
需要设置默认值 以确保主键不会冲突

2、用户register无须在login 直接可操作userInfo
刚开写时没区分login和register, userInfo的功能都依托与register 为避免重写 用户采用login登录后 将直接视为register 操作userInfo
虽然login可以开是另类的register  但还是最好区别   
   
3、Go系列：结构体标签  https://juejin.cn/post/7005465902804123679#heading-11
                        https://www.cnblogs.com/chnmig/p/11382390.html
  
  //binding tag 可用 但不好 最好专门写个tool函数处理前端输入的数据
  

  //使用匿名字段 在结构体映射为数据表时无法得到实际效果
  

  //form tag 可接收html form表单 application/x-www-form-urlencoded的类容 可以以接受json文本的内容
  

  //json tag 只可以接收json的文本内容

遇到的有意思的问题：

问题：struct结构体无法映射为数据库表  报错点err = engine.Sync2(new(model.UserLogin))

原因：报错点并非实际出错点，读取json文件中charset字段，在映射为结构体时，charset字段的tag char_set 导致映射结果中DatabaseConfig结构体Charset字段为空值
     影响代码：conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
     等价于：root:root@tcp(127.0.0.1:3306)/goWeb?charset=
     engine, err := xorm.NewEngine(database.Driver, conn) 可以正常返回数据库engine，err 且err 为nil, 此处并不会报错.

措施：tools文件夹下Config.go文件中DatabaseConfig结构体. 更改如下：Charset  string `json:"charset"`


问题：Postman Post 发送json数据 无法映射到 结构体中

原因：发送的json数据的key值和结构体中tag标签--json定义是数据不一致

措施：发送的json数据的key值和结构体中tag标签--json定义是数据保持一致


