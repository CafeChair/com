### com

- com.ColorLog(color string, info interface{})
- com.WalkDir(dirpath,suffix string) (files []string, err error)
- com.MD5Files(root string) (map[string][md5.Size]byte, error)
- com.IsFile(filename string) bool
- com.IsExist(filename string) bool
- com.ParseFile(filename string) ([]string, error)
- com.NewUUID() string
- com.WriteFile(filename, str string, w writeFile)