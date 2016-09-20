#### Problem 1
##### Sorting and visualization

1. Program utama berisi array bilangan bulat, memanggil subprogram `array-visualizer` yang menampilkan data dalam bentuk *vertical barcharts* dan di bawahnya adalah angka-angka nilainya.  
Program diam sejenak supaya pengguna bisa mengamati hasilnya dengan jelas.

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

2. Subprogram berikutnya adalah `insertion-sort-ascending` yang melakukan pengurutan dengan algoritma *insertion sort* secara *ascending*, yaitu pada setiap langkahnya jika elemen kiri lebih besar daripada elemen kanan, tukar posisi. Visualisasi setiap langkahnya ditampilkan.  
Program diam sejenak lagi.

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

3. Subprogram terakhir adalah `insertion-sort-descending`yang melakukan pengurutan dengan algoritma *insertion sort* secara *descending*, yaitu pada setiap langkahnya jika elemen kiri lebih kecil daripada elemen kanan, tukar posisi. Visualisasi setiap langkahnya juga ditampilkan.  
Program selesai.