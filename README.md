go run .

Truy cập http://localhost:8080/ để mở GraphQL Playground.

📌 1️⃣ Thêm sinh viên (Mutation)
mutation {
  addStudent(input: {
    studentId: "S001",
    name: "Nguyen Van A",
    dateOfBirth: "2001-01-01",
    gender: MALE,
    class: "10A1"
  }) {
    id
    studentId
    name
    dateOfBirth
    gender
    class
  }
}


📌 3️⃣ Lấy sinh viên theo ID (Query)

query {
  student(id: "1") {
    id
    studentId
    name
    dateOfBirth
    gender
    class
  }
}

📌 4️⃣ Tìm kiếm sinh viên theo trường (Ví dụ: tìm theo name)

query {
  searchStudents(field: "name", value: "Nguyen Van A") {
    id
    studentId
    name
    dateOfBirth
    gender
    class
  }
}