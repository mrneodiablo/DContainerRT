package images

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"main/config"
	"main/utils"
	"net/http"
	"os"
	"strings"

	"github.com/containerd/btrfs"
	uuid "github.com/nu7hatch/gouuid"
)

func InitContainerImage() error {
	cfg, _ := config.Config()

	_, err := os.Stat(cfg.PathImages)
	if err != nil {
		os.MkdirAll(cfg.PathImages, 0755)
	}

	return nil
}

func RemoveContainerImages(idImages string) error {
	cfg, _ := config.Config()
	err := btrfs.SubvolDelete(cfg.PathImages + "/" + idImages)
	return err
}


func ListContainerImages() ([]ImageInfo, error) {
	cfg, _ := config.Config()
	var out []ImageInfo
	subv, err := btrfs.SubvolList(cfg.PathImages)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	for index := 0; index < len(subv); index++ {
		info, _ := ioutil.ReadFile(cfg.PathImages + "/" + subv[index].Name  + "/img.source")

		if string(subv[index].Name) != "" && string(info) != "" {
			data := ImageInfo{
				Name:    strings.Split(string(info), ":")[0],
				Tag:     strings.Split(string(info), ":")[1],
				Id:      strings.Split(string(info), ":")[2],
				Created: subv[index].Offset,
			}
			out = append(out, data)
		}

	}

	fmt.Println(out)
	return out, errors.New("")
}

func PullContainerImage(repoName string, tags string) (string, error) {
	cfg, _ := config.Config()
	client := &http.Client{}

	//get authorization token
	var author DockerHubAuthorization
	jsonDataToken, _ := http.Get("https://auth.docker.io/token?service=registry.docker.io&scope=repository:library/" + repoName + ":pull")
	body, err := ioutil.ReadAll(jsonDataToken.Body)
	json.Unmarshal(body, &author)

	// get blobSum
	var dockerImageManifest DockerImageManifest
	req, err := http.NewRequest("GET", "https://registry-1.docker.io/v2/library/"+repoName+"/manifests/"+tags, nil)
	req.Header.Add("Authorization", "Bearer "+author.Token)
	jsonDataBlobSum, _ := client.Do(req)
	body, err = ioutil.ReadAll(jsonDataBlobSum.Body)
	//fmt.Println(string(body))
	json.Unmarshal(body, &dockerImageManifest)

	if dockerImageManifest.Errors != nil {
		fmt.Println(dockerImageManifest.Errors[0]["code"])
		return "", errors.New(dockerImageManifest.Errors[0]["code"])
	}


	imageId, _ := uuid.NewV4()

	// create subvolume
	err = btrfs.SubvolCreate(cfg.PathImages + "/img_" + imageId.String())
	if err != nil {
		return imageId.String(), err
	}

	for index := 0; index < len(dockerImageManifest.FsLayers); index++ {
		req, err := http.NewRequest("GET", "https://registry-1.docker.io/v2/library/"+repoName+"/blobs/"+dockerImageManifest.FsLayers[index]["blobSum"], nil)
		req.Header.Add("Authorization", "Bearer "+author.Token)
		jsonDataLayer, _ := client.Do(req)
		hash_tmp := strings.Split(dockerImageManifest.FsLayers[index]["blobSum"], ":")[1]
		fmt.Println("Download Layer:" + hash_tmp)

		// uuID
		err = utils.Untar("/tmp/"+imageId.String()+"/", jsonDataLayer.Body)
		if err != nil {
			return "", err
		}
	}
	ioutil.WriteFile("/tmp/"+imageId.String()+"/img.source", []byte(repoName+":"+tags+":"+"img_"+imageId.String()), 0644)
	err = utils.CopyDirectory("/tmp/"+imageId.String(), cfg.PathImages+"/img_"+imageId.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	os.RemoveAll("/tmp/" + imageId.String())
	fmt.Println(repoName + ":" + tags + " " + imageId.String())
	return imageId.String(), nil
}
