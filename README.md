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

## Ghi chú

- Nếu gặp lỗi port, hãy đổi giá trị `APP_PORT` hoặc cổng Hasura trong compose.yaml.
- Đảm bảo key JWT của Hasura đủ 32 ký tự trở lên.

---

**Chúc bạn học tập và phát triển GraphQL vui vẻ!**