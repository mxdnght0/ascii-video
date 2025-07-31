# 🎞️ ASCII Video Converter

Проект на Go, который конвертирует видеофайл в ASCII-анимацию и выводит её прямо в терминал. Используется `ffmpeg` для извлечения кадров и алгоритм преобразования изображений в ASCII-графику. Подходит для визуального представления видеороликов в текстовом виде.

## 📦 Установка

### 1. Установите зависимости

- [Go](https://golang.org/dl/) 1.18 или новее
- [ffmpeg](https://ffmpeg.org/download.html)

Убедитесь, что `ffmpeg` доступен в терминале:

```bash
ffmpeg -version
```

### 2. Клонируйте репозиторий

```bash
git clone https://github.com/yourusername/ascii-video.git
cd ascii-video
```

### 3. Установите зависимости Go

```bash
go mod tidy
```

### 4. Запустите программу

```bash
go run main.go path/to/your/video.mp4
```
