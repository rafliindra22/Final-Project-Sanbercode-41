# Final-Project-Sanbercode-Go-Batch-41
# Documentation API
# 1. Register-Admin
# POST : http://localhost:8080/segitiga-sama-sisi?hitung=luas&alas=5&tinggi=6
    1. Input :
        - raw JSON :
        {
            "name":"Rafli",
            "username":"admin",
            "password":"admin"
        }
    2. Response :
        {
        "message": "registration success",
        "username": {
            "full_access": "true",
            "name": "Rafli",
            "username": "admin"
        }
    }
# 2. Register
# POST : http://localhost:8080/segitiga-sama-sisi?hitung=keliling&alas=5&tinggi=6
    1. Input :
        - raw JSON :
        {
            "name":"Test User",
            "username":"test1",
            "password":"test1"
        }
    2. Response :
        {
        "message": "registration success",
        "username": {
            "full_access": "false",
            "name": "Test User",
            "username": "test1"
        }
    }
# 2. Login
# POST : http://localhost:8080/persegi?hitung=luas&sisi=10
    1. Input :
        - raw JSON:
        {
            "username":"user",
            "password":"user"
        }
    2. Response :
        {
        "message": "login success",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzQxMTY0MTQsImZ1bGxfYWNjZXNzIjp0cnVlLCJ1c2VyX2lkIjo0fQ.-S2X8J_HZLf_cjIUH8L0RWiJWxCPBh9llVg8534ZKso",
        "username": {
            "username": "user"
        }
    }
# 3. Brands
# GET (All) : localhost:8080/brands
    1. Response :
        {
        "result": [
            {
                "id": 1,
                "name": "samsung",
                "description": "brand ini berasal dari kore",
                "created_at": "2023-01-17T20:41:25.835167Z",
                "updated_at": "2023-01-17T20:41:25.835167Z",
                "user_id": 4
            },
            {
                "id": 2,
                "name": "xiaomi",
                "description": "brand ini dari china",
                "created_at": "2023-01-17T20:41:25.835167Z",
                "updated_at": "2023-01-18T14:58:30.040563Z",
                "user_id": 4
            },
            {
                "id": 14,
                "name": "iPhone",
                "description": "brand ini berasal dari amerika",
                "created_at": "2023-01-19T14:29:35.240632Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "user_id": 4
            },
            {
                "id": 3,
                "name": "Oppo",
                "description": "brand ini dari china",
                "created_at": "2023-01-17T20:41:25.835167Z",
                "updated_at": "2023-01-19T14:31:21.727724Z",
                "user_id": 4
            }
        ]
    }
# GET (By ID) : localhost:8080/brands/2
    1. Response :
        {
        "data": {
            "id": 2,
            "name": "xiaomi",
            "description": "brand ini dari china",
            "created_at": "2023-01-17T20:41:25.835167Z",
            "updated_at": "2023-01-18T14:58:30.040563Z",
            "user_id": 4
        }
    }
# POST (Create Brand) : http://localhost:8080/persegi-panjang?hitung=keliling&panjang=5&lebar=3
    1. Auth:
        - Admin Middleware
        - Bearer Token
    2. Input:
        - raw JSON:
        {
            "name":"iPhone",
            "description":"brand ini berasal dari amerika"
        }
    3. Response :
        {
        "data": {
            "name": "iPhone",
            "description": "brand ini berasal dari amerika"
        },
        "status": "insert success"
    }
# PUT (Update Brand) : localhost:8080/brands/update/3
    1. Auth:
        - Admin Middleware
        - Bearer Token
    2. Input:
        - raw JSON:
    3. Response :
        {
        "data": {
            "name": "Oppo",
            "description": "brand ini dari china"
        },
        "status": "Update success"
    }
# DELETE (Delete Brand) : localhost:8080/brands/delete/14
    1. Auth :
        - Public Middleware
        - Bearer Token
    2. Response :
    {
        "result": "Success Delete Brand"
    }
