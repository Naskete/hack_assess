# HackAssess

## Struct
### direct
> ID          int    `json:"id"` 
> 
> UserId      int    `json:"user_id"`
> 
> Slogan      int    `json:"slogan"`    // 产品slogan
> 
> Promotion   int    `json:"promotion"` // 线上推广文案
> 
> Idea        int    `json:"idea"`  // 活动策划
> 
> Integrity   int    `json:"integrity"` // 活动策划完整程度
> 
> Planning    int    `json:"planning"`  // 后期运营规划
> 
> Research    int    `json:"research"`  // 调研文档分析
> 
> Product     int    `json:"product"`   // 产品名称
> 
> Interactive int    `json:"interactive"`   // 产品交互文案
> 
> Total       int    `json:"total""`
> 
> GroupID     string `json:"group_id"`
> 
> Position    string `json:"position"`
> 
> Tag         string  `json:"tag" gorm:"unique"`

### product

> ID           int    `json:"id"`
> 
> UserId       int    `json:"user_id"`
> 
> Innovation   int    `json:"innovation"`   //产品创新性
> 
> Rationality  int    `json:"rationality"`  // 场景合理性
> 
> Conformity   int    `json:"conformity"`   //主题符合度
> 
> Research     int    `json:"research"` // 调研文档
> 
> Requirements int    `json:"requirements"` // 需求文档
> 
> Prototype    int    `json:"prototype"`    // 产品原型图
> 
> Competing    int    `json:"competing"`    // 竞品分析
> 
> Req          int    `json:"req"`  // 需求管理
> 
> Project      int    `json:"project"`  // 项目管理
> 
> Plan         int    `json:"plan"` // 迭代计划
> 
> Total       int    `json:"total""`
>
> GroupID     string `json:"group_id"`
> 
> Position     string `json:"position"`
> 
> Tag         string  `json:"tag" gorm:"unique"`


### Design

>ID           int    `json:"id"`
> 
> UserId       int    `json:"user_id"`
> 
> Design       int    `json:"design"`   // 设计规范
>  
> Logo         int    `json:"logo"` // Logo
> 
> Beauty       int    `json:"beauty"`   // 美观性
>
> Uniformity   int    `json:"uniformity"`   // 一致性
> 
> Friendliness int    `json:"friendliness"` // 友好性
> 
> Reliability  int    `json:"reliability"`  // 可实现性
> 
> Font         int    `json:"font"` // 字体选择
> 
> Thinking     int    `json:"thinking"` // 交互思路
> 
> Bonus        int    `json:"bonus"`    // 加分项
> 
> Total       int    `json:"total""`
>
> GroupID     string `json:"group_id"`
>
> Position     string `json:"position"`
> 
> Tag         string  `json:"tag" gorm:"unique"`

### Front

> ID            int    `json:"id"`
> 
> UserId        int    `json:"user_id"`
> 
> Function      int    `json:"function"`    //功能完整性
> 
> Fit           int    `json:"fit"`     // 与设计稿贴合程度
> 
> Layout        int    `json:"layout"`  // 布局
> 
> Encapsulation int    `json:"encapsulation"`   // 业务逻辑
> 
> Resources     int    `json:"resources"`   // 资源压缩
> 
> Docking       int    `json:"docking"`     // 与后端对接完整度
> 
> Total         int    `json:"total""`
>
> GroupID       string `json:"group_id"`
> 
> Position      string `json:"position"`
> 
> Tag         string  `json:"tag" gorm:"unique"`

### Back

> ID            int    `json:"id"`
> 
> UserId        int    `json:"user_id"`
> 
> Integrity     int    `json:"integrity"`   // 完整性
> 
> Property      int    `json:"property"`    // 性能
> 
> Scalability   int    `json:"scalability"` // 扩展性
> 
> Cert          int    `json:"cert"`    // HTTPS
> 
> Configuration int    `json:"configuration"`   // 敏感配置
> 
> Information   int    `json:"information"` // 敏感信息
> 
> Code          int    `json:"code"`    // 代码规范
> 
> Database      int    `json:"database"`    // 数据库规范
> 
> Interface     int    `json:"interface"`   // 接口文档
> 
> Document      int    `json:"document"`    // 数据库文档
> 
> Coverage      int    `json:"coverage"`    // 单元测试代码覆盖率
> 
> Solution      int    `json:"solution"`    // 现场解决方案
> 
> Total         int    `json:"total""`
>
> GroupID       string `json:"group_id"`
>
> Position      string `json:"position"`
> 
> Tag         string  `json:"tag" gorm:"unique"`

### Show

>ID          int    `json:"id"`
> 
> UserId      int    `json:"user_id"`
> 
> Performance int    `json:"performance"`   // 路演表现
> 
> Project     int    `json:"project"`   // 契合项目
> 
> Framework   int    `json:"framework"` // 逻辑框架
> 
> Emphasize   int    `json:"emphasize"` // 突出重点
> 
> Idea        int    `json:"idea"`  //  设计创意
> 
> Style       int    `json:"style"` // 统一风格
> 
> Colour      int    `json:"colour"`    // 颜色搭配
> 
> Animation   int    `json:"animation"` // 动画效果
> 
> Writing     int    `json:"writing"`   // 文案
> 
> Total       int    `json:"total""`
>
> GroupID     string `json:"group_id"`
>
> Position    string `json:"position"`
>
> Tag         string  `json:"tag" gorm:"unique"`

## Login 

Method: POST
URL: 121.199.32.101:8088/login

Data:
```json
{
    "username": <string>,
    "password": <string>
}
```
return:

```json
{
  "user_id": 100,
  "Student_id":"610911XXXX",
  "Token": "xxxx",
}
```

## Score

Method: POST
URL: 121.199.32.101:8088/:group/:id

group:组名 （front, back...
id: 组号 （1-n

以front为例

Date

```json
{
    "function":5,
    "fit":2,
    "layout":12,
    "encapsulation":10,
    "docking":19,
    "resources":18,
    "user_id":1234,
    "group_id":"2",
}
```

Return:

success:

```json
{
    "data": null,
    "message": "successful",
    "status": 1
}
```

failed
```json
{
    "data": null,
    "message": "failed to mark",
    "status": 0
}
```