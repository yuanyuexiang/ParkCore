package utils

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

// Get HttpGet
func Get(url string) (data []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	//请求头设置
	request.Header.Add("Content-Type", "application/json") //json请求
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 1 * time.Second
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("server error")
		return
	}

	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
	data = body
	return
}

// Post HttpPost
func Post(url string, postData []byte) (data []byte, err error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(string(postData)))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 50 * time.Second
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("server error")
		return
	}

	body, err := io.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println("response body ", bodystr)
	data = body
	return
}

// PostFile HttpPost
func PostFile(url string, filePath string) (data []byte, err error) {
	// 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 创建buffer
	var buffer bytes.Buffer
	// 通过包multipart实现了MIME的multipart解析
	writer := multipart.NewWriter(&buffer)
	part, err := writer.CreateFormFile("", filePath)
	if err != nil {
		panic(err)
	}
	// 拷贝文件内容到新建FormFile中
	_, err = io.Copy(part, f)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("server error")
		return
	}

	body, err := io.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println("response body ", bodystr)
	data = body
	return
}

// Put HttpPut
func Put(url string, postData []byte) (data []byte, err error) {
	request, err := http.NewRequest("PUT", url, strings.NewReader(string(postData)))
	if err != nil {
		return
	}
	//请求头设置
	request.Header.Add("Content-Type", "application/json") //json请求
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 50 * time.Second
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("server error")
		return
	}

	body, err := io.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println("response body ", bodystr)
	data = body
	return
}

// Delete HttpDelete
func Delete(url string, postData []byte) (data []byte, err error) {
	request, err := http.NewRequest("DELETE", url, strings.NewReader(string(postData)))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 1 * time.Second
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("server error")
		return
	}

	body, err := io.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println("response body ", bodystr)
	data = body
	return
}
