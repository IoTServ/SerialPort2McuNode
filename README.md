# SerialPort2McuNode
将计算机的串口直接映射到McuNode服务器
### 使用方法：
#### 需要在命令模式启动，加入必要的参数配置，参数如下：
 Usage:
  -B int
        Baud rate（波特率） (default 115200)
  -C string
        COM Port eg:COM3 (串口号) (default "COM3")
  -H string
        Host eg：eiot.club（服务器域名或ip） (default "127.0.0.1")
  -I string
        ID ,your id setting (自定义ID)eg:4567 (default "4567") 

例如：windowsAMD64 -B 115200 -H eiot.club -C COM3 -I 4567
就会将本机的COM3串口以115200波特率，id为4567映射到服务器：eiot.club
