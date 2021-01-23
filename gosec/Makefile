# $FreeBSD$

PORTNAME=	gosec
PORTVERSION=	2.6.1
DISTVERSIONPREFIX=	v
CATEGORIES=	devel security local

MAINTAINER=	none@freebsd.org
COMMENT=	Golang Security Checker

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	securego

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.sinclude "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
