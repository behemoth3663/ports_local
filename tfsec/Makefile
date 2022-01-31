#

PORTNAME=	tfsec
PORTVERSION=	1.0.6
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Static analysis powered security scanner for your terraform code

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	liamg

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'github.com/aquasecurity/${PORTNAME}/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps && doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
