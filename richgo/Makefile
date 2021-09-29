#

PORTNAME=	richgo
PORTVERSION=	0.3.9
DISTVERSIONPREFIX=	v
CATEGORIES=	devel

MAINTAINER=	none@freebsd.org
COMMENT=	Enrich `go test` outputs with text decorations

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kyoh86

#GO_TARGET=	./cmd/${PORTNAME}

#GO_BUILDFLAGS+=	-ldflags="-s -X 'github.com/aquasecurity/tfsec/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps && doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
