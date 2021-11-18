# bootstrap-example

bootstrap 系列使用示例

bootstrap-global 提供了一些统一的门面，在使用时只需导入具体的实现适配包即可。

## 使用说明

当前 bootstrap 系列提供了配置、日志的支持，具体使用如下

### 配置

当前提供了如下几类配置支持

1. yaml/yml, 通过 bootstrap-config-yaml 库支持
2. json, 通过 bootstrap-config-json 库支持
3. apollo, 通过 bootstrap-config-apollo 库支持

在项目中使用如下

1. 引入依赖, `go get -u github.com/jiurongzhao/bootstrap-config-yaml`
2. 在代码中导入适配器与门面, `_ "github.com/jiurongzhao/bootstrap-config-yaml/yaml"` 与 `"github.com/jiurongzhao/bootstrap-global/config"`
3. 在代码中初始化适配器, `config.InitGlobalInstance("yaml", "resource/app.yaml")`
4. 通过门面调用, `config.Get("foo.bb")` 与 `config.Resolve("foo.struct", &aStructInstance)`


### 日志

当前提供 zap 和 logrus 的日志适配

项目中使用步骤如下


1. 引入依赖, `go get -u github.com/jiurongzhao/bootstrap-log-logrus`
2. 在代码中导入适配器与门面, `_ "github.com/jiurongzhao/bootstrap-log-logrus/logrus"` 与 `"github.com/jiurongzhao/bootstrap-global/log"`
3. 在代码中初始化适配器, `log.InitGlobalLogger("yaml")`
4. 通过门面调用, `log.Info("statement", "any")`, `log.Warn("statement", "any")`, `log.Error("statement", "any")`


## 注意事项

1. 不支持同时使用多个适配器，最终通过门面调用的是最好初始化的适配器
2. 如依赖其他适配器，应该在代码中有明确的执行顺序