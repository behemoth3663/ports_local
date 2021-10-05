#

PORTNAME=	kubeletctl
PORTVERSION=	1.8
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Client for kubelet

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
#GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	cyberark

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#workaround for can't find vendor/gitlab.com/poldi1405/go-ansi
#pre-build: .PHONY
#	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
