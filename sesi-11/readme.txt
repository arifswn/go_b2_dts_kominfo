unit test
1. sblm aplikasi di running (debugging, development, production)
2. menguji code > function (logic, flow, dll)
    - menguji code aritmatika > penambahan
    - sample : 1 + 2 = 3
3. proses testing
4. proses debugging > sblm proses debugging > uji coba function / uji coba code
    - proses testing > lakuin tahap testing 
        > unit test, mock, testify, jmeter
5. proses development (di developer)
6. proses uji testing (development > end user | kita berumpaa menjadi end user)
6. proses production (benar" data itu real)

asumsi :
1. aritmatika.go (tidak menggangu function aritmatika yang asli) 
    - function penambahan
2. aritmatika_test.go (mencall function yang akan di test) >> fokus QA | bug hunter
    - untuk testing
note :
jika nanti web service kita di "oprek|testing|etc"
    > end point > transparansi > terlihat ?

----------------------------------------------------

1. buat folder sesi-11
2. cd sesi-11
3. mkdir hello-testing
4. cd hello-testing
5. mkdir helper
6. cd helper
7. touch calculation.go calculation_test.go
8. koding enjoy > calculation & calculation_test

running testing :
go test ./helper