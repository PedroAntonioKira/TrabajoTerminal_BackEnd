@echo off
echo Compilando codigo Go para AWS Lambda...

rem Establecer variables de entorno para la compilación
set GOOS=linux
set GOARCH=amd64

rem Compilar el archivo main.go con el nombre bootstrap
go build -tags lambda.norpc -o bootstrap main.go
if %errorlevel% neq 0 (
    echo Error al compilar el archivo Go. Verifica tu codigo.
    pause
    exit /b %errorlevel%
)

rem Crear el archivo ZIP
echo Empaquetando el archivo bootstrap en main.zip...
del main.zip
tar.exe -a -cf main.zip bootstrap
if %errorlevel% neq 0 (
    echo Error al empaquetar el archivo.
    pause
    exit /b %errorlevel%
)

echo Empaquetado exitoso. El archivo main.zip esta listo para subir a Lambda.
pause
