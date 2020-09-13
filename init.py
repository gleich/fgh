import os
import re


def ls_files(dir):
    """List all files in the current directory

    Args:
        string: Path to list from

    Returns:
        list: List of all the files

    TAKEN FROM THE FOLLOWING STACKOVERFLOW ANSWER:
    https://stackoverflow.com/a/58170903/10933841
    """
    files = []
    for item in os.listdir(dir):
        abspath = os.path.join(dir, item)
        try:
            if os.path.isdir(abspath):
                files = files + ls_files(abspath)
            else:
                files.append(abspath)
        except FileNotFoundError as err:
            print("invalid directory\n", "Error: ", err)
    return files


project_name = input("What is the name of this project?\n").lower()
project_description = input("\nWhat is the description for this project?\n")
github_username = input("\nWhat is your github username?\n")
project_author_email = input("\nWhat is your email?\n")
project_author_full_name = input("\nWhat is your first and last name?\n")
docker_username = input("\nWhat is your docker username?\n")

cleaned_files = []
for filename in ls_files("."):
    ignore = False
    if not (
        filename.startswith("./.git/")
        or "__pycache__" in filename
        or ".pytest_cache" in filename
        or filename == "./init.py"
    ):
        cleaned_files.append(filename)

for filename in cleaned_files:
    with open(filename, "r") as file:
        cleaned_lines = []
        lines = file.readlines()
        for line in lines:
            cleaned_lines.append(
                line.replace("project_name", project_name)
                .replace("project_description", project_description)
                .replace("github_username", github_username)
                .replace("project_author_email", project_author_email)
                .replace("docker_username", docker_username)
                .replace("project_author_full_name", project_author_full_name)
                .replace("p\\roject_name", "project_name")
                .replace("p\\roject_description", "project_description")
                .replace("g\\ithub_username", "github_username")
                .replace("p\\roject_author_email", "project_author_email")
                .replace("d\\ocker_username", "docker_username")
                .replace("p\\roject_author_full_name", "project_author_full_name")
            )
    with open(filename, "w") as file:
        file.write("".join(cleaned_lines))

with open("main.go", "w") as main_file:
    main_file.write(
        """package main

import (
	"github.com/Matt-Gleich/logoru"
)

func main() {
	logoru.Success("Program booted up!")
}
"""
    )

os.remove("go.mod")
os.remove("go.sum")
os.system(f"go mod init github.com/{github_username}/{project_name}")
os.system("go get -u github.com/Matt-Gleich/logoru")
os.system("go mod tidy")

print("\n\n\n\nEverything filled in!\nPlease now delete this file!")
