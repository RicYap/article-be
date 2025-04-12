Untuk menjalankan code

1. Membuat database pada MySql dengan query
CREATE TABLE IF NOT EXISTS posts (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(200) NOT NULL,
    Content TEXT NOT NULL,
    Category VARCHAR(100) NOT NULL,
    Created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    Status VARCHAR(100) NOT NULL CHECK (status IN ('Publish', 'Draft', 'Thrash'))
);

2. Jalankan program dengan command --> go run cmd/http/main.go
