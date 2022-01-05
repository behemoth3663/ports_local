#

PORTNAME=	cwl
PORTVERSION=	4.1.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Best way to tail AWS CloudWatch Logs from your terminal

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	lucagrulla
GH_PROJECT=	cw

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

post-patch: .PHONY
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor
	chmod -R u+w ${WRKDIR}/.gopath

.include <bsd.port.mk>
