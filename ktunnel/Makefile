# $FreeBSD$

PORTNAME=	ktunnel
PORTVERSION=	1.2.7
DISTVERSIONPREFIX=	v
CATEGORIES=	net

MAINTAINER=	none@freebsd.org
COMMENT=	Expose your local resources to kubernetes
                                                                   	
LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
#GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	omrikiei
PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor | grep -v '^=' >Makefile.deps && make makesum && make clean all
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
