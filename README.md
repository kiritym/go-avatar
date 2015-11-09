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
If you want to embed images, this is how you do it:
