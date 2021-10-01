#

PORTNAME=	tfsec
PORTVERSION=	0.58.12
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

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps && doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
