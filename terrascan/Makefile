# $FreeBSD$

PORTNAME=	terrascan
PORTVERSION=	1.2.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Compliance and security violations in terraform code

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

#go:modules - can't resolve difficulties
USES=		go
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	accurics

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
#.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
