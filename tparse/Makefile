#

PORTNAME=	tparse
PORTVERSION=	0.8.3
DISTVERSIONPREFIX=	v
CATEGORIES=	textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Analyzing and summarizing go test output

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mfridman

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	cd ${GO_WRKSRC} && ${SETENV} GIT_TAG=${DISTVERSIONPREFIX}${PORTVERSION} ${GO_CMD} generate -v ./...

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
