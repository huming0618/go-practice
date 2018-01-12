
### 字符选取遍历
- 选择全部 `s[:]`
- 选择到某个字节为止 `s[:5]`
- 从某个字符开始 `s[1:]`

### 字符长度
Golang中`len()`方法会直接返回字符的字节长度，而不是可见字符长度
```
len("世界") 
```
会返回6 

### 字符集函数库 `unicode/utf8`
重要函数 
- `utf8.RuneCountInString()` 返回可见字符长度
- `utf8.DecodeRuneInString(buffer)` 把字节数组转换为字符

### 遍历字符
```
func loop() {
	s := "世界"
	
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("---------------")
}
```

### Golang用字节的视角来处理字符串，这也许因为这样可以更方便地来处理网络通信。