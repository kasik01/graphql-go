# GraphQL Hasura Demo

Demo project sử dụng Go (gqlgen), PostgreSQL và Hasura GraphQL Engine.

## Yêu cầu

- Docker & Docker Compose
- (Tuỳ chọn) Go 1.22+ nếu muốn chạy app ngoài Docker

## Cấu trúc project

- `server.go`: Entry point của ứng dụng Go GraphQL
- `compose.yaml`: Docker Compose cấu hình PostgreSQL, Hasura, Go app
- `graph/`: Thư mục chứa schema và resolver GraphQL
- `.env`: Biến môi trường (DB, APP_PORT, ...)

## Cài đặt & Chạy

1. **Clone repo và cấu hình biến môi trường**
   - Tạo file `.env` với nội dung ví dụ:
     ```
     DB_USER=postgres
     DB_PASSWORD=Khoa2401_
     DB_NAME=std_mng
     DB_PORT=5432
     APP_PORT=8081
     HASURA_ADMIN_SECRET=youradminsecret
     ```
   - Đảm bảo `HASURA_GRAPHQL_JWT_SECRET` trong compose.yaml đủ 32 ký tự.

2. **Chạy Docker Compose**
   ```sh
   docker compose up --build
   ```

3. **Truy cập các dịch vụ**
   - **Go GraphQL Playground:** [http://localhost:8081/](http://localhost:8081/)
   - **Hasura Console:** [http://localhost:8080/](http://localhost:8080/)  
     Đăng nhập với `HASURA_ADMIN_SECRET` trong `.env`

## Test GraphQL API

### 1. Test với Go App

- Truy cập [http://localhost:8081/](http://localhost:8081/)
- Thử query ví dụ:
  ```graphql
  query {
    students {
      id
      name
      class
    }
  }
  ```

### 2. Test với Hasura

- Truy cập [http://localhost:8080/](http://localhost:8080/)
- Vào tab "API" để thử query/mutation trực tiếp trên database.

### 3. Thêm dữ liệu mẫu

Bạn có thể dùng Hasura Console để thêm bảng, dữ liệu, hoặc dùng mutation của Go app.

---

## Test chức năng Export Excel

### 1. Gửi query export

Sử dụng Hasura Console, Postman, Insomnia, hoặc Playground truy cập endpoint GraphQL (ví dụ: `http://localhost:8081/graph`) và gửi query sau:

```graphql
query {
  ExportStudentsByClass(class: "10A1") {
    success
    message
    data
  }
}
```
- Thay `"10A1"` bằng tên lớp bạn muốn xuất.

### 2. Nhận kết quả

- Nếu thành công, trường `data` sẽ trả về một URL download file Excel, ví dụ:
  ```
  {
    "success": true,
    "message": "Xuất Excel thành công cho lớp 10A1",
    "data": "http://localhost:8081/download/students_10A1_1717040000000.xlsx"
  }
  ```

### 3. Tải file Excel

- Truy cập URL trong trường `data` trên trình duyệt để tải file Excel về máy.

### 4. Lưu ý

- Đảm bảo thư mục `/tmp` trên server có quyền ghi file.
- Đảm bảo đã cấu hình route `/download/{filename}` trong app Go để phục vụ file tĩnh từ `/tmp`.
- Nếu muốn trả về file dạng base64 thay vì URL, hãy mở comment phần code sử dụng `WriteToBuffer` và trả về `base64Data`.

---

## Ghi chú

- Nếu gặp lỗi port, hãy đổi giá trị `APP_PORT` hoặc cổng Hasura trong compose.yaml.
- Đảm bảo key JWT của Hasura đủ 32 ký tự trở lên.

---

**Chúc bạn học tập và phát triển GraphQL vui vẻ!**