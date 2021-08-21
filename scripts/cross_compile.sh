package="github.com/simse/qc"
package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("darwin/amd64" "linux/amd64" "linux/arm64" "darwin/arm64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name='build/'$package_name'-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name -ldflags "-X github.com/simse/qc/internal/update.Version=${VERSION}" main.go

    echo 'Compiled' $output_name

    # Generate checksum
    # sha256sum $output_name >> $output_name'.sum'

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done