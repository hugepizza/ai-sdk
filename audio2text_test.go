package aiqqsdk

import "testing"

var client *Audio2TextClient

func init() {
	client = NewAudio2TextClient("appid", "appkey", "http://myserv:20019/audio")
}

func TestUploadRemote(t *testing.T) {
	tid, err := client.UploadRemote("myserv:20019/tea.wav", AudioFormatWAV)
	if err != nil {
		t.Error(err)
	}
	t.Log(tid)
}
func TestUploadLocal(t *testing.T) {
	tid, err := client.UploadLocal("/home/wanlei/python/fx/tea.wav", AudioFormatWAV)
	if err != nil {
		t.Error(err)
	}
	t.Log(tid)
}
