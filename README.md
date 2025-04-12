Untuk menjalankan project

1. Membuat database pada MySql dengan query
```bash
CREATE DATABASE IF NOT EXISTS article;
```

2. Membuat table dengan query
```bash
CREATE TABLE IF NOT EXISTS posts (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(200) NOT NULL,
    Content TEXT NOT NULL,
    Category VARCHAR(100) NOT NULL,
    Created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    Status VARCHAR(100) NOT NULL CHECK (status IN ('Publish', 'Draft', 'Thrash'))
);
```

3. Jalankan program dengan command --> go run cmd/http/main.go
