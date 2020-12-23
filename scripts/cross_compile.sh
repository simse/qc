package="github.com/simse/qc"
package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" "darwin/amd64" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name='repo/'$VERSION'/'$package_name'-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name -ldflags "-X github.com/simse/qc/cmd.Version=${VERSION}" main.go

    # Generate checksum
    sha256sum $output_name >> $output_name'.sum'

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

# Generate patches
echo "Generating binary patches for past versions..."
for d in repo/* ; do
    version_split=(${d//\// })
    folder_version=${version_split[-1]}

    if [ $folder_version != $VERSION ]; then
        echo "Generating diff for ${folder_version} -> ${VERSION}"
        mkdir -p 'repo/'$VERSION'/patch'

        for platform in "${platforms[@]}"
        do
            platform_split=(${platform//\// })
            GOOS=${platform_split[0]}
            GOARCH=${platform_split[1]}

            old_binary=$d'/qc-'$GOOS'-'$GOARCH
            new_binary='repo/'$VERSION'/qc-'$GOOS'-'$GOARCH

            if [ $GOOS = "windows" ]; then
                old_binary+='.exe'
                new_binary+='.exe'
            fi

            # Check both binaries exists
            if test -f "$old_binary" && test -f "$new_binary"; then
                diff_name='repo/'$VERSION'/patch/qc-'$GOOS'-'$GOARCH'-'$folder_version'-'$VERSION'.patch'
                bsdiff $old_binary $new_binary $diff_name
            fi
        done
    fi
done
echo "Patch generation completed"