# 1. Phones
# GET (All) :  http://localhost:8080/phones
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
# GET (Phone By ID) : localhost:8080/phones/5
    1. Ressponse : 
        {
        "data": {
            "id": 5,
            "type": "galaksi s20",
            "year": 2021,
            "brand_id": 1,
            "created_at": "2023-01-18T01:28:40.889613Z",
            "updated_at": "2023-01-18T15:35:21.82755Z",
            "editor_name": "user",
            "user_id": 4
        }
    }
# GET (Phones By Brand ID) : localhost:8080/brands/1/phones
    1. Response :
       {
        "result": [
            {
                "id": 4,
                "type": "galaksi young",
                "year": 2012,
                "brand_id": 1,
                "created_at": "2023-01-18T01:00:06.847788Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "editor_name": "user",
                "user_id": 4
            },
            {
                "id": 2,
                "type": "galaksi y",
                "year": 2017,
                "brand_id": 1,
                "created_at": "2023-01-18T01:00:06.847788Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "editor_name": "user",
                "user_id": 4
            },
            {
                "id": 3,
                "type": "galaksi s",
                "year": 2020,
                "brand_id": 1,
                "created_at": "2023-01-18T01:00:06.847788Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "editor_name": "user",
                "user_id": 4
            },
            {
                "id": 1,
                "type": "galaksi x",
                "year": 2018,
                "brand_id": 1,
                "created_at": "2023-01-18T01:00:06.847788Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "editor_name": "user",
                "user_id": 4
            },
            {
                "id": 5,
                "type": "galaksi s20",
                "year": 2021,
                "brand_id": 1,
                "created_at": "2023-01-18T01:28:40.889613Z",
                "updated_at": "2023-01-18T15:35:21.82755Z",
                "editor_name": "user",
                "user_id": 4
            }
        ]
    }
# POST (Create Phone) : localhost:8080/phone/create
    1. Auth:
        - Admin Middleware
        - Bearer Token
    2. Input:
        - raw JSON:
        {
        "type":"redmi note 11",
        "year":2020,
        "brand_id":2
    }
    3. Response :
        {
        "data": {
            "type": "redmi note 11",
            "year": 2020,
            "brand_id": 2
        },
        "status": "insert success"
    }
# PUT (Update Phone) : http://localhost:8080/phone/update/8
    1. Auth:
        - Admin Middleware
        - Bearer Token
    2. Input:
        - raw JSON:
        {
        "type":"galaksi s20",
        "year":2020,
        "brand_id":1
    }
    3. Response :
        {
        "data": {
            "type": "galaksi s20",
            "year": 2020,
            "brand_id": 1
        },
        "status": "Update success"
    }
# DELETE (Delete Phone) : http://localhost:8080/phone/delete/5
    1. Auth :
        - Public Middleware
        - Bearer Token
    2. Response :
    {
        "result": "Success Delete Phone"
    }
# 3. Specs
# GET (All) :  localhost:8080/specs
    1. Response :
        {
        "result": [
            {
                "id": 1,
                "phone_id": 1,
                "processor": "mediatek x",
                "memory": "8 GB",
                "storage": "128 GB",
                "screen": "1080x1440",
                "camera": "48 MP",
                "price": "Rp. 3.200.000,00",
                "review": "Mid-End",
                "created_at": "2023-01-18T02:08:03.423587Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "user_id": 4
            },
            {
                "id": 7,
                "phone_id": 2,
                "processor": "mediatek y",
                "memory": "6 GB",
                "storage": "64 GB",
                "screen": "720x1080",
                "camera": "24 MP",
                "price": "Rp. 2.100.000,00",
                "review": "Low-End",
                "created_at": "2023-01-18T15:59:20.39115Z",
                "updated_at": "2023-01-18T16:41:54.937016Z",
                "user_id": 4
            }
        ]
    }
