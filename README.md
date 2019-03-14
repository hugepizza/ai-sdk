# 腾讯AI

## audio2text 长语音转文字

### 音频必须符合16k或8K采样率、16bit采样位数、单声道

### 音频文件支持的格式有PCM、WAV、AMR、 SILK

### 可以使用ffmpeg调整音频的格式、声道和采样率 例:

```
ffmpeg -i test.mp3 -ac 1 -ar 8k -f wav test.wav

```

### 官方文档 `https://ai.qq.com/doc/wxasrlong.shtml`

### 目前个人用户免费 无限调用 2QPS 可以愉快玩耍