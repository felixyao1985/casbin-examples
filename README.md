#casbin 随手记

- 配置文件使用 PERM 元模型。PERM 表示策略（Policy）、效果（Effect）、请求（Request）和匹配器（Matchers）。
                
                其中定义了请求和策略来表示主体，客体和动作。
                在本例中，主体表示用户角色，客体表示访问路径，action 表示请求方法（例：GET, POST 等）。
                匹配器定义了策略是如何匹配的，可以通过直接定义主体，或者使用像 keyMatch 这样的帮助方法，它也可以匹配通配符。
                casbin 实际比这个简单的例子要强大得多，你可以用声明的方式定义各种自定义功能来达到轻松切换和维护鉴权配置的效果。

    -                 
- 策略文件就是一个简单的 csv 文件，描述了哪些角色可以访问哪些路径等（也可以使用数据库来管理）
    -       
  
            
    - 
                
                