# E-Wallet App

E-Wallet App adalah aplikasi dompet digital sederhana yang memungkinkan pengguna untuk melakukan registrasi, login, transfer, dan top-up saldo menggunakan Midtrans Sandbox.

## Fitur

- Registrasi pengguna baru.
- Login pengguna.
- Transfer saldo ke pengguna lain.
- Top-up saldo menggunakan Midtrans Sandbox.

## Persyaratan

- Go 1.20 atau yang lebih baru.
- Database PostgreSQL.
- Framework nya Fiber, dan memakai goqu.database
- Akun Midtrans dan API Key Midtrans Sandbox (https://sandbox.midtrans.com/).

## Konfigurasi

1. Clone repository ini:

    ```bash
    git clone https://github.com/reyhanhmdani/e-wallet.git
    cd e-wallet
    ```

2. Salin file `.env.example` ke `.env` dan atur konfigurasi yang diperlukan:

    ```bash
    cp .env.example .env
    ```

   Edit `.env` dan sesuaikan nilai-nilai yang diperlukan.

3. Install dependensi proyek:

    ```bash
    go mod download
    ```

4. Jalankan aplikasi:

    ```bash
    go run main.go
    ```

   Aplikasi akan berjalan di `http://localhost:5432`.

## Penggunaan

1. Registrasi Pengguna Baru:

    ```bash
    curl -X POST http://localhost:5432/register -d '{"fullname": "fullname_anda", "phone": "081234567890", "email": "email@example.com", "username": "username_anda", "password": "secretpass"}'
    ```

2. Login:

    ```bash
    curl -X POST http://localhost:5432/login -d '{"username": "username_anda, "password": "secretpass"}'
    ```

   Catatan: Salin token akses dari respons untuk digunakan di permintaan berikutnya.

3. Transfer Saldo:

    ```bash
    curl -X POST http://localhost:3000/transfer/inquiry -H 'Authorization: Bearer <token>' -d '{"receiver_username": "target_user", "amount": 100}'
    ```
   // disini untuk konfirmasi ke no siapa anda ingin transfer, maka akan muncul nanti inquiry nya yang mana untuk validasi


4. Transfer execute


    ```bash
    curl -X POST http://localhost:3000/transfer/execute -H 'Authorization: Bearer <token>' -d '{"inquiry_key": "inquiry nya" "PIN", "amount": secret_pin}'
    ```

    // disini baru memakai hasil dari inquiry tadi, dan juga harus memasukkan pin anda sendiri, supaya berhasil transfer ke no tujuan

5. Top-up Saldo:

    ```bash
    curl -X POST http://localhost:5432/topup/initialize -H 'Authorization: Bearer <token>' -d '{"amount": 50000}'
    ```

   Catatan: - Pastikan untuk menggunakan Midtrans Sandbox untuk melakukan pembayaran,
   - responnya akan menghasilkan url, yang mana akan mengarahkan ke pembayaran nya, dan cara accept pembayarannya memakai simulator payment

