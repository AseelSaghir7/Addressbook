## Address Book API server

### Prerequisite :
  1) MySQL installed.
  2) database with name "address_book" should exist.   
  3) Update config.yaml file according to your needs.
 
###Install Dependencies : 
    go mod vendor
### How to run :
    From the root of directory run the following : 
    $ go build -o ab cmd/main.go
    $ ./ab -c config.yaml

### API Docs :
####Create New Address
    /ab/v1/address/new

    Request :
    {
    "first_name" : "aseel",
    "last_name" : "saghir"
    }

    Response :

    {
    "message": "address has been created!",
    "status": 200,
    "data": {
        "first_name": "aseel",
        "last_name": "saghir",
        "phone_number": {
            "String": "",
            "Valid": false
        }
    }
}
####Get Address    
    URL : /ab/v1/address/{addressID}
    
    Request : /ab/v1/address/2
    
    Response : 
    {
    "message": "address",
    "status": 200,
    "data": {
        "id": 3,
        "first_name": "aseel",
        "last_name": "saghir",
        "phone_number": {
                "String": "",
                "Valid": false
            }
        }
    }
####Delete Address
    URL : /ab/v1/address/{addressID}
    
    Request: /ab/v1/address/3
    
    Response:
    {
        "message": "address has been removed!",
        "status": 200,
        "data": "3"
    }
####Get whole Address Book
    URL : /ab/v1/addressBook
    
    Request : /ab/v1/addressBook?sortBy=last

    Response : 
    {
    "message": "address book",
    "status": 200,
    "data": [
        {
            "id": 4,
            "first_name": "aseel",
            "last_name": "saghir",
            "phone_number": {
                "String": "",
                "Valid": false
            }
        },
        {
            "id": 3,
            "first_name": "test",
            "last_name": "test",
            "phone_number": {
                "String": "",
                "Valid": false
            }
        }
    ]
    }
####Search through Address Book
    URL : /ab/v1/addressBook
    
    Request : /ab/v1/addressBook/search?search=as
        
    Response:
    {
    "message": "address book",
    "status": 200,
    "data": [
            {
                "id": 3,
                "first_name": "aseel",
                "last_name": "saghir",
                "phone_number": {
                    "String": "",
                    "Valid": false
                }
            }
        ]
    }

### Run Test :
    From the root of directory run the following : 
    $ go test -v pkg/components/address_book/getAddressBook_test.go pkg/components/address_book/api.go pkg/components/address_book/component.go pkg/components/address_book/type.go