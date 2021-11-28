#

PORTNAME=	dstp
PORTVERSION=	0.4.0
DISTVERSIONPREFIX=	v
CATEGORIES=	net

MAINTAINER=	none@freebsd.org
COMMENT=	Run common networking tests against site

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	ycd

GO_TARGET=	./cmd/${PORTNAME}

#GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps && doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
