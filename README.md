# GoZipper
## Package for compressing directories into a ZIP file
### Install
```
# go get github.com/sriosdev/zipper@{version}
$ go get github.com/sriosdev/zipper@v1.0.1
```

### Use

#### Import:
```
import "github.com/srios/zipper"
```

#### Create ZIP:
```
// Open a directory
dir, err := os.Open("/mypath")
if err != nil {
    log.fatalln(err)
    return
}

// Make ZIP file
zipfile, err = zipper.ZipFolder(dir)
```