#

PORTNAME=	tfsec
PORTVERSION=	0.58.5
DISTVERSIONPREFIX=	v
CATEGORIES=	security

MAINTAINER=	none@freebsd.org
COMMENT=	Static analysis powered security scanner for your terraform code

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	liamg

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -X 'github.com/aquasecurity/tfsec/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
