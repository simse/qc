#/bin/python
import glob
import json
import collections
from os import path
from pkg_resources import packaging


nested_dict = lambda: collections.defaultdict(nested_dict)

manifest = {}
manifest["latest_release"] = {}
releases = []

current_latest_release = "0.0.1"

# Get all versions
for name in glob.glob('./repo/**'): 
    if ".json" not in name:
        version = name.replace("./repo/", "")
        release = {
            "version": version,
        }

        binaries = nested_dict()
        patches = nested_dict()

        # Look for all binaries in version folder
        for binary in glob.glob("./repo/" + version + "/qc-*-*"):
            if ".sum" not in binary:
                # Attempt to find checksum
                checksum = ""
                checksum_file = binary + ".sum"
                if path.exists(checksum_file):
                    with open(checksum_file, "r") as c:
                        checksum = c.read().split("  ")[0]

                binary_os = binary.split("-")[1]
                binary_arch = binary.split("-")[2].split(".")[0]

                binaries[binary_os][binary_arch] = {
                    "path": "https://repo.simse.io/qc/" + binary.replace("./repo/", ""),
                    "checksum": checksum
                }

        # Look for patches
        for patch in glob.glob("./repo/" + version + "/patch/*"):
            patch_name = patch.split("/")[-1]
            patch_url = "https://repo.simse.io/qc/" + version + "/patch/" + patch_name

            # print(patch_name)

            os = patch_name.split("-")[1]
            arch = patch_name.split("-")[2]

            from_version = patch_name.split("-")[3]
            to_version = patch_name.split("-")[4]

            patches[os][arch][from_version] = patch_url

        
        release["binaries"] = binaries
        release["patches"] = patches
        releases.append(release)

        if packaging.version.parse(version) > packaging.version.parse(current_latest_release):
            current_latest_release = version
            manifest["latest_release"] = release

manifest["releases"] = releases

with open("./repo/manifest.json", "w+") as manifest_file:
    manifest_file.write(json.dumps(manifest))