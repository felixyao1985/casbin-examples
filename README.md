#casbin 随手记

相关路由设定 请参阅 [go-study](https://github.com/felixyao1985/go-study "Title")

casbin 配置文件的规范 请参阅 [supported-models](https://casbin.org/docs/zh-CN/supported-models "Title")

- 配置文件使用 PERM 元模型。PERM 表示策略（Policy）、效果（Effect）、请求（Request）和匹配器（Matchers）。
                
        其中定义了请求和策略来表示主体，客体和动作。
        在本例中，主体表示用户角色，客体表示访问路径，action 表示请求方法（例：GET, POST 等）。
        匹配器定义了策略是如何匹配的，可以通过直接定义主体，或者使用像 keyMatch 这样的帮助方法，它也可以匹配通配符。
        casbin 实际比这个简单的例子要强大得多，你可以用声明的方式定义各种自定义功能来达到轻松切换和维护鉴权配置的效果。

    <pre><code>
    PERM 元模型配置文件 auth_model.conf
    
    [request_definition]
    
    r = sub, obj, act
    
    [policy_definition]
    p = sub, obj, act
    
    [policy_effect]
    e = some(where (p.eft == allow))
    
    [matchers] #核心权限验证部分，可以根据自己需要进行设定
    m = r.sub == p.sub && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
    </pre></code>
                   
- 策略文件就是一个简单的 csv 文件，描述了哪些角色可以访问哪些路径等（也可以使用数据库来管理）
    <pre><code>
    p, admin, /*, *
    p, anonymous, /login, *
    p, anonymous, /index, *
    p, member,/logout, *
    p, member, /home/*, *
    p, member, /home, *
    
    
    权限判断代码
    authEnforcer, err := casbin.NewEnforcerSafe("./src/casbin-examples/roles/auth_model.conf", "./src/casbin-examples/roles/policy.csv")
    sub := "member"
    obj := "/home"
    act := "GET"

    if authEnforcer.Enforce(sub, obj, act) == true {
        fmt.Println("路由成功？")
    } else {
        fmt.Println("路由失败？")
    }    
    </pre></code>
                
                