# Code Review 规范

Code Review
Gofmt
Comment Sentences
Contexts
Copying
Crypto Rand
Declaring Empty Slices
Doc Comments
Don’t Panic
Error Strings
Examples
Goroutine Lifetimes
Handle Errors
Imports
Import Blank
Import Dot
In-Band Errors
Indent Error Flow
Initialisms
Interfaces
Line Length
Mixed Caps
Named Result Parameters
Naked Returns
Package Comments
Package Names
Pass Values
Receiver Names
Receiver Type
Synchronous Functions
Useful Test Failures
Variable Names
小结
参考

## Code Review

Code Review 即代码审查，对于项目而言，其目的在于找到开发时被忽视的 Bug，以此极大地提高代码质量也可以帮助开发者们更加熟悉项目；对于 Review 者来说，Review 别人的代码，可以学习他人代码的设计，会对自己要求更高些，促使自己写出更好的代码；对于被 Review 者来说，可以尽早发现问题，及时改正；对于团队而言，Review 可以促进团队人员的沟通，让代码风格更加一致性，有利于代码的可读性和可维护性。

Go 官方也提供了 Code Review Comments，具体如下文所述。

## Gofmt

可以在你的代码中使用 Gofmt 以解决大多数代码格式问题。几乎所有的代码都使用 gofmt。

另一种方法是使用 goimports，它是 gofmt的超集，可根据需要添加（删除）行。

## Comment Sentences

参考 https://golang.org/doc/effective_go.html#commentary。注释文档声明应该是完整的句子，即使这看起来有些多余。这种方式使注释在提取到 godoc 文档时格式良好。注释应以所描述事物的名称开头，并以句点（也可以是 ！，？）结束：

// Request represents a request to run a command.
type Request struct { ...

// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) { ...
### Contexts

context.Context 类型的值包含跨 API 和进程边界的安全凭证，跟踪信息，截止时间和取消信号。比如传入 RPC 请求和 HTTP 请求一直到传出相关请求，Go 程序在整个过程的函数调用链中显式地传递 Context。

大多数使用 Context 的函数都应该接受 Context 作为函数的第一个参数：

func F(ctx context.Context, /* other arguments */) {}
从不特定于请求（request-specific）的函数可以使用 context.Background() 获取 Context，并将 err 与 Context 同时传递，即使你认为不需要。默认情况下只传递 Context ；只在你有充分的理由认为这是错误的，才能直接使用 context.Background()。

不要将 Context 成员添加到某个 struct 类型中；而是将 ctx 参数添加到该类型的方法上。一个例外情况是当前方法签名必须与标准库或第三方库中的接口方法匹配。

不要在函数签名中创建自定义 Context 类型或使用除了 Context 以外的接口。

如果要传递应用程序数据，请将其放在参数，方法接收器，全局变量中，或者如果它确实应该属于 Context，则放在 Context 的 Value 属性中。

所有的 Context 都是不可变的，因此可以将相同的 ctx 传递给多个共享相同截止日期，取消信号，安全凭据，跟踪等的调用。

### Copying

为避免意外的别名，从另一个包复制 struct 时要小心。例如，bytes.Buffer 类型包含一个 []byte 的 slice，并且作为短字符串的优化，slice 可以引用一个短字节数组。如果复制一个 Buffer，副本中的 slice 可能会对原始数组进行别名操作，从而导致后续方法调用产生令人惊讶的效果。

通常，如果 T 类型的方法与其指针类型 *T 相关联，请不要复制 T 类型的值。

### Crypto Rand

不要使用包 math/rand 来生成密钥，即使是一次性密钥。在没有种子（seed）的情况下，生成器是完全可以被预测的。使用 time.Nanoseconds() 作为种子值，熵只有几位。请使用 crypto/rand 的 Reader 作为替代，如果你倾向于使用文本，请输出成十六进制或 base64 编码：

import (
	"crypto/rand"
	// "encoding/base64"
	// "encoding/hex"
	"fmt"
)

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)  // out of randomness, should never happen
	}
	return fmt.Sprintf("%x", buf)
	// or hex.EncodeToString(buf)
	// or base64.StdEncoding.EncodeToString(buf)
}

### Declaring Empty Slices

当声明一个空 slice 时，倾向于用

var t []string
代替

t := []string{}
前者声明了一个 nil slice 值，而后者声明了一个非 nil 但是零长度的 slice。两者在功能上等同，len 和 cap 均为零，而 nil slice 是首选的风格。

