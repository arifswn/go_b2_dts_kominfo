final project sesi 12

buat aplikasi web service : 
    - flow/workflow
        > logic
        > response
        > desain/tampilan
    - relasi database
        

buat final project
    - aplikasi MyGram
        fitur :
            - menyimpan/store foto
            - membuat comment untuk foto orang lain
            - dilengkapi dengan proses crud
            - *all library
            - github.com/dgrijalva/jwt-go > untuk json web token (jwt)
            - golang.org/x/crypto > enkripsi
        database :
            - tb_user
                field :
                    - id > primary key > auto increment
                    - username > string
                    - email > string
                    - password > string
                    - age > integer
                    - created_at > date (timestamp : generate auto tanggal disaat create/store)
                    - updated_at > date (timestamp : generate auto tanggal disaat update)
                    - deleted_at > date (timestamp : generate auto tanggal disaat deleted)
            - tb_photo
                field : 
                    - id > primary key > auto increment
                    - title > string
                    - caption > string
                    - photo_url > string*
                    - type_photo > enum > file_default | file_url
                    - user_id > foreign key ke tb_user : id_user
                    - created_at > date (timestamp : generate auto tanggal disaat create/store)
                    - updated_at > date (timestamp : generate auto tanggal disaat update)
                    - deleted_at > date (timestamp : generate auto tanggal disaat deleted)
            - query sample :
                select tu.id as id_user, tu.username, tu.email, tu.created_at as create_user,
                    tp.id as id_photo, tp.title, tp.caption, tp.photo_url, tp.created_at as create_photo
                    from tb_user tu
                join tb_photo tp on tp.user_id=tu.id_user                
        flow crud :
            - delete :
                1. boleh menghapus data dengan cara 
                    "deleted_at date > update date"
                    "status_delete > 0=blm dihapus, 1=dihapus"
                    select * from tb_user where status_delete = "0"
                2. hapus record berdasarkan id_user
            - penyimpanan photo
                1. by url > get url (online)
                2. by file photo > extension png|PNG|jpg|JPG|jpeg|JPEG
                    yg di save di dalam tb_photo > dengan field photo_url > nama_file.extension
                    sample : bola.jpeg
            - file yang akan di upload > simpan ke dalam assets atau uploads
                - assets > resource > template/css/photo/favicon
                - uploads > resource > file upload yang tadi dikirim ditaruh di dalam folder uploads
            - struktur folder apps
                mod go > library (go get *)
                main.go > index default > first run app
                assets > favicon / response code 404/500/dll
                uploads > mygram > file photo
            - get data url photo > http://localhost:5000/uploads/mygram/bola.jpeg
            - logic :
                if photo_url == "file_url" then
                    // get data url photo online
                    https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTK-HIl85fO5Dkp-4qUDI2276RrQ_QDObuviaCpY_aVDg&s
                else
                    // get data url photo uploads
                    http://localhost:5000/uploads/mygram/bola.jpeg
            - fokus point tugas :
                - sesuai di module.
                    > improve            
        referense :
            - gin gonic & orm gorm
                        
flow : (crud)
    1. user register > generate crypto > hash password > simpan data ke tb user
    2. user login > generate token > token dengan type auth bearer > respon token
    3. user input sosial media > simpan data ke tb sosial media
    4. user input poto > masukan data photo (title, caption, photo_url) 
        dan upload foto (form-data > postman) > kirim data > simpan data ke tb photo
    note untuk user : melakukan comment terhadap poto
        user punya data 
            1. sudah melakukan register
            2. sudah login
            3. get data poto seluruh user
    5. user input data comment > 
        1. terlebih dahulu get data photo
        2. list data comment di dapat dari get poto (where id photo)
        3. user chat (input comment ke dalam photo)
        4. improve
            
asumsi :
    1. get data comment pertama > arif : wah keren banget pemandangan malam minggu
    2. reply > elang : iya nih yuk nobar di malam minggu
    3. reply > puspa : yuk yuk sambil ngopi nyambi bahasan koding
    
    ig :
    1. get data comment pertama > arif : wah keren banget pemandangan malam minggu
        1. reply > elang : iya nih yuk nobar di malam minggu
            1. reply > arif : wah boleh mas, kapan ya kira bisa bareng ngopi
            2. reply > elang : siap di atur waktunya
        2. reply > puspa : yuk yuk tmn bareng,, sambil ngopi nyambi bahasan koding