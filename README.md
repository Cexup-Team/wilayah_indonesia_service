# Servis Wilayah Indonesia
Kode dan Data Wilayah Administrasi Pemerintahah Indonesia sesuai Permendagri No 72 Tahun 2019 (dgn pembaruan data berdasarkan  Kepmendagri No.146.1-4717 Tahun 2020).

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/cahyadsn/wilayah.svg)](https://github.com/Cexup-Team/wilayah_indonesia_service/issues)


## Project Structure

    .
    ├── app                        # Configuration file (include `connection`)
    ├── controller                 # Controller files
    ├── helper                     # Helper function
    ├── model                      # Model files (alternatively `entities`)
    ├     └── domain                
    ├     └── web                
    ├── repository                 # Repository files (alternatively `persistence`)
    ├── services                   # Services files (`business logic`)
    ├── main
    ├── readme
    └── tests
     
    

### Database

This service using "MySql" to stored wilayah indonesia has been searched by user.


### Endpoint

To connected to this endpoint please ensure your base_url is "localhost:3000" or you can adjust your own port in main.go file

Get provinsi  ( GET )
```sh
{base_url}/provinsi
```

Show by wilayah_id ( GET )
```sh
{base_url}/wilayah/:wilayah_id
```

