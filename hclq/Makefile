# $FreeBSD$

PORTNAME=	hclq
PORTVERSION=	0.5.3
#DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Command-line processor for HashiCorp Command Language files

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
#GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	mattolenik
GH_TAGNAME=	e55c48c

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#workaround for can't find vendor/gitlab.com/poldi1405/go-ansi
#pre-build: .PHONY
#	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor

.sinclude "${.CURDIR}/Makefile.deps"

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/dads-after-port-mk.html
.include <bsd.port.pre.mk>
GO_BUILDFLAGS:=${GO_BUILDFLAGS:C|-ldflags=(.*)|-ldflags="\\1 -X github.com/${GH_ACCOUNT}/${PORTNAME}/cmd.version=${PORTVERSION}"|}
.include <bsd.port.post.mk>
