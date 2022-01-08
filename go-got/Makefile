#

PORTNAME=	got
PORTVERSION=	0.6.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local ftp
PKGNAMEPREFIX=	go-
DISTNAME=	got

MAINTAINER=	none@freebsd.org
COMMENT=	Simple and fast concurrent downloader

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

GO_MODULE=	github.com/melbahja/${PORTNAME}
GH_ACCOUNT=	${GO_MODULE:H:T}
GH_PROJECT=	${GO_MODULE:T}

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include <bsd.port.mk>
