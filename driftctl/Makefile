#

PORTNAME=	driftctl
PORTVERSION=	0.15.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Detect, track and alert on infrastructure drift

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.md

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	cloudskiff

#GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/pkg/version.version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/build.env=release'"

#_OPTIONS_extract=	501:pre-post-extract
#pre-post-extract: .PHONY
#	${MV} ${WRKDIR}/go-commons-0.7.0 ${WRKDIR}/gruntwork-cli-0.7.0

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
