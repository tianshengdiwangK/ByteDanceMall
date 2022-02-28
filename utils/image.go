package utils

import (
	"github.com/google/uuid"
	"io"
	settings "login_register_demo/utils/setting"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

func SaveImage(file *multipart.FileHeader, folder string) (f string, err error) {
	// 1.打开源文件
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// 2.创建目标文件
	ext := path.Ext(file.Filename) // path包与操作系统无关，使用“/”作为分隔符
	filename := uuid.NewString() + ext
	dst := filepath.Join(settings.StaticResource.ImagePath, folder, filename) // filepath包生成的路径是操作系统相关的
	out, err := create(dst)
	if err != nil {
		return
	}
	defer out.Close()

	// 3.拷贝
	_, err = io.Copy(out, src)
	if err != nil {
		return
	}
	// 4.返回图片相对路径
	f = path.Join("/img", folder, filename)
	return
}

// 当文件所在目录不存在时，先创建目录，再创建文件
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
