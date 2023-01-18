# Final-Project-Sanbercode-Go-Batch-41
# Documentation API
# 1. Register-Admin
# GET Luas : http://localhost:8080/segitiga-sama-sisi?hitung=luas&alas=5&tinggi=6
    1. Response :
        {
        "output": {
            "alas": 5,
            "tinggi": 6,
            "hasil": 15
        }
    }
# 2. Register
# GET Keliling : http://localhost:8080/segitiga-sama-sisi?hitung=keliling&alas=5&tinggi=6
    1. Response :
        {
        "output": {
            "alas": 5,
            "tinggi": 6,
            "hasil": 15
        }
    }
# 2. Login
# GET Luas : http://localhost:8080/persegi?hitung=luas&sisi=10
    1. Response :
        {
        "output": {
            "sisi": 10,
            "hasil": 100
        }
    }
# 3. Brands
# GET (All) : http://localhost:8080/persegi?hitung=keliling&sisi=10
    1. Response :
        {
        "output": {
            "sisi": 10,
            "hasil": 40
        }
    }
# GET (By ID) : http://localhost:8080/persegi-panjang?hitung=luas&panjang=5&lebar=3
    1. Response :
        {
        "output": {
            "panjang": 5,
            "lebar": 3,
            "hasil": 15
        }
    }
# POST (Create Brand) : http://localhost:8080/persegi-panjang?hitung=keliling&panjang=5&lebar=3
    1. Response :
        {
        "output": {
            "panjang": 5,
            "lebar": 3,
            "hasil": 16
        }
    }
# PUT (Update Brand) : http://localhost:8080/lingkaran?hitung=luas&jariJari=7
    1. Response :
        {
        "output": {
            "jari_jari": 7,
            "hasil": 153.93804002589985
        }
    }
# DELETE (Delete Brand) : http://localhost:8080/lingkaran?hitung=keliling&jariJari=7
    1. Response :
        {
        "output": {
            "jari_jari": 7,
            "hasil": 153.93804002589985
        }
    }
# 1. Phones
# GET (All) :  http://localhost:8080/categories
    1. Response :
        {
            "result": [
            {
                "id": 1,
                "name": "pemrograman",
                "created_at": "2023-01-15T01:05:02.484004Z",
                "updated_at": "2023-01-15T02:00:57.470563Z"
            },
            {
                "id": 2,
                "name": "bahasa",
                "created_at": "2023-01-15T01:05:15.906159Z",
                "updated_at": "2023-01-15T01:05:15.906159Z"
            },
            {
                "id": 3,
                "name": "sosial",
                "created_at": "2023-01-15T01:26:22.705704Z",
                "updated_at": "2023-01-15T01:26:22.705704Z"
            },
            {
                "id": 4,
                "name": "olahraga",
                "created_at": "2023-01-15T01:32:00.475199Z",
                "updated_at": "2023-01-15T01:32:00.475199Z"
            },
            {
                "id": 5,
                "name": "matematika",
                "created_at": "2023-01-15T01:35:04.567262Z",
                "updated_at": "2023-01-15T01:35:04.567262Z"
            },
            {
                "id": 6,
                "name": "komputer",
                "created_at": "2023-01-15T01:36:44.708672Z",
                "updated_at": "2023-01-15T01:36:44.708672Z"
            },
            {
                "id": 7,
                "name": "kimia",
                "created_at": "2023-01-15T01:38:20.262782Z",
                "updated_at": "2023-01-15T15:34:05.823863Z"
            }
        ]
    }
