package requests

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
)

// maxDropboxUpload is 2MiB
const (
	maxDropboxUpload = 2 * 1024 * 1024
	fileUploadPath   = "/class_sessions/"
)

func (s WebmWriter) getVideoFileSharableLink() (string, error) {
	if models.Config.DropboxToken != "" {
		dropBoxUploader, err := newDropboxUploader(s.FileName)
		if err != nil {
			log.Errorln("unable to initialize dropbox uploader", err)
			return "", err
		}

		link, err := dropBoxUploader.dropboxFileUploader()
		if err != nil {
			log.Errorln("unable to get dropbox sharable link", err)
			return "", err
		}

		log.Infoln("file uploaded to file server")

		return link, nil
	}

	return "", nil
}

func (s WebmWriter) uploadToDB() {
	defer func() {
		if err := os.Remove(s.FileName); err != nil {
			log.Errorln("unable to remove file", err)
		}
	}()

	if err := uploadFileGridFS(s.FileName); err != nil {
		log.Errorln("error saving file to DB", err)
		return
	}

	log.Println("File uploaded to DB")
}


// newDropboxUploader uploads file to /class_session path on dropbox
func newDropboxUploader(filePath string) (*dropboxUploader, error) {
	config := dropbox.Config{Token: values.Config.DropboxToken}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileFullPath := fileUploadPath + fileStat.Name()

	fileUploadInfo := files.NewCommitInfo(fileFullPath)
	fileUploadInfo.Mode.Tag = "overwrite"
	fileUploadInfo.ClientModified = time.Now().UTC().Round(time.Second)

	return &dropboxUploader{
		uploadClient:       files.New(config),
		sharableLinkClient: sharing.New(config),
		fileUploadInfo:     fileUploadInfo,

		file:         file,
		fileFullPath: fileFullPath,
		fileSize:     fileStat.Size(),
	}, nil
}

func (d *dropboxUploader) dropboxFileUploader() (string, error) {
	defer func() {
		if err := os.Remove(d.file.Name()); err != nil {
			log.Println("unable to remove file", err)
		}
	}()

	if err := d.uploadFile(); err != nil {
		// Upload file to server to be later uploaded by administrators.
		log.Println("could not upload file, saving to database", err)
		uploadFileGridFS(d.file.Name())
		return "", err
	}

	return d.getSharableLink()
}

func (d *dropboxUploader) uploadFile() error {
	var err error
	if d.fileSize > maxDropboxUpload {
		err = d.UploadFileChunked()
	} else {
		_, err = d.uploadClient.Upload(d.fileUploadInfo, d.file)
	}

	return err
}

func (d *dropboxUploader) UploadFileChunked() error {
	uploadSession, err := d.uploadClient.UploadSessionStart(files.NewUploadSessionStartArg(), &io.LimitedReader{R: d.file, N: maxDropboxUpload})
	if err != nil {
		log.Println("could not start upload session", err)
		return err
	}

	sentChunk := int64(maxDropboxUpload)

	for (d.fileSize - sentChunk) > maxDropboxUpload {
		arg := files.NewUploadSessionAppendArg(files.NewUploadSessionCursor(uploadSession.SessionId, uint64(sentChunk)))

		if err := d.uploadClient.UploadSessionAppendV2(arg, &io.LimitedReader{R: d.file, N: maxDropboxUpload}); err != nil {
			log.Println("could not append to upload session", err)
			return err
		}

		log.Println(sentChunk, "sent to path", d.fileFullPath, "remaining", d.fileSize)
		sentChunk += maxDropboxUpload
	}

	cursor := files.NewUploadSessionCursor(uploadSession.SessionId, uint64(sentChunk))
	args := files.NewUploadSessionFinishArg(cursor, d.fileUploadInfo)

	if _, err := d.uploadClient.UploadSessionFinish(args, d.file); err != nil {
		log.Println("could not finish upload session", err)
		return err
	}

	return nil
}

func (d *dropboxUploader) getSharableLink() (string, error) {
	linkArg := sharing.NewCreateSharedLinkWithSettingsArg(d.fileFullPath)
	res, err := d.sharableLinkClient.CreateSharedLinkWithSettings(linkArg)
	if err != nil {
		return "", err
	}

	switch sl := res.(type) {
	case *sharing.FileLinkMetadata:
		return sl.Url, nil
	}

	return "", models.ErrFileUpload
}
