package m_file

import "whacos/models"

type File struct {
	models.Model
	Name    string `json:"name"`
	KeyId   string `json:"keyId"`
	Suffix  string `json:"suffix"`
	Size    int    `json:"size"`
	Address string `json:"address"`
}
