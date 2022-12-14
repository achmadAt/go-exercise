package main

import (
	//"os"

	"fmt"
	"image-processing/exif"
	"log"

	//"os/exec"
	"runtime"

	"github.com/ncruces/go-exiftool"
	// go get -u github.com/disintegration/imaging
	//"github.com/rwcarlsen/goexif/exif" // go get -u github.com/rwcarlsen/goexif/exif
	//"github.com/rwcarlsen/goexif/mknote"
	// go get -u github.com/sirupsen/logrus
)

func main() {
	fmt.Println("hello")
	// data, err := ioutil.ReadFile("download.jpeg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// cmd := exec.Command("chmod", "600", "/home/alfazari/go-exercise/image-processing/excutable/exiftool_unix.tgz")
	// _, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err := exif.SetupPaths()
	if err != nil {
		log.Fatal(err, "yo")
	}
	//exiftool.Exec = "/home/alfazari/go-exercise/image-processing/excutable/exiftool_unix.tgz"
	// err := exif.SetupPaths()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	exiftserver, err := exiftool.NewServer("-ignoreMinorErrors")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer exiftserver.Close()
	exfiutil := exif.NewExifUtil(exiftserver)
	err2 := exfiutil.AddOrientation("download.jpeg")
	if err2 != nil {
		log.Fatal(err)
		return
	}
	exfiutil.RawOrientation("image.jpeg")
	fmt.Println(runtime.GOOS)
	fmt.Println(exfiutil.RawOrientation("download.jpeg"))
	// filtered, err := exifremove.Remove(data)
	// if err != nil {
	// 	fmt.Printf("* " + err.Error() + "\n")
	// }
	// serveImage(filtered)
	// fmt.Println(filtered)

	// file, err := os.Open("img.JPG")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// //image, err := jpeg.Decode(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// exif.RegisterParsers(mknote.All...)
	// im, err := exif.Decode(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(im.Raw)
	//serveImage(im.Raw)
	//imaging.Rotate(image, float64(180), color.Gray16{})
	//imaging.Fit(image, 200, 200, imaging.Lanczos)
	// reverseOrientation(image, "2")
}

// func serveImage(imgByte []byte) {

// 	img, _, err := image.Decode(bytes.NewReader(imgByte))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	out, _ := os.Create("img2.jpeg")
// 	defer out.Close()

// 	var opts jpeg.Options
// 	opts.Quality = 20

// 	err = jpeg.Encode(out, img, &opts)
// 	//jpeg.Encode(out, img, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func ReadImage(fpath string) *image.Image {
// 	var img image.Image
// 	var err error
// 	// deal with image
// 	ifile, err := os.Open(fpath)
// 	if err != nil {
// 		logrus.Warnf("could not open file for image transformation: %s", fpath)
// 		return nil
// 	}
// 	defer ifile.Close()
// 	filetype, err := GetSuffix(fpath)
// 	if err != nil {
// 		return nil
// 	}
// 	if filetype == "jpg" || filetype == "jpeg" {
// 		img, err = jpeg.Decode(ifile)
// 		if err != nil {
// 			return nil
// 		}
// 	} else if filetype == "png" {
// 		img, err = png.Decode(ifile)
// 		if err != nil {
// 			return nil
// 		}
// 	} else if filetype == "gif" {
// 		img, err = gif.Decode(ifile)
// 		if err != nil {
// 			return nil
// 		}
// 	}
// 	// deal with exif
// 	efile, err := os.Open(fpath)
// 	if err != nil {
// 		logrus.Warnf("could not open file for exif decoder: %s", fpath)
// 	}
// 	defer efile.Close()
// 	x, err := exif.Decode(efile)
// 	if err != nil {
// 		if x == nil {
// 			// ignore - image exif data has been already stripped
// 		}
// 		logrus.Errorf("failed reading exif data in [%s]: %s", fpath, err.Error())
// 	}
// 	if x != nil {
// 		orient, _ := x.Get(exif.Orientation)
// 		if orient != nil {
// 			logrus.Infof("%s had orientation %s", fpath, orient.String())
// 			img = reverseOrientation(img, orient.String())
// 		} else {
// 			logrus.Warnf("%s had no orientation - implying 1", fpath)
// 			img = reverseOrientation(img, "1")
// 		}
// 		imaging.Save(img, fpath)
// 	}
// 	return &img
// }

// // reverseOrientation amply`s what ever operation is necessary to transform given orientation
// // to the orientation 1
// func reverseOrientation(img image.Image, o string) *image.NRGBA {
// 	switch o {
// 	case "1":
// 		return imaging.Clone(img)
// 	case "2":
// 		return imaging.FlipV(img)
// 	case "3":
// 		return imaging.Rotate180(img)
// 	case "4":
// 		return imaging.Rotate180(imaging.FlipV(img))
// 	case "5":
// 		return imaging.Rotate270(imaging.FlipV(img))
// 	case "6":
// 		return imaging.Rotate270(img)
// 	case "7":
// 		return imaging.Rotate90(imaging.FlipV(img))
// 	case "8":
// 		return imaging.Rotate90(img)
// 	}
// 	logrus.Errorf("unknown orientation %s, expect 1-8", o)
// 	return imaging.Clone(img)
// }

// func GetSuffix(name string) (string, error) {
// 	if !strings.Contains(name, ".") {
// 		return name, errors.New("file names without file type suffix are not supported")
// 	}
// 	split := strings.Split(name, ".")
// 	return strings.ToLower(strings.TrimSpace(split[len(split)-1])), nil
// }