请注意，在部分场景下，首选非零但零长度的切片，例如编码 JSON 对象时（前者编码为 null，而后者则可以正确编码为 JSON array[]）。

在设计 interface 时，避免区分 nil slice 和 非 nil，零长度的 slice，因为这会导致细微的编程错误。

有关 Go 中对于 nil 的更多讨论，请参阅 Francesc Campoy 的演讲 Understanding Nil。

### Doc Comments

所有的顶级导出的名称都应该有 doc 注释，重要的未导出类型或函数声明也应如此。有关注释约束的更多信息，请参阅 https://golang.org/doc/effective_go.html#commentary。

Don’t Panic

请参阅 https://golang.org/doc/effective_go.html#errors。不要将 panic 用于正常的错误处理。使用 error 和多返回值。

### Error Strings

错误信息字符串不应大写（除非以专有名词或首字母缩略词开头）或以标点符号结尾，因为它们通常是在其他上下文后打印的。即使用 fmt.Errorf("something bad") 而不要使用 fmt.Errorf("Something bad")，因此 log.Printf("Reading %s: %v", filename, err) 的格式中将不会出现额外的大写字母。否则这将不适用于日志记录，因为它是隐式的面向行，而不是在其他消息中组合。

### Examples

添加新包时，请包含预期用法的示例：可运行的示例，或是演示完整调用链的简单测试。

阅读有关 testable Example() functions 的更多信息。

### Goroutine Lifetimes

当你生成 goroutines 时，要清楚它们何时或是否会退出。

通过阻塞 channel 的发送或接收可能会引起 goroutines 的内存泄漏：即使被阻塞的 channel 无法访问，垃圾收集器也不会终止 goroutine。

即使 goroutines 没有泄漏，当它们不再需要时却仍然将其留在内存中会导致其他细微且难以诊断的问题。往已经关闭的 channel 发送数据将会引发 panic。在“结果不被需要之后”修改仍在使用的输入仍然可能导致数据竞争。并且将 goroutines 留在内存中任意长时间将会导致不可预测的内存使用。

请尽量让并发代码足够简单，从而更容易地确认 goroutine 的生命周期。如果这不可行，请记录 goroutines 退出的时间和原因。

### Handle Errors

请参阅 https://golang.org/doc/effective_go.html#errors。不要使用 _ 变量丢弃 error。如果函数返回 error，请检查它以确保函数成功。处理 error，返回 error，或者在真正特殊的情况下使用 panic。

### Imports

避免包重命名导入，防止名称冲突；好的包名称不需要重命名。如果发生命名冲突，则更倾向于重命名最接近本地的包或特定于项目的包。

包导入按组进行组织，组与组之间有空行。标准库包始终位于第一组中。

package main

import (
	"fmt"
	"hash/adler32"
	"os"

	"appengine/foo"
	"appengine/user"

	"github.com/foo/bar"
	"rsc.io/goversion/version"
)
goimports 会为你做这件事。

### Import Blank

仅出于副作用而导入的软件包（使用语法 import _“ pkg”）应仅在程序的主 main package 或需要它们的测试中导入。

### Import Dot

部分包由于循环依赖，不能作为测试包的一部分进行测试时，以 . 形式导入它们可能很有用：

package foo_test

import (
    "bar/testutil" // also imports "foo"
    . "foo"
)
在这种情况下，测试文件不能位于 foo 包中，因为它使用的 bar/testutil 依赖于 foo 包。所以我们使用 import . 形式使得测试文件伪装成 foo 包的一部分，即使它不是。除了这种情况，不要在程序中使用 import .。它将使程序更难阅读——因为不清楚如 Quux 这样的名称是否是当前包中或导入包中的顶级标识符。

### In-Band Errors

在 C 和类 C 语言中，通常使函数返回 -1 或 null 之类的值用来发出错误信号或缺少结果：

// Lookup returns the value for key or "" if there is no mapping for key.
func Lookup(key string) string

// Failing to check a for an in-band error value can lead to bugs:
Parse(Lookup(key))  // returns "parse failure for value" instead of "no value for key"
Go 对多返回值的支持提供了一种更好的解决方案。函数应返回一个附加值以指示其他返回值是否有效，而不是要求客户端检查 in-band 错误值。此附加值可能是一个 error，或者在不需要解释时可以是布尔值。它应该是最终的返回值。

// Lookup returns the value for key or ok=false if there is no mapping for key.
func Lookup(key string) (value string, ok bool)
这可以防止调用者错误地使用返回结果：

