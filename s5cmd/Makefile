# $FreeBSD$

PORTNAME=	s5cmd
PORTVERSION=	1.2.0
DISTVERSIONPREFIX=	v
CATEGORIES=	net

MAINTAINER=	none@freebsd.org
COMMENT=	Parallel S3 and local filesystem execution tool

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	peak

#GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

post-patch: .PHONY
	${REINPLACE_CMD} 's|Version[[:space:]]*=[[:space:]]*"v0\.0\.0"|Version = "v${PORTVERSION}"|' \
		${WRKSRC}/version/version.go

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.sinclude "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
