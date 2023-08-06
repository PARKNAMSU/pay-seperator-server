#!/bin/bash

export PROJECT_VERSION=$(echo $(sed -n -e '1,1p' VERSION) | cut -d ' ' -f2)
export READ_ME=$(sed -n -e '2,$p' VERSION) 

echo "previous ${PROJECT_VERSION}"

export MAJOR=$(echo $PROJECT_VERSION | cut -d '.' -f1)
export MINOR=$(echo $PROJECT_VERSION | cut -d '.' -f2)
export PATCH=$(echo $PROJECT_VERSION | cut -d '.' -f3)

if [ "$1" == "" ] || [ "$1" == "--patch" ]; then
    PATCH="$(($PATCH + 1))"
elif [ "$1" == "--minor" ]; then
    MINOR="$(($MINOR + 1))"
elif [ "$1" == "--major" ]; then
    MAJOR="$(($MAJOR + 1))"
else
    echo "invalid command"
    exit 0
fi


# if [ "$PATCH" -ge "10" ]; then 
#     PATCH=0
#     MINOR="$(($MINOR + 1))"
# fi

# if [ "$MINOR" -ge "10" ]; then
#     echo "here2"
#     MINOR=0
#     MAJOR="$(($MAJOR + 1))"
# fi

PROJECT_VERSION="V ${MAJOR}.${MINOR}.${PATCH}"

echo "now ${PROJECT_VERSION}"

echo $PROJECT_VERSION > ./VERSION

cat << EOF >> ./VERSION
$READ_ME
EOF

# git add .
# git commit -m "${PROJECT_VERSION}"