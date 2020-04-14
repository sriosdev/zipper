# GoZipper
## Package for compressing directories into a ZIP file
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