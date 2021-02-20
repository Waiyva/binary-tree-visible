- - 
# <span id="head1"> BTPrinter</span>

Print binary tree in extremely small area for go language



[ BTPrinter](#head1)

- [ About](#head2)
- [ Installation](#head3)
- [ Example](#head4)
- [ Others](#head5)



## <span id="head2"> About</span>

To practise algorithm quicker, code Visual binary tree for ready



## <span id="head3"> Installation</span>

1. use tool **go get**

> go get github.com/waiyva/binary-tree-printer

2. download code to your $GOPATH



## <span id="head4"> Example</span>

```go
import "github.com/waiyva/binary-tree-printer"

func main() {
  btprinter.PrintTree("1,2,3")
}
```

```
  1  
 / \ 
2   3
```







```go
import "github.com/waiyva/binary-tree-printer"

func main() {
  btprinter.PrintTree("1,2,3,4,5,#,#,6,7,8,1,#,#,#,#,#,#,2,3,4,5,6,7,8,9,10,11,12,13,14,15")
}
```

```
        1              
       / \             
      2   3            
     / \               
    /   \              
   4     5             
  / \   / \            
 6   7 8   1           
          / \          
         /   \         
        /     \        
       /       \       
      /         \      
     2           3     
    / \         / \    
   /   \       /   \   
  4     5     6     7  
 / \   / \   / \   / \ 
8   9 10 11 12 13 14 15

```







## <span id="head5"> Others</span>

I 'm go beginner, so there may be some nonstandard code here. Welcome to make issue.

Ideas come from Java version https://github.com/afkbrb/binary-tree-printer

Thanks!