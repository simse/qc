docker build -t simse/qc:dev docker/
clear
docker run  -v /home/simon/projects/qc/samples:/samples -v /home/simon/projects/qc:/app -v /home/simon/projects/qc/.cache:/root/go simse/qc:dev go run main.go $1