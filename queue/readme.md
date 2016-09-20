#### Problem 2
##### Unique Queue

Implementasi dari *interface* `Queue` adalah sebuah *struct* berisi satu *field*, yaitu `Key`, yang menyimpan nilai-nilai item dalam antrian (*queue*) dengan ukuran tetap (*fixed size*) untuk banyaknya item yang bisa disimpan.

###### Method-methodnya:
* `Push` untuk menambahkan item baru ke dalam antrian dan mengusir item paling tua jika slot penuh.
* `Pop` untuk mendapatkan item paling tua supaya dikeluarkan dari antrian, kemudian menggeser antrian.
* `Contains` untuk memeriksa keberadaan sebuah item di dalam antrian.
* `Len` untuk mendapatkan banyaknya item di dalam antrian.
* `Keys` untuk mendapatkan semua item di dalam antrian.

###### Contoh perubahan antrian:
1. **`Push`**  
**Dari kosong:**  
`[<kosong>, <kosong>, <kosong>, <kosong>]`  
Push A!  
`[A, <kosong>, <kosong>, <kosong>]`  
Push B!  
`[A, B, <kosong>, <kosong>]`  
**Dari penuh:**  
`[A, B, C, D]`  
Push E!  
`[B, C, D, E]`  
Push F!  
`[C, D, E, F]`
2. **`Pop`**  
**Dari penuh:**  
`[C, D, E, F]`  
Pop!  
`[D, E, F, <kosong>]`  
Pop!  
`[E, F, <kosong>, <kosong>]`
