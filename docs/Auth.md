# Authorization and authentication in Intuit Payements SDK

You need to setup authorization so your app can connect to our APIs and give customers a way to start "user consent" flow and grant it permission to access their data.

## Pre requisite

- Make a Developer account on Intuit

## How to generate Authorization Code

Keep in my that you can use this code only for once to generate the access token for the user and every `access token expires in 3600 seconds`

1. Go to [OAuth 2.0 Playeground](https://developer.intuit.com/app/developer/playground)

2. Select the workspace and app from the dropdown menu

    choose `sandbox` for the development mode and `production` for the production mode in the app dropdown menu

3. Select the scopes in the scopes option (you can choose one or more scopes) and then click on `Get authorization code` button

4. You will get Authorization Code and Realm ID

5. Click on Get Tokens button to get the tokens like `access token` and `refresh token` with the expiration time of these tokens.

6. You can use the access token to take the advantage of the benefits of the Intuit Api


## Sample Request to get the Access token


```sh
curl -X POST 'https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer' \
-H 'Accept: application/json' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-H 'Authorization: REPLACE_WITH_AUTHORIZATION_HEADER (details below)' \
-d 'grant_type=authorization_code' \
-d 'code=REPLACE_WITH_AUTHORIZATION_CODE' \
-d 'redirect_uri=REPLACE_WITH_REDIRECT_URI'
```

`Sample output`

```json
{
"token_type": "bearer",
"expires_in": 3600,
"refresh_token":"Q311488394272qbajGfLBwGmVsbF6VoNpUKaIO5oL49aXLVJUB",
"x_refresh_token_expires_in":15551893,
"access_token":"eJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGGlyIn0..KM1_Fezsm6BUSaqqfTedaA.
dBUCZWiVmjH8CdpXeh_pmaM3kJlJkLEqJlfmavwGQDThcf94fbj9nBZkjEPLvBcQznJnEmltCIvsTGX0ue_w45h7_
yn1zBoOb-1QIYVE0E5TI9z4tMUgQNeUkD1w-X8ECVraeOEecKaqSW32Oae0yfKhDFbwQZnptbPzIDaqiduiM_q
EFcbAzT-7-znVd09lE3BTpdMF9MYqWdI5wPqbP8okMI0l8aa-UVFDH9wtli80zhHb7GgI1eudqRQc0sS9zWWb
I-eRcIhjcIndNUowSFCrVcYG6_kIj3uRUmIV-KjJUeXdSV9kcTAWL9UGYoMnTPQemStBd2thevPUuvKrPdz3ED
ft-RVRLQYUJSJ1oA2Q213Uv4kFQJgNinYuG9co_qAE6A2YzVn6A8jCap6qGR6vWHFoLjM2TutVd6eOeYoL2bb7jl
QALEpYGj4E1h3y2xZITWvnmI0CEL_dYQX6B3QTO36TDaVl9WnTaCCgAcP6bt70rFlPYbCjOxLoI6qFm5pUwGLLp
67JZ36grc58k7NIyKJ8dLJUL_Q9r1WoUvw.ZS298t_u7dSlkfajxLfO9Q"
}
```



## References

- [Authorization and authentication](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization)
- [Practice authorization in the OAuth Playground](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0-playground)