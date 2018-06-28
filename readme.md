## 原则
 1. 依赖倒置原则（DIP)
    1. 高层模块不应依赖于低层模块，两者应该依赖于抽象。
    1. 抽象不不应该依赖于实现，实现应该依赖于抽象。
```
            Interface    +-------------- Interface
	  +	     +                       +
          | 	     |                       |
	Impl1       Impl2                  Impl
```

## 模式
 1. 控制反转（IoC）
    1. 依赖对象不在被依赖模块的类中直接通过new来获取, 将依赖对象的获得交给第三方（系统）来控制
## 实现
 1. 依赖注入（DI)
    1. [依赖对象]的<创建> <绑定> 转移到[被依赖对象类]的外部来实现
       1. db = new(mysql); record.setDb(db); record.save()
       1. 在换db 时， 只需要在顶层代码换掉创建 db
       1. mydql 继承 db， record 依赖db的接口
       1. 所谓注入：形象的讲是 db 对象不再 record 对象内部构造，而是外部设置进去(不一直是setDb，也可以是构造函数的参数, 或调用save 时的参数) [构造注入， 属性注入，接口注入]
 1. IoC容器
    1. 依赖对象构建， 放入容器， 通过映射依赖关系管理

## mocks

### 原理
 1. 通过依赖注入的方式实现的代码，可以通过mock的方式来单元测试
 1. 控制依赖对象的行为， 尽量增加单元测试的覆盖

### go pkg

 1. github.com/stretchr/testify/mock
    1. 用我们需要的mock 的功能了
    1. 需要为我们需要mock的接口写代码
 1. mockery
    1. 自动生成上述代码
    1. 安装 https://github.com/vektra/mockery

### 单元测试
 1. 选择注入点
 1. 用mockery 生成 mocks
 1. 编写测试
 1. 看覆盖率

### go test
 1. mockery -name=InterfaceName
 1. go test $PKG -covermode=count -coverprofile /tmp/$OUTPUT.cov
 1. go tool cover -html=/tmp/$OUTPUT.cov  -o /tmp/$OUTPUT.html
