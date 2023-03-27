# GO ReceiverFunction
ReceiverFunction is function that assosiated with only the specific object. you can NOT directly call the function.

example here [link](#head1234)


## Practice to create a receiver function
1. create a object called bill which has feild "name", "items", "tip".
2. create a function called "newBill" which has argument of name of the bill (string) and inside "newBill" add `{"pie":5.99,"cake":3.99}` in items feild then return bill
3. printout newBill with argument as "your bill". (you can see result in [#2](https://github.com/hikaori/Go/pull/2) and [added items](https://github.com/hikaori/Go/commit/e3b84a61bba0b909ac029cecc36edf5505c99816))
4. create a function called "format" that assosiated with only bill object. another way of saying, you can NOT call just `format()`. you need to call `mybill.format()`. 
**"format" function is a receiver function**

 <a name="head1234">[example of calling format function(receiver function)]</a>
 
 <img width="286" alt="Screen Shot 2023-03-26 at 10 21 11 PM" src="https://user-images.githubusercontent.com/23109342/227847790-7d04f0bf-e8ca-4876-8750-523fba2660b1.png"> <img width="281" alt="Screen Shot 2023-03-26 at 10 21 26 PM" src="https://user-images.githubusercontent.com/23109342/227847753-73993b3c-ab1c-4050-a955-41cf5fd70cad.png">
 
5. make "format" function displays all bill items and format like below. (you can see result in [#3](https://github.com/hikaori/Go/pull/3)

<img width="365" alt="Screen Shot 2023-03-26 at 11 14 36 PM" src="https://user-images.githubusercontent.com/23109342/227856689-d7c8e6e2-450c-4b34-81d6-5fa9db911323.png">

