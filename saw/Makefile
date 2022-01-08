#

PORTNAME=	saw
PORTVERSION=	0.2.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Fast, multi-purpose tool for AWS CloudWatch Logs

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	TylerBrock

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${PORTNAME}
PLIST_FILES=	bin/${PORTNAME}

#USES=go:modules but there is no go.mod
pre-build: .PHONY
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
