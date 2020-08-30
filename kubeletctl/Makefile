# $FreeBSD$

PORTNAME=	kubeletctl
PORTVERSION=	1.5
DISTVERSIONPREFIX=	v
CATEGORIES=	ftp

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

.sinclude "${.CURDIR}/Makefile.deps"
.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
