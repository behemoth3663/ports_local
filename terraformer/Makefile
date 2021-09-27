#

PORTNAME=	terraformer
PORTVERSION=	0.8.17
#DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

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
	pwd="$${PWD}" && trap "cd \"$${pwd}\"" INT EXIT && cd ${WRKSRC} && \
	${GO_CMD} mod download && \
	${GO_CMD} build ${GO_BUILDFLAGS:N-mod*} -o ${GO_WRKDIR_BIN}/${PORTNAME}

.include <bsd.port.mk>
