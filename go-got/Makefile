#

PORTNAME=	got
PORTVERSION=	0.7.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local ftp
PKGNAMEPREFIX=	go-
DISTNAME=	got

MAINTAINER=	none@freebsd.org
COMMENT=	Simple and fast concurrent downloader
WWW=		https://github.com/melbahja/got

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
USE_GITHUB=	yes
USE_GITLAB=	nodefault

GH_ACCOUNT=	melbahja

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