# GET (Phone By Brand ID) : http://localhost:8080/categories
    1. Input :
        - raw JSON
        {
            "name": "sejarah"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
        {
        "data": {
            "id": 8,
            "name": "sejarah",
            "created_at": "2023-01-15T18:31:04.843844Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "status": "success"
    }
    # GET (By ID) : http://localhost:8080/categories
    1. Input :
        - raw JSON
        {
            "name": "sejarah"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
        {
        "data": {
            "id": 8,
            "name": "sejarah",
            "created_at": "2023-01-15T18:31:04.843844Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "status": "success"
    }
# PUT (Update Phone) : http://localhost:8080/categories/8
    1. Input :
        - raw JSON
        {
            "name": "fisika"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
    {
        "result": "Success Update Category"
    }
    # POST (Create Phone) : http://localhost:8080/categories/8
    1. Input :
        - raw JSON
        {
            "name": "fisika"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
    {
        "result": "Success Update Category"
    }
# DELETE (Delete Phone) : http://localhost:8080/categories/8
    1. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    2. Response :
    {
        "result": "Success Delete Category"
    }
# 3. Specs
# GET (All) :  http://localhost:8080/categories
    1. Response :
        {
            "result": [
            {
                "id": 1,
                "name": "pemrograman",
                "created_at": "2023-01-15T01:05:02.484004Z",
                "updated_at": "2023-01-15T02:00:57.470563Z"
            },
            {
                "id": 2,
                "name": "bahasa",
                "created_at": "2023-01-15T01:05:15.906159Z",
                "updated_at": "2023-01-15T01:05:15.906159Z"
            },
            {
                "id": 3,
                "name": "sosial",
                "created_at": "2023-01-15T01:26:22.705704Z",
                "updated_at": "2023-01-15T01:26:22.705704Z"
            },
            {
                "id": 4,
                "name": "olahraga",
                "created_at": "2023-01-15T01:32:00.475199Z",
                "updated_at": "2023-01-15T01:32:00.475199Z"
            },
            {
                "id": 5,
                "name": "matematika",
                "created_at": "2023-01-15T01:35:04.567262Z",
                "updated_at": "2023-01-15T01:35:04.567262Z"
            },
            {
                "id": 6,
                "name": "komputer",
                "created_at": "2023-01-15T01:36:44.708672Z",
                "updated_at": "2023-01-15T01:36:44.708672Z"
            },
            {
                "id": 7,
                "name": "kimia",
                "created_at": "2023-01-15T01:38:20.262782Z",
                "updated_at": "2023-01-15T15:34:05.823863Z"
            }
        ]
    }
# GET (By ID) : http://localhost:8080/categories
    1. Input :
        - raw JSON
        {
            "name": "sejarah"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
        {
        "data": {
            "id": 8,
            "name": "sejarah",
            "created_at": "2023-01-15T18:31:04.843844Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "status": "success"
    }
# PUT (Update Phone) : http://localhost:8080/categories/8
    1. Input :
        - raw JSON
        {
            "name": "fisika"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
    {
        "result": "Success Update Category"
    }
    # POST (Create Phone) : http://localhost:8080/categories/8
    1. Input :
        - raw JSON
        {
            "name": "fisika"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
    {
        "result": "Success Update Category"
    }
# DELETE (Delete Phone) : http://localhost:8080/categories/8
    1. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    2. Response :
    {
        "result": "Success Delete Category"
    }
# 5. Comments
# GET (All) :  http://localhost:8080/categories
    1. Response :
        {
            "result": [
            {
                "id": 1,
                "name": "pemrograman",
                "created_at": "2023-01-15T01:05:02.484004Z",
                "updated_at": "2023-01-15T02:00:57.470563Z"
            },
            {
                "id": 2,
                "name": "bahasa",
                "created_at": "2023-01-15T01:05:15.906159Z",
                "updated_at": "2023-01-15T01:05:15.906159Z"
            },
            {
                "id": 3,
                "name": "sosial",
                "created_at": "2023-01-15T01:26:22.705704Z",
                "updated_at": "2023-01-15T01:26:22.705704Z"
            },
            {
                "id": 4,
                "name": "olahraga",
                "created_at": "2023-01-15T01:32:00.475199Z",
                "updated_at": "2023-01-15T01:32:00.475199Z"
            },
            {
                "id": 5,
                "name": "matematika",
                "created_at": "2023-01-15T01:35:04.567262Z",
                "updated_at": "2023-01-15T01:35:04.567262Z"
            },
            {
                "id": 6,
                "name": "komputer",
                "created_at": "2023-01-15T01:36:44.708672Z",
                "updated_at": "2023-01-15T01:36:44.708672Z"
            },
            {
                "id": 7,
                "name": "kimia",
                "created_at": "2023-01-15T01:38:20.262782Z",
                "updated_at": "2023-01-15T15:34:05.823863Z"
            }
        ]
    }
# GET (By Coment & Spec By ID) : http://localhost:8080/categories
    1. Input :
        - raw JSON
        {
            "name": "sejarah"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
        {
        "data": {
            "id": 8,
            "name": "sejarah",
            "created_at": "2023-01-15T18:31:04.843844Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "status": "success"
    }
# PUT (Update Phone) : http://localhost:8080/categories/8
    1. Input :
        - raw JSON
        {
            "name": "fisika"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
    {
        "result": "Success Update Category"
    }
    # POST (Create Phone) : http://localhost:8080/categories/8
    1. Input :
        - raw JSON
        {
            "name": "fisika"
        }
    2. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    3. Response :
    {
        "result": "Success Update Category"
    }
# DELETE (Delete Phone) : http://localhost:8080/categories/8
    1. Auth :
        - Basic  Auth:
            username : "admin",
            password : "password"
    2. Response :
    {
        "result": "Success Delete Category"
    }