Parse(Lookup(key))  // compile-time error
并鼓励更健壮和可读性强的代码：

value, ok := Lookup(key)
if !ok  {
    return fmt.Errorf("no value for %q", key)
}
return Parse(value)
此规则适用于公共导出函数，但对于未导出函数也很有用。

返回值如 nil，“”，0 和 -1 在他们是函数的有效返回结果时是可接收的，即调用者不需要将它们与其他值做分别处理。

某些标准库函数（如 “strings” 包中的函数）会返回 in-band 错误值。这大大简化了字符串操作，代价是需要程序员做更多事。通常，Go 代码应返回表示错误的附加值。

### Indent Error Flow

尝试将正常的代码路径保持在最小的缩进处，优先处理错误并缩进。通过允许快速可视化扫描正常路径来提高代码的可读性。例如，不要写：

if err != nil {
    // error handling
} else {
    // normal code
}
相反，书写以下代码：

if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code
如果 if 语句具有初始化语句，例如：

if x, err := f(); err != nil {
    // error handling
    return
} else {
    // use x
}
那么这可能需要将短变量声明移动到新行：

x, err := f()
if err != nil {
    // error handling
    return
}
// use x
### Initialisms

名称中的单词是首字母或首字母缩略词（例如 “URL” 或 “NATO” ）需要具有相同的大小写规则。例如，“URL” 应显示为 “URL” 或 “url” （如 “urlPony” 或 “URLPony” ），而不是 “Url”。举个例子：ServeHTTP 不是 ServeHttp。对于具有多个初始化 “单词” 的标识符，也应当显示为 “xmlHTTPRequest” 或 “XMLHTTPRequest”。

当 “ID” 是 “identifier” 的缩写时，此规则也适用于 “ID” ，因此请写 “appID” 而不是“appId”。

由协议缓冲区编译器生成的代码不受此规则的约束。人工编写的代码比机器编写的代码要保持更高的标准。

### Interfaces

Go 接口通常属于使用 interface 类型值的包，而不是实现这些值的包。实现包应返回具体（通常是指针或结构）类型：这样一来可以将新方法添加到实现中，而无需进行大量重构。

不要在 API 的实现者端定义 “for mocking” 接口；相反，设计 API 以便可以使用真实实现的公共 API 进行测试。

在使用接口之前不要定义接口：如果没有真实的使用示例，很难看出接口是否是必要的，更不用说它应该包含哪些方法。

package consumer  // consumer.go

type Thinger interface { Thing() bool }

func Foo(t Thinger) string { … }
package consumer // consumer_test.go

type fakeThinger struct{ … }
func (t fakeThinger) Thing() bool { … }
…
if Foo(fakeThinger{…}) == "x" { … }
// DO NOT DO IT!!!
package producer

type Thinger interface { Thing() bool }

type defaultThinger struct{ … }
func (t defaultThinger) Thing() bool { … }

func NewThinger() Thinger { return defaultThinger{ … } }
相反，返回一个具体的类型，让消费者模拟生产者实现。

### package producer

type Thinger struct{ … }
func (t Thinger) Thing() bool { … }

func NewThinger() Thinger { return Thinger{ … } }
Line Length

Go 代码中没有严格的行长度限制，但避免使用造成阅读障碍的长行。类似的，如果长行的可读性更好，不要为了缩短行而添加换行符——例如，行组成是重复的。

大多数情况下，当人们 “不自然地” 自动换行（wrap lines）时（在函数调用或函数声明的中间，或多或少，比如，虽然有一些例外），如果它们有合理数量的参数并且变量名称较短时，自动换行将是不必要的。长行似乎与长名称有关，避免名称过长有很大帮助。

换句话说，换行是因为你所写的语义（作为一般规则）而不是因为行的长度。如果您发现这会产生太长的行，那么更改名称或语义，可能也会得到一个好结果。

实际上，这与关于函数应该有多长的建议完全相同。没有 “永远不会有超过 N 行的函数” 这样的规则，但是程序中肯定会存在行数太多，功能过于微弱的函数，而解决方案是改变这个函数边界的位置，而不是执着在行数上。

### Mixed Caps

请参阅 https://golang.org/doc/effective_go.html#mixed-caps。即使 Go 中混合大小写的规则打破了其他语言的惯例，也是适用的。例如，未导出的常量写成 maxLength 而不是 MaxLength 或 MAX_LENGTH。

