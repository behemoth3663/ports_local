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

GO_BUILDFLAGS+=	-ldflags="-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=v${PORTVERSION}"

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
