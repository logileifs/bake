OS=$1
ARCH=$2
VERSION=$3
COMMIT=$4
BUILD=$5
DATE=$6

GOOS=$OS GOARCH=$ARCH go build -o bake-"$OS"-"$ARCH"-"$VERSION" \
	-ldflags "-X main.version=$VERSION \
	-X main.commit=$COMMIT \
	-X main.build=$BUILD \
	-X main.date=$DATE"
