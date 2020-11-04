package main

func main() {
  // 1. ยิง Request เก็บไว้ใน response
  url := ""
  response, err := http.Get(url)
  err != nil {
    log.Println(err)
  }
  
  // 2. ปิดตัวอ่าน เขียน response
  defer response.Body.Close()
  
  // 3. สร้างตัวแปรรับ body ตัวนึง ในที่นี้เราไม่รู้ type ของ response เลยใส่เเป็น interface{}
  var body interface{}
  
  // 4. Decode ตัว request
  json.NewDecoder(response.Body).Decode(&body)
  
  // 5. ตัว body จะมีข้อมูลจริงๆ ออกมา
  log.Println(body)
  
}
