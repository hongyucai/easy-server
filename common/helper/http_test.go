package helper

import "testing"

func Test_Get(t *testing.T) {
	h := NewHttpSend(GetUrlBuild("http://127.0.0.1/test.php", map[string]string{"name": "xiaochuan"}))
	_, err := h.Get()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Test_Post(t *testing.T) {
	h := NewHttpSend("http://127.0.0.1/test.php")
	h.SetHeader(map[string]string{"cs": "cs"})
	h.SetBody(map[string]interface{}{"name": "xiaochuan"})
	_, err := h.Post()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Test_Json(t *testing.T) {
	h := NewHttpSend("http://127.0.0.1/test.php")
	h.SetSendType("JSON")
	h.SetBody(map[string]interface{}{"name": "xiaochuan"})
	_, err := h.Post()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Benchmark_GET(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewHttpSend(GetUrlBuild("http://127.0.0.1/test.php", map[string]string{"name": "xiaochuan"}))
		_, err := h.Get()
		if err != nil {
			b.Error("请求错误:", err)
		} else {
			b.Log("正常返回")
		}
	}
}
func test_postFile()  {
	target_url := "http://localhost/upload"
	filename := "./example.pdf"
	PostFile(target_url,filename,"media")
}