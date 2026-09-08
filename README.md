# backend-nrp

Repo tugas mata kuliah **Pengembangan Backend Dasar**, dibuat dari template [`webdev-if-its/backend-template`](https://github.com/webdev-if-its/backend-template). Ganti judul di atas jadi nama repo kalian sendiri (`backend-nrp`, contoh: `backend-5025201012`).

## Aturan Umum

- Tugas tiap pertemuan disimpan di folder `pertemuan-XX/` pada repo ini.
- Commit message wajib menyebut level yang dicapai: `pertemuan-XX: level N selesai`.
- Deadline push: sebelum pertemuan berikutnya dimulai.
- Semua level dicek otomatis lewat `go test` — baca `pertemuan-XX/SOAL.md` tiap minggu untuk detail levelnya.

## Mengambil Pertemuan Baru Tiap Minggu

Repo ini **tidak otomatis sinkron** dengan template dosen. Begitu ada pertemuan baru, jalankan (ganti `pertemuan-02` sesuai minggu berjalan):

```bash
git fetch https://github.com/webdev-if-its/backend-template.git main
git checkout FETCH_HEAD -- pertemuan-02
```

Perintah ini **aman dijalankan kapan pun** — tidak akan menimpa folder pertemuan lain yang sudah kalian kerjakan, karena hanya mengambil folder yang disebutkan. Setelah itu, commit folder barunya seperti biasa.

Kalau dosen memperbaiki sesuatu di pertemuan yang sudah dirilis (mis. ada bug di test), biasanya cukup ambil ulang file yang diperbaiki saja, bukan seluruh folder — akan diumumkan file mana yang berubah.

---

Bagian di bawah ini **isi bertahap** sesuai level yang sedang kalian kerjakan (lihat `pertemuan-01/SOAL.md`) — heading-nya dicek otomatis, jangan diganti namanya.

## Identitas
- Nama: Krisna Anugrah Arianto Heru Putro
- NRP: 5053241002
- Kelas: M

## Commit vs Push
Commit: git commit adalah sebuah aksi dimana aksi tersebut menyimpan perubahan kode secara lokal atau di komputer.
Push: git push adalah sebuah aksi dimana aksi tersebut "mengunggah" hasil commit ke repository

## Reproducibility
Reproducibility adalah kondisi dimana terdapat perbedaan versi Go diantara anggota tim. Kondisi ini dapat menjadi masalah jika seseorang menjalankan kode Go versi terbaru, namun perangkatnya masih menggunakan Go dengan versi dibawahnya. Hal tersebut dapat menyebabkan kegagalan build, dikarenakan ada beberapa library baru yang belum ada di versi Go lama.

## Catatan Merge Conflict
Conflict terjadi di fungsi CetakInfo pada file main.go. Conflict tersebut terjadi dikarenakan terdapat fungsi yang sama namun dengan instruksi yang berbeda. Pada branch fitur-sapaan, fungsi CetakInfo menambahkan fungsi Sapa(nama). Sedangkan pada branch main, fungsi CetakInfo telah diubah urutannya. Disini saya menyelesaikan dengan menghapus code dari branch main, sehingga fungsi CetakInfo berasal dari branch fitur-sapaan.

## Kenapa .gitignore Penting
.gitignore penting karena dengan .gitignore membantu mencegah file file yang tidak diperlukan serta file sensitif dalam direktori lokal ikut di upload ke repository online. Salah satu contohnya adalah ketika ada file build IDE yang ikut ter-commit dan di push ke repository. Jika ada anggota tim lain yang melakukan pull dari repository tersebut, maka ketika programnya dijalankan, akan menemui error. Karena, terdapat perbedaan konfigurasi environtment antara perangkat tim lain dengan file build IDE yang ikut ter-pull. Hal tersebut dapat menghambat kerja dikarenakan, harus melakukan debugging terkait masalah konfigurasi environtment yang seharusnya bisa dicegah dengan menggunakan .gitignore.

## Refleksi
Bagi saya, soal yang membingungkan ada di soal 8. Soal yang membahas terkait branching, merging, dan conflict. Karena ini pertama kali bagi saya melakukan skenario branching, merging, dan conflict. Namun, setelah mencobanya secara langsung saya mulai paham dan conflict membantu mencegah kode yang tumpang tindih ketika hendak di merge.