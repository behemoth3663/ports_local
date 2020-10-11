# $FreeBSD$

PORTNAME=	tfsec
PORTVERSION=	0.30.1
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

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor >Makefile.deps && make makesum && mkcal
.sinclude "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
