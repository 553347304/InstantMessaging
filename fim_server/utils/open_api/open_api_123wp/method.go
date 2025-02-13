package open_api_123wp

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/https"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/valid"
	"io"
	"mime/multipart"
)

type fileResponse struct {
	ID   uint   `json:"id"`   // 文件ID
	Name string `json:"name"` // 文件名
	Type int    `json:"type"`
}

func (t *trayService) upload(h *multipart.FileHeader, parentFileID uint, data []byte) error {
	r := https.Post(t.url.Upload, t.header, conv.Json().Marshal(map[string]any{
		"parentFileID": parentFileID,
		"filename":     h.Filename,
		"etag":         valid.MD5().Hash(string(data)),
		"size":         h.Size,
	}))
	if r.Error != nil {
		return logs.Error(r.Error)
	}
	var scan baseResponse[struct {
		FileID      int    `json:"fileID"`
		Reuse       bool   `json:"reuse"`
		PreuploadID string `json:"preuploadID"`
		SliceSize   int    `json:"sliceSize"`
	}]
	if !conv.Json().Unmarshal(r.Body, &scan) {
		return logs.Error("参数绑定错误", string(r.Body))
	}
	if scan.Data.Reuse {
		return nil
	}

	t.preuploadID = scan.Data.PreuploadID
	return nil
}
func (t *trayService) getUploadUrl(sliceNo int) error {
	r := https.Post(t.url.GetUploadUrl, t.header, conv.Json().Marshal(map[string]any{
		"preuploadID": t.preuploadID,
		"sliceNo":     sliceNo,
	}))
	if r.Error != nil {
		return logs.Error(r.Error)
	}
	var scan baseResponse[struct {
		PresignedURL string `json:"presignedURL"`
		IsMultipart  bool   `json:"isMultipart"`
	}]
	if !conv.Json().Unmarshal(r.Body, &scan) {
		return logs.Error("参数绑定错误")
	}
	t.presignedURL = scan.Data.PresignedURL
	return nil
}
func (t *trayService) uploadSlice(data []byte) error {
	r := https.Put(t.presignedURL, map[string]string{"Platform": "open_platform"}, data)
	return logs.Error(r.Error)
}
func (t *trayService) uploadSucceed() error {
	r := https.Post(t.url.UploadSucceed, t.header, conv.Json().Marshal(map[string]any{
		"preuploadID": t.preuploadID,
	}))
	return logs.Error(r.Error)
}

func (t *trayService) GetURL(fileID uint) (string, error) {
	r := https.Get(t.url.GetURL, t.header, map[string]any{
		"fileID": fileID,
	})
	if r.Error != nil {
		return "", logs.Error(r.Error)
	}
	var scan baseResponse[struct {
		Url string `json:"url"`
	}]
	if !conv.Json().Unmarshal(r.Body, &scan) {
		return "", logs.Error("参数绑定错误", string(r.Body))
	}

	if scan.Code != 0 {
		return "", logs.Error(scan.Message)
	}
	return scan.Data.Url, nil
}
func (t *trayService) List(parentFileID uint, page int, search string) ([]fileResponse, error) {
	r := https.Get(t.url.FileList, t.header, map[string]any{
		"parentFileId":   parentFileID,
		"page":           page,
		"limit":          100,
		"orderBy":        "file_id",
		"orderDirection": "desc",
		"trashed":        0,
		"searchData":     search,
	})
	logs.Info(r.Response.Request.URL)
	if r.Error != nil {
		return nil, logs.Error(r.Error)
	}
	var scan baseResponse[struct {
		Total    int `json:"total"`
		FileList []struct {
			FileID       uint   `json:"fileID"`
			Filename     string `json:"filename"`
			ParentFileId int    `json:"parentFileId"`
			ParentName   string `json:"parentName"`
			Type         int    `json:"type"`
			Etag         string `json:"etag"`
			Size         int    `json:"size"`
			ContentType  string `json:"contentType"`
			Category     int    `json:"category"`
			Hidden       bool   `json:"hidden"`
			Status       int    `json:"status"`
			PunishFlag   int    `json:"punishFlag"`
			S3KeyFlag    string `json:"s3KeyFlag"`
			StorageNode  string `json:"storageNode"`
			CreateAt     string `json:"createAt"`
			UpdateAt     string `json:"updateAt"`
			Thumbnail    string `json:"thumbnail"`
			DownloadUrl  string `json:"downloadUrl"`
		} `json:"fileList"`
	}]
	if !conv.Json().Unmarshal(r.Body, &scan) {
		return nil, logs.Error("参数绑定错误", string(r.Body))
	}

	var response []fileResponse
	for _, file := range scan.Data.FileList {
		response = append(response, fileResponse{
			ID:   file.FileID,
			Name: file.Filename,
			Type: file.Type,
		})
	}

	return response, nil
}
func (t *trayService) Upload(h *multipart.FileHeader, parentFileID uint) error {
	open, _ := h.Open()
	data, err := io.ReadAll(open)
	if err != nil {
		return logs.Error("读取错误")
	}

	err = t.upload(h, parentFileID, data)
	if err != nil {
		return err
	}

	if t.preuploadID == "" {
		return nil
	}

	err = t.getUploadUrl(1)
	if err != nil {
		return err
	}
	err = t.uploadSlice(data)
	if err != nil {
		return err
	}
	err = t.uploadSucceed()
	if err != nil {
		return err
	}

	return nil
}
func (t *trayService) Mkdir(parentID uint64, name string) (uint64, error) {

	r := https.Post(t.url.Mkdir, t.header, conv.Json().Marshal(map[string]any{
		"name":     name,
		"parentID": parentID,
	}))

	if r.Error != nil {
		return 0, logs.Error(r.Error)
	}

	var scan baseResponse[struct {
		DirID uint64 `json:"dirID"`
	}]
	if !conv.Json().Unmarshal(r.Body, &scan) {
		return 0, logs.Error("参数绑定错误", string(r.Body))
	}
	return scan.Data.DirID, nil
}
