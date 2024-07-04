<div align="center">
  <a href="https://github.com/dimassagngsptr/react-client-blanja.git">
      <img src="https://github.com/dimassagngsptr/blanja-kelompok-1/blob/master/apps/web/src/assets/logo/Group%201158.png" width="300"/>
  </a>

  <h1 align="center">Blanja Innovation</h1>

  <p align="center">
     Blanja Implementation
    <br />
    <br />
    <a href="https://innovation-blanja.vercel.app/" target="_blank">View Website Demo</a>
    ·
     <a href="https://github.com/Azhar-54/golang-server-blanja" target="_blank">View Back-End Repo</a>
  </p>
</div>

## Table Of Contents 

- [Table of Contents](#table-of-contents))
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Contributors](#contributors)
- [Contributing](#contributing)
- [Contact](#contact)
- [Documentation](#documentation)
- [Acknowledgements](#acknowledgements)




## About The Project

Blanja is a sales application specializing in clothing items, including jackets, pants, t-shirts, and shirts. The application is built using React.js for the client side and Golang for developing the RESTful API. With the assistance of the Gofiber framework, Blanja offers functionalities such as (search, sorting, filtering) products, authentication, authorization, add-to-cart, customer address addition, email verification, and checkout features utilizing the payment gateway from Midtrans.


### Built With

[Golang](https://go.dev/)

[Gofiber](https://gofiber.io/)

[GoORM](https://gorm.io/)

[Postgre Sql](https://www.postgresql.org/)

[Cloudinary](https://cloudinary.com/)

[Midtrans](https://midtrans.com/)

[JSON Web Token](https://jwt.io/)

[Bluemonday](https://github.com/microcosm-cc/bluemonday)

[NetSMTP](https://pkg.go.dev/net/smtp)

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Ensure you have the following installed on your local machine:

- Go
- PostgreSQL
- Ngrok 

### Installation

1. Clone Repo

   ```sh
     git clone https://github.com/Azhar-54/golang-server-blanja.git
   ```

2. Go to workspace

   ```sh
      cd golang-server-blanja
   ```

    Open VsCode

    ```sh
      code .
    ```

3. Setup enviroment ```.env``` 

   ```sh

      URL="host=your_hostname_db user=your_username_db password=your_password_db dbname=your_db port=your_port sslmode=optional; TimeZone=optional"
      
      JWT_KEY=your_jwt_key
   
      BASE_URL=base_url_from_your_api_development || default localhost:3000
   
      CLOUDINARY_URL="cloudinary://your_cloudinary_url"
   
      API_SECRET=cloudinary_api_secret
   
      API_KEY=cloudinay_api_key
   
      CLOUD_NAME=cloudinary_name
   
      SMTP_USERNAME=your_email_smtp
   
      SMTP_PASSWORD=your_pass_smtp
   
      SERVER_KEY=midtrans_server_key
   
      CLIENT_KEY=midtrans_client_key
   
      WEB_URL=http://yourwebsite
    ```

3. Running the project

   install depedencies


    ```sh
      go run main.go
    ```

    ```sh
      air
    ```

4. Turn on your Ngrok and forward your port

   ```sh
    ngrok http 3000

    The default port is 3000
   ```
  
The server will start on port 3000 by default. You can use Postman to interact with the endpoints in [Documentation](#documentation).

## Usage

To use this project, follow the instructions below for understanding the project structure and how to use the provided API documentation.

### Project Structure

```
golang-server-blanja/
 .air.toml
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   main.go
│
├───src
│   ├───configs
│   │       db.go
│   │
│   ├───controllers
│   │       addressController.go
│   │       authController.go
│   │       cartController.go
│   │       categoryController.go
│   │       productColorController.go
│   │       productController.go
│   │       productSizeController.go
│   │       storeController.go
│   │       transactionController.go
│   │       userController.go
│   │
│   ├───helpers
│   │       countData.go
│   │       generateJwt.go
│   │       generateRandomNumber.go
│   │       migration.go
│   │       transformFilePath.go
│   │       uploadFile.go
│   │       validationPassword.go
│   │       validator.go
│   │       xssClean.go
│   │
│   ├───middlewares
│   │       jwt.go
│   │
│   ├───models
│   │       Address.go
│   │       Cart.go
│   │       Category.go
│   │       Product.go
│   │       ProductColor.go
│   │       ProductImage.go
│   │       ProductSize.go
│   │       Store.go
│   │       Transaction.go
│   │       User.go
│   │
│   ├───routes
│   │       main.go
│   │
│   ├───services
│   │       cloudinary.go
│   │       midtrans.go
│   │       smtp.go
│   │
│   └───templates
│           email_template.html
│
└───tmp
        main.exe
```

### ERD

To view the Entity Relationship Diagram you can visit the following link [ERD Blanja Innovation](https://drawsql.app/teams/dimas-team-8/diagrams/blanja-kelompok-1)

### Documentation

Access the API documentation for the **Blanja** project. Use this documentation to test endpoints and understand the structure and functionality of the available APIs in this project.

[![Blanja Innovation API Postman Documentation](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/29288801/2sA3duECJC)


### Project Related

- [`Front-End Project Repository Link`](https://github.com/Azhar-54/blanja-kelompok-1)

- [`Front-End Demonstration Link`](https://innovation-blanja.vercel.app/)

- [`Back-End Demonstration Link`](https://golang-server-blanja-production.up.railway.app/)


## Contributors

Here are the contributors to the project and their respective roles:

<table>
<tr>
<td> <img width="360" height="360" src="https://github.com/dimassagngsptr/blanja-kelompok-1/blob/58b4f42682a91658138b4fb75bf68669028b1348/apps/web/src/assets/screenshot/dimas.jpg" /> </td>
<td>

**Dimas Ageng Saputro**: Project Manager, Fullstack Developer </br>
Github   :  [github.com/dimassagngsptr](https://github.com/dimassagngsptr)</br>
Linkedin : [linkedin.com/in/dimasagengsaputro](https://www.linkedin.com/in/dimasagengsaputro/)

</td>
</tr>
<tr>
<td> <img width="360" height="360" src="https://github.com/dimassagngsptr/blanja-kelompok-1/blob/58b4f42682a91658138b4fb75bf68669028b1348/apps/web/src/assets/screenshot/ghaffar.JPG" /> </td>
<td>

**Abdul Ghaffar Sidik**: Fullstack Developer </br>
Github   :  [github.com/ghofarsidik](https://github.com/ghofarsidik)</br>
Linkedin : [linkedin.com/in/abdul-ghaffar-sidik](https://www.linkedin.com/in/abdul-ghaffar-sidik/)

</td>
</tr>
<tr>
<td> <img width="360" height="360" src="https://github.com/dimassagngsptr/blanja-kelompok-1/blob/58b4f42682a91658138b4fb75bf68669028b1348/apps/web/src/assets/screenshot/azhar.png" /> </td>
<td>

**Burhanuddin Azhar**: Frontend Developer </br>
Github   :  [github.com/Azhar-54](https://github.com/Azhar-54) </br>
Linkedin : [linkedin.com/in/burhanudin-azhar](https://www.linkedin.com/in/burhanudin-azhar/)

</td>
</tr>




</table>


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Contact

If you have any questions or inquiries regarding this project, feel free to contact:

email :burhanuddinazhar2@gmail.com

linkedin : [https://www.linkedin.com/in/burhanudin-azhar/] (https://www.linkedin.com/in/burhanudin-azhar/)

## Acknowledgements

Feel free to check it out:

- [Choose an Open Source License](https://choosealicense.com/)
- [GitHub Pages](https://pages.github.com/)
