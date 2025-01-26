@echo off
echo Compilando codigo Go para AWS Lambda...

rem Establecer variables de entorno para la compilación
set GOOS=linux
set GOARCH=amd64

rem Compilar el archivo main.go
go build -o main
if %errorlevel% neq 0 (
    echo Error al compilar el archivo Go. Verifica tu codigo.
    pause
    exit /b %errorlevel%
)

rem Crear el archivo ZIP
echo Empaquetando el archivo main en main.zip...
powershell Compress-Archive -Path main -DestinationPath main.zip -Force
if %errorlevel% neq 0 (
    echo Error al empaquetar el archivo.
    pause
    exit /b %errorlevel%
)

echo Empaquetado exitoso. El archivo main.zip esta listo para subir a Lambda.
pause