# GET (By ID) : localhost:8080/specs/1
    1. Response :
        {
        "data": {
            "id": 1,
            "phone_id": 1,
            "processor": "mediatek x",
            "memory": "8 GB",
            "storage": "128 GB",
            "screen": "1080x1440",
            "camera": "48 MP",
            "price": "Rp. 3.200.000,00",
            "review": "Mid-End",
            "created_at": "2023-01-18T02:08:03.423587Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "user_id": 4
        }
    }
# PUT (Update Spec) : localhost:8080/spec/update/7
    1. Input :
        - raw JSON
        {
            "phone_id":2,
            "processor":"mediatek y",
            "memory":"4 GB",
            "storage":"64 GB",
            "screen":"720x1080",
            "camera":"24 MP",
            "price":"Rp. 2.100.000,00",
            "review":"Low-End"
        }
    2. Auth :
        - Admin Middleware
        - Bearer Token
    3. Response :
    {
        "data": {
            "phone_id": 2,
            "processor": "mediatek y",
            "memory": "4 GB",
            "storage": "64 GB",
            "screen": "720x1080",
            "camera": "24 MP",
            "price": "Rp. 2.100.000,00",
            "review": "Low-End"
        },
        "status": "Update success"
    }
# DELETE (Delete Phone) : localhost:8080/spec/delete/7
    1. Auth :
        - Admin Middleware
        - Bearer Token
    2. Response :
    {
        "result": "Success Delete Phone"
    }
# 5. Comments
# GET (All) :  http://localhost:8080/comments
    1. No Auth
    2. Response :
        {
        "result": [
            {
                "id": 1,
                "phone_id": 1,
                "name": "Test User",
                "comment": "MEMANG KEREN",
                "created_at": "2023-01-18T02:30:46.492526Z",
                "user_id": 5
            },
            {
                "id": 2,
                "phone_id": 2,
                "name": "Test User",
                "comment": "kurang menarik",
                "created_at": "2023-01-18T02:33:20.333542Z",
                "user_id": 5
            },
            {
                "id": 3,
                "phone_id": 3,
                "name": "Test User",
                "comment": "lumayan lah",
                "created_at": "2023-01-18T02:34:50.001975Z",
                "user_id": 5
            },
            {
                "id": 4,
                "phone_id": 4,
                "name": "Test User",
                "comment": "terlalu jadul",
                "created_at": "2023-01-18T02:37:44.865057Z",
                "user_id": 5
            }
        ]
    }
# GET (By Coment & Spec By PhoneID) : localhost:8080/phones/1/spec-comment
    1. No Auth 
    2. Response :
        {
        "comment": {
            "result": [
                {
                    "id": 1,
                    "phone_id": 1,
                    "name": "Test User",
                    "comment": "MEMANG KEREN",
                    "created_at": "2023-01-18T02:30:46.492526Z",
                    "user_id": 5
                }
            ]
        },
        "spec": {
            "result": [
                {
                    "id": 1,
                    "phone_id": 1,
                    "processor": "mediatek x",
                    "memory": "8 GB",
                    "storage": "128 GB",
                    "screen": "1080x1440",
                    "camera": "48 MP",
                    "price": "Rp. 3.200.000,00",
                    "review": "Mid-End",
                    "created_at": "2023-01-18T02:08:03.423587Z",
                    "updated_at": "0001-01-01T00:00:00Z",
                    "user_id": 4
                }
            ]
        }
    }
# PUT (Create Comment) : localhost:8080/comment/create
    1. Input :
        - raw JSON
        {
            "phone_id":4,
            "comment":"terlalu jadul"
        }
    2. Auth :
        - Public Middleware
        - Bearer  TOken
    3. Response :
    {
        "data": {
                "phone_id": 4,
                "comment": "terlalu jadul"
        },
         "status": "comment success"
    }
# DELETE (Delete Comment) : localhost:8080/comment/delete/4
    1. Auth :
        - Public Middleware
        - Bearer Token
    2. Response :
    {
    "result": "Success Delete Comment"
    }