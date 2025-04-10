package article

import (
	"article/pkg/errors"
	"bytes"
	"context"
	"log"

	"github.com/jung-kurt/gofpdf"
	"article/internal/entity/article"
)

func (s Service) GetAllUser(ctx context.Context) ([]article.User, error) {

	users, err := s.data.GetAllUser(ctx)
	if err != nil {
		return users, errors.Wrap(err, "[SERVICE][GetAllUser]")
	}

	return users, err
}

func (s Service) GeneratePDF(ctx context.Context) ([]byte, error) {

	docPdf := bytes.Buffer{}

	users, err := s.data.GetAllUser(ctx)
	if err != nil {
		return docPdf.Bytes(), errors.Wrap(err, "[SERVICE][GeneratePDF]")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")

	cellWidth := 47.5
	cellHeight := 10.0

	pdf.AddPage()
	pdf.SetFont("Arial", "", 20)

	pdf.CellFormat(190, 20, "Daftar User", "0", 0, "C", false, 0, "")
	pdf.Ln(30)

	pdf.SetFont("Arial", "", 12)

	pdf.CellFormat(cellWidth, 15, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(cellWidth, 15, "Username", "1", 0, "C", false, 0, "")
	pdf.CellFormat(cellWidth, 15, "Role", "1", 0, "C", false, 0, "")
	pdf.CellFormat(cellWidth, 15, "Phone Number", "1", 1, "C", false, 0, "")

	for _, user := range users {
		pdf.CellFormat(cellWidth, cellHeight, user.Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, user.UserName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, user.RoleName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, user.PhoneNumber, "1", 1, "C", false, 0, "")
	}

	err = pdf.Output(&docPdf)
	if err != nil {
		log.Fatalf("Error creating PDF: %s", err)
	}

	return docPdf.Bytes(), err
}
