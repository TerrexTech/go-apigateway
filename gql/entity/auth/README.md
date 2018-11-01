## Usage examples
---

### Mutations

* #### UserInsert
  ```graphql
  mutation{
    authRegister(
      userName:"testUserName",
      password:"testPassword",
      firstName:"testFirstName",
      lastName:"testLastName",
      email:"testEmail",
      role:"employee"
    ){
      accessToken,
      refreshToken
    }
  }
  ```

* #### UserQuery
  ```graphql
  {
    authLogin(
      userName:"testUserName",
      password:"testPassword"
    ){
      accessToken,
      refreshToken
    }
  }
  ```