另见当前页面的 Initialisms 一节。

### Named Result Parameters

考虑一下 godoc 中会是什么样子。命名结果参数如：

func (n *Node) Parent1() (node *Node)
func (n *Node) Parent2() (node *Node, err error)
将会造成口吃现象（stutter）; 最好这样使用：

func (n *Node) Parent1() *Node
func (n *Node) Parent2() (*Node, error)
另一方面，如果函数返回两个或三个相同类型的参数，或者如果从上下文中不清楚返回结果的含义，那么在某些上下文中添加命名可能很有用。但是不要仅仅为了避免在函数内做结果参数的声明（var 或者 :=）而命名结果参数；这以牺牲不必要的 API 冗长性为代价，换取了一个微小的实现简洁性。

func (f *Foo) Location() (float64, float64, error)
不如以下代码清晰：

// Location returns f's latitude and longitude.
// Negative values mean south and west, respectively.
func (f *Foo) Location() (lat, long float64, err error)
如果函数行数较少，那么非命名结果参数是可以的。一旦它是一个中等规模的函数，请明确返回值。推论：仅仅因为它使得能够直接使用预命名返回而命名结果参数是不值得的。文档的清晰度总比在函数中的一行两行更重要。

最后，在某些情况下，您需要命名结果参数，以便在延迟闭包中更改它，这也是可以的。

### Naked Returns

请参阅当前页面 Named Result Parameters 一节。

### Package Comments

与 godoc 呈现的所有注释一样，包注释必须出现在 package 声明的临近位置，无空行。

// Package math provides basic constants and mathematical functions.
package math
/*
Package template implements data-driven templates for generating textual
output such as HTML.
....
*/
package template
对于 “package main” 注释，在二进制文件名称之后可以使用其他样式的注释(如果它们放在前面，则可以大写)，例如，对于你可以编写 seedgen 目录中下的 main 包注释：

// Binary seedgen ...
package main
或是

// Command seedgen ...
package main
或是

// Program seedgen ...
package main
或是

// The seedgen command ...
package main
或是

// The seedgen program ...
package main
或是

// Seedgen ..
package main
以上是相应示例，它们的合理变体也是可以接受的。

请注意，以小写单词开头的句子不属于包注释的可接受选项，因为注释是公开可见的，应该用适当的英语书写，包括将句子的第一个单词的首字母大写。当二进制文件名称是第一个单词时，即使它与命令行调用的拼写不严格匹配，也需要对其进行大写。

有关评论约定的更多信息，请参阅 https://golang.org/doc/effective_go.html#commentary。

### Package Names

包中名称的所有引用都将使用包名完成，因此您可以从标识符中省略该名称。例如，如果有一个 chubby 包，你不应该定义类型名称为 ChubbyFile ，否则使用者将写为 chubby.ChubbyFile。而是应该命名类型名称为 File，使用时将写为 chubby.File。避免使用无意义的包名称，如 util，common，misc，api，types 和 interfaces。有关更多信息，请参阅 http://golang.org/doc/effective_go.html#package-names 和 http://blog.golang.org/package-names。

### Pass Values

不要只是为了节省几个字节就将指针作为函数参数传递。如果一个函数在整个过程中只引用它的参数 x 作为 x，那么这个参数不应该是一个指针。此常见实例包括将指针传递给 string（string）或是指向接口值（*io.Reader）的指针。在这两种情况下，值本身都是固定大小，可以直接传递。这个建议不适用于大型 struct ，甚至不适用于可能生长的小型 struct。

### Receiver Names

方法接收者的名称应该反映其身份；通常，其类型的一个或两个字母缩写就足够了（例如 “client” 的 “c” 或 “cl” ）。不要使用通用名称，例如 “me”，“this” 或 “self”，这是面向对象语言的典型标识符，它们更强调方法而不是函数。名称不必像方法论证那样具有描述性，因为它的作用是显而易见的，不起任何记录目的。名称可以非常短，因为它几乎出现在每种类型的每个方法的每一行上；familiarity admits brevity。使用上也要保持一致：如果你在一个方法中叫将接收器命名为“c”，那么在其他方法中不要把它命名为“cl”。

### Receiver Type

选择到底是在方法上使用值接收器还是使用指针接收器可能会很困难，尤其是对于 Go 新手程序员。如有疑问，请使用指针接收器，但有时候值接收器是有意义的，通常是出于效率的原因，例如小的不变结构或基本类型的值。以下是一些有用的指导：

- 如果接收器是 map，func 或 chan，则不要使用指向它们的指针。如果接收器是 slice 并且该方法不重新切片或不重新分配切片，则不要使用指向它的指针。
- 如果该方法需要改变接收器的值，则接收器必须是指针。
- 如果接收器是包含 sync.Mutex 或类似同步字段的 struct，则接收器必须是避免复制的指针。
- 如果接收器是大型结构或数组，则指针接收器更有效。多大才算大？假设它相当于将其包含的所有元素作为参数传递给方法。如果感觉太大，那么对接收器来说也太大了。
函数或方法可以改变接收器吗（并发调用或调用某方法时继续调用相关方法或函数）？在调用方法时，值类型会创建接收器的副本，因此外部更新将不会应用于此接收器。如果必须在原始接收器中看到更改效果，则接收器必须是指针。
- 如果接收器是 struct，数组或 slice，并且其任何元素是指向可能改变的对象的指针，则更倾向于使用指针接收器，因为它将使读者更清楚地意图。
- 如果接收器是一个小型数组或 struct，那么它自然是一个值类型（例如，类似于 time.Time 类型），对于没有可变字段，没有指针的类型，或者只是一个简单的基本类型，如 int 或 string，值接收器是合适的。值接收器可以减少可以生成的垃圾量；如果将值作为参数传递给值类型方法，则可以使用堆栈上的副本而不需要在堆上进行分配。（编译器试图避免这种分配，但它不能总是成功）因此，在没有进行分析之前，不要选择值接收器类型。
最后，如有疑问，请使用指针接收器。

### Synchronous Functions

相比异步函数更倾向于同步函数——直接返回结果的函数，或是在返回之前已完成所有回调或 channel 操作的函数。

同步函数让 goroutine 在调用中本地化，能够更容易地推断其生命周期并避免泄漏和数据竞争。同步函数也更容易测试：调用者可以传递输入并检查输出，而无需轮询或同步。

如果调用者需要更多的并发性，他们可以定义和调用单独的 goroutine 中的函数来轻松实现。但是在调用者端删除不必要的并发性是非常困难的——有时是不可能的。

### Useful Test Failures

失败的测试也应该提供有用的消息，说明错误，展示输入内容，实际内容以及预期结果。编写一堆 assertFoo 帮助程序可能很吸引人，但请确保您的帮助程序能产生有用的错误消息。假设调试失败测试的人不是你，也不是你的团队。典型的 Go 失败测试如：

if got != tt.want {
    t.Errorf("Foo(%q) = %d; want %d", tt.in, got, tt.want) // or Fatalf, if test can't test anything more past this point
}
请注意，此处的命令是 实际结果!=预期结果，并且错误消息也使用该命令格式。然而一些测试框架鼓励倒写输出格式，如 预期结果 != 实际结果，“预期结果为 0，实际结果为 x”，等等。但是 Go 没有这样做。

如果这看起来像是打了很多字，你可能想写一个表驱动的测试。

在使用具有不同输入的测试帮助程序时以消除失败测试歧义的另一种常见技术是使用不同的 TestFoo 函数包装每个调用者，而测试名称也根据对应的输入命名：

func TestSingleValue(t *testing.T) { testHelper(t, []int{80}) }
func TestNoValues(t *testing.T) { testHelper(t, []int{}) }
在任何情况下，你都有责任向可能会在将来调试你的代码的开发者提供有用的消息。

### Variable Names

Go 中的变量名称应该短而不是长。对于范围域中的局部变量尤其如此。例如为 line count 定义 c 变量，为 slice index 定义 i 变量。

基本规则：范围域中，越晚使用的变量，名称必须越具有描述性。对于方法接收器，一个或两个字母就足够了。诸如循环索引和读取器（Reader）之类的公共变量可以是单个字母（i，r）。更多不寻常的事物和全局变量则需要更具描述性的名称。

## 小结

Code Review 跟项目代码的可读性、可维护性息息相关，也是团队知识沉淀的一个重要实践，大家都应该重视起来。Go 语言社区也非常重视 Code Review。强烈建议使用 Go 语言编程的同学们，都应该至少完整地阅读一遍官方的代码规范指南，它既是我们在写代码时应该遵守的规则，也是在代码审查时需要注意的规范。

## 参考

https://github.com/golang/go/wiki/CodeReviewComments
https://golang.org/doc/effective_go.html
https://draveness.me/golang-101/
