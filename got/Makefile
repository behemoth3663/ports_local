# $FreeBSD$

PORTNAME=	got
PORTVERSION=	0.2.0
DISTVERSIONPREFIX=	v
CATEGORIES=	ftp

MAINTAINER=	none@freebsd.org
COMMENT=	Simple and fast concurrent downloader

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	melbahja

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.sinclude "${.CURDIR}/Makefile.deps"
.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
