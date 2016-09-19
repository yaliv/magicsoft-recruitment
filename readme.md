## Magicsoft Recruitment Test 2016

#### Instruksi
* Setup *pre-requisite* *compiler*, *package* jika diperlukan.
* Solusi dari seluruh soal bisa diselesaikan dengan menggunakan *standard package*, penggunaan *external package* dibolehkan dengan catatan mampu menjelaskan dengan alasan logis.
* Buat *markdown file* untuk menuliskan *design detail*, penjelasan atau alasan dari implementasi. 
* Solusi dari setiap soal harus bisa di *compile* dan *run* untuk validasi, dipersilahkan menggunakan direktori terpisah atau pendekatan lain.
* *Commit* dan *push* hasil ke branch dengan nama masing-masing, misal "bayu".

#### Soal

1. *Design* dan implementasikan sebuah *program* atau *subprogram* yang akan menampilan visualisasi *data array* sederhana dalam bentuk *vertical barcharts*, dan sebagai tambahan tampilkan setiap nilai data di sumbu *horizontal*.
    
    ```
    INPUT: Numerical array
    [1, 4, 5, 6, 8, 2]

    OUTPUT: Vertical Barcharts

            |   
            |   
          | |  
        | | |   
      | | | |  
      | | | |  
      | | | | |
    | | | | | | 
    1 4 5 6 8 2 

    ```
2. Implementasikan algoritma *insertion sort*, dan gunakan *subprogram* (1) untuk memvisualisasikan setiap langkah/*steps* *sorting* 

    ```
    INPUT: Numerical array

    [1, 4, 5, 6, 8, 2]

    OUTPUT:

    - Sorted array (ascending)
    - Steps visualization

            |   
            |   
          | |  
        | | |   
      | | | |   
      | | | |   
      | | | | | 
    | | | | | | 
    1 4 5 6 8 2 

              | 
              | 
          |   | 
        | |   | 
      | | |   | 
      | | |   | 
      | | | | | 
    | | | | | | 
    1 4 5 6 2 8 

    ... dan seterusnya ...

    ```

3. Modifikasi *subprogram* (2) untuk *reverse sorting* dan lakukan juga visualisasi dengan *subprogram* (1)