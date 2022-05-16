#

PORTNAME=	terraformer
PORTVERSION=	0.8.20
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Generate terraform files from existing infrastructure

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	GoogleCloudPlatform

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

do-build: .PHONY
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod download && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} build ${GO_BUILDFLAGS:N-mod*} -o ${GO_WRKDIR_BIN}/${PORTNAME}

.include <bsd.port.mk>
