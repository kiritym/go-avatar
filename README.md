# go-avatar
This will create a default and unique avatar based on the user name. You can specify the pixel size for different size.
## Usage
- ```git clone https://github.com/kiritym/go-avatar.git```
- ```cd go-avatar```
- ```go build```
- ```./go-avatar```
- Now you can hit the below URL with your user-name to get your avatar:
```http://localhost:8080/user-name```
- For various size, you can pass a size parameter as:
```http://localhost:8080/user-name?size=15```
- Below are few examples

```size = 10 ```

![alt tag](https://raw.githubusercontent.com/kiritym/go-avatar/master/images/kiritym-10.png)

```size = 20 or Default size ```

![alt tag](https://raw.githubusercontent.com/kiritym/go-avatar/master/images/kiritym.png)

```size = 25 ```

![alt tag](https://raw.githubusercontent.com/kiritym/go-avatar/master/images/kiritym-25.png)

```size = 30 ```

![alt tag](https://raw.githubusercontent.com/kiritym/go-avatar/master/images/kiritym-30.png)
