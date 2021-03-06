
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有2016 Go Ethereum作者
//此文件是Go以太坊库的一部分。
//
//Go-Ethereum库是免费软件：您可以重新分发它和/或修改
//根据GNU发布的较低通用公共许可证的条款
//自由软件基金会，或者许可证的第3版，或者
//（由您选择）任何更高版本。
//
//Go以太坊图书馆的发行目的是希望它会有用，
//但没有任何保证；甚至没有
//适销性或特定用途的适用性。见
//GNU较低的通用公共许可证，了解更多详细信息。
//
//你应该收到一份GNU较低级别的公共许可证副本
//以及Go以太坊图书馆。如果没有，请参见<http://www.gnu.org/licenses/>。

//包geth包含用于以太坊的简化移动API。
//
//此包的作用域是*不*以允许写入自定义以太坊客户端
//从Go Ethereum中提取片段，而不是允许在上面写本地dapps
//移动平台。在使用或扩展此包时请记住这一点！
//
//API限制
//
//由于gomobile无法在go和android/ios之间桥接任意类型，因此
//公开的API需要手动包装成简化的类型，并自定义
//构造器和getter/setter，以确保它们可以被有效地使用。
//也从Java/Objc。
//
//考虑到这一点，请尝试限制此包的范围，并仅添加
//移动支持不起作用的基本要素，尤其是手动
//否则，同步代码将非常困难。从长远来看，我们可能会考虑
//正在编写自定义库生成器，但这些生成器现在已超出范围。
//
//内容方面，此包中的每个文件都对应于整个go包
//从Go以太坊存储库。请遵守此范围以防止
//无法维护的包。
//
//包装指南：
//
//要暴露的每种类型都应该包装成它自己的平面结构，
//它内部包含一个字段：原始的Go以太坊版本。
//这是必需的，因为GoMobile目前无法公开命名类型。
//
//每当方法参数或返回类型是自定义结构时，指针
//变量应始终用作语言之间的值类型
//边界可能有奇怪的行为。
//
//类型切片应转换为单个乘法类型包装
//使用“大小”、“获取”和“设置”方法执行切片。进一步切片操作
//不应提供以限制远程代码的复杂性。数组应该是
//尽量避免，因为它们会使边界检查复杂化。
//
//如果一个方法有多个返回值（例如某个返回值+一个错误），那么
//在objc中作为输出参数生成。为了避免产生奇怪的名字，比如
//对于它们，如果元组，请始终为输出变量分配名称。
//
//注意，恐慌*不能*跨越语言边界，反而会导致
//进程中无法检测的分段故障。对于错误处理，只使用错误
//返回，这可能是唯一的或第二个返回。
package geth
