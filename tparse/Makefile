PORTNAME=	tparse
PORTVERSION=	0.12.2
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Analyzing and summarizing go test output
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mfridman

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
