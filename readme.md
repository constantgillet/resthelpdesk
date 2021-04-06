# RestHelpdesk 📘🆘✌

RestHelpdesk is a small rest API made with Golang for having your own helpdesk API and integrating with your app or web app

## Features & planned features

- [x] Articles & knowledges base with livesearch
- [ ] Tickets support with smtp integration
- [ ] LiveChat
- [ ] React Client
- [ ] Admin dashboard

## Requirements

- Golang (this project has been made on go version go1.16.3)
- MariaDB or MySQL database

## Installation & Launch 

Enter this command into the project folder to install go modules
```bash
go install
```

Import databse & set the environnements variables

```bash
PORT=8000 # Your app port, default is 8000
DATABASE_URL=mysql://root:@127.0.0.1:3306/resthelpdesk # Your mysql connection url , mysql://username:database_password@host:database_port/database_name
```

Enter this command into the project folder

```bash
 go run main.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
This project is licensed under the [MIT](https://choosealicense.com/licenses/mit/) License - see the LICENSE.md file for details