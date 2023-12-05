package usecase

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
)

type Document struct {
	docPort port.Document
}

func NewDocument(i *do.Injector) (*Document, error) {
	docPort := do.MustInvoke[port.Document](i)
	return &Document{
		docPort: docPort,
	}, nil
}

func (d *Document) FindContent(id edinet.DocumentId, fileType edinet.FileType) (edinet.DocumentFile, error) {
	// 全ファイルを取得
	if fileType == edinet.ALL {
		var allFiles []edinet.DocumentFile
		fts := edinet.AllFileType()
		for _, ft := range fts {
			doc, err := d.docPort.Get(id, ft)
			if err != nil {
			}
			allFiles = append(allFiles, doc)
		}
		content, err := createZip(allFiles)
		if err != nil {
			return edinet.DocumentFile{}, err
		}

		response := edinet.DocumentFile{
			Name:       fmt.Sprintf("ALL_%s", id.String()),
			Extension:  ".zip",
			DocumentId: id,
			Content:    content,
		}

		return response, nil
	}

	document, err := d.docPort.Get(id, fileType)
	if err != nil {
		return edinet.DocumentFile{}, err
	}
	return document, nil
}

func createZip(files []edinet.DocumentFile) ([]byte, error) {
	var zipDataBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipDataBuffer)

	for _, file := range files {
		zipEntry, err := zipWriter.Create(file.NameWithExtension())
		if err != nil {
			return nil, err
		}
		_, err = zipEntry.Write(file.Content)
		if err != nil {
			return nil, err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return zipDataBuffer.Bytes(), nil
}
