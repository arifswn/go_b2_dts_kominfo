map
--
Sama seperti tipe data array dan slice, map juga berfungsi 
untuk menyimpan satu atau lebih data namun,
map disimpan sebagai key:value pairs (pasang key dan value) .

Jika teman-teman pernah mempelajari bahasa Javascript dan PHP, 
maka map pada bahasa Go mirip,
seperti tipe data object pada Javascript dan mirip 
seperti tipe data associative array pada PHP.

Semua key dan value memiliki tipe data yang static, 
sehingga semua key harus memiliki  tipe data yang sama 
begitu pula juga dengan tipe data value nya.

Kemudian setiap key pada sebuah map harus unik 
namun value nyatidak mesti unik.Map juga termasuk 
salah satu tipe data yang masuk dalam kategori 
reference type sama seperti tipe data slice.

Berarti jika ada suatu map yang berusaha 
untuk meng-copy map lainnya, dan map tersebut
mengganti value padasuatu key, maka value dari map lainnya
tersebut juga akan ikut terganti.
Zero value dari tipe data map adalah nil.
