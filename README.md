# DUMBWAYS_B42 - Integration Task

## ENV for Front-End

```env
REACT_APP_API_BASE_URL=<your backend url>
REACT_APP_MIDTRANS_CLIENT_KEY=<your midtrans client key>
```

## ENV for Back-End

```env
# database url
DATABASE_URL=<your database url>
# example -> db_user:db_password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local

# secret key for jwt
SECRET_KEY=<your secret key>

# origin set up
ORIGIN_ALLOWED=<domain that allow to hit/use your endpoint>

# port for run backend
PORT=<port that you want to run your api server>

# server key midtrans
SERVER_KEY=<your midtrans server key>

# setup email for gomail
CONFIG_SMTP_HOST=<host of mail server, you can get it from email provider like gmail and another>
CONFIG_SMTP_PORT=<port of mail server, you can get it from email provider like gmail and another>
CONFIG_SENDER_NAME=<your name + your email>
CONFIG_AUTH_EMAIL=<email that used to send mail>
CONFIG_AUTH_PASSWORD=<email password>

# redirect url on verification link
REDIRECT_URL_ON_VERIFICATION=<your front-end/client url>
```
