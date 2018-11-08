## Usage examples
---

### Mutations

* #### UserInsert
  ```graphql
  mutation{
    authRegister(
      userName:"danpie3",
      password:"danpie",
      firstName:"obnoxious",
      lastName:"potato",
      email:"explosion",
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
      userName: "danpie",
      password:"danpie"
    ){
      accessToken,
      refreshToken
    }
  }
  ```
